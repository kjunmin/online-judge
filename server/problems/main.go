package main

import (
	"context"
	"log"
	"net"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kjunmin/online-judge/server/common"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:3000")
)

func main() {
	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	defer l.Close()

	store := NewStore()
	service := NewService(store)
	NewGRPCHandler(grpcServer)

	service.CreateProblem(context.Background())

	log.Println("GRPC Server started at ", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
