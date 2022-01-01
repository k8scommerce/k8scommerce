out_dir='../services/rpc'
protos_dir='../protos'

# define the RPC services
services='cart customer email inventory othersbought payment product shipping similarproducts store user warehouse'

# define the proto import paths
go_opts=''
for service in $services; do
  go_opts="${go_opts}-go_opt=M${service}.proto=github.com/k8scommerce/k8scommerce/services/rpc/${service}/pb/${service} "
done

for service in $services; do
  # create the directory if it doesn't exist
  if [ ! -d "${out_dir}/${service}" ]; then
    mkdir "${out_dir}/${service}"
  fi

  # if no go.mod file exists create one with the correct module name
  # this ensures when the goctl rpc proto command builds the microservice
  # with the correct import paths
  if [ ! -f "${out_dir}/${service}/go.mod" ]; then
    touch "${out_dir}/${service}/go.mod"
    echo "module github.com/k8scommerce/k8scommerce/services/rpc/client\n\ngo 1.17" >"${out_dir}/${service}/go.mod"
  fi

  # adding the $go_ops created above enables the proto imports paths
  # are properly rewritten to their correct destinations
  goctl rpc proto \
    -proto_path=$protos_dir \
    -src "${protos_dir}/${service}.proto" \
    -dir "${out_dir}/${service}" \
    $go_opts \
    "${service}.proto"
done
