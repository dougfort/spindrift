#! /bin/bash
# create the binary for spindrift

set -e
set -x

pushd $GOPATH/src/github.com/dougfort/spindrift
go install -race
popd
