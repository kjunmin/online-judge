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
	GetProblemsList(context.Context) ([]*pb.ProblemsListItem, error)
	GetProblemById(context.Context, string) error
}

type ProblemsService interface {
	CreateProblem(context.Context, *pb.CreateProblemRequest) error
	GetProblemsList(context.Context, *pb.GetProblemsListRequest) ([]*pb.ProblemsListItem, error)
	GetProblemById(context.Context, *pb.GetProblemByIdRequest) error
}
