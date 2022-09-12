package main

import (
	"net"
	"os"

	"grpc_using_go/book"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	port, ok := os.LookupEnv("PORT")
	if ok {
		log.WithFields(log.Fields{
			"PORT": port,
		}).Info("PORT env var defined")
	} else {
		port = "9000"
		log.WithFields(log.Fields{
			"PORT": port,
		}).Warn("PORT env var NOT defined. Going with default")
	}

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Failed to listen")
	}

	s := book.Server{}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	book.RegisterBookServiceServer(grpcServer, &s)

	log.Info("gRPC server started at ", port)
	if err := grpcServer.Serve(l); err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Failed to serve")
	}
}
