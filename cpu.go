package main

import (
	"io/ioutil"
	"runtime"
	"strconv"
	"fmt"
)

func getCpu() string {
	var load float32
	var loadavg, err = ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		return "ERR"
	}

	_, err = fmt.Sscanf(string(loadavg), "%f", &load)
	if err != nil {
		return "ERR"
	}


	fresult := load*100.0/float32(runtime.NumCPU())

	result := strconv.FormatFloat(float64(fresult), 'f', -1, 32)

	return fmt.Sprintf("load: %v", result);
}
