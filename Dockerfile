# syntax=docker/dockerfile:1

##
## Build
##

FROM golang:1.19-bullseye AS build

WORKDIR /app

COPY . ./
RUN go mod download
RUN go build -o /grpc_using_go

##
## Deploy
##

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /grpc_using_go /grpc_using_go
EXPOSE 9000

USER nonroot:nonroot
ENTRYPOINT [ "/grpc_using_go" ]
