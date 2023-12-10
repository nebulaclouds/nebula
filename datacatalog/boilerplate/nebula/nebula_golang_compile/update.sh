#!/usr/bin/env bash

# WARNING: THIS FILE IS MANAGED IN THE 'BOILERPLATE' REPO AND COPIED TO OTHER REPOSITORIES.
# ONLY EDIT THIS FILE FROM WITHIN THE 'NEBULAORG/BOILERPLATE' REPOSITORY:
# 
# TO OPT OUT OF UPDATES, SEE https://github.com/nebulaclouds/boilerplate/blob/master/Readme.rst

set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

echo "     - generating ${DIR}/nebula_golang_compile.sh"
sed -e "s/{{REPOSITORY}}/${REPOSITORY}/g" ${DIR}/nebula_golang_compile.Template > ${DIR}/nebula_golang_compile.sh
