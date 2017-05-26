package models

import (
	"container/list"
	"encoding/base64"
	"encoding/json"
	"github.com/riveryang/sysgo/aes"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"net"
	"strings"
)

type PCStat struct {
	NetStats     []NetStat              `json:"net"`
	CpuInfoStats []cpu.InfoStat         `json:"cpu"`
	DiskStats    []disk.PartitionStat   `json:"disk"`
	MemStat      *mem.VirtualMemoryStat `json:"mem"`
	HostStat     *host.InfoStat         `json:"host"`
}

type NetStat struct {
	Ip  string `json:"ip"`
	Mac string `json:"mac"`
}

func NewPcStat() (*PCStat, error) {
	stat := PCStat{}
	if netStats, err := initNetStat(); err != nil {
		return nil, err
	} else {
		stat.NetStats = netStats
	}

	if cpuInfoStats, err := cpu.Info(); err != nil {
		return nil, err
	} else {
		stat.CpuInfoStats = cpuInfoStats
	}

	if parts, err := disk.Partitions(true); err != nil {
		return nil, err
	} else {
		stat.DiskStats = parts
	}

	if hostInfo, err := host.Info(); err != nil {
		return nil, err
	} else {
		stat.HostStat = hostInfo
	}

	if vm, err := mem.VirtualMemory(); err != nil {
		return nil, err
	} else {
		stat.MemStat = vm
	}

	return &stat, nil
}

func initNetStat() ([]NetStat, error) {
	itfs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var netStats = list.New()
	for _, itf := range itfs {
		addrs, err := itf.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					netStats.PushBack(NetStat{Ip: ipnet.IP.String(), Mac: strings.ToUpper(itf.HardwareAddr.String())})
				}
			}
		}
	}

	var tmp = make([]NetStat, netStats.Len())
	var index = 0
	for e := netStats.Front(); e != nil; e = e.Next() {
		tmp[index] = e.Value.(NetStat)
		index++
	}

	return tmp, nil
}

func (ps *PCStat) Encrypt() (string, error) {
	data, err := json.Marshal(ps)
	if err != nil {
		return "", err
	}

	encrypt, err := aes.AesEncrypt(data, aes.DefaultKey())
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encrypt), nil
}
