#!/bin/bash

# get the root path of the directory this file resides
# this enables this script to be called from any path
# https://gist.github.com/olegch/1730673
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

out_dir=$ROOT/../services/rpc
protos_dir=$ROOT/../protos

# define the RPC services
services='cart catalog customer email inventory othersbought payment shipping similarproducts store user warehouse'

# define the proto import paths
go_opts=''
for service in $services; do
  go_opts="${go_opts}--go_opt=M${service}.proto=k8scommerce/services/rpc/${service}/pb/${service} "
done

for service in $services; do
  # adding the $go_ops created above enables the proto imports paths
  # are properly rewritten to their correct destinations
  goctl rpc proto \
    -proto_path=$protos_dir \
    -src "${protos_dir}/${service}.proto" \
    -dir "${out_dir}/${service}" \
    $go_opts \
    "${service}.proto"



  # $go_opts
  # goctl rpc protoc \
  #   "${protos_dir}/${service}.proto" \
  #   $go_opts \
  #   --proto_path=$protos_dir \
  #   --go_out=plugins=grpc:"${out_dir}/${service}" \
  #   --zrpc_out="${out_dir}/${service}"
done
