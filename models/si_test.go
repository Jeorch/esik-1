package models

import (
	"fmt"
	"testing"
)

func TestSystemInfo_ExtractSystemInfo(t *testing.T) {
	si := new(SystemInfo)
	err := si.ExtractSystemInfo()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(si)
}
