define NEBULA_BINARY_BUILD
mkdir -p images/tar/$(1)

docker buildx build \
	--build-arg NEBULACONSOLE_VERSION=$(NEBULACONSOLE_VERSION) \
	--platform linux/$(1) \
	--tag nebula-binary:sandbox \
	--output type=docker,dest=images/tar/$(1)/nebula-binary.tar \
	../..

endef

.PHONY: nebula
nebula: NEBULACONSOLE_VERSION := latest
nebula:
	$(foreach arch,amd64 arm64,$(call NEBULA_BINARY_BUILD,$(arch)))

.PHONY: manifests
manifests:
	mkdir -p manifests
	helm dependency update ../../charts/nebula-sandbox
	helm dependency update ../../charts/nebulaagent
	kustomize build \
		--enable-helm \
		--load-restrictor=LoadRestrictionsNone \
		kustomize/complete > manifests/complete.yaml
	kustomize build \
		--enable-helm \
		--load-restrictor=LoadRestrictionsNone \
		kustomize/dev > manifests/dev.yaml
	kustomize build \
		--enable-helm \
		--load-restrictor=LoadRestrictionsNone \
		kustomize/complete-agent > manifests/complete-agent.yaml

.PHONY: build
build: nebula manifests
	[ -n "$(shell docker buildx ls | awk '/^nebula-sandbox / {print $$1}')" ] || \
		docker buildx create --name nebula-sandbox \
		--driver docker-container --driver-opt image=moby/buildkit:master \
		--buildkitd-flags '--allow-insecure-entitlement security.insecure' \
		--platform linux/arm64,linux/amd64
	docker buildx build --builder nebula-sandbox --allow security.insecure --load \
		--tag nebula-sandbox:latest .

# Port map
# 6443 - k8s API server
# 30000 - Docker Registry
# 30001 - DB
# 30002 - Minio
# 30003 - Buildkit
# 30080 - Nebula Proxy
.PHONY: start
start: NEBULA_SANDBOX_IMAGE := nebula-sandbox:latest
start: NEBULA_DEV := False
start:
	[ -n "$(shell docker volume ls --filter name=^nebula-sandbox$$ --format {{.Name}})" ] || \
		docker volume create nebula-sandbox
	[ -n "$(shell docker ps --filter name=^nebula-sandbox$$ --format {{.Names}})" ] || \
		docker run --detach --rm --privileged --name nebula-sandbox \
			--add-host "host.docker.internal:host-gateway" \
			--env NEBULA_DEV=$(NEBULA_DEV) \
			--env K3S_KUBECONFIG_OUTPUT=/.kube/kubeconfig \
			--volume $(PWD)/.kube:/.kube \
			--volume $(HOME)/.nebula/sandbox:/var/lib/nebula/config \
			--volume nebula-sandbox:/var/lib/nebula/storage \
			--publish "6443":"6443" \
			--publish "30000:30000" \
			--publish "30001:30001" \
			--publish "30002:30002" \
			--publish "30003:30003" \
			--publish "30080:30080" \
			$(NEBULA_SANDBOX_IMAGE)
		export KUBECONFIG=$(PWD)/.kube/kubeconfig
.PHONY: kubeconfig
.SILENT: kubeconfig
kubeconfig:
	sed -i -e "/server:/ s/: .*/: https:\/\/127.0.0.1:$(shell docker port nebula-sandbox | grep ^6443 | awk '{print $$3}' | awk -F: '{print $$2}')/" .kube/kubeconfig
	echo "export KUBECONFIG=$(PWD)/.kube/kubeconfig"

.PHONY: stop
stop:
	docker stop nebula-sandbox

.PHONY: console
console:
	open http://localhost:30080/console

.venv:
	python -m venv .venv
	. .venv/bin/activate && pip install nebulakit

.PHONY: test
test: export NEBULACTL_CONFIG := test/config.yaml
test: .venv
	. .venv/bin/activate && pynebula run --remote test/wf.py wf --n 500 --mean 42 --sigma 2
	. .venv/bin/activate && pynebula run --remote test/raw_container_wf.py wf
	kubectl create secret generic user-info \
		--namespace nebulasnacks-development \
		--from-literal=user_secret=sosecret \
		--dry-run=client --output yaml | kubectl apply --filename -
	. .venv/bin/activate && pynebula run --remote test/secrets_wf.py wf
