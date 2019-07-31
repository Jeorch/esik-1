package models

import (
	"fmt"
	"testing"
)

func TestGetDiskUsage(t *testing.T) {
	dfUsage, err := getDiskUsage("/")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("dfUsage", dfUsage)
}
