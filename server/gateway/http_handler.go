package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjunmin/online-judge/server/common"
	pb "github.com/kjunmin/online-judge/server/common/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	client pb.ProblemsServiceClient
}

func NewHandler(client pb.ProblemsServiceClient) *handler {
	return &handler{client: client}
}

func (h *handler) registerRoutes(mux *mux.Router) {
	mux.HandleFunc("/problem/create", h.HandleCreateProblem).Methods("POST")
	mux.HandleFunc("/problem/{problem_id:[0-9]+}", h.HandleGetProblem).Methods("GET")
	mux.HandleFunc("/healthcheck", h.HandleHealthCheck).Methods("GET")
}

func (h *handler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	common.WriteJSONToHTTPResponse(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *handler) HandleGetProblem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	problemId := vars["problem_id"]
	log.Printf("Problem id %v", problemId)

	common.WriteJSONToHTTPResponse(w, http.StatusOK, map[string]string{"Problem ID": problemId})
}

func (h *handler) HandleCreateProblem(w http.ResponseWriter, r *http.Request) {
	var problem *pb.Problem
	if err := common.ReadJSONFromHTTPRequest(r, &problem); err != nil {
		common.WriteErrorToHTTPResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateProblem(problem); err != nil {
		common.WriteErrorToHTTPResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	problem, err := h.client.CreateProblem(r.Context(), &pb.CreateProblemRequest{
		Problem: problem,
	})
	rStatus := status.Convert(err)
	if rStatus != nil {
		if rStatus.Code() != codes.InvalidArgument {
			common.WriteErrorToHTTPResponse(w, http.StatusBadRequest, rStatus.Message())
			return
		}
		common.WriteErrorToHTTPResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.WriteJSONToHTTPResponse(w, http.StatusOK, problem)
}

func validateProblem(problem *pb.Problem) error {
	if problem.Title == "" {
		return errors.New("problem must have a title")
	}

	if problem.Description == "" {
		return errors.New("problem must have a description")
	}

	return nil
}
