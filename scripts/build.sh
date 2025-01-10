#!/bin/bash

REPO_DIR="$( cd "$(dirname $( dirname "${BASH_SOURCE[0]}" ))" &> /dev/null && pwd )"
# flag for ci-only behavior (CI is auto-populated with 'true' in github actions)
ensure_docker_write_permissions=$CI

# default values
skipcompilation=false

while getopts d:s: flag
do
    case "${flag}" in
        d) distribution=${OPTARG};;
        s) skipcompilation=${OPTARG};;
    esac
done

if [[ -z $distribution ]]; then
    echo "Distribution to build not provided. Use '-d' to specify the name of the distribution to build. Ex.:"
    echo "$0 -d nr-otel-collector"
    exit 1
fi

if [[ "$skipcompilation" = true ]]; then
    echo "Skipping the compilation, we'll only generate the sources."
fi

echo "Distribution to build: $distribution";

pushd "${REPO_DIR}/distributions/${distribution}" > /dev/null
ocb_config="manifest.yaml"
output_dir=$(yq '.dist.output_path' "${ocb_config}")
echo "Output dir: $(pwd)/${output_dir}"
if [[ -d "${output_dir}" ]]; then
    # cleanup build dir as reruns of workflows seem to reuse the same filesystem
    rm -rf "${output_dir}"
fi
mkdir "${output_dir}"
if [[ "$ensure_docker_write_permissions" == "true" ]]; then
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

if [[ "$ensure_docker_write_permissions" == "true" ]]; then
    # change owner of output dir back to the 'build' user to allow access for following steps
    sudo chown -R $(id -u):$(id -g) "${output_dir}"
fi

popd > /dev/null
