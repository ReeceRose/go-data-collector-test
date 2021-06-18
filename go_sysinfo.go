package main

import (
	"github.com/elastic/go-sysinfo"
)

type GoSysInfo struct {
}

func (g GoSysInfo) HostInfo() string {
	host, err := sysinfo.Host()
	if err != nil {
		panic(err)
	}
	return host.Info().OS.Type //TODO: add more
}

func (g GoSysInfo) CPU() string {
	return ""
}

func (g GoSysInfo) Memory() string {
	return ""
}

func (g GoSysInfo) Processes() string {
	return ""
}

func (g GoSysInfo) OS() string {
	return ""
}
