#!/bin/bash

# https://github.com/zeromicro/goctl-swagger

# install swagger-markdown
# npm install -g swagger-markdown > /dev/null 2>&1

out_dir='../docs'
api_dir='../endpoint-definitions'

# define the gateway services
services='admin client'

for service in $services; do
  # make the directory if one does not exist
  mkdir -p $out_dir/swagger/$service
  mkdir -p $out_dir/endpoints/$service

  # generate a complete json description
  goctl api plugin -plugin goctl-swagger="swagger -filename ${out_dir}/swagger/${service}.json" -api "${api_dir}/${service}.api" -dir "${api_dir}" >/dev/null 2>&1
  swagger-markdown -i "${out_dir}/swagger/${service}.json" -o "${out_dir}/swagger/${service}.md" >/dev/null 2>&1

  service_api_file_content=$(cat "${api_dir}/${service}.api")
  # echo "$service_api_file_content"

  # perl -0777ne'print "$1\n" while /ModelProxy\("([\w*.]+)"\);/g' test.txt
  shared_imports=$(echo -e "$service_api_file_content" | perl -0777ne 'print "\n    ","$1" while /("shared\/.*"?)/g;')
  api_tpl="import($shared_imports\n    IMPORT_API_FILE\n)"

  for api_file in $(ls ../endpoint-definitions/$service); do
    filename=$(basename -- "$api_file")
    extension="${filename##*.}"
    filename="${filename%.*}"

    # create a tmp file with the parts required for generating the service document
    svc_tpl_path="${api_dir}/tmp/${service}"
    #
    mkdir -p $svc_tpl_path
    svc_filename="${svc_tpl_path}/${filename}.api"
    svc_tpl=$(echo "${api_tpl}" | sed "s/IMPORT_API_FILE/\"${service}\/${filename}.api\"/")
    echo -e "${svc_tpl}" >"${svc_filename}"

    goctl api plugin -plugin goctl-swagger="swagger -filename ${out_dir}/swagger/${service}/${filename}.json" -api "${svc_filename}" -dir "${api_dir}" >/dev/null 2>&1
    # swagger-markdown -i "${out_dir}/swagger/${service}/${filename}.json" -o "${out_dir}/endpoints/${service}/${filename}.md" > /dev/null 2>&1
    swagger generate markdown --quiet --spec="${out_dir}/swagger/${service}/${filename}.json" --output="${out_dir}/endpoints/${service}/${filename}.md"

    # cat "${out_dir}/endpoints/${service}/${filename}.md" >> "../../k8scommerce.github.io/content/en/docs/rest-gateway-endpoints/v1/${service}-gateway/${filename}.md"

    # set the path to the current doc file
    doc_file="../../k8scommerce.github.io/content/en/docs/rest-gateway-endpoints/v1/${service}-gateway/${filename}.md"

    # get the newly generated markdown
    gen_md=$(cat "${out_dir}/endpoints/${service}/${filename}.md")

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

# go build -o swagtest -ldflags "-s -w" . && goctl api plugin -plugin ./swagtest="swagger -filename ../k8scommerce/docs/swagger/client/cart.json" -api ../k8scommerce/endpoint-definitions/client/cart.api -dir . && swagger-markdown -i ../k8scommerce/docs/swagger/client/cart.json -o ../k8scommerce/docs/endpoints/client/cart.md
