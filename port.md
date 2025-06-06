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

   docker tag 936ac163fa29 harith2001/coco-core:latest
   docker tag 3a43b5fe1b16 harith2001/coco-grading:latest
   docker push harith2001/coconut-peat-supply-chain_core_system-core:latest
# kube config 
cd /mnt/c/Users/DELL/Desktop/Coconut-Peat-Supply-chain_core_system/kube-config


kubectl apply -f grading-plugin.yaml
kubectl apply -f cutting-plugin.yaml

kubectl get pods
kubectl get services

kubectl get pods -o wide
kubectl logs -f grading-plugin-54f48cfbdb-wbxmh

kubectl port-forward svc/prometheus-kube-prometheus-prometheus 9090:9090 -n default
kubectl port-forward svc/prometheus-grafana 3000:80 -n default

View the file directly
kubectl exec -it <pod-name> -n <namespace> -- cat /path/to/file

To open a terminal in the pod:
kubectl exec -it <pod-name> -n <namespace> -- sh

# kube dash board
kubectl -n kubernetes-dashboard create token admin-user
kubectl proxy
http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/workloads?namespace=default

# hpa
kubectl get hpa
kubectl top pods
kubectl top nodes 

# blockchain 
abigen --sol ShipmentTracker.sol --pkg tracking --out Tracking.go
