#!/usr/bin/env bash


# So that when a command fails, bash exits.
set -o errexit

# This will make the script fail, when accessing an unset variable.
set -o nounset

# the return value of a pipeline is the value of the last (rightmost) command
# to exit with a non-zero status, or zero if all commands in the pipeline exit
# successfully.
set -o pipefail

# This helps in debugging your scripts. TRACE=1 ./script.sh
if [[ "${TRACE-0}" == "1" ]]; then set -o xtrace; fi



if [[ "${1-}" =~ ^-*h(elp)?$ ]]; then
    echo 'Usage: ./check_distro_names.sh
Checks that defined distributions have unique names.
'
    exit
fi

# check that any distribution directory has the upstream collectors name
check_distributions_dirs() {
    base_distributions=("otelcol-contrib" "otelcol")

    DISTRIBUTIONS=$(ls ./distributions)

    for distribution in ${DISTRIBUTIONS}
    do
        if [[ " ${base_distributions[*]} " =~ " ${distribution} " ]]; then
            echo "[ERROR]: Distribution already exisits, change your distribution name"
            exit 1
        fi
    done
}

check_distributions_mainfests() {
    distributions_name=("otelcol-contrib" "otelcol")

    MANIFESTS=$(find . -type f -name "manifest.yaml" | sort)

    for manifest_file in ${MANIFESTS}
    do
        distribution=$(awk '/name:/{print $2}' ${manifest_file})
        if [[ " ${distributions_name[*]} " =~ " ${distribution} " ]]; then
            echo "[ERROR]: Distribution already exisits on a manifest file, change your distribution name"
            exit 1
        fi
        distributions_name+=($distribution)
    done

}

main() {
    check_distributions_dirs
    check_distributions_mainfests
    echo "[SUCCEED]: All distributions have unique names"
}

main "$@"
