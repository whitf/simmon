@echo off

go run .\agent.go .\config.go .\cpuWin.go .\diskWin.go .\heartbeat.go .\memWin.go .\netComm.go .\process.go --conf .\EXAMPLE-simmon-agent.conf
