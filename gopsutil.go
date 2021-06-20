package main

import (
	gCPU "github.com/shirou/gopsutil/v3/cpu"
	gHost "github.com/shirou/gopsutil/v3/host"
	gMem "github.com/shirou/gopsutil/v3/mem"
)

type GoPSUtil struct {
}

func (g GoPSUtil) HostInfo() HostInfo {
	var hostInfo HostInfo
	host, err := gHost.Info()
	if err != nil {
		panic(err)
	}
	hostInfo.OSInfo = OS{
		Hostname:             host.Hostname,
		Uptime:               host.Uptime,
		BootTime:             host.BootTime,
		Procs:                host.Procs,
		OS:                   host.OS,
		Platform:             host.Platform,
		PlatformFamily:       host.PlatformFamily,
		PlatformVersion:      host.PlatformVersion,
		KernelVersion:        host.KernelVersion,
		KernelArch:           host.KernelArch,
		VirtualizationSystem: host.VirtualizationSystem,
		VirtualizationRole:   host.VirtualizationRole,
		HostID:               host.HostID,
	}
	users, err := gHost.Users()
	if err != nil {
		panic(err)
	}
	var hostUsers []User
	for _, user := range users {
		tempUser := User{
			User:     user.User,
			Terminal: user.Terminal,
			Host:     user.User,
			Started:  user.Started,
		}
		hostUsers = append(hostUsers, tempUser)
	}

	temps, err := gHost.SensorsTemperatures()
	if err != nil {
		panic(err)
	}

	var hostTemps []Temperature
	for _, temp := range temps {
		tempTemp := Temperature{
			SensorKey:   temp.SensorKey,
			Temperature: temp.Temperature,
			High:        temp.High,
			Critical:    temp.Critical,
		}
		hostTemps = append(hostTemps, tempTemp)
	}

	hostInfo.Users = hostUsers
	hostInfo.Temperatures = hostTemps
	return hostInfo
}

func (g GoPSUtil) CPU() CPU {
	cpuInfo, err := gCPU.Info()
	if err != nil {
		panic(err)
	}

	var cpuInfoStat []CPUInfoStat
	for _, cpu := range cpuInfo {
		cpuInfoStat = append(cpuInfoStat, CPUInfoStat(cpu))
	}

	return CPU{
		InfoStat: cpuInfoStat,
	}
}

func (g GoPSUtil) Memory() Memory {
	virtualMemory, _ := gMem.VirtualMemory()
	swapMemory, _ := gMem.SwapMemory()
	return Memory{
		GoPSUtilMemory: GoPSUtilMemory{
			GoPSUtilVirtualMemory: GoPSUtilVirtualMemory(*virtualMemory),
			GoPSUtilSwapMemory: GoPSUtilSwapMemory{ // not sure why GoPSUtilSwapMemory(*swapMemory) isn't working here...
				Total:       swapMemory.Total,
				Used:        swapMemory.Used,
				Free:        swapMemory.Free,
				UsedPercent: swapMemory.UsedPercent,
				Sin:         swapMemory.Sin,
				Sout:        swapMemory.Sout,
				PgIn:        swapMemory.PgIn,
				PgFault:     swapMemory.PgFault,
				PgMajFault:  swapMemory.PgMajFault,
			},
		},
	}
}

func (g GoPSUtil) Processes() []Process {
	return []Process{}
}
