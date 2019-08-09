# Install protoc compiler from apt repository
### sudo apt install protobuf-compiler

# Install python packages to genrate python packages from .proto file
### pip install grpcio grpcio-tools

# Install python packages to genrate python packages from .proto file
### go get -u google.golang.org/grpc
### go get -u github.com/golang/protobuf/proto
### go get -u github.com/golang/protobuf/protoc-gen-go

# Genrate go package from .proto file with grpc extension 
### protoc --go_out=plugins=grpc:. models/calculator.proto

# Genrate python package from .proto file with grpc package support 
### python -m  grpc_tools.protoc -I. --python_out=. --grpc_python_out=. ./models/calculator.proto
