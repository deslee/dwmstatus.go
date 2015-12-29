package main

import (
	"io/ioutil"
	"fmt"
)

func getBattery() string {
	var tmp, err = ioutil.ReadFile("/sys/class/power_supply/BAT0/capacity")

	if (err != nil) {
		return "ERR"
	}


	return fmt.Sprintf("Bat: %v", stripLast(string(tmp)) + "%")
}
