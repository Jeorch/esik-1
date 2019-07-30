package main

import (
	"esik/models"
	"fmt"
)

func main() {
	fmt.Println("Esik Start.")
	// cpu.
	cpuInfo, err := models.GetCpuStatus()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("cpuinfo", cpuInfo)
}
