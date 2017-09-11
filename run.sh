#!/bin/bash -e

env GOPATH=$GOPATH:/srv/ go build
mv weasel bin/
bin/weasel -port 8082 -config="/srv/src/weasel/conf.d"