#!/bin/bash

REPO_DIR="$( cd "$(dirname $( dirname "${BASH_SOURCE[0]}" ))" &> /dev/null && pwd )"

# default values
skipcompilation=false

while getopts d:s: flag
do
    case "${flag}" in
        d) distributions=${OPTARG};;
        s) skipcompilation=${OPTARG};;
    esac
done

if [[ -z $distributions ]]; then
    echo "List of distributions to build not provided. Use '-d' to specify the names of the distributions to build. Ex.:"
    echo "$0 -d nr-otel-collector"
    exit 1
fi

if [[ "$skipcompilation" = true ]]; then
    echo "Skipping the compilation, we'll only generate the sources."
fi

echo "Distributions to build: $distributions";

for distribution in $(echo "$distributions" | tr "," "\n")
do
    pushd "${REPO_DIR}/distributions/${distribution}" > /dev/null
    ocb_config="manifest.yaml"
    relative_output_dir=$(yq '.dist.output_path' "${ocb_config}")
    echo "Output dir: $(pwd)/${relative_output_dir}"
    mkdir -p ${relative_output_dir}

    # TODO: submit PR to include git in ocb image (also delete custom Dockerfile from repo root)
#    builder_version=$(yq '.dist.otelcol_version' "${ocb_config}")
#    builder_image="otel/opentelemetry-collector-builder:${builder_version}"
#    docker pull "${builder_image}"
    builder_image="ocb-with-git:1"
    docker build --tag "${builder_image}" . -f "${REPO_DIR}/scripts/Dockerfile"

    # TODO: submit PR to update ocb docs which suggest to use '/build' instead
    #  - docs: https://github.com/open-telemetry/opentelemetry-collector/tree/main/cmd/builder#official-release-docker-image
    #  - workdir defined in Dockerfile: https://github.com/open-telemetry/opentelemetry-collector-releases/blob/main/cmd/builder/Dockerfile#L11
    #  - default is a /tmp dir: https://github.com/open-telemetry/opentelemetry-collector/blob/main/cmd/builder/internal/builder/config.go#L97
    container_work_dir=$(docker image inspect -f '{{.Config.WorkingDir}}' ${builder_image})
    container_path_ocb_config="${container_work_dir}/${ocb_config}"

    echo "Building: $distribution"
    docker run \
      -v "$(pwd)/${ocb_config}:${container_path_ocb_config}" \
      -v "$(pwd)/${relative_output_dir}:${container_work_dir}/${relative_output_dir}" \
      "${builder_image}" \
      --config "${container_path_ocb_config}" \
      --skip-compilation=${skipcompilation}

    popd > /dev/null
done
