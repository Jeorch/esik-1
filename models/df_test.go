package models

import (
	"fmt"
	"syscall"
	"testing"
)

func TestGetDiskInfo(t *testing.T) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/", &fs)
	if err != nil {
		panic(err.Error())
	}
	all := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)
	fmt.Println("all \t", all)
	fmt.Println("free \t", free)
	fmt.Println("usage \t", PercentFormat(float64(all - free)/float64(all) * 100))
}

func TestGetDiskUsage(t *testing.T) {
	dfUsage, err := getDiskUsage("/")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("dfUsage", dfUsage)
}
