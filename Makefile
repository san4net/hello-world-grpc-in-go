BINARY_NAME=grpc-commons

projectID=grpc-commons-prj

all: clean build

clean:
	go clean
	rm -f bin/${BINARY_NAME}

build:
	protoc -I . --go_out . --go_opt paths=import --go-grpc_out . --go-grpc_opt paths=import ./proto/*.proto