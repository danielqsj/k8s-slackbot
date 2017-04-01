#!/bin/sh
# Finalize the build

set -x
set -e

# Final cleaning
rm -rf /app/k8s-slackbot/build

rm -rf /tmp/go
rm -rf /usr/local/go