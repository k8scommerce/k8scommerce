exit 1

out_dir='../services/rpc'
api_dir='../protos'

# define the RPC services
services='admin client'

for service in $services; do
  goctl api go -api admin.api 
    -dir .




  goctl api go \
    -api=$api_dir \
    -src "${api_dir}/${service}.proto" \
    -dir "${out_dir}/${service}" \
    $go_opts \
    "${service}.proto"
done
