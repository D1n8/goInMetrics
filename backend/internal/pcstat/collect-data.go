package pcstat

import (
	"backend/internal/pcstat/cpu"
	"backend/internal/pcstat/mem"
	"encoding/json"
	"log"
)

type collectedDataStruct struct {
	Cpu	json.RawMessage	`json:"cpu"`
	Mem	json.RawMessage	`json:"mem"`
}

func SendDataToServer() ([]byte, error) {
	collectedData := collectedDataStruct {
		Cpu: json.RawMessage(cpu.CpuStatistic()),
		Mem: json.RawMessage(mem.MemStatistic()),
	}
	jsonData, err := json.Marshal(collectedData)
	if err != nil {
		// FIX: Remove log.Fatal
		log.Fatal(err)
	}
	return []byte(jsonData), nil
}
