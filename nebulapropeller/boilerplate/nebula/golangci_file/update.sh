#!/usr/bin/env bash

set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

# Clone the .golangci file
echo "     - copying ${DIR}/.golangci to the root directory."
cp ${DIR}/.golangci.yml ${DIR}/../../../.golangci.yml
