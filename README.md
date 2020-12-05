[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity)
[![Ask Me Anything !](https://img.shields.io/badge/Ask%20me-anything-1abc9c.svg)](https://GitHub.com/Naereen/ama)
[![made-for-VSCode](https://img.shields.io/badge/Made%20for-VSCode-1f425f.svg)](https://code.visualstudio.com/)
[![GitHub forks](https://img.shields.io/github/forks/saswatamcode/grpc_using_go?)](https://GitHub.com/saswatamcode/grpc_using_go/network/)
[![GitHub stars](https://img.shields.io/github/stars/saswatamcode/grpc_using_go?)](https://GitHub.com/saswatamcode/grpc_using_go/stargazers/)
[![GitHub issues](https://img.shields.io/github/issues/saswatamcode/grpc_using_go.svg)](https://GitHub.com/saswatamcode/grpc_using_go/issues/)
[![Open Source Love svg1](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)

# gRPC Using Go
A simple demo of a gRPC server using Go.

## To Run
- Clone into repo
- Run `go get`
- Run `go run main.go`


## Test RPC using gRPCurl
- Start gRPC server
- Run `grpcurl -plaintext localhost:9000 describe` to describe all services and methods
- Run `grpcurl -d '{"name": "To Kill a Mockingbird", "isbn": 12345}' -plaintext localhost:9000 book.BookService/GetBook`