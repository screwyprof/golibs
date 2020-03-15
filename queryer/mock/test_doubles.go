package mock

import "context"

type TestQuery struct {
	ID int
}

type TestReport struct {
	Value int
}

type ConcreteQueryerStub struct {
	Rep int
	Err error
}

func (s *ConcreteQueryerStub) Run(ctx context.Context, q TestQuery, r *TestReport) error {
	r.Value = s.Rep
	return s.Err
}

type GenericQueryerSpy struct {
	WasCalled bool
}

func (s *GenericQueryerSpy) RunQuery(ctx context.Context, q interface{}, r interface{}) error {
	s.WasCalled = true
	return nil
}
