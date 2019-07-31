package models

import (
	"fmt"
	"github.com/luosangnanka/goinfo"
	"testing"
)

func TestGetMemoryInfo(t *testing.T) {

	// mem.
	mem, err := goinfo.Memory()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("mem", mem)

	memUsage := PercentFormat(float64(mem.Mem.Used / mem.Mem.Total) * 100)
	fmt.Println("memUsage", memUsage)

}

func TestGetMemoryUsage(t *testing.T) {
	memUsage, err := getMemoryUsage()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("memUsage", PercentFormat(memUsage))
}
