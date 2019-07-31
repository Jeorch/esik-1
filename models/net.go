package models

import (
	"fmt"
	"io/ioutil"
	"strings"
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
