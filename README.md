### simmon - simple monitoring

simmon is a simple monitoring project, intended to provide basic monitoring 

simmon is an ongoing work in progress intended primarily for general learning and may be broken at any given time.

### Prerequisites:
#### General
github.com/gorilla/mux
github.com/lib/pq
github.com/mattn/go-sqlite3

#### simmon-agent
mpstat (part of the sysstat package)

### Build
A Jenkins pipeline script can be found in /build/Jenkinsfile.groovy

See the build file for additional information on automated builds as well as the necessary commands to build and package individual parts.

### Installation
Packages built by jenkins can be unzip and installed with the provided script (install.sh).

### Configuration
Example configuration files will be copied into /etc/simmon/ at the time of installation.

Values in configuration files should be configured as necessary.

Agents (and the web server) need access to one or more nodes on the configured ports in order to "report" QoS (Quality of Service) statistics.

### Use
Start each part as necessary.

> nohup ./simmon-node >> /dev/null 2>&1 &

> nohup ./simmon-web >> /dev/null 2>&1 &

> nohup ./simmon-agent >> /dev/null 2>&1 &

Access to information from the web server should be available at the configured host:port.