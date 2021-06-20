package main

import "github.com/elastic/go-sysinfo"

type GoSysInfo struct {
}

func (g GoSysInfo) HostInfo() HostInfo {
	//var hostInfo HostInfo
	host, err := sysinfo.Host()
	info := host.Info()
	if err != nil {
		panic(err)
	}
	os := OS{
		Type:     info.OS.Type,
		Family:   info.OS.Family,
		Name:     info.OS.Name,
		Version:  info.OS.Version,
		Major:    info.OS.Major,
		Minor:    info.OS.Minor,
		Patch:    info.OS.Patch,
		Build:    info.OS.Build,
		Codename: info.OS.Codename,
		Platform: info.OS.Platform,
	}

	hostInfo := HostInfo{
		Architecture:      info.Architecture,
		BootTime:          info.BootTime,
		Containerized:     info.Containerized,
		Hostname:          info.Hostname,
		IPs:               info.IPs,
		KernelVersion:     info.KernelVersion,
		MACs:              info.MACs,
		OS:                &os,
		Timezone:          info.Timezone,
		TimezoneOffsetSec: info.TimezoneOffsetSec,
		UniqueID:          info.UniqueID,
	}
	return hostInfo
}

func (g GoSysInfo) CPU() CPU {
	return CPU{}
}

func (g GoSysInfo) Memory() Memory {
	return Memory{}
}

func (g GoSysInfo) Processes() []Process {
	return []Process{}
}
