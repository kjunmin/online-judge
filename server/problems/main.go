package main

import (
	"fmt"
	"log"
	"net"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kjunmin/online-judge/server/common"
	"google.golang.org/grpc"
)

func main() {
	cfg := common.GetConfig()
	grpcAddr := fmt.Sprintf(":%s", cfg.GRPCPort)
	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	defer l.Close()

	store := NewStore()
	service := NewService(store)
	NewGRPCHandler(grpcServer, service)

	log.Println("GRPC Server started at ", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
