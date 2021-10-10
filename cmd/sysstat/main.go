package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func formatNumber(num uint64) uint64 {
	return uint64(num / 1024 / 1014)
}

func main() {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Total: %vM, Free: %vM, Used: %.2f%%\n", formatNumber(v.Total), formatNumber(v.Free), v.UsedPercent)

	cs, err := cpu.Counts(true)
	if err != nil {
		log.Fatalln(err)
	}

	cp, _ := cpu.Percent(1*time.Second, false)

	fmt.Printf("Cpu count: %d, Cpu Percent: %v\n", cs, cp)
}
