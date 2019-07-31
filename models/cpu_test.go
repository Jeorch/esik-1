package models

import (
	"fmt"
	"testing"
)

func TestGetCpuUsage(t *testing.T) {
	// cpu.
	cpuUsage, err := getCpuUsage()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("cpuUsage", PercentFormat(cpuUsage))
}
