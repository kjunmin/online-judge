package main

import "context"

type service struct {
	store ProblemsStore
}

func NewService(store ProblemsStore) *service {
	return &service{store: store}
}

func (s *service) CreateProblem(context.Context) error {
	return nil
}
