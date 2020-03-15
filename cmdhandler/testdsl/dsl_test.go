package testdsl

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/screwyprof/golibs/cmdhandler/mock"
)

func TestDSL(t *testing.T) {
	t.Run("ensure command handler returns an error", func(t *testing.T) {
		want := errors.New("some error")
		commandHandler := &mock.ConcreteCommandHandlerStub{Err: want}

		Test(t)(
			Given("TestCommand", commandHandler.Handle),
			When(context.Background(), mock.TestCommand{}),
			ThenFailWith(want),
		)
	})

	t.Run("ensure query runner returns valid result", func(t *testing.T) {
		commandHandler := &mock.ConcreteCommandHandlerStub{}

		Test(t)(
			Given("TestCommand", commandHandler.Handle),
			When(context.Background(), mock.TestCommand{}),
			ThenOk(),
		)
	})

	t.Run("ensure test fails if invalid concrete command handler given", func(t *testing.T) {
		invalidQueryRunner := struct{}{}

		tester := &SpyTester{T: t}
		Test(tester)(
			Given("InvalidRunner", invalidQueryRunner),
			When(context.Background(), mock.TestCommand{}),
			ThenOk(),
		)

		if !tester.wasCalled {
			t.Fatalf("\033[31minvalid query runner given: %v\033[39m\n\n", tester.err)
		}
	})
}

type SpyTester struct {
	*testing.T
	wasCalled bool
	err       string
}

func (s *SpyTester) Fatalf(format string, args ...interface{}) {
	s.wasCalled = true
	s.err = fmt.Sprintf(format, args...)
}
