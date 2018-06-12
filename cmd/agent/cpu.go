package main

/**
 * cmd/agent/cpu.go
 * simmon - simmple monitoring
 *
 * cpu monitoring
 */

import (
	"os/exec"
	"strconv"
	"time"
)

func cpuPerc(t int, options []ServiceOption) int {
	return 0
}

func load(t int, options []ServiceOption) int {
	loadThreshold, _ := strconv.ParseFloat(options[0].Value, 64)

	for {
		time.Sleep(time.Duration(t) * time.Second)

		awk := exec.Command("awk", "{ print $9 }")
		uptime := exec.Command("uptime", "")

		pipe, _ := uptime.StdoutPipe()
		defer pipe.Close()

		awk.Stdin = pipe

		uptime.Start()
		res, _ := awk.Output()

		load5, _ := strconv.ParseFloat(string(res[:(len(res)-2)]), 64)

		msg := "Expecting CPU load of " + options[0].Value + " or less, found " + string(res[:(len(res)-2)]) + "."

		q := QoS {
			Agent: c.Name,
			Date: time.Now().String(),
			Msg: msg,
			QoSType: "CPULOAD",
		}

		if loadThreshold > load5 {
			q.Alert = false
			q.AlertLevel = "Clean"
			q.AlertValue = "Clean"
		} else {
			q.Alert = true
			q.AlertLevel = "Alert"
			q.AlertValue = "CPU Load Alert"
		}

		q.report()
	}

	return 0
}


