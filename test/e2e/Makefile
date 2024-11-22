KIND_CLUSTER_NAME=e2etest
K8S_CONTEXT_NAME=kind-${KIND_CLUSTER_NAME}
IMAGE_REPO=newrelic/nr-otel-collector

################
#### ASSERT ####
################
REQUIRED_BINS := go kind docker helm kubectl awk
$(foreach bin,$(REQUIRED_BINS),\
	$(if $(shell command -v $(bin) 2> /dev/null),$(info Found `$(bin)`),$(error Please install `$(bin)`)))

.PHONY: assert_image-tag-present
assert_image-tag-present:
	@[ "${IMAGE_TAG}" ] || ( echo ">> env var IMAGE_TAG is not set"; exit 1 )

################
###### CI ######
################
.PHONY=ci_create-cluster-if-not-exists
ci_create-cluster-if-not-exists:
	kind get clusters | grep -q ${KIND_CLUSTER_NAME} || kind create cluster --name ${KIND_CLUSTER_NAME} --config kind-config.yaml

.PHONY=ci_load-image
ci_load-image: ci_create-cluster-if-not-exists assert_image-tag-present
	kind load docker-image ${IMAGE_REPO}:${IMAGE_TAG} --name ${KIND_CLUSTER_NAME}

.PHONY=ci_test
ci_test: ci_load-image
	E2E_TEST__K8S_CONTEXT_NAME=${K8S_CONTEXT_NAME} \
	E2E_TEST__IMAGE_REPO=${IMAGE_REPO} \
	E2E_TEST__IMAGE_TAG=${IMAGE_TAG} \
	go test ./... -count=1

################
#### LOCAL #####
################
.PHONY=local_build-image
local_build-image:
	GORELEASER_CURRENT_TAG=0.0.999 goreleaser --snapshot --clean --skip-sign

.PHONY=local_test
# hardcode an image you build with local_build-image, then run this target to run tests against it
local_test: IMAGE_TAG=insert_locally_available_tag_to_test_here
local_test: ci_test

.PHONY=local_helm-cleanup
local_helm-cleanup:
	kubectl config use-context ${K8S_CONTEXT_NAME} && \
	helm list -A --no-headers | awk 'NF {system("helm uninstall -n " $2 " " $1)}'