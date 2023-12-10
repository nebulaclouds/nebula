#!/bin/bash

# Note that this file is meant to be run on OSX by a user with the necessary GitHub privileges.
# This script
#    a) clones the two Nebula repositories from which additional RSTs not in this nebula repo, need to be generated.
#       namely nebulakit, and nebulaidl
#    b) runs a docker image to parse through the cloned repos, and creates the RSTs in the _rsts/ folder, which has
#       been added to gitignore.

set -ex

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
BASEDIR="${DIR}/.."

# Set up a temp directory
mkdir ${BASEDIR}/_repos || true
REPOS_DIR=`mktemp -d "${BASEDIR}/_repos/XXXXXXXXX"`

# Clone all repos
echo "Cloning Nebulaidl"
git clone https://github.com/lyft/nebulaidl.git --single-branch --branch v${NEBULAIDL_VERSION} ${REPOS_DIR}/nebulaidl
echo "Cloning Nebulakit"
git clone https://github.com/lyft/nebulakit.git --single-branch --branch v${NEBULAKIT_VERSION} ${REPOS_DIR}/nebulakit

# Generate documentation by running script inside the generation container
docker run --rm -t -e NEBULAKIT_VERSION=${NEBULAKIT_VERSION} -v ${BASEDIR}:/base -v ${REPOS_DIR}:/repos -v ${BASEDIR}/_rsts:/_rsts ghcr.io/nuclyde-io/docbuilder:e70cfafe3397c3b23d11183973c0f44e0024f025 /base/docs_infra/in_container_rst_generation.sh

# Cleanup
rm -rf ${REPOS_DIR}/* || true
