package main

import (
	"fmt"
	"encoding/json"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/cpu"
	"strings"
)

func jsonHostInfo() []byte {
	hostinfo, _ := host.HostInfo()
	jsonHostInfo , _ := json.Marshal(hostinfo)
	return jsonHostInfo
}

func jsonCPUInfo() []byte {
	type CpuInfo struct {
		Info []cpu.CPUInfoStat
		Times []cpu.CPUTimesStat
		Time []cpu.CPUTimesStat
	}
	c := &CpuInfo{}
	CPUTime, _ := cpu.CPUTimes(false)
	CPUTimes, _ := cpu.CPUTimes(true) 	
	CPUInfo, _ := cpu.CPUInfo() 	
	c.Info = CPUInfo
	c.Time =  CPUTime
	c.Times =  CPUTimes
	jsonPartitions , _ := json.Marshal(c)
	return jsonPartitions
}

func jsonDiskInfo() []byte {
	
	type DiskStat struct {
		Partition disk.DiskPartitionStat
		Usage     disk.DiskUsageStat     
		Counters  disk.DiskIOCountersStat  
	}

	var disks []DiskStat
	partitions, _ := disk.DiskPartitions(false)
	c := &DiskStat{}
	IOCounters, _ := disk.DiskIOCounters()

	for _, singleDisk := range partitions{
		if ! strings.HasPrefix(singleDisk.Device, "/") { 
			continue
		}
		usage, _ := disk.DiskUsage(singleDisk.Mountpoint)
		c.Partition = singleDisk
		c.Usage = *usage
		c.Counters = IOCounters[strings.Split(singleDisk.Device, "/")[2]]
		disks = append(disks, *c)
	}
	jsonPartitions , _ := json.Marshal(disks)
	return jsonPartitions
}

func main() {
	fmt.Println(string(jsonHostInfo()))
	fmt.Println(string(jsonDiskInfo()))
	fmt.Println(string(jsonCPUInfo()))
}
