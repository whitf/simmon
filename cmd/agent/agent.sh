#!/usr/bin/bash

echo "starting cmd/agent"

go run agent.go config.go cpu.go disk.go heartbeat.go mem.go netComm.go process.go
