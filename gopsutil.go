package main

import (
	"github.com/shirou/gopsutil/v3/host"
)

type GoPSUtil struct {
}

func (g GoPSUtil) HostInfo() string {
	host, err := host.Info()
	if err != nil {
		panic(err)
	}
	return host.OS
}

func (g GoPSUtil) CPU() string {
	return ""
}

func (g GoPSUtil) Memory() string {
	return ""
}

func (g GoPSUtil) Processes() string {
	return ""
}

func (g GoPSUtil) OS() string {
	return ""
}
