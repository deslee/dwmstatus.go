package main

import (
	"os/exec"
	"strings"
	"fmt"
)

var disks = map[string]string {
	"/dev/sda2": "/",
}

func getDisks() string {

	var tmp, err = exec.Command("df", "-h").Output()

	if err != nil {
		return "Disks: ERR"
	}

	var results []string
	var output = string(tmp)
	var lines = strings.Split(output, "\n")

	for _,line := range lines {
		var fields = strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		var disk = fields[0]
		var perc = fields[4]
		if _, ok := disks[disk]; ok {
			name := disks[disk]
			results = append(results, fmt.Sprintf("%v:%v", name, perc))
		}
	}


	return strings.Join(results, ",")
}
