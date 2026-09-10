package cpu

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/cpu"
)

func getCpuInfo() {
	cpuStat, err := cpu.Info()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cpuStat)
}
