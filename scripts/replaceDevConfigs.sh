#!/usr/bin/bash

echo ""
echo "Replacing .conf files in /etc/simmon/ (requires write permissions)."

cp ../configs/agent/EXAMPLE-simmon-agent.conf /etc/simmon/simmon-agent.conf
cp ../configs/node/EXAMPLE-simmon-node.conf /etc/simmon/simmon-node.conf
cp ../configs/web/EXAMPLE-simmon-web.conf /etc/simmon/simmon-web.conf

echo ""

ls -alh /etc/simmon/