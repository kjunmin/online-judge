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
func (h *grpcHandler) GetProblem(context.Context, *pb.GetProblemRequest) (*pb.Problem, error) {
	o := &pb.Problem{}
	log.Println("Problem retrieved")
	return o, nil
}
