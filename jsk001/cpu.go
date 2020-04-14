package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"time"
)

func getCpuInfo() {
	cpuInfos, _ := cpu.Info()

	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

// mem info
func getMemInfo() {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("mem info:%v\n", memInfo)
}
func getCpuLoad() {
	info, _ := load.Avg()
	fmt.Printf("%v\n", info)
}
func maincpu() {
	//getCpuInfo();
	//getCpuLoad();
	getMemInfo()
}
