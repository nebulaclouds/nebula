#!/usr/bin/env bash

set -ex

NEBULAKIT_TAG=$(curl --silent "https://api.github.com/repos/nebulaclouds/nebulakit/releases/latest" | jq -r .tag_name | sed 's/^v//')
NEBULACONSOLE_TAG=$(curl --silent "https://api.github.com/repos/nebulaclouds/nebulaconsole/releases/latest" | jq -r .tag_name)

# bump latest release of nebula component in kustomize
grep -rlZ "newTag:[^P]*# NEBULAADMIN_TAG" ./kustomize/overlays | xargs -I {} sed -i "s/newTag:[^P]*# NEBULAADMIN_TAG/newTag: ${VERSION} # NEBULAADMIN_TAG/g" {}
grep -rlZ "newTag:[^P]*# DATACATALOG_TAG" ./kustomize/overlays | xargs -I {} sed -i "s/newTag:[^P]*# DATACATALOG_TAG/newTag: ${VERSION} # DATACATALOG_TAG/g" {}
grep -rlZ "newTag:[^P]*# NEBULACONSOLE_TAG" ./kustomize/overlays | xargs -I {} sed -i "s/newTag:[^P]*# NEBULACONSOLE_TAG/newTag: ${NEBULACONSOLE_TAG} # NEBULACONSOLE_TAG/g" {}
grep -rlZ "newTag:[^P]*# NEBULAPROPELLER_TAG" ./kustomize/overlays | xargs -I {} sed -i "s/newTag:[^P]*# NEBULAPROPELLER_TAG/newTag: ${VERSION} # NEBULAPROPELLER_TAG/g" {}

# bump latest release of nebula component in helm
sed -i "s,tag:[^P]*# NEBULAADMIN_TAG,tag: ${VERSION} # NEBULAADMIN_TAG," ./charts/nebula/values.yaml
sed -i "s,tag:[^P]*# NEBULAADMIN_TAG,tag: ${VERSION} # NEBULAADMIN_TAG," ./charts/nebula-core/values.yaml

sed -i "s,tag:[^P]*# NEBULASCHEDULER_TAG,tag: ${VERSION} # NEBULASCHEDULER_TAG," ./charts/nebula/values.yaml
sed -i "s,tag:[^P]*# NEBULASCHEDULER_TAG,tag: ${VERSION} # NEBULASCHEDULER_TAG," ./charts/nebula-core/values.yaml

sed -i "s,tag:[^P]*# DATACATALOG_TAG,tag: ${VERSION} # DATACATALOG_TAG," ./charts/nebula/values.yaml
sed -i "s,tag:[^P]*# DATACATALOG_TAG,tag: ${VERSION} # DATACATALOG_TAG," ./charts/nebula-core/values.yaml

sed -i "s,tag:[^P]*# NEBULACONSOLE_TAG,tag: ${NEBULACONSOLE_TAG} # NEBULACONSOLE_TAG," ./charts/nebula/values.yaml
sed -i "s,tag:[^P]*# NEBULACONSOLE_TAG,tag: ${NEBULACONSOLE_TAG} # NEBULACONSOLE_TAG," ./charts/nebula-core/values.yaml

sed -i "s,tag:[^P]*# NEBULAPROPELLER_TAG,tag: ${VERSION} # NEBULAPROPELLER_TAG," ./charts/nebula/values.yaml
sed -i "s,tag:[^P]*# NEBULAPROPELLER_TAG,tag: ${VERSION} # NEBULAPROPELLER_TAG," ./charts/nebula-core/values.yaml

sed -i "s,image:[^P]*# NEBULACOPILOT_IMAGE,image: cr.nebula.org/nebulaclouds/nebulacopilot:${VERSION} # NEBULACOPILOT_IMAGE," ./charts/nebula/values.yaml
sed -i "s,image:[^P]*# NEBULACOPILOT_IMAGE,image: cr.nebula.org/nebulaclouds/nebulacopilot:${VERSION} # NEBULACOPILOT_IMAGE," ./charts/nebula-core/values.yaml
sed -i "s,tag:[^P]*# NEBULACOPILOT_TAG,tag: ${VERSION} # NEBULACOPILOT_TAG," ./charts/nebula-binary/values.yaml

sed -i "s,tag:[^P]*# NEBULAAGENT_TAG,tag: ${NEBULAKIT_TAG} # NEBULAAGENT_TAG," ./charts/nebulaagent/values.yaml
