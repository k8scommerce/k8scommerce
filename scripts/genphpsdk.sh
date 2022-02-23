#!/bin/bash

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

version=v1
out_dir=$(cd $ROOT/../.. && pwd)
json_dir=$(cd $ROOT/../docs/swagger/$version && pwd)
config_dir=$(cd $ROOT/.. && pwd)
flavor=php

openapi-generator generate -i ${json_dir}/admin.json -g $flavor -o $out_dir/php_sdk -c $config_dir/open-generator-admin-config.yaml --enable-post-process-file

code $out_dir/php_sdk

# cd $out_dir/php_sdk
# flutter pub get && flutter pub run build_runner build --delete-conflicting-outputs
# cd $ROOT