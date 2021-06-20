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

	hostInfo := HostInfo{
		// Users: No support?
		// Temperatures: No support?
		// OSInfo: This is OS below
		Architecture:  info.Architecture,
		BootTime:      info.BootTime,
		Containerized: info.Containerized,
		Hostname:      info.Hostname,
		IPs:           info.IPs,
		KernelVersion: info.KernelVersion,
		MACs:          info.MACs,
		OS: &OS{
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
		},
		Timezone:          info.Timezone,
		TimezoneOffsetSec: info.TimezoneOffsetSec,
		UniqueID:          info.UniqueID,
	}
	return hostInfo
}

func (g GoSysInfo) CPU() CPU {
	host, err := sysinfo.Host()
	if err != nil {
		panic(err)
	}
	cpuTime, err := host.CPUTime()
	if err != nil {
		panic(err)
	}

	return CPU{
		Sysifno: SysInfoCPU{
			User:    cpuTime.User,
			System:  cpuTime.System,
			Idle:    cpuTime.Idle,
			IOWait:  cpuTime.IOWait,
			IRQ:     cpuTime.IRQ,
			Nice:    cpuTime.Nice,
			SoftIRQ: cpuTime.SoftIRQ,
			Steal:   cpuTime.Steal,
		},
	}
}

func (g GoSysInfo) Memory() Memory {
	host, err := sysinfo.Host()
	if err != nil {
		panic(err)
	}
	memory, err := host.Memory()
	if err != nil {
		panic(err)
	}
	return Memory{
		SysInfoMemory: SysInfoMemory(*memory),
	}
}

func (g GoSysInfo) Processes() []Process {
	return []Process{}
}
