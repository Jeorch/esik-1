package models

import (
	"errors"
	"esik/kits"
	"fmt"
	"strconv"
	"strings"
)

func getDiskUsage(mountPoint string) (result float64, err error) {
	dfInfo, err := kits.GetCmdInfo(gDiskCmd)
	if err != nil {
		return
	}
	s := strings.SplitAfter(dfInfo, "\n")

	for _, v := range s {
		cc := strings.Fields(v)
		if len(cc) == 0 {
			return -1, errors.New("Read \"gDiskCmd\" failed because first line is empty.")
		}
		//暂时仅监控根目录的存储情况
		if cc[5] == mountPoint {
			result, err = strconv.ParseFloat(cc[4][:len(cc[4])-1], 64)
			return
		}
	}
	return -1, errors.New(fmt.Sprint("Not found ", mountPoint, " mountPoint."))
}