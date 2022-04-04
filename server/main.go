package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/abdulmoeid7112/inventory-backend/config"
	"github.com/abdulmoeid7112/inventory-backend/db"
	"github.com/abdulmoeid7112/inventory-backend/db/firestore"
	"github.com/abdulmoeid7112/inventory-backend/db/mongo"
	pb "github.com/abdulmoeid7112/inventory-backend/pb"
	"github.com/abdulmoeid7112/inventory-backend/server/services"
)

func main() {
	var store db.DataStore

	listen, err := net.Listen("tcp", viper.GetString(config.ServerHost)+":"+viper.GetString(config.ServerPort))
	if err != nil {
		panic(fmt.Sprintf("Error opening tcp port \"%v\": \"%v\".\n", viper.GetString(config.ServerPort), err))
	}

	done := make(chan bool)

	grpcServer := grpc.NewServer()

	if viper.GetString(config.DataFactory) == "mongo" {
		store, err = mongo.NewMongoClient(db.Option{})
		if err != nil {
			panic(fmt.Sprintf("failed to create mongo client: %+v", err))
		}
	} else {
		store, err = firestore.NewFirestoreClient(db.Option{})
		if err != nil {
			panic(fmt.Sprintf("failed to create firestore client: %+v", err))
		}
	}

	pb.RegisterInventoryServiceServer(grpcServer, services.NewService(store))

	log().Infof("Server started on %v:%v", viper.GetString(config.ServerHost), viper.GetString(config.ServerPort))

	go gracefulShutdown(context.TODO(), grpcServer, done)

	grpcServer.Serve(listen)

	<-done
	log().Info("Server stopped gracefully")
}

func gracefulShutdown(_ context.Context, server *grpc.Server, done chan<- bool) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTSTP, syscall.SIGTERM)

	<-quit

	log().Info("server is shutting down...")

	server.Stop()

	close(done)
}

func log() *logger.Entry {
	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp: true,
	})

	return logger.WithFields(logger.Fields{
		"package": "server",
	})
}
