package main

import (
	"context"

	pb "github.com/kjunmin/online-judge/server/common/api"
)

type ProblemsStore interface {
	Create(context.Context) error
}

type ProblemsService interface {
	CreateProblem(context.Context, *pb.CreateProblemRequest) error
	GetProblem(context.Context, *pb.GetProblemRequest) error
}
