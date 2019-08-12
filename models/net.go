package models

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type NetStatus struct {
	Receive  int64  `json:"receive"`
	Transmit int64  `json:"transmit"`
}

// String format the NetStatus struct.
func (ns *NetStatus) String() (net string) {
	if ns == nil {
		return
	}

	return fmt.Sprintf("receive:%s, transmit:%s", ByteSize(ns.Receive).String(), ByteSize(ns.Transmit).String())
}

type NetSpeed struct {
	Receive  float64  `json:"receive"`
	Transmit float64  `json:"transmit"`
}

// String format the NetStatus struct.
func (ns *NetSpeed) String() (net string) {
	if ns == nil {
		return
	}

	return fmt.Sprintf("receive:%2.fkb/s, transmit:%2.fkb/s", ns.Receive/1024, ns.Transmit/1024)
}

func getNetSpeed() (netSpeed *NetSpeed, err error) {

	s1, err := getNetStatus()
	bmerror.PanicError(err)
	netTimeMsStr := os.Getenv("ESIK_NET_DURATION_MS")
	if netTimeMsStr == "" {
		netTimeMsStr = "100"	//100ms
	}
	netTimeMs, err := strconv.ParseInt(netTimeMsStr, 10, 64)
	if err != nil {
		panic(err.Error())
	}
	time.Sleep(time.Duration(netTimeMs) * time.Millisecond)
	s2, err := getNetStatus()
	bmerror.PanicError(err)
	receiveSpeed := float64(s2.Receive - s1.Receive)/0.1	//默认速度以kb/s为单位
	transmitSpeed := float64(s2.Transmit - s1.Transmit)/0.1	//默认速度以kb/s为单位

	return &NetSpeed{Receive: receiveSpeed, Transmit: transmitSpeed}, err
}

func getNetStatus() (netStatus *NetStatus, err error) {

	b, err := ioutil.ReadFile(gNetFile)
	if err != nil {
		return
	}
	s := strings.SplitAfter(string(b), "\n")
	length := len(s)
	var totalReceive int64
	var totalTransmit int64
	for i := 2; i < length; i++ {
		t := strings.Fields(s[i])
		if len(t) == 17 {
			receive := string2Int64(t[1])
			transmit := string2Int64(t[10])
			totalReceive += receive
			totalTransmit += transmit
		}
	}
	return &NetStatus{Receive: totalReceive, Transmit: totalTransmit}, err
}
