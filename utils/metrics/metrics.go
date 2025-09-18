package metrics

import (
	"fmt"
	"net"
	"os/exec"
	"sort"
	"strings"
	"time"

	"github.com/shirou/gopsutil/process"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	host "github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

// -------------------- Structs --------------------
type CPUInfo struct {
	Model   string    `json:"model"`
	Average float64   `json:"average"`
	Cores   []float64 `json:"cores"`
}

type GPUInfo struct {
	Name       string  `json:"name"`
	Utilization int     `json:"utilization_pct"`
	MemTotal   int     `json:"mem_total_mb"`
	MemUsed    int     `json:"mem_used_mb"`
}

type MemoryInfo struct {
	Total     uint64  `json:"total"`
	Used      uint64  `json:"used"`
	Free      uint64  `json:"free"`
	UsedPct   float64 `json:"used_pct"`
	SwapTotal uint64  `json:"swap_total"`
	SwapUsed  uint64  `json:"swap_used"`
	SwapPct   float64 `json:"swap_pct"`
}

type DiskInfo struct {
	Mountpoint string  `json:"mountpoint"`
	Device     string  `json:"device"`
	Fstype     string  `json:"fstype"`
	Total      uint64  `json:"total"`
	Used       uint64  `json:"used"`
	UsedPct    float64 `json:"used_pct"`
}

type TempInfo struct {
	Sensor string  `json:"sensor"`
	Temp   float64 `json:"temp"`
}

type NetInfo struct {
	Interface string `json:"interface"`
	Addr      string `json:"addr"`
	Flags     string `json:"flags"`
}

type SystemStats struct {
	Timestamp 	string      `json:"timestamp"`
	Hostname  	string      `json:"hostname"`
	OS        	string      `json:"os"`
	Kernel    	string      `json:"kernel"`
	Uptime    	string      `json:"uptime"`
	Load1     	float64     `json:"load1"`
	Load5     	float64     `json:"load5"`
	Load15    	float64     `json:"load15"`
	CPU       	CPUInfo     `json:"cpu"`
	Memory    	MemoryInfo  `json:"memory"`
	Disks     	[]DiskInfo  `json:"disks"`
	GPUs		[]GPUInfo 	`json:"gpus"`
	Temps     	[]TempInfo  `json:"temps"`
	Networks  	[]NetInfo   `json:"networks"`
	ProcCount 	int         `json:"proc_count"`
}

// -------------------- Collectors --------------------
func collectCPU() CPUInfo {
	info := CPUInfo{}
	perCore, _ := cpu.Percent(0, true)
	avg, _ := cpu.Percent(0, false)
	cpuInfo, _ := cpu.Info()
	if len(cpuInfo) > 0 {
		info.Model = cpuInfo[0].ModelName
	}
	if len(avg) > 0 {
		info.Average = avg[0]
	}
	info.Cores = perCore
	return info
}

func collectNvidiaGPU() ([]GPUInfo, error) {
	// Query: name, util, mem.total, mem.used
	cmd := exec.Command("nvidia-smi",
		"--query-gpu=name,utilization.gpu,memory.total,memory.used",
		"--format=csv,noheader,nounits")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	gpus := []GPUInfo{}

	for _, line := range lines {
		parts := strings.Split(line, ", ")
		if len(parts) != 4 {
			continue
		}
		var g GPUInfo
		g.Name = parts[0]
		fmt.Sscanf(parts[1], "%d", &g.Utilization)
		fmt.Sscanf(parts[2], "%d", &g.MemTotal)
		fmt.Sscanf(parts[3], "%d", &g.MemUsed)
		gpus = append(gpus, g)
	}

	return gpus, nil
}

func collectMemory() MemoryInfo {
	vm, _ := mem.VirtualMemory()
	sm, _ := mem.SwapMemory()
	return MemoryInfo{
		Total:     vm.Total,
		Used:      vm.Used,
		Free:      vm.Free,
		UsedPct:   vm.UsedPercent,
		SwapTotal: sm.Total,
		SwapUsed:  sm.Used,
		SwapPct:   sm.UsedPercent,
	}
}

func collectDisks() []DiskInfo {
	result := []DiskInfo{}
	parts, _ := disk.Partitions(true)
	for _, p := range parts {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}
		result = append(result, DiskInfo{
			Mountpoint: p.Mountpoint,
			Device:     p.Device,
			Fstype:     p.Fstype,
			Total:      usage.Total,
			Used:       usage.Used,
			UsedPct:    usage.UsedPercent,
		})
	}
	return result
}

func collectTemps() []TempInfo {
	result := []TempInfo{}
	temps, _ := host.SensorsTemperatures()
	for _, t := range temps {
		result = append(result, TempInfo{
			Sensor: t.SensorKey,
			Temp:   t.Temperature,
		})
	}
	return result
}

func collectNet() []NetInfo {
	out := []NetInfo{}
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		addrs, _ := iface.Addrs()
		if len(addrs) == 0 {
			continue
		}
		flags := iface.Flags.String()
		for _, a := range addrs {
			out = append(out, NetInfo{
				Interface: iface.Name,
				Addr:      a.String(),
				Flags:     flags,
			})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Interface < out[j].Interface })
	return out
}

// -------------------- Master Collector --------------------
func Pull() SystemStats {
	h, _ := host.Info()
	ld, _ := load.Avg()
	pids, _ := process.Pids()
	gpus, _ := collectNvidiaGPU()

	if gpus == nil {
    	gpus = []GPUInfo{} // ensures JSON is []
	}

	return SystemStats{
		Timestamp: time.Now().Format(time.RFC3339),
		Hostname:  h.Hostname,
		OS:        h.Platform + " " + h.PlatformVersion,
		Kernel:    h.KernelVersion,
		Uptime:    (time.Duration(h.Uptime) * time.Second).String(),
		Load1:     ld.Load1,
		Load5:     ld.Load5,
		Load15:    ld.Load15,
		CPU:       collectCPU(),
		Memory:    collectMemory(),
		Disks:     collectDisks(),
		Temps:     collectTemps(),
		GPUs: 		gpus,
		Networks:  collectNet(),
		ProcCount: len(pids),
	}
}


