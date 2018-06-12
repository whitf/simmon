#!/usr/bin/env groovy

/**
 * simmon/build/Jenkinsfile.groovy
 * simmon - simple monitoring
 *
 */

import groovy.json.JsonSlurperClassic

node {
	try {
		notifyBuild('STARTED')

		withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"]) {

			env.PATH="${GOPATH}/bin:$PATH"

			Date date = new Date()
			String buildTime = date.format("yyyyMMdd-HHmm")

			stage('Pre Test') {
				echo 'Pulling Dependencies'

				sh 'go version'

				sh 'go get github.com/gorilla/mux'
				sh 'go get github.com/lib/pq'
				sh 'go get github.com/mattn/go-sqlite3'

				sh 'mkdir -p $GOPATH/archive/'

			}

			stage('Checkout') {
				echo 'Checking out SCM...'
				sh 'cd $GOPATH'

				sh 'mkdir -p $GOPATH/src/simmon/'
				ws("$GOPATH/src/simmon/") {
					checkout scm
				}

				//checkout scm
				sh 'pwd && ls -alh'
			}

			stage('Test') {
				echo 'Test stage'

			}
			

			stage('Build and Package Simmon Node') {
				def json = readFile(file: "$GOPATH/src/simmon/configs/node/EXAMPLE-simmon-node.conf")
				def data = new JsonSlurperClassic().parseText(json)

				sh """cd $GOPATH/src/simmon/cmd/node/ && go build -ldflags '-s'"""

				sh 'mkdir -p $GOPATH/pkg/simmon/node'

				sh 'mkdir -p $GOPATH/pkg/simmon/node/etc/simmon/'
				sh 'cp $GOPATH/src/simmon/configs/node/EXAMPLE-simmon-node.conf $GOPATH/pkg/simmon/node/etc/simmon/simmon-node.conf'

				sh 'mkdir -p $GOPATH/pkg/simmon/node/var/log/simmon/node/'
				
				sh 'mkdir -p $GOPATH/pkg/simmon/node/opt/simmon/node/'
				sh 'cp $GOPATH/src/simmon/cmd/node/node $GOPATH/pkg/simmon/node/opt/simmon/node/simmon-node'

				sh 'cp $GOPATH/src/simmon/scripts/node/* $GOPATH/pkg/simmon/node/'

				sh "cd $GOPATH/pkg/simmon/node && tar -zcvf $GOPATH/archive/simmon-node-${data.version}-${buildTime}.tar.gz ."

			}

			stage('Build and Package Simmon Agent') {
				def json = readFile(file: "$GOPATH/src/simmon/configs/agent/EXAMPLE-simmon-agent.conf")
				def data = new JsonSlurperClassic().parseText(json)

				sh """cd $GOPATH/src/simmon/cmd/agent/ && go build  -ldflags '-s' agent.go config.go cpu.go disk.go heartbeat.go mem.go netComm.go process.go"""

				sh 'mkdir -p $GOPATH/pkg/simmon/agent/etc/simmon/'
				sh 'cp $GOPATH/src/simmon/configs/agent/EXAMPLE-simmon-agent.conf $GOPATH/pkg/simmon/agent/etc/simmon/simmon-agent.conf'

				sh 'mkdir -p $GOPATH/pkg/simmon/agent/var/log/simmon/agent/'

				sh 'mkdir -p $GOPATH/pkg/simmon/agent/opt/simmon/agent/'
				sh 'cp $GOPATH/src/simmon/cmd/agent/agent $GOPATH/pkg/simmon/agent/opt/simmon/agent/simmon-agent'

				sh 'cp $GOPATH/src/simmon/scripts/agent/* $GOPATH/pkg/simmon/agent/'

				sh "cd $GOPATH/pkg/simmon/agent && tar -zcvf $GOPATH/archive/simmon-agent-${data.version}-${buildTime}.tar.gz ."

			}

			stage('Build and Package Simmon Web Server') {
				def json = readFile(file: "$GOPATH/src/simmon/configs/web/EXAMPLE-simmon-web.conf")
				def data = new JsonSlurperClassic().parseText(json)

				sh """cd $GOPATH/src/simmon/cmd/web/ && go build -ldflags '-s'"""

				sh 'mkdir -p $GOPATH/pkg/simmon/web/etc/simmon/'
				sh 'cp $GOPATH/src/simmon/configs/web/EXAMPLE-simmon-web.conf $GOPATH/pkg/simmon/web/etc/simmon/simmon-web.conf'

				sh 'mkdir -p $GOPATH/pk/simmon/web/var/log/simmon/web'

				sh 'mkdir -p $GOPATH/pkg/simmon/web/opt/simmon/web/static/'
				sh 'cp $GOPATH/src/simmon/cmd/web/web $GOPATH/pkg/simmon/web/opt/simmon/web/simmon-web'
				sh 'cp $GOPATH/src/simmon/cmd/web/alerts.html $GOPATH/pkg/simmon/web/opt/simmon/web/'
				sh 'cp $GOPATH/src/simmon/cmd/web/static/* $GOPATH/pkg/simmon/web/opt/simmon/web/static/'

				sh 'cp $GOPATH/src/simmon/scripts/web/* $GOPATH/pkg/simmon/web/'

				sh "cd $GOPATH/pkg/simmon/web && tar -zcvf $GOPATH/archive/simmon-web-${data.version}-${buildTime}.tar.gz ."

			}

			stage('Archive') {
				echo 'Archive finished binaries.'

				archive '$GOPATH/archive/*'
			}

		}

	} catch(e) {
		currentBuild.result = "FAILED"

		throw e

	} finally {
		notifyBuild(currentBuild.result)

		def bs = currentBuild.result ?: 'SUCCESSFUL'
		if(bs == 'SUCCESSFUL') {
			echo 'build success'
		}
	}

}


def notifyBuild(String buildStatus = 'STARTED') {
	buildStatus = buildStatus ?: 'SUCCESSSFUL'

	def colorName = 'RED'
	def colorCode = '#FF0000'

}
