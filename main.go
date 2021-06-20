package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type DataCollector interface {
	HostInfo() HostInfo
	CPU() CPU
	Memory() Memory
	Processes() []Process
}

type HostInfo struct {
	// gopsutil
	OSInfo OS
	// go-sysinfo
	Users             []User
	Temperatures      []Temperature
	Architecture      string    `json:"architecture"`            // Hardware architecture (e.g. x86_64, arm, ppc, mips).
	BootTime          time.Time `json:"boot_time"`               // Host boot time.
	Containerized     *bool     `json:"containerized,omitempty"` // Is the process containerized.
	Hostname          string    `json:"name"`                    // Hostname
	IPs               []string  `json:"ip,omitempty"`            // List of all IPs.
	KernelVersion     string    `json:"kernel_version"`          // Kernel version.
	MACs              []string  `json:"mac"`                     // List of MAC addresses.
	OS                *OS       `json:"os"`                      // OS information.
	Timezone          string    `json:"timezone"`                // System timezone.
	TimezoneOffsetSec int       `json:"timezone_offset_sec"`     // Timezone offset (seconds from UTC).
	UniqueID          string    `json:"id,omitempty"`            // Unique ID of the host (optional).

}

type User struct {
	User     string `json:"user"`
	Terminal string `json:"terminal"`
	Host     string `json:"host"`
	Started  int    `json:"started"`
}

type Temperature struct {
	SensorKey   string  `json:"sensorKey"`
	Temperature float64 `json:"temperature"`
	High        float64 `json:"sensorHigh"`
	Critical    float64 `json:"sensorCritical"`
}

type CPU struct {
}

type Memory struct {
}

type Process struct {
}

type OS struct {
	// gopsutil Host InfoStat (describes the host status)
	Hostname             string `json:"hostname"`
	Uptime               uint64 `json:"uptime"`
	BootTime             uint64 `json:"bootTime"`
	Procs                uint64 `json:"procs"`           // number of processes
	OS                   string `json:"os"`              // ex: freebsd, linux
	PlatformFamily       string `json:"platformFamily"`  // ex: debian, rhel
	PlatformVersion      string `json:"platformVersion"` // version of the complete OS
	KernelVersion        string `json:"kernelVersion"`   // version of the OS kernel (if available)
	KernelArch           string `json:"kernelArch"`      // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
	VirtualizationSystem string `json:"virtualizationSystem"`
	VirtualizationRole   string `json:"virtualizationRole"` // guest or host
	HostID               string `json:"hostId"`             // ex: uuid

	// go-sysinfo OSINfo
	Type     string `json:"type"`               // OS Type (one of linux, macos, unix, windows).
	Family   string `json:"family"`             // OS Family (e.g. redhat, debian, freebsd, windows).
	Name     string `json:"name"`               // OS Name (e.g. Mac OS X, CentOS).
	Version  string `json:"version"`            // OS version (e.g. 10.12.6).
	Major    int    `json:"major"`              // Major release version.
	Minor    int    `json:"minor"`              // Minor release version.
	Patch    int    `json:"patch"`              // Patch release version.
	Build    string `json:"build,omitempty"`    // Build (e.g. 16G1114).
	Codename string `json:"codename,omitempty"` // OS codename (e.g. jessie).

	// Duplicates
	// Platform string // OS platform (e.g. centos, ubuntu, windows). GO-SYSINFO
	// Platform  string // ex: ubuntu, linuxmint GOPSUTIL
	Platform string `json:"platform"`
}

func main() {
	fmt.Println("Go System Information Data Collector Test")
	var collector DataCollector

	collector = GoSysInfo{}
	hostInfo, _ := json.Marshal(collector.HostInfo())
	fmt.Println(string(hostInfo))

	collector = GoPSUtil{}
	hostInfo, _ = json.Marshal(collector.HostInfo())
	fmt.Println(string(hostInfo))
}
