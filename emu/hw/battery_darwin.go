package hw

import (
	"os/exec"
	"regexp"
	"strconv"
	"sync/atomic"
	"time"
)

var lastStatus int32

func init() {
	ReadBatteryStatus = readBatteryStatusMac
	lastStatus = 100

	go func() {
		rx, err := regexp.Compile("(\\d+)%")
		if err != nil {
			panic(err)
		}

		for {
			out, err := exec.Command("pmset", "-g", "batt").Output()
			if err == nil {
				matches := rx.FindSubmatch(out)
				if len(matches) > 1 {
					val, _ := strconv.Atoi(string(matches[1]))
					atomic.StoreInt32(&lastStatus, int32(val))
				}
			}
			time.Sleep(5 * time.Second)
		}
	}()
}

func readBatteryStatusMac() int {
	return int(atomic.LoadInt32(&lastStatus))
}
