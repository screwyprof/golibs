package interactor_test

import (
	"context"
	"errors"

	"github.com/screwyprof/golibs/interactor"
)

var errSomeErr = errors.New("some error")

type TestRequest struct {
	id int
}

func (r TestRequest) Type() string {
	return "TestRequest"
}

type TestResponse struct {
	result int
}

type ConcreteUseCase struct {
	res int
	err error
}

func (i ConcreteUseCase) RunUseCase(ctx context.Context, req TestRequest, res *TestResponse) error {
	res.result = i.res

	return i.err
}

type GeneralUseCaseSpy struct {
	wasCalled bool
}

func (s *GeneralUseCaseSpy) Run(ctx context.Context, req interactor.Request, res interactor.Response) error {
	s.wasCalled = true

	return nil
}
