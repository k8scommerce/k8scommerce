#!/bin/bash

# get the root path of the directory this file resides
# this enables this script to be called from any path
# https://gist.github.com/olegch/1730673
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

version=v1
out_dir=$ROOT/../services/api
api_dir=$ROOT/../endpoint-definitions

# define the RPC services
services='admin client'

for service in $services; do
  goctl api go -api "${api_dir}/$version/${service}.api" -dir "${out_dir}/${service}"
done
