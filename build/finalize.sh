#!/bin/sh
# Finalize the build

set -x
set -e

# Final cleaning
rm -rf /app/build

rm -rf /tmp/go
rm -rf /usr/local/go