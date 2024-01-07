#!/usr/bin/env bash

set -eu

CONFIG_FILE=$1; shift
EXTRA_FLAGS=( "$@" )

python ./boilerplate/nebula/end2end/run-tests.py $NEBULASNACKS_VERSION $NEBULASNACKS_PRIORITIES $CONFIG_FILE ${EXTRA_FLAGS[@]}
