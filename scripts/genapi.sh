#!/bin/bash

version='v1'
out_dir='../services/api'
api_dir='../endpoint-definitions'

# define the RPC services
services='admin client'

for service in $services; do
  goctl api go -api "${api_dir}/$version/${service}.api" -dir "${out_dir}/${service}"
done
