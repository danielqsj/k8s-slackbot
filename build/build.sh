#!/bin/sh
set -x
set -e

# Set temp environment vars
export GOPATH=/tmp/go
export PATH=/usr/local/go/bin:${PATH}:${GOPATH}/bin
export GO15VENDOREXPERIMENT=1

# Install build deps
apk --no-cache --no-progress add --virtual build-deps build-base linux-pam-dev

# Build k8s-slackbot
mkdir -p ${GOPATH}/src/github.com/danielqsj/
ln -s /app/build ${GOPATH}/src/github.com/danielqsj/k8s-slackbot
cd ${GOPATH}/src/github.com/danielqsj/k8s-slackbot
make k8s-slackbot
chmod +x k8s-slackbot
mv k8s-slackbot /app/

# Cleanup GOPATH
rm -r $GOPATH

# Remove build deps
apk --no-progress del build-deps