#!/usr/bin/env bash

set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

cp ${DIR}/pull_request_template.md ${DIR}/../../../pull_request_template.md
