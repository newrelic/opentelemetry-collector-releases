DISTRIBUTIONS ?= "nr-otel-collector"

ci: build

build:
	@./scripts/build.sh -d "${DISTRIBUTIONS}" -c true

build-local:
	@./scripts/build.sh -d "${DISTRIBUTIONS}" -c false

generate-sources:
	@./scripts/build.sh -d "${DISTRIBUTIONS}" -s true -c true


