KIND_CLUSTER_NAME ?=e2etest
K8S_CONTEXT_NAME=kind-${KIND_CLUSTER_NAME}
IMAGE_REPO=newrelic/nr-otel-collector
THIS_MAKEFILE_DIR := $(realpath $(dir $(realpath $(lastword $(MAKEFILE_LIST)))))
ROOT_DIR := $(realpath $(THIS_MAKEFILE_DIR)/../..)


################
#### ASSERT ####
################
REQUIRED_BINS := go kind docker helm kubectl awk
$(foreach bin,$(REQUIRED_BINS),\
	$(if $(shell command -v $(bin) 2> /dev/null),$(info Found required dependency `$(bin)`),$(error Please install `$(bin)`)))

.PHONY: assert_image-tag-present
assert_image-tag-present:
	@[ "${IMAGE_TAG}" ] || ( echo ">> env var IMAGE_TAG is not set"; exit 1 )

.PHONY: assert_cluster-exists
assert_cluster-exists:
	kind get clusters | grep -q ${KIND_CLUSTER_NAME} || ( echo ">> cluster ${KIND_CLUSTER_NAME} does not exist"; exit 1 )

################
###### CI ######
################
.PHONY: ci_load-image
ci_load-image: assert_cluster-exists assert_image-tag-present
	kind load docker-image ${IMAGE_REPO}:${IMAGE_TAG} --name ${KIND_CLUSTER_NAME}

.PHONY: ci_build-load-mocked-otlp-image
ci_build-load-mocked-otlp-image: BUILD_PATH=${THIS_MAKEFILE_DIR}/mocked_otlp/http
ci_build-load-mocked-otlp-image: assert_cluster-exists
	docker build -t mocked_otlp:latest -f ${BUILD_PATH}/Dockerfile ${BUILD_PATH}
	kind load docker-image mocked_otlp:latest --name ${KIND_CLUSTER_NAME}

.PHONY: ci_test-fast
ci_test-fast: ci_load-image ci_build-load-mocked-otlp-image
	cd ${THIS_MAKEFILE_DIR} && \
	E2E_TEST__K8S_CONTEXT_NAME=${K8S_CONTEXT_NAME} \
	E2E_TEST__IMAGE_TAG=${IMAGE_TAG} \
	go test ./... -count=1 -short

.PHONY: ci_test-slow
ci_test-slow: ci_load-image ci_build-load-mocked-otlp-image
	cd ${THIS_MAKEFILE_DIR} && \
	E2E_TEST__K8S_CONTEXT_NAME=${K8S_CONTEXT_NAME} \
	E2E_TEST__IMAGE_TAG=${IMAGE_TAG} \
	E2E_TEST__NR_BACKEND_URL=${NR_BACKEND_URL} \
	E2E_TEST__NR_INGEST_KEY=${NR_INGEST_KEY} \
	E2E_TEST__NR_API_KEY=${NR_API_KEY} \
	E2E_TEST__NR_ACCOUNT_ID=${NR_ACCOUNT_ID} \
	E2E_TEST__NR_API_BASE_URL=${NR_API_BASE_URL} \
	go test ./... -count=1

################
#### LOCAL #####
################
.PHONY: local_create-cluster-if-not-exists
local_create-cluster-if-not-exists:
	kind get clusters | grep -q ${KIND_CLUSTER_NAME} \
	|| kind create cluster --name ${KIND_CLUSTER_NAME}

.PHONY: local_build-image
local_build-image:
	cd $(ROOT_DIR) && goreleaser --snapshot --clean --skip-sign

.PHONY: local_test-fast
# hardcode an image you build with local_build-image, then run this target to run tests against it
local_test: IMAGE_TAG=insert_locally_available_tag_to_test_here
local_test: local_create-cluster-if-not-exists ci_test-fast

.PHONY: local_test-slow
local_test-slow: IMAGE_TAG=insert_locally_available_tag_to_test_here
local_test-slow: local_create-cluster-if-not-exists ci_test-slow

.PHONY: local_helm-cleanup
local_helm-cleanup:
	kubectl config use-context ${K8S_CONTEXT_NAME} && \
	helm list -A --no-headers | awk 'NF {system("helm uninstall -n " $$2 " " $$1)}'
