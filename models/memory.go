package models

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getMemoryUsage() (float64, error) {

	b, err := ioutil.ReadFile(gMemFile)
	if err != nil {
		return -1, err
	}
	s := strings.SplitAfter(string(b), "\n")
	var m = make([]ByteSize, 0)

	for _, v := range s {
		if v == "" {
			continue
		}
		mm := strings.Split(v, ":")
		if len(mm) < 2 {
			err = fmt.Errorf("mem info fields has no enough fields")
			return -1, err
		}
		info := strings.Replace(mm[1], "kB", "", -1)
		info = strings.TrimSpace(info)
		m = append(m, ByteSize(string2Float64(info)*1024))
	}
	if len(m) < 14 {
		err = fmt.Errorf("mem info fields has no enough fields")
		return -1, err
	}

	memUsage := float64((m[0] - m[2]) / m[0]) * 100
	return memUsage, err
}
