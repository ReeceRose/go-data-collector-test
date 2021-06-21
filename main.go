package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Go System Information Data Collector Test")
	var collector DataCollector

	collector = GoSysInfo{}
	hostInfo, _ := json.Marshal(collector.HostInfo())
	cpuInfo, _ := json.Marshal(collector.CPU())
	memoryInfo, _ := json.Marshal(collector.Memory())
	processesInfo, _ := json.Marshal(collector.Processes())
	diskInfo, _ := json.Marshal(collector.Disk())
	fmt.Println("go-sysinfo: HOST INFORMATION")
	fmt.Println(string(hostInfo))
	fmt.Println("go-sysinfo: CPU INFORMATION")
	fmt.Println(string(cpuInfo))
	fmt.Println("go-sysinfo: MEMORY INFORMATION")
	fmt.Println(string(memoryInfo))
	fmt.Println("go-sysinfo: PROCESSES INFORMATION")
	fmt.Println(string(processesInfo))
	fmt.Println("go-sysinfo: DISK INFORMATION")
	fmt.Println(string(diskInfo))

	collector = GoPSUtil{}
	hostInfo, _ = json.Marshal(collector.HostInfo())
	cpuInfo, _ = json.Marshal(collector.CPU())
	memoryInfo, _ = json.Marshal(collector.Memory())
	processesInfo, _ = json.Marshal(collector.Processes())
	diskInfo, _ = json.Marshal(collector.Disk())
	fmt.Println("gopsutil: HOST INFORMATION")
	fmt.Println(string(hostInfo))
	fmt.Println("gopsutil: CPU INFORMATION")
	fmt.Println(string(cpuInfo))
	fmt.Println("gopsutil: MEMORY INFORMATION")
	fmt.Println(string(memoryInfo))
	fmt.Println("go-sysinfo: PROCESSES INFORMATION")
	fmt.Println(string(processesInfo))
	fmt.Println("go-sysinfo: DISK INFORMATION")
	fmt.Println(string(diskInfo))
}
