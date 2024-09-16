package main

import (
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

// func (h *grpcHandler) CreateProblem(ctx context.Context, p *pb.CreateProblemRequest) (*pb.Problem, error) {
// 	o := &pb.Problem{}
// 	log.Printf("New problem created! Problem %v", p)
// 	return o, nil
// }
// func (h *grpcHandler) GetProblem(context.Context, *pb.GetProblemRequest) (*pb.Problem, error) {
// 	o := &pb.Problem{}
// 	log.Println("Problem retrieved")
// 	return o, nil
// }
