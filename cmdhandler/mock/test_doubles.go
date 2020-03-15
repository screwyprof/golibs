package mock

import "context"

type TestCommand struct {
	ID string
}

type ConcreteCommandHandlerStub struct {
	Err error
}

func (s *ConcreteCommandHandlerStub) Handle(ctx context.Context, command TestCommand) error {
	return s.Err
}

type ConcreteCommandHandlerSpy struct {
	WasCalled bool
}

func (s *ConcreteCommandHandlerSpy) Handle(ctx context.Context, command interface{}) error {
	s.WasCalled = true
	return nil
}

type GenericCommandHandlerSpy struct {
	WasCalled bool
}

func (s *GenericCommandHandlerSpy) Handle(ctx context.Context, command interface{}) error {
	s.WasCalled = true
	return nil
}
