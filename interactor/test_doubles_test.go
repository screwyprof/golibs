package interactor_test

import "context"

type TestRequest struct {
	id int
}

type TestResponse struct {
	result int
}

type ConcreteInteractorStub struct {
	res int
	err error
}

func (i ConcreteInteractorStub) RunUseCase(ctx context.Context, req TestRequest, res *TestResponse) error {
	res.result = i.res
	return i.err
}

type GeneralInteractorSpy struct {
	wasCalled bool
}

func (s *GeneralInteractorSpy) RunUseCase(ctx context.Context, req interface{}, res interface{}) error {
	s.wasCalled = true
	return nil
}
