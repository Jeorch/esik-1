package models

import (
	"fmt"
	"github.com/luosangnanka/goinfo"
	"testing"
)

func TestGetNetInfo(t *testing.T) {
	netInfo, err := goinfo.Net()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("netInfo", netInfo)

}

func TestGetNetStatus(t *testing.T) {
	netStatus, err := getNetStatus()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("netStatus", netStatus)
}
