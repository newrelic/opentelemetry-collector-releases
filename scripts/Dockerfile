# TODO: file issue to include git into ocb image and use that instead
# https://github.com/open-telemetry/opentelemetry-collector-releases/blob/main/cmd/builder/Dockerfile#
FROM otel/opentelemetry-collector-builder:0.112.0

USER root
# go:alpine intentionally does not include git: https://github.com/docker-library/golang/pull/345#issuecomment-736012250
# but for a convenient build process without GOPROXY it is needed
RUN apk --update add git

# same entrypoint as base image
ENTRYPOINT [ "ocb" ]