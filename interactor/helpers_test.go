package interactor_test

import (
	"context"
	"errors"
)

var errSomeErr = errors.New("some error")

type TestRequest struct {
	id int
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

func (s *GeneralUseCaseSpy) RunUseCase(ctx context.Context, req interface{}, res interface{}) error {
	s.wasCalled = true

	return nil
}
