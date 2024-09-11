package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjunmin/online-judge/server/common"
	pb "github.com/kjunmin/online-judge/server/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	GRPCPort = "3000"
	HTTPPort = "8080"
)

func main() {
	cfg := common.GetConfig()
	problemsServiceAddr := fmt.Sprintf(":%s", cfg.GRPCPort)
	httpAddr := fmt.Sprintf(":%s", cfg.HTTPPort)

	conn, err := grpc.NewClient(problemsServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	log.Printf("Dialing problems service at %v", problemsServiceAddr)

	c := pb.NewProblemsServiceClient(conn)

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	handler := NewHandler(c)
	handler.registerRoutes(subrouter)

	fmt.Printf("Started and listening on %s\n", httpAddr)
	if err := http.ListenAndServe(httpAddr, router); err != nil {
		panic(err)
	}
}
