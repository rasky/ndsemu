package main

import (
	"flag"
	"fmt"
	"ndsemu/e2d"
	"ndsemu/emu/hw"
	log "ndsemu/emu/logger"
	"ndsemu/homebrew"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/debug"
	"runtime/pprof"
	"strings"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type CpuNum int

const (
	CpuNds9 CpuNum = 0
	CpuNds7 CpuNum = 1
)

/*
 * NDS9: ARM946E-S, architecture ARMv5TE, 66Mhz
 * NDS7: ARM7TDMI, architecture ARMv4T, 33Mhz
 *
 */

const cFirmwareDefault = "bios/firmware.bin"

var (
	skipBiosArg  = flag.Bool("s", false, "skip bios and run immediately")
	flagDebug    = flag.Bool("debug", false, "run with debugger")
	cpuprofile   = flag.String("cpuprofile", "", "write cpu profile to file")
	flagLogging  = flag.String("log", "", "enable logging for specified modules")
	flagJit      = flag.Bool("jit", false, "use JIT for emulation (unstable, eats memory)")
	flagVsync    = flag.Bool("vsync", true, "run at normal speed (60 FPS)")
	flagFirmware = flag.String("firmware", cFirmwareDefault, "specify the firwmare file to use")
	flagHbrewFat = flag.String("homebrew-fat", "", "FAT image to be mounted for homebrew ROM")

	nds7     *NDS7
	nds9     *NDS9
	KeyState = make([]uint8, 256)
)

func init() {
	// For now, disable GC during JIT. This is required because the Go runtime
	// doesn't like generated code on the Go stack without a stackmap.
	// Use a hack because we need to do this before any goroutine is started
	for _, arg := range os.Args {
		if arg == "-jit" || arg == "-jit=true" {
			debug.SetGCPercent(-1)
		}
	}
}

func main() {
	sdl.Main(main1)
}

func main1() {
	flag.Parse()

	// Check whether there is a local firmware copy, otherwise
	// create one (to handle read/write)
	if (*flagFirmware)[0] != '/' {
		bindir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		*flagFirmware = filepath.Join(bindir, *flagFirmware)
	}

	if _, err := os.Stat(*flagFirmware); err != nil {
		log.ModEmu.FatalZ("cannot open firmware").Error("err", err).End()
	}

	firstboot := false
	fwsav := *flagFirmware + ".sav"
	if _, err := os.Stat(fwsav); err != nil {
		fw, err := os.ReadFile(*flagFirmware)
		if err != nil {
			log.ModEmu.FatalZ("cannot load firwmare:").Error("err", err).End()
		}
		err = os.WriteFile(fwsav, fw, 0777)
		if err != nil {
			log.ModEmu.FatalZ("cannot save firwmare:").Error("err", err).End()
		}
		firstboot = true
	}

	Emu = NewNDSEmulator(fwsav, *flagJit)

	// Check if the NDS ROM is homebrew. If so, directly load it into slot2
	// like PassMe does.
	if len(flag.Args()) > 0 {

		if hbrew, _ := homebrew.Detect(flag.Arg(0)); hbrew {
			if err := Emu.Hw.Sl2.MapCartFile(flag.Arg(0)); err != nil {
				log.ModEmu.FatalZ(err.Error()).End()
			}
			if len(flag.Args()) > 1 {
				log.ModEmu.FatalZ("slot2 ROM specified but slot1 ROM is homebrew").End()
			}
			// FIXME: also load the ROM in slot1. Theoretically, for a full
			// Passme emulation, the ROM in slot1 should be patched by PassMe,
			// but it looks like the firmware we're using doesn't need it.
			if err := Emu.Hw.Gc.MapCartFile(flag.Arg(0)); err != nil {
				log.ModEmu.FatalZ(err.Error()).End()
			}

			// See if we are asked to load a FAT image as well. If so, we concatenate it
			// to the ROM, and then do a DLDI patch to make libfat find it.
			if *flagHbrewFat != "" {
				if err := Emu.Hw.Sl2.HomebrewMapFatFile(*flagHbrewFat); err != nil {
					log.ModEmu.FatalZ(err.Error()).End()
				}

				if err := homebrew.FcsrPatchDldi(Emu.Hw.Sl2.Rom); err != nil {
					log.ModEmu.FatalZ(err.Error()).End()
				}
			}

			// Activate IDEAS-compatibile debug output on both CPUs
			// (use a special SWI to write messages in console)
			homebrew.ActivateIdeasDebug(nds9.Cpu)
			homebrew.ActivateIdeasDebug(nds7.Cpu)
		} else if strings.HasSuffix(flag.Arg(0), ".nds") {
			// Map Slot1 cart file (NDS ROM)
			if err := Emu.Hw.Gc.MapCartFile(flag.Arg(0)); err != nil {
				log.ModEmu.FatalZ(err.Error()).End()
			}

			// Map save file for Slot1
			if err := Emu.Hw.Bkp.MapSaveFile(flag.Arg(0) + ".sav"); err != nil {
				log.ModEmu.FatalZ(err.Error()).End()
			}

			// If specified, map Slot2 cart file (GBA ROM)
			if len(flag.Args()) > 1 {
				if err := Emu.Hw.Sl2.MapCartFile(flag.Arg(1)); err != nil {
					log.ModEmu.FatalZ(err.Error()).End()
				}
			}

			if *flagHbrewFat != "" {
				log.ModEmu.FatalZ("cannot specify -homebrew-fat for non-homebrew ROM").End()
			}
		} else if strings.HasSuffix(flag.Arg(0), ".gba") {
			if err := Emu.Hw.Sl2.MapCartFile(flag.Arg(0)); err != nil {
				log.ModEmu.FatalZ(err.Error()).End()
			}
			if len(flag.Args()) > 1 {
				log.ModEmu.FatalZ("cannot specify multiple ROMs after GBA rom").End()
			}
			if *flagHbrewFat != "" {
				log.ModEmu.FatalZ("cannot specify -homebrew-fat for non-homebrew ROM").End()
			}
		} else {
			log.ModEmu.FatalZ("unrecognized ROM type").String("rom", flag.Arg(0)).End()
		}
	}

	if err := Emu.Hw.Ff.MapFirmwareFile(fwsav); err != nil {
		log.ModEmu.FatalZ(err.Error()).End()
	}
	if firstboot {
		Emu.Hw.Rtc.ResetDefaults()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		f, err := os.Create("ram.dump")
		if err == nil {
			f.Write(Emu.Mem.Ram[:])
			f.Close()
		}
		f, err = os.Create("wram.dump")
		if err == nil {
			f.Write(Emu.Hw.Mc.wram[:])
			f.Write(Emu.Mem.Wram[:])
			f.Close()
		}
		for i := 0; i < len(Emu.Hw.Mc.vram); i++ {
			char := 'a' + i
			f, err = os.Create(fmt.Sprintf("vram-%c.dump", char))
			if err == nil {
				f.Write(Emu.Hw.Mc.vram[i][:])
				f.Close()
			}
		}
		f, err = os.Create("vram-bg-a.dump")
		if err == nil {
			v := Emu.Hw.Mc.VramLinearBank(0, e2d.VramLinearBG, 0)
			v.Dump(f)
			v = Emu.Hw.Mc.VramLinearBank(0, e2d.VramLinearBG, 256*1024)
			v.Dump(f)
			f.Close()
		}
		f, err = os.Create("vram-bg-b.dump")
		if err == nil {
			v := Emu.Hw.Mc.VramLinearBank(1, e2d.VramLinearBG, 0)
			v.Dump(f)
			f.Truncate(128 * 1024)
			f.Close()
		}
		f, err = os.Create("vram-bgextpal-a.dump")
		if err == nil {
			v := Emu.Hw.Mc.VramLinearBank(0, e2d.VramLinearBGExtPal, 0)
			v.Dump(f)
			f.Close()
		}
		f, err = os.Create("vram-bgextpal-b.dump")
		if err == nil {
			v := Emu.Hw.Mc.VramLinearBank(1, e2d.VramLinearBGExtPal, 0)
			v.Dump(f)
			f.Close()
		}

		f, err = os.Create("oam.dump")
		if err == nil {
			f.Write(Emu.Mem.OamRam[:])
			f.Close()
		}

		f, err = os.Create("texture.dump")
		if err == nil {
			texbank := Emu.Hw.Mc.VramTextureBank()
			for i := 0; i < 16; i++ {
				f.Write(texbank.Slots[i])
			}
			f.Close()
		}

		f, err = os.Create("texpal.dump")
		if err == nil {
			texbank := Emu.Hw.Mc.VramTexturePaletteBank()
			for i := 0; i < 8; i++ {
				f.Write(texbank.Slots[i])
			}
			f.Close()
		}

		if *cpuprofile != "" {
			pprof.StopCPUProfile()
		}
		os.Exit(1)
	}()

	if *skipBiosArg {
		if err := InjectGamecard(Emu.Hw.Gc, Emu.Mem); err != nil {
			fmt.Println(err)
			return
		}

		// Shared wram: map everything to ARM7
		Emu.Hw.Mc.WramCnt.Write8(0, 3)

		// Set post-boot flag to 1
		nds9.misc.PostFlg.Value = 1
		nds7.misc7.PostFlg.Value = 1

		nds9.Irq.Ime.Value = 0x1
		nds7.Irq.Ime.Value = 0x1
		nds9.Irq.Ie.Value = uint32(IrqIpcRecvFifo | IrqTimers | IrqVBlank)
		nds7.Irq.Ie.Value = uint32(IrqIpcRecvFifo | IrqTimers | IrqVBlank)

		// VRAM: map everything in "LCDC mode"
		Emu.Hw.Mc.VramCntA.Write8(0, 0x80)
		Emu.Hw.Mc.VramCntB.Write8(0, 0x80)
		Emu.Hw.Mc.VramCntC.Write8(0, 0x80)
		Emu.Hw.Mc.VramCntD.Write8(0, 0x80)
		Emu.Hw.Mc.VramCntE.Write8(0, 0x80)
		Emu.Hw.Mc.VramCntF.Write8(0, 0x80)
		Emu.Hw.Mc.VramCntG.Write8(0, 0x80)
		Emu.Hw.Mc.VramCntH.Write8(0, 0x80)
		Emu.Hw.Mc.VramCntI.Write8(0, 0x80)

		// Gamecard: skip directly to key2 status
		Emu.Hw.Gc.stat = gcStatusKey2

		nds9.Cp15.ConfigureControlReg(0x52078, 0x00FF085)
	}

	if *flagDebug {
		Emu.StartDebugger()
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.ModEmu.FatalZ(err.Error())
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if *flagLogging != "" {
		var modmask log.ModuleMask
		for _, modname := range strings.Split(*flagLogging, ",") {
			if modname == "all" {
				modmask |= log.ModuleMaskAll
			} else if m, found := log.ModuleByName(modname); found {
				modmask |= m.Mask()
			} else {
				log.ModEmu.FatalZ("invalid module name").String("name", modname).End()
			}
		}
		log.EnableDebugModules(modmask)
	}

	hwout := hw.NewOutput(hw.OutputConfig{
		Title:             "NDSEmu - Nintendo DS Emulator",
		Width:             256,
		Height:            192 + 90 + 192,
		FramePerSecond:    60,
		NumBackBuffers:    3,
		EnforceSpeed:      *flagVsync,
		AudioFrequency:    cAudioFreq,
		AudioChannels:     2,
		AudioSampleSigned: true,
	})
	hwout.EnableVideo(true)
	hwout.EnableAudio(true)

	var fprof *os.File
	profiling := 0

	KeyState = hw.GetKeyboardState()
	for hwout.Poll() {
		if KeyState[hw.SCANCODE_P] != 0 {
			time.Sleep(1 * time.Second)
		}
		if KeyState[hw.SCANCODE_L] != 0 && profiling == 0 {
			fprof, _ = os.Create("profile.dump")
			pprof.StartCPUProfile(fprof)
			profiling = Emu.framecount
		}
		if profiling > 0 && profiling <= Emu.framecount-120 {
			pprof.StopCPUProfile()
			fprof.Close()
			fprof = nil
			profiling = 0
			log.ModEmu.Warnf("profile dumped")
		}

		x, y, btn := hwout.GetMouseState()
		y -= 192 + 90
		pendown := btn&hw.MouseButtonLeft != 0
		Emu.Hw.Key.SetPenDown(pendown)
		Emu.Hw.Tsc.SetPen(pendown, x, y)

		v, a := hwout.BeginFrame()
		exit := Emu.RunOneFrame(v, ([]int16)(a))
		hwout.EndFrame(v, a)
		if exit {
			fmt.Println("System was powered off")
			break
		}
	}
}
