#!/usr/bin/bash

echo "starting cmd/node"

go run config.go dbPostgres.go dbSqlite.go node-heartbeat.go alerts.go qos.go node.go agent.go nodeApi.go
