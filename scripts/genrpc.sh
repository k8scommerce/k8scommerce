out_dir='../services/rpc'
protos_dir='../protos'

services='cart customer email inventory othersbought payment product shipping similarproducts store user warehouse'

go_opts=''
for service in $services; do
  go_opts="${go_opts}-go_opt=M${service}.proto=github.com/k8s-commerce/k8s-commerce/services/rpc/${service}/pb/${service} "
done

for service in $services; do
  goctl rpc proto \
    -proto_path=$protos_dir \
    -src "${protos_dir}/${service}.proto" \
    -dir "${out_dir}/${service}" \
    $go_opts \
    "${service}.proto"
done