package main

type GoSysInfo struct {
}

func (g GoSysInfo) HostInfo() HostInfo {
	// host, err := sysinfo.Host()
	// if err != nil {
	// 	panic(err)
	// }
	return HostInfo{}
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
