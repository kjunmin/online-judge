package main

import (
	"context"
	"log"

	pb "github.com/kjunmin/online-judge/server/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedProblemsServiceServer

	service ProblemsService
}

func NewGRPCHandler(grpcServer *grpc.Server, service ProblemsService) {
	handler := &grpcHandler{service: service}
	pb.RegisterProblemsServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateProblem(ctx context.Context, p *pb.CreateProblemRequest) (*pb.Problem, error) {
	err := h.service.CreateProblem(ctx, p)
	if err != nil {
		log.Printf("Error in create problem grpc handler with problem %v. err %v", p.Problem, err)
		return nil, err
	}
	log.Printf("New problem created! Problem %v", p.Problem)
	return p.Problem, nil
}
func (h *grpcHandler) GetProblemsList(ctx context.Context, r *pb.GetProblemsListRequest) (*pb.GetProblemsListResponse, error) {
	problemsList, err := h.service.GetProblemsList(ctx, r)
	if err != nil {
		log.Printf("Error in get problems list grpc handler. err %v", err)
		return nil, err
	}
	return &pb.GetProblemsListResponse{ProblemsList: problemsList}, nil
}
func (h *grpcHandler) GetProblemById(context.Context, *pb.GetProblemByIdRequest) (*pb.Problem, error) {
	o := &pb.Problem{}
	log.Println("Problem retrieved")
	return o, nil
}
