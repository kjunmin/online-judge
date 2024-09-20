package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	pb "github.com/kjunmin/online-judge/server/common/api"
)

type ProblemsStore interface {
	CreateProblemsTable(context.Context) (*types.TableDescription, error)
	TableExists(context.Context) (bool, error)
	AddProblem(context.Context, *pb.Problem) error
}

type ProblemsService interface {
	CreateProblem(context.Context, *pb.CreateProblemRequest) error
	GetProblem(context.Context, *pb.GetProblemRequest) error
}
