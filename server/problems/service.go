package main

import (
	"context"

	pb "github.com/kjunmin/online-judge/server/common/api"
)

type service struct {
	store ProblemsStore
}

func NewService(store ProblemsStore) *service {
	return &service{store: store}
}

func (s *service) CreateProblem(context.Context, *pb.CreateProblemRequest) error {
	return nil
}

func (s *service) GetProblem(context.Context, *pb.GetProblemRequest) error {
	return nil
}
