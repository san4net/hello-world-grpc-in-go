# hello-world-grpc-in-go
sample hello world grpc implementation in Go lang

## setting up grpc
- protoc compiler installation
- https://grpc.io/docs/protoc-installation/
- Go plugins for the protocol compiler:
  Install the protocol compiler plugins for Go using the following commands:
  - `$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
  - `$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`

## generate server and client stub using protoc command
 use below command to generate client & server stub
- protoc -I . --go_out . --go_opt paths=import --go-grpc_out . --go-grpc_opt paths=import ./proto/*.proto
 file will be generated under folder gen/go/hello

## update go mod with below dependencies
- `go get google.golang.org/grpc`
- `go get google.golang.org/protobuf`

## server
 we need to implement grpc server by embedding struct UnimplementedHelloServiceServer 
 reason `UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.`

## client
 we first create hello.HelloServiceClient with url as "localhost:8080"
 Thereafter call hello.HelloServiceClient.SayHello and get reply from server




