package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	pb "github.com/kjunmin/online-judge/server/common/api"
	"google.golang.org/protobuf/proto"
)

type service struct {
	store ProblemsStore
}

func NewService(store ProblemsStore) *service {
	ctx := context.Background()
	if exists, _ := store.TableExists(ctx); !exists {
		store.CreateProblemsTable(ctx)
	}
	return &service{store: store}
}

func (s *service) CreateProblem(ctx context.Context, p *pb.CreateProblemRequest) error {
	problemWithId := pb.Problem{
		ProblemID:   proto.String(uuid.NewString()),
		Title:       p.Problem.Title,
		Tags:        p.Problem.Tags,
		Description: p.Problem.Description,
	}

	err := s.store.AddProblem(ctx, &problemWithId)
	if err != nil {
		log.Printf("Error creating problem in service %v. Error %v\n", problemWithId, err)
	}
	return err
}

func (s *service) GetProblemsList(ctx context.Context, r *pb.GetProblemsListRequest) ([]*pb.ProblemsListItem, error) {
	problemsList, err := s.store.GetProblemsList(ctx)
	if err != nil {
		log.Printf("Error getting problems list in service. Error %v\n", err)
		return nil, err
	}
	return problemsList, nil
}

func (s *service) GetProblemById(context.Context, *pb.GetProblemByIdRequest) error {
	return nil
}
