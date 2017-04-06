#! /bin/bash
# run the binary for spindrift

set -e
set -x

for (( i=0 ; i < 2 ; i++ ))
do
	address="127.0.0.1:$(( 9000+i ))"
	seed="127.0.0.1:$(( 9000 ))"
	$GOPATH/bin/spindrift -address=$address -seed-peers=$seed &
done
