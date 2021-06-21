package main

import "time"

type DataCollector interface {
	HostInfo() HostInfo
	CPU() CPU
	Memory() Memory
	Processes() []Process
	Disk() Disk
}

type GoPSUtilHost struct {
	GoPSUtilOS   GoPSUtilOS    `json:"osInfo"`
	Users        []User        `json:"users"`
	Temperatures []Temperature `json:"temperatures"`
}

type GoSysInfoHost struct {
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

type HostInfo struct {
	GoPSUtilHost  GoPSUtilHost  `json:"goPSUtilHost"`
	GoSysInfoHost GoSysInfoHost `json:"goSysInfoHost"`
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
	User    time.Duration `json:"user,omitempty"`
	System  time.Duration `json:"system,omitempty"`
	Idle    time.Duration `json:"idle,omitempty"`
	IOWait  time.Duration `json:"iowait,omitempty"`
	IRQ     time.Duration `json:"irq,omitempty"`
	Nice    time.Duration `json:"nice,omitempty"`
	SoftIRQ time.Duration `json:"soft_irq,omitempty"`
	Steal   time.Duration `json:"steal,omitempty"`
}

type GoPSUtilCPU struct {
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

type CPU struct {
	SysInfoCPU  SysInfoCPU  `json:"sysInfoCPU"`
	GoPSUtilCPU GoPSUtilCPU `json:"goPSUtilCPU"`
}

type SysInfoMemory struct {
	Total        uint64            `json:"total_bytes"`         // Total physical memory.
	Used         uint64            `json:"used_bytes"`          // Total - Free
	Available    uint64            `json:"available_bytes"`     // Amount of memory available without swapping.
	Free         uint64            `json:"free_bytes"`          // Amount of memory not used by the system.
	VirtualTotal uint64            `json:"virtual_total_bytes"` // Total virtual memory.
	VirtualUsed  uint64            `json:"virtual_used_bytes"`  // VirtualTotal - VirtualFree
	VirtualFree  uint64            `json:"virtual_free_bytes"`  // Virtual memory that is not used.
	Metrics      map[string]uint64 `json:"raw,omitempty"`
}

type GoPSUtilMemory struct {
	GoPSUtilVirtualMemory GoPSUtilVirtualMemory `json:"virtualMemory"`
	GoPSUtilSwapMemory    GoPSUtilSwapMemory    `json:"swapMemory"`
}

type GoPSUtilVirtualMemory struct {
	Total          uint64  `json:"total"`
	Available      uint64  `json:"available"`
	Used           uint64  `json:"used"`
	UsedPercent    float64 `json:"usedPercent"`
	Free           uint64  `json:"free"`
	Active         uint64  `json:"active"`
	Inactive       uint64  `json:"inactive"`
	Wired          uint64  `json:"wired"`
	Laundry        uint64  `json:"laundry"`
	Buffers        uint64  `json:"buffers"`
	Cached         uint64  `json:"cached"`
	WriteBack      uint64  `json:"writeBack"`
	Dirty          uint64  `json:"dirty"`
	WriteBackTmp   uint64  `json:"writeBackTmp"`
	Shared         uint64  `json:"shared"`
	Slab           uint64  `json:"slab"`
	Sreclaimable   uint64  `json:"sreclaimable"`
	Sunreclaim     uint64  `json:"sunreclaim"`
	PageTables     uint64  `json:"pageTables"`
	SwapCached     uint64  `json:"swapCached"`
	CommitLimit    uint64  `json:"commitLimit"`
	CommittedAS    uint64  `json:"committedAS"`
	HighTotal      uint64  `json:"highTotal"`
	HighFree       uint64  `json:"highFree"`
	LowTotal       uint64  `json:"lowTotal"`
	LowFree        uint64  `json:"lowFree"`
	SwapTotal      uint64  `json:"swapTotal"`
	SwapFree       uint64  `json:"swapFree"`
	Mapped         uint64  `json:"mapped"`
	VmallocTotal   uint64  `json:"vmallocTotal"`
	VmallocUsed    uint64  `json:"vmallocUsed"`
	VmallocChunk   uint64  `json:"vmallocChunk"`
	HugePagesTotal uint64  `json:"hugePagesTotal"`
	HugePagesFree  uint64  `json:"hugePagesFree"`
	HugePageSize   uint64  `json:"hugePageSize"`
}

type GoPSUtilSwapMemory struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
	Sin         uint64  `json:"sin"`
	Sout        uint64  `json:"sout"`
	PgIn        uint64  `json:"pgIn"`
	PgFault     uint64  `json:"pgFault"`
	PgMajFault  uint64  `json:"pgMajFault"`
}

type Memory struct {
	SysInfoMemory  SysInfoMemory  `json:"sysInfoMemory"`
	GoPSUtilMemory GoPSUtilMemory `json:"goPSUtilMemory"`
}

type Process struct {
	SysInfoProcess  SysInfoProcess  `json:"sysInfoProcess"`
	GoPSUtilProcess GoPsUtilProcess `json:"goPSUtilProcess"`
}

type SysInfoProcess struct {
	ProcessInfo ProcessInfo `json:"processInfo"`
	UserInfo    UserInfo    `json:"userInfo"`
}

type ProcessInfo struct {
	Name      string    `json:"name"`
	PID       int       `json:"pid"`
	Exe       string    `json:"exe"`
	StartTime time.Time `json:"start_time"`
	// There are more but I've cut them out...
}

type UserInfo struct {
	UID string `json:"uid"`
	// There are more but I've cut them out...
}

type GoPsUtilProcess struct {
	Pid    int32 `json:"pid"`
	name   string
	status string
	// There are more but I've cut them out...
}

type GoPSUtilOS struct {
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
	Platform             string `json:"platform"`           // Platform  string // ex: ubuntu, linuxmint
}

type GoSysInfoOS struct {
	Type     string `json:"type"`               // OS Type (one of linux, macos, unix, windows).
	Family   string `json:"family"`             // OS Family (e.g. redhat, debian, freebsd, windows).
	Name     string `json:"name"`               // OS Name (e.g. Mac OS X, CentOS).
	Version  string `json:"version"`            // OS version (e.g. 10.12.6).
	Major    int    `json:"major"`              // Major release version.
	Minor    int    `json:"minor"`              // Minor release version.
	Patch    int    `json:"patch"`              // Patch release version.
	Build    string `json:"build,omitempty"`    // Build (e.g. 16G1114).
	Codename string `json:"codename,omitempty"` // OS codename (e.g. jessie).
	Platform string `json:"platform"`           // Platform string // OS platform (e.g. centos, ubuntu, windows).
}

type OS struct {
	GoPPUtilOS  GoPSUtilOS  `json:"goPSUtilOS"`
	GoSysInfoOS GoSysInfoOS `json:"goSysInfoOS"`
}

type GoPsUtilDisk struct {
	Path              string  `json:"path"`
	Fstype            string  `json:"fstype"`
	Total             uint64  `json:"total"`
	Free              uint64  `json:"free"`
	Used              uint64  `json:"used"`
	UsedPercent       float64 `json:"usedPercent"`
	InodesTotal       uint64  `json:"inodesTotal"`
	InodesUsed        uint64  `json:"inodesUsed"`
	InodesFree        uint64  `json:"inodesFree"`
	InodesUsedPercent float64 `json:"inodesUsedPercent"`
}

type Disk struct {
	GoPsUtilDisk GoPsUtilDisk `json:"goPSUtilDisk"`
}
