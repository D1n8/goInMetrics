package cpu

import (
	"encoding/json"
	"log"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/v3/host"
)


type Cpu struct {
	Model		string		`json:"model"`
	LoadPercent	float64		`json:"load_percent"`
	Temperature	int		`json:"temperature"`
	Frequency	freqStat	`json:"frequency"`
}

type freqStat struct {
	MaxFreqMhz 	float64 	`json:"max_mhz"`
	WorkFreqMhz 	float64		`json:"work_mhz"`
}

func getCpuModel() string {
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Printf("%s%v", "Backend error | getCpuModel | :", err)
		return "Undefined"
	}
		
	var model string = "Undefined"
	for _, item := range cpuInfo {
		if item.ModelName != "" {
			model = item.ModelName
		}
	}
	return model
}

func getCpuPercent() float64 {
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		log.Printf("%s%v", "Backend error | getCpuPercent | :", err)
		return 0.0
	}
	return cpuPercent[0]
}

func getCpuTemp() int {
	tempStat, err := host.SensorsTemperatures()
	if err != nil {
		log.Printf("%s%v", "Backend error | getCpuTemp | :", err)
		return 0
	}
	var temp = 0
	for _, item := range tempStat {
		if item.Temperature != 0 {
			temp = int(item.Temperature)
			break
		}
	}
	return  temp
}

func getCpuFreq() (freqStat) {

	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Printf("%s%v", "Backend error | getCpuModel | :", err)
		return freqStat {
			MaxFreqMhz: 0.0,
			WorkFreqMhz: 0.0,
		}
	}

	var maxFreq float64 = 0.0
	for _, item := range cpuInfo {
		if item.Mhz != 0.0 {
			maxFreq = item.Mhz
			break
		}
	}

	return freqStat {
		MaxFreqMhz: maxFreq,
		// FIX: Remove hardcoded value 0.0
		WorkFreqMhz: 0.0,
	}
}


func CpuStatistic() string {
	structFill := Cpu{
		Model: 		getCpuModel(),
		LoadPercent:	getCpuPercent(),
		Temperature:	getCpuTemp(),
		Frequency: 	getCpuFreq(),
	}

	jsonData, err := json.Marshal(structFill)
	if err != nil {
		// FIX: REMOVE log.Fatal | Change it to some error value, but without exiting the whole proccess
		log.Fatal(err)
	}
	// NOTE: REMOVE OUTPUT from this func
	return string(jsonData)
}
