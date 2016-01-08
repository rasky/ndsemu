package hw

import (
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

var lastStatus int

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
					lastStatus, _ = strconv.Atoi(string(matches[1]))
				}
			}
			time.Sleep(5 * time.Second)
		}
	}()
}

func readBatteryStatusMac() int {
	return lastStatus
}
