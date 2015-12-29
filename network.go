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
		var tmp, err = ioutil.ReadFile(fmt.Sprintf("/sys/class/net/%v/carrier", dev))

		s := stripLast(string(tmp))
		
		value := "ERR"
		if (err == nil) {
			switch s {
			case "1":
				value = "up"
				break;
			case "0":
				value = "down"
				break;
			}
		}

		r := fmt.Sprintf("%v(%v)", dev, value)


		result = append(result, r)
	}


	return strings.Join(result, ",")
}
