#!/usr/bin/env bash

# WARNING: THIS FILE IS MANAGED IN THE 'BOILERPLATE' REPO AND COPIED TO OTHER REPOSITORIES.
# ONLY EDIT THIS FILE FROM WITHIN THE 'LYFT/BOILERPLATE' REPOSITORY:
#
# TO OPT OUT OF UPDATES, SEE https://github.com/lyft/boilerplate/blob/master/Readme.rst

set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

# TODO: load all images
docker tag ${IMAGE} "nebulaadmin:test"
kind load docker-image nebulaadmin:test

# start nebulaadmin and dependencies
kubectl apply -f "${DIR}/k8s/integration.yaml"

# This is a separate function so that we can potentially reuse in the future when we have more than one test
function wait_for_nebula_deploys() {
    SECONDS=0
    echo ""
    echo "waiting for nebula deploy to complete..."
    echo ""

    # wait for nebulaadmin deployment to complete
    kubectl -n nebula rollout status deployment nebulaadmin
    echo ""

    echo "Nebula deployed in $SECONDS seconds."
}

wait_for_nebula_deploys

## get the name of the nebulaadmin pod
POD_NAME=$(kubectl get pods -n nebula --field-selector=status.phase=Running -o go-template="{{range .items}}{{.metadata.name}}:{{end}}" | tr ":" "\n" | grep nebulaadmin)

echo $POD_NAME

# launch the integration tests
kubectl exec -it -n nebula "$POD_NAME" -- make -C /go/src/github.com/nebulaclouds/nebulaadmin integration
