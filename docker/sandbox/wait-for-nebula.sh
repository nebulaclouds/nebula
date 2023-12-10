#!/bin/sh

set -e

echo "Waiting for Nebula to become ready..."

NEBULA_TIMEOUT=${NEBULA_TIMEOUT:-600}

# Ensure cluster is up and running. We don't need a timeout here, since the container
# itself will exit with the appropriate error message if the kubernetes cluster is not
# up within the specified timeout.
until k3s kubectl explain deployment &> /dev/null; do sleep 1; done

# Wait for Nebula namespace to be created. This is necessary for the next step.
timeout $NEBULA_TIMEOUT sh -c "until k3s kubectl get namespace nebula &> /dev/null; do sleep 1; done"  || ( echo >&2 "Timed out while waiting for the Nebula namespace to be created"; exit 1 )

# Wait for Nebula deployment to be created. This is necessary for the next step.
timeout $NEBULA_TIMEOUT sh -c "until k3s kubectl rollout status deployment datacatalog -n nebula  &> /dev/null; do sleep 1; done"  || ( echo >&2 "Timed out while waiting for the datacatalog rollout to be created"; exit 1 )
timeout $NEBULA_TIMEOUT sh -c "until k3s kubectl rollout status deployment nebulaadmin -n nebula  &> /dev/null; do sleep 1; done"  || ( echo >&2 "Timed out while waiting for the nebulaadmin rollout to be created"; exit 1 )
timeout $NEBULA_TIMEOUT sh -c "until k3s kubectl rollout status deployment nebulaconsole -n nebula  &> /dev/null; do sleep 1; done"  || ( echo >&2 "Timed out while waiting for the nebulaconsole rollout to be created"; exit 1 )
timeout $NEBULA_TIMEOUT sh -c "until k3s kubectl rollout status deployment nebulapropeller -n nebula  &> /dev/null; do sleep 1; done"  || ( echo >&2 "Timed out while waiting for the nebulapropeller rollout to be created"; exit 1 )
timeout $NEBULA_TIMEOUT sh -c "until k3s kubectl rollout status deployment nebula-deps-contour-contour -n nebula  &> /dev/null; do sleep 1; done"  || ( echo >&2 "Timed out while waiting for the nebulapropeller rollout to be created"; exit 1 )

# Wait for nebula deployment
k3s kubectl wait --for=condition=available deployment/datacatalog deployment/nebulaadmin deployment/nebulaconsole deployment/nebulapropeller deployment/nebula-deps-contour-contour -n nebula --timeout=10m || ( echo >&2 "Timed out while waiting for the Nebula deployment to start"; exit 1 )

echo "Nebula is ready! Nebula UI is available at http://localhost:30081/console."
