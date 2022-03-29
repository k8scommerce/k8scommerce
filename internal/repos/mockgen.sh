#!/bin/bash

# ensure mockgen executable exists
if [ -z $(command -v mockgen) ]; then
    echo "mockgen is required - exiting"
    exit 1
fi

# get the root path of the directory this file resides
# https://gist.github.com/olegch/1730673
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

dest="${ROOT}/mocks/repos.go" # destination path & filename
src="github.com/k8scommerce/k8scommerce/internal/repos" # package path
iface="Repo" # the interface name to be mocked

echo "Generating mock for ${iface}"
mockgen -destination=$dest $src $iface
echo "Finished generating mocks"

exit 0
