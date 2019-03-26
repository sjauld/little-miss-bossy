package main

import (
	"testing"
)

func TestCPUSample(t *testing.T) {
	err, _ := cpuSample()
	if err != nil {
		t.Error(err)
	}
}
