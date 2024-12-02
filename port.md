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
