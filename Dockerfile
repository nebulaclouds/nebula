ARG NEBULACONSOLE_VERSION=latest
FROM ghcr.io/nebulaclouds/nebulaconsole:${NEBULACONSOLE_VERSION} AS nebulaconsole


FROM --platform=${BUILDPLATFORM} golang:1.19.13-bookworm AS nebulabuilder

ARG TARGETARCH
ENV GOARCH "${TARGETARCH}"
ENV GOOS linux

WORKDIR /nebulaclouds/build

COPY datacatalog datacatalog
COPY nebulaadmin nebulaadmin
COPY nebulacopilot nebulacopilot
COPY nebulaidl nebulaidl
COPY nebulaplugins nebulaplugins
COPY nebulapropeller nebulapropeller
COPY nebulastdlib nebulastdlib

COPY go.mod go.sum ./
RUN go mod download
COPY cmd cmd
COPY --from=nebulaconsole /app/ cmd/single/dist
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,target=/root/go/pkg/mod \
    go build -tags console -v -o dist/nebula cmd/main.go


FROM debian:bookworm-slim

ARG NEBULA_VERSION
ENV NEBULA_VERSION "${NEBULA_VERSION}"

ENV DEBCONF_NONINTERACTIVE_SEEN true
ENV DEBIAN_FRONTEND noninteractive

# Install core packages
RUN apt-get update && apt-get install --no-install-recommends --yes \
        ca-certificates \
        tini \
    && rm -rf /var/lib/apt/lists/*

# Copy compiled executable into image
COPY --from=nebulabuilder /nebulaclouds/build/dist/nebula /usr/local/bin/

# Set entrypoint
ENTRYPOINT [ "/usr/bin/tini", "-g", "--", "/usr/local/bin/nebula" ]
