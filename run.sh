#!/bin/bash -e

env GOPATH=$GOPATH:/srv/ go build server.go
mv server bin/weasel
bin/weasel -port 8082 -config="/srv/src/weasel/conf.d"