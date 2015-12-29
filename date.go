package main

import (
	"time"
)

func getDate() string {
	return time.Now().Format("Mon Jan 2 3:04")
}
