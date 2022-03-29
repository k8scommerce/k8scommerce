#!/bin/bash

# get the root path of the directory this file resides
# this enables this script to be called from any path
# https://gist.github.com/olegch/1730673
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

base=$ROOT/../../
out_dir=$ROOT/../services/rpc
protos_dir=$ROOT/../protos
# $ROOT/
# $ROOT/
# define the RPC services
services='cart catalog customer inventory othersbought payment shipping similarproducts store user warehouse'

# define the proto import paths
go_opts=''
for service in $services; do
  go_opts="${go_opts}--go_opt=M${service}.proto=github.com/k8scommerce/k8scommerce/services/rpc/${service}/pb/${service} "
done

for service in $services; do
  goctl rpc protoc \
    "${protos_dir}/${service}.proto" \
    $go_opts \
    --go_opt paths=source_relative \
    --proto_path=$protos_dir \
    --go_out=$out_dir/$service/pb/$service \
    --go-grpc_out=$out_dir/$service \
    --zrpc_out=$out_dir/$service
done
