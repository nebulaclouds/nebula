#!/usr/bin/env bash

set -e

echo "Generating Nebula Configuration Documents"
CUR_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
ROOT_DIR=${CUR_DIR}/..
OUTPUT_DIR="${ROOT_DIR}"/docs/deployment/configuration/generated
GOBIN=${GOPATH:-~/go}/bin

make -C datacatalog compile
mv datacatalog/bin/datacatalog ${GOBIN}/datacatalog
make -C nebulaadmin compile
mv nebulaadmin/bin/nebulaadmin ${GOBIN}/nebulaadmin
make -C nebulaadmin compile_scheduler
mv nebulaadmin/bin/nebulascheduler ${GOBIN}/scheduler
make -C nebulapropeller compile_nebulapropeller
mv nebulapropeller/bin/nebulapropeller ${GOBIN}/nebulapropeller

output_config () {
CONFIG_NAME=$1
COMPONENT=$2
COMMAND=$3
OUTPUT_PATH=${OUTPUT_DIR}/${COMMAND}_config.rst

if [ -z "$CONFIG_NAME" ]; then
  log_err "output_config CONFIG_NAME value not specified in arg1"
  return 1
fi

if [ -z "$COMPONENT" ]; then
  log_err "output_config COMPONENT value not specified in arg2"
  return 1
fi

echo ".. _$COMPONENT-config-specification:

#########################################
Nebula $CONFIG_NAME Configuration
#########################################
" > "${OUTPUT_PATH}"

$GOBIN/$COMMAND config docs >> "${OUTPUT_PATH}"
}

output_config "Admin" nebulaadmin nebulaadmin
output_config "Propeller" nebulapropeller nebulapropeller
output_config "Datacatalog" nebuladatacatalog datacatalog
output_config "Scheduler" nebulascheduler scheduler
