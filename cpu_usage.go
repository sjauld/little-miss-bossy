package main

import (
	"time"
)

type cpuDelta struct {
	start  *sample
	finish *sample
}

func newCPUDelta(t time.Duration) (error, *cpuDelta) {
	d := &cpuDelta{}

	err, start := cpuSample()
	if err != nil {
		return err, nil
	}
	d.start = start

	time.Sleep(t)

	err, finish := cpuSample()
	if err != nil {
		return err, nil
	}
	d.finish = finish

	return nil, d
}

func (d *cpuDelta) idleTime() uint64 {
	return d.finish.idle - d.start.idle
}

func (d *cpuDelta) totalTime() uint64 {
	return d.finish.total - d.start.total
}

func (d *cpuDelta) idlePercent() float64 {
	return 100 * float64(d.idleTime()) / float64(d.totalTime())
}

func (d *cpuDelta) utilisation() float64 {
	return 100 - d.idlePercent()
}
