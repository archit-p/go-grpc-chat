.PHONY: help clean build

help:
	@echo "use \`make <target> where <target> is:\'"
	@echo "    help: display this help"
	@echo "    clean: clean the build directory"
	@echo "    proto: generate protobuf code"
	@echo "    build: create the executables"

clean:
	rm -f chat/chat.pb.go 1>/dev/null 2>&1
	rm -f server 1>/dev/null 2>&1
	rm -f client 1>/dev/null 2>&1

proto:
	go get github.com/golang/protobuf/protoc-gen-go
	protoc --go_out=plugins=grpc:. chat/chat.proto

server:
	go build -o server server.go

client:
	go build -o client client.go

build: proto server client
