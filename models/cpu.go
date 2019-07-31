package models

import (
	"errors"
	"io/ioutil"
	"strings"
	"time"
)

func getCpuUsage() (float64, error) {

	idle1, total1, err := getIdleTotal()
	if err != nil {
		panic(err.Error())
	}
	time.Sleep(1 * time.Second)
	idle2, total2, err := getIdleTotal()
	if err != nil {
		panic(err.Error())
	}
	idle := float64(idle2 - idle1)
	total := float64(total2 - total1)
	result := (total - idle) / total * 100

	return result, nil
}

func getIdleTotal() (idle int64, total int64, err error) {
	b, err := ioutil.ReadFile(gCPUFile)
	if err != nil {
		return -1, -1, err
	}
	s := strings.SplitAfter(string(b), "\n")

	cc := strings.Fields(s[0])
	if len(cc) == 0 {
		return -1, -1, errors.New("Read \"gCPUFile\" failed because first line is empty.")
	}
	if strings.HasPrefix(cc[0], "cpu") {
		if len(cc) < 8 {
			return -1, -1, errors.New("cpu info fields has no enough fields")
		}
		user := string2Int64(cc[1])
		nice := string2Int64(cc[2])
		system := string2Int64(cc[3])
		idle = string2Int64(cc[4])
		iowait := string2Int64(cc[5])
		irq := string2Int64(cc[6])
		softirq := string2Int64(cc[7])
		total = sum(user, nice, system, idle, iowait, irq, softirq)
	}
	return
}
