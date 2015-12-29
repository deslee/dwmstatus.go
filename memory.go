package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)

func updateMemUse() string {
	var file, err = os.Open("/proc/meminfo")
	if err != nil {
		return "ERR"
	}
	defer file.Close()

	// done must equal the flag combination (0001 | 0010 | 0100 | 1000) = 15
	var total, used, done = 0, 0, 0
	for info := bufio.NewScanner(file); done != 15 && info.Scan(); {
		var prop, val = "", 0
		if _, err = fmt.Sscanf(info.Text(), "%s %d", &prop, &val); err != nil {
			return "ERR"
		}
		switch prop {
		case "MemTotal:":
			total = val
			used += val
			done |= 1
		case "MemFree:":
			used -= val
			done |= 2
		case "Buffers:":
			used -= val
			done |= 4
		case "Cached:":
			used -= val
			done |= 8
		}
	}

	return strconv.Itoa(used*100/total)
}

func getMemory() string {
	return fmt.Sprintf("Mem: %v%%", updateMemUse())
}
