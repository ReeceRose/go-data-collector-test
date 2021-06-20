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
	OSInfo       OS
	Users        []User
	Temperatures []Temperature
	// go-sysinfo
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

type SysInfoCPU struct {
	// go-sysinfo CPUTimes struct: CPU timing status for a process

	User    time.Duration `json:"user,omitempty"`
	System  time.Duration `json:"system,omitempty"`
	Idle    time.Duration `json:"idle,omitempty"`
	IOWait  time.Duration `json:"iowait,omitempty"`
	IRQ     time.Duration `json:"irq,omitempty"`
	Nice    time.Duration `json:"nice,omitempty"`
	SoftIRQ time.Duration `json:"soft_irq,omitempty"`
	Steal   time.Duration `json:"steal,omitempty"`
}

type CPU struct {
	Sysifno SysInfoCPU `json:"Sysinfo,omitempty"`
	// gopsutil
	CPU       string        `json:"cpu"`
	User      float64       `json:"user"`
	System    float64       `json:"system"`
	Idle      float64       `json:"idle"`
	Nice      float64       `json:"nice"`
	Iowait    float64       `json:"iowait"`
	Irq       float64       `json:"irq"`
	Softirq   float64       `json:"softirq"`
	Steal     float64       `json:"steal"`
	Guest     float64       `json:"guest"`
	GuestNice float64       `json:"guestNice"`
	InfoStat  []CPUInfoStat `json:"infoStat"`
}

type CPUInfoStat struct {
	CPU        int32    `json:"cpu"`
	VendorID   string   `json:"vendorId"`
	Family     string   `json:"family"`
	Model      string   `json:"model"`
	Stepping   int32    `json:"stepping"`
	PhysicalID string   `json:"physicalId"`
	CoreID     string   `json:"coreId"`
	Cores      int32    `json:"cores"`
	ModelName  string   `json:"modelName"`
	Mhz        float64  `json:"mhz"`
	CacheSize  int32    `json:"cacheSize"`
	Flags      []string `json:"flags"`
	Microcode  string   `json:"microcode"`
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
	cpuInfo, _ := json.Marshal(collector.CPU())
	fmt.Println("go-sysinfo: HOST INFORMATION")
	fmt.Println(string(hostInfo))
	fmt.Println("go-sysinfo: CPU INFORMATION")
	fmt.Println(string(cpuInfo))

	collector = GoPSUtil{}
	hostInfo, _ = json.Marshal(collector.HostInfo())
	cpuInfo, _ = json.Marshal(collector.CPU())
	fmt.Println("gopsutil: HOST INFORMATION")
	fmt.Println(string(hostInfo))
	fmt.Println("gopsutil: CPU INFORMATION")
	fmt.Println(string(cpuInfo))
}
