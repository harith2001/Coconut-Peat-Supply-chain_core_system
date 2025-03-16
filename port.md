# Plugins and Port Table

This table lists services and their corresponding ports.

| Service Name | Port Number |
|--------------|-------------|
| core         | 50051       |
| grading      | 50052       |
| cutting      | 50053       |

# proto command
protoc --go_out=. --go-grpc_out=. proto/plugin.proto

# docker command
docker build -t grading_plugin -f grading.dockerfile .
docker run -p 50052:50052 grading_plugin
docker images
docker ps

# run main
docker run -p 50051:50051 `
>>   -v /var/run/docker.sock:/var/run/docker.sock `
>>   -v "C:/Program Files/Docker/Docker/resources/bin/docker.exe:/usr/bin/docker" `
>>   main:latest
>> 

# docker build all
docker-compose up --build

# create a network
docker network create my_network
docker run --network my_network

# docker push 
docker login
docker images 
docker tag --- 
docker push
coconut-peat-supply-chain_core_system-core                 latest              
   2a114460ad36   9 seconds ago    1.07GB
# kube config 
cd /mnt/c/Users/DELL/Desktop/Coconut-Peat-Supply-chain_core_system/kube-config

kubectl apply -f core-system.yaml
kubectl apply -f grading-plugin.yaml
kubectl apply -f cutting-plugin.yaml

kubectl get pods
kubectl get services

kubectl get pods -o wide
kubectl logs <pod-name>

