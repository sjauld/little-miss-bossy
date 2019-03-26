package main

import (
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	cpuStatsHeader = "cpu"
	cpuStatsPath   = "/proc/stat"
)

var (
	errCPUSampleNotObtained = errors.New("Failed to obtain a CPU sample")
)

func cpuSample() (error, *sample) {
	s := &sample{}

	data, err := ioutil.ReadFile(cpuStatsPath)
	if err != nil {
		return err, nil
	}

	for _, row := range strings.Split(string(data), "\n") {
		ok, err := s.processProcRow(row)
		if ok {
			return nil, s
		}

		if err != nil {
			log.Printf("[WARNING] something bad happened: %v", err)
		}
	}

	return errCPUSampleNotObtained, nil
}

type sample struct {
	idle  uint64
	total uint64
}

// processProcRow checks if the row provided is the cpu totals row. If so it
// calculates the stats we're after
//
// returns true as the first output if we're finished processing
func (s *sample) processProcRow(row string) (ok bool, err error) {
	cols := strings.Fields(row)
	if len(cols) == 0 || cols[0] != cpuStatsHeader {
		return
	}

	ok = true

	// add every column to the total time, and just the 5th column to idle time
	for i, col := range cols {
		// skip the header
		if i == 0 {
			continue
		}

		var val uint64

		val, err = strconv.ParseUint(col, 10, 64)
		if err != nil {
			return
		}

		s.total += val

		if i == 4 {
			s.idle = val
		}
	}

	return
}
