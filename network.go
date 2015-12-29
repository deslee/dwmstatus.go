package main

import (
	"io/ioutil"
	"strings"
	"fmt"
)

var netDevs = [...]string{"wlp1s0"}

func getNetwork() string {

	var result []string
	for _,dev := range netDevs {
		var tmp, err = ioutil.ReadFile(fmt.Sprintf("/sys/class/net/%v/operstate", dev))

		value := "ERR"
		if (err == nil) {
			s := stripLast(string(tmp))
			value = s
		}

		r := fmt.Sprintf("%v(%v)", dev, value)


		result = append(result, r)
	}


	return strings.Join(result, ",")
}
