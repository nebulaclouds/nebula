#!/bin/sh

set -euo pipefail

# Apply cgroup v2 hack
cgroup-v2-hack.sh

trap 'pkill -P $$' EXIT
NEBULA_TIMEOUT=${NEBULA_TIMEOUT:-600}

monitor() {
    while : ; do
        for pid in $@ ; do
            kill -0 $pid &> /dev/null || exit 1
        done

        sleep 1
    done
}

# Start docker daemon
echo "Starting Docker daemon..."
file="/var/run/docker.pid"
if [ -f "$file" ] ; then
    rm "$file"
fi
dockerd &> /var/log/dockerd.log &
DOCKERD_PID=$!
timeout "$NEBULA_TIMEOUT" sh -c "until docker info &> /dev/null; do sleep 1; done" || ( echo >&2 "Timed out while waiting for dockerd to start"; exit 1 )
echo "Done."

# Start k3s
echo "Starting k3s cluster..."
KUBERNETES_API_PORT=${KUBERNETES_API_PORT:-6443}
k3s server --docker --no-deploy=traefik --no-deploy=servicelb --no-deploy=local-storage --no-deploy=metrics-server --https-listen-port=${KUBERNETES_API_PORT} &> /var/log/k3s.log &
K3S_PID=$!
timeout "$NEBULA_TIMEOUT" sh -c "until k3s kubectl get node $HOSTNAME &> /dev/null; do sleep 1; done" || ( echo >&2 "Timed out while waiting for the Kubernetes cluster to start"; exit 1 )
k3s kubectl wait node $HOSTNAME --for condition=Ready --timeout ${NEBULA_TIMEOUT}s &> /dev/null || ( echo >&2 "Timed out while waiting for the Kubernetes cluster to be ready"; exit 1 )
echo "Done."

echo "Deploying Nebula..."
NEBULA_VERSION=${NEBULA_VERSION:-latest}
if [[ $NEBULA_VERSION = "latest" ]]
then
  NEBULA_VERSION=$(curl --silent "https://api.github.com/repos/nebulaclouds/nebula/releases/latest" | jq -r .tag_name)
fi

# Deploy nebula
helm repo add nebulaclouds https://nebulaclouds.github.io/nebula

echo "Deploying Nebula-deps..."
version=""
charts="/nebulaclouds/share/nebula-deps"

if [[ $NEBULA_TEST = "release" ]]
then
  version="--version $NEBULA_VERSION"
  charts="nebulaclouds/nebula-deps"
fi

if [[ $NEBULA_TEST = "local" ]]
then
  helm dep update $charts
fi

helm upgrade -n nebula --create-namespace nebula-deps $charts --kubeconfig /etc/rancher/k3s/k3s.yaml --install $version --set webhook.enabled=false,contour.enabled=true --wait

version=""
charts="/nebulaclouds/share/nebula-core"

if [[ $NEBULA_TEST = "release" ]]
then
  version="--version $NEBULA_VERSION"
  charts="nebulaclouds/nebula-core"
fi

if [[ $NEBULA_TEST = "local" ]]
then
  helm dep update $charts
fi

helm upgrade -n nebula --create-namespace nebula-core $charts --kubeconfig /etc/rancher/k3s/k3s.yaml --install $version -f /nebulaclouds/share/nebula-core/values-sandbox.yaml

wait-for-nebula.sh

# With nebulactl sandbox --source flag, we mount the root volume to user source dir that will create helm & k8s cache specific directory.
# In Linux, These file belongs to root user that is different then current user
# In this case during fast serialization, Pynebula will through error because of permission denied
rm -rf /root/.cache /root/.kube /root/.config

# Monitor running processes. Exit when the first process exits.
monitor ${DOCKERD_PID} ${K3S_PID}
