// +build linux
package main

/**
 * mem.go
 * memory monitoring function
 */

import (
	"log"
	"syscall"
	"strconv"
	"time"
)

type MemoryThreshold struct {
	memoryFreePercent			float64
	swapFreePercent				float64
}

var mt MemoryThreshold


func memCheck(e error) {
	if e != nil {
		log.Fatal("Error accessing memory statistics", e)
	}
}

func memfree(t int, options []ServiceOption) int {
	for {
		time.Sleep(time.Duration(t) * time.Second)
		
		in := &syscall.Sysinfo_t{}
		err := syscall.Sysinfo(in)
		memCheck(err)

		var k uint64
		k = 1024

		freeram := in.Freeram/k
		totalram := in.Totalram/k
		alert := true
		alertLevel := ""

		if float64(freeram) < mt.memoryFreePercent * float64(totalram) {
			alert = true
			alertLevel = "Warn"
		} else {
			alertLevel = "Clean"
			alert = false
		}

		alertValue := strconv.FormatFloat((float64(freeram) * float64(100))/float64(totalram), 'f', 2, 64)
		msg := strconv.FormatUint(freeram, 10) + "k free memory of " + strconv.FormatUint(totalram, 10) + "k total (" + alertValue + "% free)."
	
		q := QoS {
			Date: time.Now().String(),
			QoSType: "MEMFREE",
			Msg: msg,
			Alert: alert,
			AlertLevel: alertLevel,
			AlertValue: alertValue,
			Agent: c.Name,
		}
		
		q.report()
	}

	return 0
}

func swapfree(t int, options []ServiceOption) int {
	for {
		// Build data from monitoring service.
		time.Sleep(time.Duration(t) * time.Second)

		in := &syscall.Sysinfo_t{}
		err := syscall.Sysinfo(in)
		memCheck(err)

		var k uint64
		k = 1024

		freeswap := in.Freeswap/k
		totalswap := in.Totalswap/k
		alert := false
		alertLevel := ""

		if float64(freeswap) < mt.swapFreePercent * float64(totalswap) {
			alert = true
			alertLevel = "Warn"
		} else {
			alertLevel = "Clean"

		}

		alertValue := strconv.FormatFloat((float64(freeswap) * float64(100))/float64(totalswap), 'f', 2, 64)
		msg := strconv.FormatUint(freeswap, 10) + "k free swap of " + strconv.FormatUint(totalswap, 10) + "k total (" + alertValue + "% free)."

		q := QoS {
			Date: time.Now().String(),
			QoSType: "SWAPFREE",
			Msg: msg,
			Alert: alert,
			AlertLevel: alertLevel,
			AlertValue: alertValue,
			Agent: c.Name,
		}
		q.report()
	}

	return 0
}
