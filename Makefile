BINARY_NAME=grpc_using_go

PROTOC_GEN_GO=$(GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC=$(GOPATH)/bin/protoc-gen-go-grpc

$(PROTOC_GEN_GO):
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

$(PROTOC_GEN_GO_GRPC):
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

pb.go: book.proto | $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC)
	protoc book.proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:.

compile: pb.go

build: compile
	go get
	go build -o $(BINARY_NAME)

run: compile
	go get
	go run main.go

run-docker:
	DOCKER_BUILDKIT=1 docker build -t $(BINARY_NAME) .
	docker run --rm -p 9000:9000 $(BINARY_NAME) || true

clean:
	rm -f $(BINARY_NAME)
	rm -f book/book.pb.go book/book_grpc.pb.go
