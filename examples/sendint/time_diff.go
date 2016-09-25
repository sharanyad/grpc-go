package main

import (
	"fmt"
	"time"
)

func main() {
	str1 := "2016-09-24 23:13:24.357665074 -0500 CDT"
	//str2 := "Mon Jan 2 15:04:05 -0700 MST 2006"
	time1, _ := time.Parse(time.RFC3339, str1)
	//time2, _ := time.Parse(str2, str1)
	fmt.Printf("\n time1: %v", time1)
	//fmt.Printf("\nDuration: %v", time2.Sub(time1).Nanoseconds())
}
