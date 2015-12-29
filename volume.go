package main

import (
	"os/exec"
	"fmt"
)

func getVolume() string {
	var cmd,err = exec.Command("pamixer", "--get-volume").Output()
	if err != nil {
		return "ERR"
	}

	return fmt.Sprintf("Vol: %v", stripLast(string(cmd)))
}
