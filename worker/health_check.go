package worker

import (
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

//enum define status
const (
	WorkerError   = -1
	WorkerIdle    = 0
	WorkerRunning = 1
	WorkerLocked  = 2
)

type Info struct {
	Mem    float64 `json:"memory_usage"`
	CPU    float64 `json:"cpu_usage"`
	Status int     `json:"status"`
}

var oldCPUStat *cpu.Stats

func HealthCheck() Info {
	return Info{
		Mem:    MemUsage(),
		CPU:    CpuUsage(),
		Status: 0,
	}
}

func MemUsage() float64 {
	mem, err := memory.Get()
	if err != nil {
		logger.Error(err)
		return 0
	}

	return float64(mem.Used/mem.Total) * 100
}

func CpuUsage() float64 {
	cpu, err := cpu.Get()
	if err != nil {
		logger.Error("Get cpu info error: ", err)
		return 0
	}
	if oldCPUStat == nil {
		oldCPUStat = cpu
		return 0
	}

	total := cpu.Total - oldCPUStat.Total
	usage := (cpu.User - oldCPUStat.User) / total * 100

	oldCPUStat = cpu

	return float64(usage)
}
