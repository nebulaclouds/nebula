#!/usr/bin/env bash

set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

cp ${DIR}/CODE_OF_CONDUCT.md ${DIR}/../../../CODE_OF_CONDUCT.md
