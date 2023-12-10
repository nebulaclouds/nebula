#!/usr/bin/env bash

set -e

# Create dir structure
mkdir -p release

# Copy all deployment manifest in release directory
for file in ./deployment/**/nebula_generated.yaml; do
    if [ -f "$file" ]; then
        result=${file/#"./deployment/"}
        result=${result/%"/nebula_generated.yaml"}
        cp $file "./release/nebula_${result}_manifest.yaml"
    fi
done

grep -rlZ "version:[^P]*# VERSION" ./charts/nebula/Chart.yaml | xargs -0 sed -i "s/version:[^P]*# VERSION/version: ${VERSION} # VERSION/g"
sed "s/v0.1.10/${VERSION}/g" ./charts/nebula/README.md  > temp.txt && mv temp.txt ./charts/nebula/README.md

grep -rlZ "version:[^P]*# VERSION" ./charts/nebula-core/Chart.yaml | xargs -0 sed -i "s/version:[^P]*# VERSION/version: ${VERSION} # VERSION/g"
sed "s/v0.1.10/${VERSION}/g" ./charts/nebula-core/README.md  > temp.txt && mv temp.txt ./charts/nebula-core/README.md

grep -rlZ "version:[^P]*# VERSION" ./charts/nebula-deps/Chart.yaml | xargs -0 sed -i "s/version:[^P]*# VERSION/version: ${VERSION} # VERSION/g"
sed "s/v0.1.10/${VERSION}/g" ./charts/nebula-deps/README.md  > temp.txt && mv temp.txt ./charts/nebula-core/README.md

grep -rlZ "version:[^P]*# VERSION" ./charts/nebula-binary/Chart.yaml | xargs -0 sed -i "s/version:[^P]*# VERSION/version: ${VERSION} # VERSION/g"
sed "s/v0.1.10/${VERSION}/g" ./charts/nebula-binary/README.md > temp.txt && mv temp.txt ./charts/nebula-core/README.md

helm dep update ./charts/nebula
helm dep update ./charts/nebula-deps

# bump latest release of nebula component in helm
sed -i "s,tag:[^P]*# NEBULAADMIN_TAG,tag: ${VERSION} # NEBULAADMIN_TAG," ./charts/nebula/values.yaml
sed -i "s,tag:[^P]*# NEBULAADMIN_TAG,tag: ${VERSION} # NEBULAADMIN_TAG," ./charts/nebula-core/values.yaml
sed -i "s,repository:[^P]*# NEBULAADMIN_IMAGE,repository: cr.nebula.org/nebulaclouds/nebulaadmin-release # NEBULAADMIN_IMAGE," ./charts/nebula/values.yaml
sed -i "s,repository:[^P]*# NEBULAADMIN_IMAGE,repository: cr.nebula.org/nebulaclouds/nebulaadmin-release # NEBULAADMIN_IMAGE," ./charts/nebula-core/values.yaml

sed -i "s,tag:[^P]*# NEBULASCHEDULER_TAG,tag: ${VERSION} # NEBULASCHEDULER_TAG," ./charts/nebula/values.yaml
sed -i "s,tag:[^P]*# NEBULASCHEDULER_TAG,tag: ${VERSION} # NEBULASCHEDULER_TAG," ./charts/nebula-core/values.yaml
sed -i "s,repository:[^P]*# NEBULASCHEDULER_IMAGE,repository: cr.nebula.org/nebulaclouds/nebulascheduler-release # NEBULASCHEDULER_IMAGE," ./charts/nebula/values.yaml
sed -i "s,repository:[^P]*# NEBULASCHEDULER_IMAGE,repository: cr.nebula.org/nebulaclouds/nebulascheduler-release # NEBULASCHEDULER_IMAGE," ./charts/nebula-core/values.yaml

sed -i "s,tag:[^P]*# DATACATALOG_TAG,tag: ${VERSION} # DATACATALOG_TAG," ./charts/nebula/values.yaml
sed -i "s,tag:[^P]*# DATACATALOG_TAG,tag: ${VERSION} # DATACATALOG_TAG," ./charts/nebula-core/values.yaml
sed -i "s,repository:[^P]*# DATACATALOG_IMAGE,repository: cr.nebula.org/nebulaclouds/datacatalog-release # DATACATALOG_IMAGE," ./charts/nebula/values.yaml
sed -i "s,repository:[^P]*# DATACATALOG_IMAGE,repository: cr.nebula.org/nebulaclouds/datacatalog-release # DATACATALOG_IMAGE," ./charts/nebula-core/values.yaml

sed -i "s,tag:[^P]*# NEBULACONSOLE_TAG,tag: ${VERSION} # NEBULACONSOLE_TAG," ./charts/nebula/values.yaml
sed -i "s,tag:[^P]*# NEBULACONSOLE_TAG,tag: ${VERSION} # NEBULACONSOLE_TAG," ./charts/nebula-core/values.yaml
sed -i "s,repository:[^P]*# NEBULACONSOLE_IMAGE,repository: cr.nebula.org/nebulaclouds/nebulaconsole-release # NEBULACONSOLE_IMAGE," ./charts/nebula/values.yaml
sed -i "s,repository:[^P]*# NEBULACONSOLE_IMAGE,repository: cr.nebula.org/nebulaclouds/nebulaconsole-release # NEBULACONSOLE_IMAGE," ./charts/nebula-core/values.yaml

sed -i "s,tag:[^P]*# NEBULAPROPELLER_TAG,tag: ${VERSION} # NEBULAPROPELLER_TAG," ./charts/nebula/values.yaml
sed -i "s,tag:[^P]*# NEBULAPROPELLER_TAG,tag: ${VERSION} # NEBULAPROPELLER_TAG," ./charts/nebula-core/values.yaml
sed -i "s,repository:[^P]*# NEBULAPROPELLER_IMAGE,repository: cr.nebula.org/nebulaclouds/nebulapropeller-release # NEBULAPROPELLER_IMAGE," ./charts/nebula/values.yaml
sed -i "s,repository:[^P]*# NEBULAPROPELLER_IMAGE,repository: cr.nebula.org/nebulaclouds/nebulapropeller-release # NEBULAPROPELLER_IMAGE," ./charts/nebula-core/values.yaml

sed -i "s,image:[^P]*# NEBULACOPILOT_IMAGE,image: cr.nebula.org/nebulaclouds/nebulacopilot-release:${VERSION} # NEBULACOPILOT_IMAGE," ./charts/nebula/values.yaml
sed -i "s,image:[^P]*# NEBULACOPILOT_IMAGE,image: cr.nebula.org/nebulaclouds/nebulacopilot-release:${VERSION} # NEBULACOPILOT_IMAGE," ./charts/nebula-core/values.yaml

sed -i "s,repository:[^P]*# NEBULA_IMAGE,repository: cr.nebula.org/nebulaclouds/nebula-binary-release # NEBULA_IMAGE," ./charts/nebula-binary/values.yaml
sed -i "s,tag:[^P]*# NEBULA_TAG,tag: ${VERSION} # NEBULA_TAG," ./charts/nebula-binary/values.yaml
sed -i "s,repository:[^P]*# NEBULACOPILOT_IMAGE,repository: cr.nebula.org/nebulaclouds/nebulacopilot-release # NEBULACOPILOT_IMAGE," ./charts/nebula-binary/values.yaml
sed -i "s,tag:[^P]*# NEBULACOPILOT_TAG,tag: ${VERSION} # NEBULACOPILOT_TAG," ./charts/nebula-binary/values.yaml
