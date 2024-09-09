package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjunmin/online-judge/backend/service/runner"
)

type APIServer struct {
	addr string
}

func InitApiServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	runnerHandler := runner.NewHandler()
	runnerHandler.RegisterRoutes(subrouter)

	log.Printf("Server started listening on %s...", s.addr)

	if err := http.ListenAndServe(s.addr, router); err != nil {
		return err
	}
	return nil
}
