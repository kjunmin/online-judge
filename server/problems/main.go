package main

import (
	"fmt"
	"log"
	"net"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kjunmin/online-judge/server/common"
	"google.golang.org/grpc"
)

var (
	GRPCPort = common.EnvString("GRPC_PORT", "4000")
)

func main() {
	grpcAddr := fmt.Sprintf(":%s", GRPCPort)
	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	defer l.Close()

	store := NewStore("Problems")
	service := NewService(store)
	NewGRPCHandler(grpcServer, service)

	log.Println("GRPC Server started at ", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
