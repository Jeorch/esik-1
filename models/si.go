package models

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmlog"
	"os"
	"time"
)

var (
	gTimeFarmat = "2006-01-02 15:04:05"
	gCPUFile    = "/proc/stat"
	gMemFile    = "/proc/meminfo"
	gDiskCmd    = "df -lh | grep /"
	gNetFile    = "/proc/net/dev"
)

type SystemInfo struct {
	Time      string
	Hostname  string
	Ip        string
	CpuUsage  float64
	MemUsage  float64
	DiskUsage float64
	NetStatus  *NetStatus
}

// String format the SystemInfo struct.
func (si *SystemInfo) String() (net string) {
	if si == nil {
		return
	}

	return fmt.Sprintf("Time:%s, Hostname:%s, Ip:%s, CpuUsage:%.2f%%, MemUsage:%.2f%%, DiskUsage:%.2f%%, NetStatus:%v",
		si.Time, si.Hostname, si.Ip, si.CpuUsage, si.MemUsage, si.DiskUsage, si.NetStatus)
}

func (si *SystemInfo) ExtractSystemInfo() (err error) {

	si.Time = time.Now().Format(gTimeFarmat)
	si.Hostname, err = os.Hostname()
	if err != nil {
		bmlog.StandardLogger().Error("Cause by: ", err.Error())
		return
	}
	si.Ip, err = GetIntranetIp()
	if err != nil {
		bmlog.StandardLogger().Error("Cause by: ", err.Error())
		return
	}
	si.CpuUsage, err = getCpuUsage()
	if err != nil {
		bmlog.StandardLogger().Error("Cause by: ", err.Error())
		return
	}
	si.MemUsage, err = getMemoryUsage()
	if err != nil {
		bmlog.StandardLogger().Error("Cause by: ", err.Error())
		return
	}
	mountPoint := os.Getenv("ESIK_MOUNT_POINT")
	if mountPoint == "" {
		mountPoint = "/"
	}
	si.DiskUsage, err = getDiskUsage(mountPoint)
	if err != nil {
		bmlog.StandardLogger().Error("Cause by: ", err.Error())
		return
	}
	si.NetStatus, err = getNetStatus()
	if err != nil {
		bmlog.StandardLogger().Error("Cause by: ", err.Error())
		return
	}
	return
}
