DISTRIBUTIONS ?= "nr-otel-collector"

ci: build

build:
	@./scripts/build.sh -d "${DISTRIBUTIONS}"

generate-sources:
	@./scripts/build.sh -d "${DISTRIBUTIONS}" -s true


