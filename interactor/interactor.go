package interactor

import (
	"context"
)

type Request interface {
	Type() string
}

// UseCaseRunner runs a use-case and returns the corresponding result setting res value by reference.
type UseCaseRunner interface {
	RunUseCase(ctx context.Context, req Request, res interface{}) error
}

// UseCaseRunnerFn defines a use case runner signature.
type UseCaseRunnerFn func(ctx context.Context, req Request, res interface{}) error

// RunUseCase runs a use-case and returns the corresponding result setting res value by reference.
func (u UseCaseRunnerFn) RunUseCase(ctx context.Context, req Request, res Request) error {
	return u(ctx, req, res)
}
