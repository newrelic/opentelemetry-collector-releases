#!/bin/bash

# This script will create a new manifest from an existing -rc.
# During the pre-release process we tag all the images using -rc suffix which is not required for the released images.

set -e

print_usage() {
  printf -- "Usage: %s\n" "$(basename "${0}")"
  printf -- "-i: Docker image name\n"
  printf -- "-v: Version of the release\n"
  printf -- "-h: Print help page\n"
}

while getopts 'i:v:ph' flag
do
    case "${flag}" in
        h)
          print_usage
          exit 0
        ;;
        i)
         image_name="${OPTARG}"
         continue
        ;;
        v)
         version="${OPTARG}"
         continue
        ;;
        *)
          print_usage
          exit 1
        ;;
    esac
done

# Get the list of docker images contained by the rc manifest.
images=$(docker manifest inspect "${image_name}":"${version}"-rc | jq --arg image "${image_name}" '.manifests[] | ($image+"@"+.digest)' | tr -d \")

printf "Images:\n%s\n" "${images}"

# Create and push a two new manifests latest and versioned without -rc suffix.
docker manifest create "${image_name}:latest" $images
docker manifest create "${image_name}:${version}" $images

docker manifest push "${image_name}:latest"
docker manifest push "${image_name}:${version}"
