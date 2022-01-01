. ./swarm.down.sh

docker swarm init --advertise-addr 127.0.0.1
docker service create --name registry --publish published=5000,target=5000 registry:2

sleep 5

http http://localhost:5000/v2/


# build and push the images
make -C ./db docker

make -C ./services/rpc/product build
make -C ./services/rpc/product docker

make -C ./services/api/client build
make -C ./services/api/client docker

docker compose push --ignore-push-failures


docker stack deploy --compose-file docker-compose.swarm.yaml ecomm
docker service ls 

# docker-compose build --no-cache