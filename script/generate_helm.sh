#!/usr/bin/env bash

set -ex

echo "Generating Helm"

HELM_SKIP_INSTALL=${HELM_SKIP_INSTALL:-false}

if [ "${HELM_SKIP_INSTALL}" != "true" ]; then
    curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
fi

helm version

# All the values files to be built
DEPLOYMENT_CORE=${1:-eks gcp}

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
HELM_CAPABILITIES="-a rbac.authorization.k8s.io/v1 -a networking.k8s.io/v1/Ingress -a apiextensions.k8s.io/v1/CustomResourceDefinition"

helm dep update ${DIR}/../charts/nebula-deps/
helm dep update ${DIR}/../charts/nebula-core/
helm dep update ${DIR}/../charts/nebula-binary/
helm dep update ${DIR}/../charts/nebula-sandbox/
helm dep update ${DIR}/../charts/nebula/

helm template nebula -n nebula ${DIR}/../charts/nebula/ -f ${DIR}/../charts/nebula/values.yaml ${HELM_CAPABILITIES} --debug > ${DIR}/../deployment/sandbox/nebula_helm_generated.yaml

for deployment in ${DEPLOYMENT_CORE}; do
    helm template nebula -n nebula ${DIR}/../charts/nebula-core/ -f ${DIR}/../charts/nebula-core/values.yaml -f ${DIR}/../charts/nebula-core/values-${deployment}.yaml ${HELM_CAPABILITIES} > ${DIR}/../deployment/${deployment}/nebula_helm_generated.yaml
    helm template nebula -n nebula ${DIR}/../charts/nebula-core/ -f ${DIR}/../charts/nebula-core/values.yaml -f ${DIR}/../charts/nebula-core/values-${deployment}.yaml  -f ${DIR}/../charts/nebula-core/values-controlplane.yaml ${HELM_CAPABILITIES} > ${DIR}/../deployment/${deployment}/nebula_helm_controlplane_generated.yaml
    helm template nebula -n nebula ${DIR}/../charts/nebula-core/ -f ${DIR}/../charts/nebula-core/values.yaml -f ${DIR}/../charts/nebula-core/values-${deployment}.yaml -f ${DIR}/../charts/nebula-core/values-dataplane.yaml ${HELM_CAPABILITIES} > ${DIR}/../deployment/${deployment}/nebula_helm_dataplane_generated.yaml
done

# Generate manifest AWS Scheduler
helm template nebula -n nebula ${DIR}/../charts/nebula-core/ -f ${DIR}/../charts/nebula-core/values.yaml -f ${DIR}/../charts/nebula-core/values-eks.yaml -f ${DIR}/../charts/nebula-core/values-eks-override.yaml ${HELM_CAPABILITIES} --debug > ${DIR}/../deployment/eks/nebula_aws_scheduler_helm_generated.yaml

# Generate manifest deps chart
helm template nebula -n nebula ${DIR}/../charts/nebula-deps/ ${HELM_CAPABILITIES} --debug > ${DIR}/../deployment/sandbox/nebula_sandbox_deps_helm_generated.yaml

# Generate manifest single binary chart
helm template nebula -n nebula ${DIR}/../charts/nebula-binary/ ${HELM_CAPABILITIES} --debug > ${DIR}/../deployment/sandbox-binary/nebula_sandbox_binary_helm_generated.yaml

# Generate manifest nebula agent
helm template nebula -n nebula ${DIR}/../charts/nebulaagent/ ${HELM_CAPABILITIES} --debug > ${DIR}/../deployment/agent/nebula_agent_helm_generated.yaml


echo "Generating helm docs"
if  command -v helm-docs &> /dev/null
then
    rm $(which helm-docs)
fi

GO111MODULE=on go install github.com/norwoodj/helm-docs/cmd/helm-docs@latest

${GOPATH:-~/go}/bin/helm-docs -c ${DIR}/../charts/

# This section is used by GitHub workflow to ensure that the generation step was run
if [ -n "$DELTA_CHECK" ]; then
  DIRTY=$(git status --porcelain)
  if [ -n "$DIRTY" ]; then
    echo "FAILED: helm code updated without committing generated code."
    echo "Ensure make helm has run and all changes are committed."
    DIFF=$(git diff)
    echo "diff detected: $DIFF"
    DIFF=$(git diff --name-only)
    echo "files different: $DIFF"
    exit 1
  else
    echo "SUCCESS: Generated code is up to date."
  fi
fi
