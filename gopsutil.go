package main

import (
	gHost "github.com/shirou/gopsutil/v3/host"
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
	return CPU{}
}

func (g GoPSUtil) Memory() Memory {
	return Memory{}
}

func (g GoPSUtil) Processes() []Process {
	return []Process{}
}
