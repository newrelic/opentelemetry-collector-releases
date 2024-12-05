#!/bin/bash

REPO_DIR="$( cd "$(dirname $( dirname "${BASH_SOURCE[0]}" ))" &> /dev/null && pwd )"

# default values
skipcompilation=false

while getopts d:s:c: flag
do
    case "${flag}" in
        d) distributions=${OPTARG};;
        s) skipcompilation=${OPTARG};;
        c) ci=${OPTARG};;
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
    output_dir=$(yq '.dist.output_path' "${ocb_config}")
    echo "Output dir: $(pwd)/${output_dir}"
    mkdir "${output_dir}"
    if [[ "$ci" == "true" ]]; then
        # ci is running into permission issues when trying to write to output dir on host
        # ocb dockerfile user/group id is 10001 (https://github.com/open-telemetry/opentelemetry-collector-releases/blob/main/cmd/builder/Dockerfile#L6)
        sudo chown 10001:10001 "${output_dir}"
    fi

    builder_version=$(yq '.dist.otelcol_version' "${ocb_config}")
    builder_image="otel/opentelemetry-collector-builder:${builder_version}"
    docker pull "${builder_image}"

    container_work_dir=$(docker image inspect -f '{{.Config.WorkingDir}}' ${builder_image})
    container_path_ocb_config="${container_work_dir}/${ocb_config}"

    echo "Building: $distribution"
    docker run \
      -v "$(pwd)/${ocb_config}:${container_path_ocb_config}" \
      -v "$(pwd)/${output_dir}:${container_work_dir}/${output_dir}:rw" \
      "${builder_image}" \
      --config "${container_path_ocb_config}" \
      --skip-compilation=${skipcompilation}

    popd > /dev/null
done
