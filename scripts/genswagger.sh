#!/bin/bash

# https://github.com/zeromicro/goctl-swagger

# install swagger-markdown
npm install -g swagger-markdown > /dev/null 2>&1

out_dir='../docs'
api_dir='../endpoint-definitions'

# define the gateway services
services='admin client'

for service in $services; do
  # make the directory if one does not exist
  mkdir -p $out_dir/swagger/$service
  mkdir -p $out_dir/endpoints/$service

  goctl api plugin -plugin goctl-swagger="swagger -filename ${out_dir}/swagger/${service}.json" -api "${api_dir}/${service}.api" -dir "${api_dir}" > /dev/null 2>&1
  swagger-markdown -i "${out_dir}/swagger/${service}.json" -o "${out_dir}/swagger/${service}.md" > /dev/null 2>&1

  for api_file in $(ls ../endpoint-definitions/$service); do
    filename=$(basename -- "$api_file")
    extension="${filename##*.}"
    filename="${filename%.*}"
    
    goctl api plugin -plugin goctl-swagger="swagger -filename ${out_dir}/swagger/${service}/${filename}.json" -api "${api_dir}/${service}/${api_file}" -dir "${api_dir}"> /dev/null 2>&1
    swagger-markdown -i "${out_dir}/swagger/${service}/${filename}.json" -o "${out_dir}/endpoints/${service}/${filename}.md" > /dev/null 2>&1
  done
done
