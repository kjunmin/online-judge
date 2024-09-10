package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kjunmin/online-judge/server/common"
	pb "github.com/kjunmin/online-judge/server/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr            = common.EnvString("HTTP_ADDR", ":8080")
	problemsServiceAddr = "localhost:3000"
)

func main() {
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
