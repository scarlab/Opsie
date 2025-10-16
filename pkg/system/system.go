package system

import (
	"fmt"
	"net"
	"sort"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	host "github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

// -------------------- Structs --------------------

type cpuInfo struct {
	Model          string    `json:"model"`
	Cores          uint16    `json:"cores"`
	Threads        uint16    `json:"threads"`
	Average        float64   `json:"average"`
	AveragePerCore []float64 `json:"average_per_core"`
}

type memoryInfo struct {
	Total     uint64  `json:"total"`
	Used      uint64  `json:"used"`
	Free      uint64  `json:"free"`
	UsedPct   float64 `json:"used_pct"`
	SwapTotal uint64  `json:"swap_total"`
	SwapUsed  uint64  `json:"swap_used"`
	SwapPct   float64 `json:"swap_pct"`
}

type diskInfo struct {
	Mountpoint string  `json:"mountpoint"`
	Device     string  `json:"device"`
	Fstype     string  `json:"fstype"`
	Total      uint64  `json:"total"`
	Used       uint64  `json:"used"`
	UsedPct    float64 `json:"used_pct"`
}

type tempInfo struct {
	Sensor string  `json:"sensor"`
	Temp   float64 `json:"temp"`
}

type netInfo struct {
	Interface string `json:"interface"`
	Addr      string `json:"addr,omitempty"`
	Flags     string `json:"flags,omitempty"`
	RxBytes   uint64 `json:"rx_bytes"`
	TxBytes   uint64 `json:"tx_bytes"`
}

type SystemInfo struct {
	Hostname  string `json:"hostname"`
	OS        string `json:"os"`
	Kernel    string `json:"kernel"`
	Arch      string `json:"arch"`
	IPAddress string `json:"ip_address"`
	Cores     uint16 `json:"cores"`
	Threads   uint16 `json:"threads"`
	Memory    uint64 `json:"memory"`
}

type SystemMetrics struct {
	Timestamp  	int64       	`json:"timestamp"`
	Uptime     	uint64      	`json:"uptime"`
	Load1      	float64     	`json:"load1"`
	Load5      	float64     	`json:"load5"`
	Load15     	float64     	`json:"load15"`
	CPU        	cpuInfo     	`json:"cpu"`
	Memory 		memoryInfo		`json:"memory"`
	Disk       	[]diskInfo  	`json:"disk"`
	NetIO      	[]netInfo   	`json:"net_io"`
	ProcCount  	int         	`json:"proc_count"`
	Temps      	[]tempInfo  	`json:"temps,omitempty"`
}

// -------------------- Collectors --------------------

func collectCPU() cpuInfo {
	info := cpuInfo{}
	perCore, _ := cpu.Percent(0, true)
	avg, _ := cpu.Percent(0, false)
	cpuInfoList, _ := cpu.Info()
	cores, _ := cpu.Counts(false)
	threads, _ := cpu.Counts(true)

	if len(cpuInfoList) > 0 {
		info.Model = cpuInfoList[0].ModelName
	}
	if len(avg) > 0 {
		info.Average = avg[0]
	}
	info.AveragePerCore = perCore
	info.Cores = uint16(cores)
	info.Threads = uint16(threads)

	return info
}

func collectMemory() memoryInfo {
	vm, _ := mem.VirtualMemory()
	sm, _ := mem.SwapMemory()
	toMB := func(b uint64) uint64 { return b / 1024 / 1024 }

	return memoryInfo{
		Total:     toMB(vm.Total),
		Used:      toMB(vm.Used),
		Free:      toMB(vm.Free),
		UsedPct:   vm.UsedPercent,
		SwapTotal: toMB(sm.Total),
		SwapUsed:  toMB(sm.Used),
		SwapPct:   sm.UsedPercent,
	}
}

func collectDisks() []diskInfo {
	result := []diskInfo{}
	parts, _ := disk.Partitions(true)
	toMB := func(b uint64) uint64 { return b / 1024 / 1024 }

	for _, p := range parts {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}
		result = append(result, diskInfo{
			Mountpoint: p.Mountpoint,
			Device:     p.Device,
			Fstype:     p.Fstype,
			Total:      toMB(usage.Total),
			Used:       toMB(usage.Used),
			UsedPct:    usage.UsedPercent,
		})
	}
	return result
}

func collectTemps() []tempInfo {
	result := []tempInfo{}
	temps, _ := host.SensorsTemperatures()
	for _, t := range temps {
		result = append(result, tempInfo{
			Sensor: t.SensorKey,
			Temp:   t.Temperature,
		})
	}
	return result
}

func collectNet() []netInfo {
	out := []netInfo{}
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		addrs, _ := iface.Addrs()
		if len(addrs) == 0 {
			continue
		}
		flags := iface.Flags.String()
		for _, a := range addrs {
			out = append(out, netInfo{
				Interface: iface.Name,
				Addr:      a.String(),
				Flags:     flags,
			})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Interface < out[j].Interface })
	return out
}

func getPrimaryIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		if ip == nil || ip.IsLoopback() {
			continue
		}
		if ipv4 := ip.To4(); ipv4 != nil {
			return ipv4.String(), nil
		}
	}
	return "", fmt.Errorf("no valid IP found")
}

// -------------------- System Info --------------------

func Info() SystemInfo {
	h, _ := host.Info()
	IPAddress, _ := getPrimaryIP()
	cpu := collectCPU()
	memory := collectMemory()

	return SystemInfo{
		Hostname:  h.Hostname,
		OS:        h.Platform + " " + h.PlatformVersion,
		Kernel:    h.KernelVersion,
		Arch:      h.KernelArch,
		IPAddress: IPAddress,
		Cores:     cpu.Cores,
		Threads:   cpu.Threads,
		Memory:    memory.Total,
	}
}

// -------------------- System Metrics --------------------

func Metrics() SystemMetrics {
	cpu := collectCPU()
	mem := collectMemory()
	disks := collectDisks()
	netio := collectNet()
	temps := collectTemps()

	// Process count
	procs, _ := process.Processes()
	procCount := len(procs)

	// Load averages
	loadStat, _ := load.Avg()

	h, _ := host.Info()

	return SystemMetrics{
		Timestamp:  time.Now().Unix(),
		Uptime:     uint64(h.Uptime),
		Load1:      loadStat.Load1,
		Load5:      loadStat.Load5,
		Load15:     loadStat.Load15,
		Memory:    	mem,
		CPU:        cpu,
		Disk:       disks,
		NetIO:      netio,
		ProcCount:  procCount,
		Temps:      temps,
	}
}
