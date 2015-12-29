package main

import (
	"os/exec"
	"strings"
	"time"
)

func setRoot(text string) {
	exec.Command("xsetroot", "-name", text).Run();
}

func tick() {
	var result []string
	result = append(result, getBattery())
	result = append(result, getMemory())
	result = append(result, getDisks())
	result = append(result, getNetwork())
	result = append(result, getVolume())
	result = append(result, getCpu())
	result = append(result, getDate())
	setRoot(strings.Join(result, " | "))
}

func main() {
	for {
		tick()
		now := time.Now()
		time.Sleep(now.Truncate(time.Second).Add(time.Second).Sub(now))
	}
}
