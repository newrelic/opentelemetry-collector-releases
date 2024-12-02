#!/bin/bash

set -e

# fetch the history (including tags) from within a shallow clone like CI-GHA
# suppress error when the repository is a complete one.
git fetch --prune --unshallow 2> /dev/null || true

tag=$(git describe --tags --abbrev=0)

prev_tag_commit=$(git rev-list --tags --skip=1 --max-count=1)
prev_tag_git=$(git describe --tags --abbrev=0 "${prev_tag_commit}")
prev_tag=${PREVIOUS_TAG:-${prev_tag_git}}

# Expected tag format following semantic versioning conventions, i.e major.minor.patch
regex="^([0-9]+\.[0-9]+\.[0-9]+)$"

if [[ "${tag}" =~ ${regex} ]]
then
    version="${BASH_REMATCH[1]}"
else
    printf "Bad tag format: %s isn't following semantic versioning convention: 'major.minor.patch'\n" "${tag}" >&2
    exit 1
fi

if [[ "${prev_tag}" =~ ${regex} ]]
then
    prev_version="${BASH_REMATCH[1]}"
else
    printf "Bad tag format: %s isn't following semantic versioning convention: 'major.minor.patch'\n" "${prev_tag}" >&2
    exit 1
fi

printf "Version: %s\n" "${version}"
printf "Previous version: %s\n" "${prev_version}"

# Set the variables for later use in the GHA pipeline
{
    echo "NR_VERSION=${version}"
    echo "NR_RELEASE_TAG=${tag}"

    echo "PREVIOUS_NR_VERSION=${prev_version}"
    echo "PREVIOUS_NR_RELEASE_TAG=${prev_tag}"
} >> "$GITHUB_ENV"

check_manifest_versions() {
    expected_version="${version}"
    MANIFESTS=$(find ./distributions -type f -name "manifest.yaml" | sort)
    echo "Expect version ${expected_version} for all detected manifests:\n${MANIFESTS}"

    for manifest_file in ${MANIFESTS}
    do
        version_in_manifest=$(awk '/[^_]+version:/{print $2}' ${manifest_file})
        if [[ "${expected_version}" != "${version_in_manifest}" ]]; then
            echo "Invalid version in manifest: Expected ${expected_version}, but got ${version_in_manifest} in '${manifest_file}'" >&2
            exit 1
        fi
    done
}
check_manifest_versions