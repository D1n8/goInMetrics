package mem

import (
	"encoding/json"
	"log"

	"github.com/shirou/gopsutil/v3/mem"
)

type Mem struct {
	Total 		float64 	`json:"total"`
	Used		float64 	`json:"used"`
	UsedPercent	float64	`json:"used_percent"`
	Avail		float64 	`json:"avail"`
	SwapTotal	float64	`json:"swap_total"`
	SwapUsed	float64 	`json:"swap_used"`
}

func convertBytetoMb(byteIn uint64) float64 {
	return float64(byteIn)/1024/1024
	
}

func getMemStat() Mem {
	memStat, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("%s%v", "Backend error | getMemStat | :", err)
		return Mem {
			Total: 		0.0,
			Used: 		0.0,
			UsedPercent: 	0.0,
			Avail: 		0.0,
			SwapTotal: 	0.0,
			SwapUsed: 	0.0,
		}
	}
	return Mem {
		Total: 		convertBytetoMb(memStat.Total),
		Used: 		convertBytetoMb(memStat.Used),
		UsedPercent: 	memStat.UsedPercent,
		Avail: 		convertBytetoMb(memStat.Available),
		SwapTotal: 	convertBytetoMb(memStat.SwapTotal),
		SwapUsed: 	convertBytetoMb(memStat.SwapTotal - memStat.SwapFree),
	}

}

func memStatistic() string {
	// Json values. All values return in Mb size
	jsonData, err := json.Marshal(getMemStat())
	if err != nil {
		// FIX: REMOVE log.Fatal | Change it to some error value, but without exiting the whole proccess
		log.Fatal(err)
	}
	return string(jsonData)
}
