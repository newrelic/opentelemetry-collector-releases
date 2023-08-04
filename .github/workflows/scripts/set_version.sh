#!/bin/bash

set -e

# fetch the history (including tags) from within a shallow clone like CI-GHA
# supress error when the repository is a complete one.
git fetch --prune --unshallow 2> /dev/null || true

tag=$(git describe --tags --abbrev=0)

prev_tag_commit=$(git rev-list --tags --skip=1 --max-count=1)
prev_tag=$(git describe --tags --abbrev=0 "${prev_tag_commit}")

# Expected tag format <distro>-<version> e.g. distro_name-major.minor.patch
regex="^(.*)-([0-9]+\.[0-9]+\.[0-9]+)$"

if [[ "${tag}" =~ ${regex} ]]
then
    distro="${BASH_REMATCH[1]}"
    version="${BASH_REMATCH[2]}"
else
    printf "Bad tag format: %s doesn't match expected pattern 'distro_name-major.minor.patch'\n" "${tag}" >&2
    exit 1
fi

if [[ "${prev_tag}" =~ ${regex} ]]
then
    prev_distro="${BASH_REMATCH[1]}"
    prev_version="${BASH_REMATCH[2]}"
else
    printf "Bad tag format: %s doesn't match expected pattern 'distro_name-major.minor.patch'\n" "${prev_tag}" >&2
    exit 1
fi

printf "Distribution name: %s, Version name: %s\n" "${distro}" "${version}"
printf "Previous distribution name: %s, Previous version name: %s\n" "${prev_distro}" "${prev_version}"

# Set the variables for later use in the GHA pipeline
{
    echo "NR_DISTRO=${distro}"
    echo "NR_VERSION=${version}"
    echo "NR_RELEASE_TAG=${tag}"

    echo "PREVIOUS_NR_DISTRO=${prev_distro}"
    echo "PREVIOUS_NR_VERSION=${prev_version}"
    echo "PREVIOUS_NR_RELEASE_TAG=${prev_tag}"
} >> "$GITHUB_ENV"

# Assert manifest distro and version
manifest_file="./distributions/${distro}/manifest.yaml"

if [ ! -f "${manifest_file}" ]; then
    printf "Manifest file for the distribution: '%s' extracted from the tag: '%s' wasn't found in %s\n" "${distro}" "${tag}" "${manifest_file}" >&2
    exit 1
fi

# #TODO: Instead of asserting we could replace it in manifest file to avoid manual steps.
manifest_version=$(yq .dist.version "${manifest_file}")

if [ "${manifest_version}" != "${version}" ]; then
    printf "Wrong manifest version: expected '%s' but was %s\n" "${version}" "${manifest_version}" >&2
    exit 1
fi

manifest_distro=$(yq .dist.name "${manifest_file}")

if [ "${manifest_distro}" != "${distro}" ]; then
    printf "Wrong manifest version: expected '%s' but was %s\n" "${distro}" "${manifest_distro}" >&2
    exit 1
fi

# Rename the tag locally to have a semantic versioning format
# which is required by the packaging step.
git tag "${version}" "${tag}"
