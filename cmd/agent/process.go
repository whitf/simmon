package main

/**
 * cmd/agent/process.go
 * simmon - simple monitoring
 * 
 * process monitoring
 */



import (
 	"bytes"
 	"os/exec"
 	"strconv"
	"time"
)

func process(t int, options []ServiceOption) int {
	// Monitor specific process

	var processName, processCount = "", ""

	for i := range options {
		switch options[i].Key {
			case "processName":
				processName = options[i].Value

			case "processCount":
				processCount = options[i].Value
		}
	}

	for {
		time.Sleep(time.Duration(t) * time.Second)
		
		grep := exec.Command("grep", processName)
		ps := exec.Command("ps", "cax")

		pipe, _ := ps.StdoutPipe()
		defer pipe.Close()

		grep.Stdin = pipe

		ps.Start()
		res, _ := grep.Output()

		lcount := bytes.Count(res, []byte("\n"))

		msg := "Expecting " + processCount + " instance(s) of process " + processName + ", found " + strconv.Itoa(lcount) + "."

		q := QoS {
			Agent: c.Name,
			Date: time.Now().String(),
			Msg: msg,
			QoSType: "PROCESS-" + processName,
		}

		pcount, _ := strconv.Atoi(processCount)

		if bytes.Count(res, []byte("\n")) < pcount {
			q.Alert = true
			q.AlertLevel = "Alert"
			q.AlertValue = "Process Alert"
		} else {
			q.Alert = false
			q.AlertLevel = "Clean"
			q.AlertValue = "Clean"
		}

		q.report()
	}

	return 0
}


