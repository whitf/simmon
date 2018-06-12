#!/usr/bin/bash

mkdir -p /etc/simmon
cp ./etc/simmon/simmon-web.conf /etc/simmon/

mkdir -p /opt/simmon/web/static
cp ./opt/simmon/web/simmon-web /opt/simmon/web/simmon-web
cp ./opt/simmon/web/alerts.html /opt/simmon/web/
cp ./opt/simmon/web/static/* /opt/simmon/web/static/

