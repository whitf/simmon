@echo off

go run .\node.go .\agent.go .\alerts.go .\config.go .\dbPostgres.go .\dbSqlite.go .\node-heartbeat.go .\nodeApi.go .\qos.go  --conf ..\..\configs\node\EXAMPLE-simmon-node.conf
