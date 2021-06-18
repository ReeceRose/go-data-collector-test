package main

import "fmt"

type DataCollector interface {
	HostInfo() string
	CPU() string
	Memory() string
	Processes() string
	OS() string
}

func main() {
	fmt.Println("Go System Information Data Collector Test")
	var collector DataCollector
	collector = GoSysInfo{}
	fmt.Println(collector.HostInfo())
}
