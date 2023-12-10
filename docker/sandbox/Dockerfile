FROM alpine:3.13.5 AS base

# Install dependencies
RUN apk add --no-cache openssl

# Make directory to store artifacts
RUN mkdir -p /nebulaclouds/bin /nebulaclouds/share

# Install k3s
ARG K3S_VERSION="v1.21.1%2Bk3s1"
ARG TARGETARCH

RUN case $TARGETARCH in \
    amd64) export SUFFIX=;; \
    arm64) export SUFFIX=-arm64;; \
    aarch64)  export SUFFIX=-arm64;; \
    # TODO: Check if we need to add case fail
    esac; \
    wget -q -O /nebulaclouds/bin/k3s https://github.com/k3s-io/k3s/releases/download/${K3S_VERSION}/k3s${SUFFIX} \
    && chmod +x /nebulaclouds/bin/k3s

# Install Helm
ARG HELM_VERSION="v3.6.3"

RUN wget -q -O /nebulaclouds/bin/get_helm.sh https://raw.githubusercontent.com/helm/helm/${HELM_VERSION}/scripts/get-helm-3 && \
    chmod 700 /nebulaclouds/bin/get_helm.sh && \
    sh /nebulaclouds/bin/get_helm.sh --version ${HELM_VERSION} && \
    mv /usr/local/bin/helm /nebulaclouds/bin/helm && \
    rm /nebulaclouds/bin/get_helm.sh


# Install nebulactl
RUN wget -q -O - https://raw.githubusercontent.com/nebulaclouds/nebulactl/master/install.sh | BINDIR=/nebulaclouds/bin sh -s

# Copy nebula chart
COPY charts/nebula/ /nebulaclouds/share/nebula
COPY charts/nebula-core/ /nebulaclouds/share/nebula-core
COPY charts/nebula-deps/ /nebulaclouds/share/nebula-deps

# Copy scripts
COPY docker/sandbox/kubectl docker/sandbox/cgroup-v2-hack.sh docker/sandbox/wait-for-nebula.sh /nebulaclouds/bin/

FROM docker:20.10.14-dind-alpine3.15 AS dind

# Install dependencies
RUN apk add --no-cache bash git make tini curl jq

# Copy artifacts from base
COPY --from=base /nebulaclouds/ /nebulaclouds/

# Copy entrypoints
COPY docker/sandbox/nebula-entrypoint-default.sh /nebulaclouds/bin/nebula-entrypoint.sh

ARG NEBULA_VERSION="latest"
ENV NEBULA_VERSION "${NEBULA_VERSION}"

ARG NEBULA_TEST="release"
ENV NEBULA_TEST "${NEBULA_TEST}"

# Update PATH variable
ENV PATH "/nebulaclouds/bin:${PATH}"

# Declare volumes for k3s
VOLUME /var/lib/kubelet
VOLUME /var/lib/rancher/k3s
VOLUME /var/lib/cni
VOLUME /var/log

# Expose Nebula ports
EXPOSE 30081 30082 30084 30088


ENTRYPOINT ["tini", "nebula-entrypoint.sh"]
