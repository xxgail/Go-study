package sys

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/wailsapp/wails"
	"math"
	"time"
)

type Stats struct {
	log *wails.CustomLogger
}

type CPUUsage struct {
	Average int `json:"avg"`
}

func (s *Stats) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Stats")
	return nil
}

func (s *Stats) GetCPUUsage() *CPUUsage {
	percent, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		s.log.Errorf("unable to get cpu stats: %s", err.Error())
		return nil
	}

	return &CPUUsage{
		Average: int(math.Round(percent[0])),
	}
}
