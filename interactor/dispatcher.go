package interactor

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
)

// ErrNotFound indicates that the use case runner was not registered for the given request.
var ErrNotFound = errors.New("use case runner not found")

// Dispatcher dispatches the given request to a pre-registered use case runner.
type Dispatcher struct {
	runners   map[string]UseCaseRunnerFn
	runnersMu sync.RWMutex
}

// NewDispatcher creates a new instance of Dispatcher.
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		runners: make(map[string]UseCaseRunnerFn),
	}
}

// RunUseCase runs a use-case and returns the corresponding result setting res value by ref.
// Implements UseCaseRunner interface.
func (d *Dispatcher) RunUseCase(ctx context.Context, req interface{}, res interface{}) error {
	reqType := reflect.TypeOf(req).Name()
	runner, ok := d.runners[reqType]
	if !ok {
		return fmt.Errorf("interactor for %s request is not found: %w", reqType, ErrNotFound)
	}
	return runner(ctx, req, res)
}

// RegisterUseCaseRunner registers a use case runner.
// TODO: add guards for an empty runner name and a nil runner.
func (d *Dispatcher) RegisterUseCaseRunner(useCaseRunnerFnName string, useCaseRunnerFn UseCaseRunnerFn) {
	d.runnersMu.Lock()
	defer d.runnersMu.Unlock()

	d.runners[useCaseRunnerFnName] = useCaseRunnerFn
}
