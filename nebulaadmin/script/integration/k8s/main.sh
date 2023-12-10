#!/usr/bin/env bash
set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

echo ""
echo "waiting up to 5 minutes for kubernetes to start..."

K8S_TIMEOUT="300"

SECONDS=0
while ! systemctl is-active --quiet multi-user.target; do
  sleep 2
  if [ "$SECONDS" -gt "$K8S_TIMEOUT" ]; then
    echo "ERROR: timed out waiting for kubernetes to start."
    exit 1
  fi
done

echo "kubernetes started in $SECONDS seconds."
echo ""

# Load the locally-built nebulaadmin image
docker load -i /images/nebulaadmin

# Start nebulaadmin and dependencies
kubectl create -f "${DIR}/integration.yaml"

# In debug mode, run bash instead of running the tests
if [ -n "$DOCKERNETES_DEBUG" ]; then
  bash
fi

# Wait for nebulaadmin deployment to complete
kubectl -n nebula rollout status deployment nebulaadmin

# Get the name of the nebulaadmin pod
POD_NAME=$(kubectl get pods -n nebula -o go-template="{{range .items}}{{.metadata.name}}:{{end}}" | tr ":" "\n" | grep nebulaadmin)

# Launch the integration tests
kubectl exec -it -n nebula "$POD_NAME" -- make -C /go/src/github.com/nebulaclouds/nebulaadmin integration
