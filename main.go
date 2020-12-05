package main

import (
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	var port string
	var ok bool
	port, ok = os.LookupEnv("PORT")
	if ok {
		log.WithFields(log.Fields{
			"PORT": port,
		}).Info("PORT env var defined")

	} else {
		port = "9000"
		log.WithFields(log.Fields{
			"PORT": port,
		}).Warn("PORT env var not defined. Going with default")

	}

	_, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Failed to listen")
	}
}
