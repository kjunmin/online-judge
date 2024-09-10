package main

import "context"

type ProblemsStore interface {
}

type ProblemsService interface {
	CreateProblem(context.Context) error
}
