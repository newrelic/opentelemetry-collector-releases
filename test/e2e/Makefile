REQUIRED_BINS := go kind docker helm
$(foreach bin,$(REQUIRED_BINS),\
    $(if $(shell command -v $(bin) 2> /dev/null),$(info Found `$(bin)`),$(error Please install `$(bin)`)))

KIND_CLUSTER_NAME=e2etest
K8S_CONTEXT_NAME=kind-${KIND_CLUSTER_NAME}
IMAGE_REPO=newrelic/nr-otel-collector
#TODO: determine programatically
# with tls
IMAGE_TAG=0.0.999-SNAPSHOT-c2e5ff6-rc-amd64


.PHONY=create-cluster
create-cluster-if-not-exists:
	kind get clusters | grep -q ${KIND_CLUSTER_NAME} || kind create cluster --name ${KIND_CLUSTER_NAME} --config kind-config.yaml

.PHONY=build-image
build-image:
	GORELEASER_CURRENT_TAG=0.0.999 goreleaser --snapshot --clean --skip-sign --timeout 2h

.PHONY=load-image
load-image: create-cluster-if-not-exists
	kind load docker-image ${IMAGE_REPO}:${IMAGE_TAG} --name ${KIND_CLUSTER_NAME}

.PHONY=test
test: load-image
	E2E_TEST__K8S_CONTEXT_NAME=${K8S_CONTEXT_NAME} \
	E2E_TEST__IMAGE_REPO=${IMAGE_REPO} \
	E2E_TEST__IMAGE_TAG=${IMAGE_TAG} \
	go test ./... -count=1 -v
# TODO: make this more generic
cleanup:
	helm uninstall --kube-context ${K8S_CONTEXT_NAME} -n hostmetrics collector-startup