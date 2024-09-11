package main

import (
	"errors"
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
	mux.HandleFunc("/create/problem", h.HandleCreateProblem)
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
