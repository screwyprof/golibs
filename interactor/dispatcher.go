package interactor

import (
	"context"
	"errors"
	"fmt"
)

// ErrNotFound indicates that the use case runner was not registered for the given request.
var ErrNotFound = errors.New("use case runner not found")

// Dispatcher dispatches the given request to a pre-registered use case runner.
type Dispatcher struct {
	runners map[string]UseCaseRunnerFn
}

// NewDispatcher creates a new instance of Dispatcher.
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		runners: make(map[string]UseCaseRunnerFn),
	}
}

// RunUseCase runs a use-case and returns the corresponding result setting res value by ref.
// Implements UseCaseRunner interface.
func (d *Dispatcher) RunUseCase(ctx context.Context, req Request, res interface{}) error {
	reqType := req.Type()

	runner, ok := d.runners[reqType]
	if !ok {
		return fmt.Errorf("%w: request type: %s", ErrNotFound, reqType)
	}

	return runner(ctx, req, res)
}

// RegisterUseCaseRunner registers a use case runner.
func (d *Dispatcher) RegisterUseCaseRunner(useCaseRunnerFnName string, useCaseRunnerFn UseCaseRunnerFn) {
	d.runners[useCaseRunnerFnName] = useCaseRunnerFn
}
