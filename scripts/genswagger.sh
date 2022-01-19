#!/bin/bash

# https://github.com/zeromicro/goctl-swagger

# get the root path of the directory this file resides
# this enables this script to be called from any path
# https://gist.github.com/olegch/1730673
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

version=v1
out_dir=$ROOT/../docs
api_dir=$ROOT/../endpoint-definitions

# define the gateway services
services='admin client'

# create a tmp directory to work in
tmp_dir=$(mktemp -d 2>/dev/null || mktemp -d -t 'k8sswagtemp')

for service in $services; do
  # make the directory if one does not exist
  mkdir -p $out_dir/swagger/$version/$service
  mkdir -p $out_dir/endpoints/$version/$service

  # create a service and shared directory in our
  # tmp directory for individual service swagger and markdown generation
  svc_tpl_path="${tmp_dir}/${service}"
  mkdir -p $svc_tpl_path
  mkdir -p "${svc_tpl_path}/shared"

  # copy the shared folder to the svc_tpl_path
  cp -R $api_dir/$version/shared/. $svc_tpl_path/shared

  # generate a complete json description
  goctl api plugin -plugin goctl-swagger="swagger -filename ${out_dir}/swagger/${version}/${service}.json" -api "${api_dir}/${version}/${service}.api" -dir "${api_dir}" >/dev/null 2>&1
  service_api_file_content=$(cat "${api_dir}/${version}/${service}.api") >/dev/null 2>&1
  shared_imports=$(echo -e "$service_api_file_content" | perl -0777ne 'print "\n    ","$1" while /("shared\/.*"?)/g;') >/dev/null 2>&1

  # create an api template to be used in the service api generation
  api_tpl="import($shared_imports\n    IMPORT_API_FILE\n)"

  # loop through each service .api file and generate a swagger file and markdown for each
  for api_file in $(ls ../endpoint-definitions/$version/$service); do

    cp -R $api_dir/$version/$service/. $svc_tpl_path/$service

    # extract the file from the path
    filename=$(basename -- "$api_file")
    extension="${filename##*.}"
    filename="${filename%.*}"

    # create an .api template for the service
    svc_filename="${svc_tpl_path}/${filename}.api"
    svc_tpl=$(echo "${api_tpl}" | sed "s/IMPORT_API_FILE/\"${service}\/${filename}.api\"/")
    echo -e "${svc_tpl}" >"${svc_filename}"

    # generate the swagger file
    goctl api plugin -plugin goctl-swagger="swagger -filename ${out_dir}/swagger/${version}/${service}/${filename}.json" -api "${svc_filename}" -dir "${api_dir}" >/dev/null 2>&1

    # generate the markdown file from the swagger file
    swagger generate markdown --quiet --spec="${out_dir}/swagger/${version}/${service}/${filename}.json" --output="${out_dir}/endpoints/${version}/${service}/${filename}.md" >/dev/null 2>&1

    # set the path to the current doc file
    doc_file="../../k8scommerce.github.io/content/en/docs/rest-gateway-endpoints/${version}/${service}-gateway/${filename}.md"

    # get the newly generated markdown
    gen_md=$(cat "${out_dir}/endpoints/${version}/${service}/${filename}.md")

    # remove everything before 'All endpoints'
    cleaned_gen_md=${gen_md#*\#\#\ All\ endpoints}

    # create a title for the page
    title=$(echo "${filename} Endpoints" | awk '{for(i=1;i<=NF;i++){ $i=toupper(substr($i,1,1)) substr($i,2) }}1')

    # combine the title and cleaned gen markdown
    clean_md=$(printf "\n## ${title}${cleaned_gen_md}")

    # get the doc file that currently exists
    doc_md=$(cat "${doc_file}")

    # extract the title from the current doc page
    title=$(echo -e "$gen_md" | perl -0777 -ne '/[#]{1}\s+(.*?)\n/ && print $1;')

    # extract the header from the current doc page
    header=$(echo -e "${doc_md}" | perl -0777 -ne '/[-]{3}(.*?)[-]{3}/s && print "---",$1,"---";')

    # update the header with the documentation title
    header=$(echo -e "$header" | perl -0777 -ne '/(.*)description: (.*?)[-]{3}/s && print $1, "description: >", "\n  ", "NEW_TITLE", "\n---","\n";')

    # write the header
    header=$(echo "${header}" | sed "s/NEW_TITLE/${title}/")

    # replace the current documenation file with the new header
    echo "${header}" >"${doc_file}"

    # append the cleaned, new markdown
    echo "${clean_md}" >>"${doc_file}"

  done
done

# remove the tmp directory
rm -rf $tmp_dir
