package main

/**
 * disk.go
 * disk monitoring functions
 * 
 */

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)


func diskIO(t int, options[]ServiceOption) int {
	return 0
}

func diskFree(t int, options[]ServiceOption) int {

	var volumeName, thresholdUnits = "", ""
	var freeThreshold = 0.00

	for i := range options {
		switch options[i].Key {
			case "volumeName":
				volumeName = options[i].Value

			case "thresholdUnits":
				thresholdUnits = options[i].Value

			case "freeThreshold":
				freeThreshold, _ = strconv.ParseFloat(options[i].Value, 64)
		}

	}

	var fieldSelect = 0
	if thresholdUnits == "%" {
		fieldSelect = 4
	} else {
		fieldSelect = 3
	}

	for {
		time.Sleep(time.Duration(t) * time.Second)

		grep := exec.Command("grep", "-v", "Filesystem")
		df := exec.Command("df", "-h", volumeName)

		pipe, _ := df.StdoutPipe()
		defer pipe.Close()

		grep.Stdin = pipe
		df.Start()

		res, rErr := grep.Output()

		if rErr == nil {

			s := strings.Fields(string(res))

			freeSpace, _ := strconv.ParseFloat(s[fieldSelect][:len(s[fieldSelect])-1], 64)
			if thresholdUnits == "%" {
				freeSpace = 100.0 - freeSpace
			}

			msg := "Expecting free space of at least " + strconv.FormatFloat(freeThreshold, 'f', 2, 64) + thresholdUnits + ", found " + strconv.FormatFloat(freeSpace, 'f', 2, 64) + thresholdUnits + "."

			q := QoS {
				Agent: c.Name,
				Date: time.Now().String(),
				Msg: msg,
				QoSType: "DISKFREE",
			}
	

			if freeThreshold < freeSpace {
				q.Alert = false
				q.AlertLevel = "Clean"
				q.AlertValue = "Clean"
			} else {
				q.Alert = true
				q.AlertLevel = "Alert"
				q.AlertValue = "Disk Free Alert"
			}

			q.report()
		} else {
			log.Println("Error parsing output for diskFree.")
			log.Println(" Skipping this reporting interval and trying again.")
		}
	}

	return 0
}