package testdsl

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/screwyprof/golibs/queryer/mock"
)

func TestDSL(t *testing.T) {
	t.Run("ensure query runner returns an error", func(t *testing.T) {
		want := errors.New("some error")
		queryRunner := &mock.ConcreteQueryerStub{Err: want}

		Test(t)(
			Given("TestQuery", queryRunner.Run),
			When(context.Background(), mock.TestQuery{}, &mock.TestReport{}),
			ThenFailWith(want),
		)
	})

	t.Run("ensure query runner returns valid result", func(t *testing.T) {
		want := &mock.TestReport{Value: 123}
		queryRunner := &mock.ConcreteQueryerStub{Rep: 123}

		Test(t)(
			Given("TestQuery", queryRunner.Run),
			When(context.Background(), mock.TestQuery{}, &mock.TestReport{}),
			Then(want),
		)
	})

	t.Run("ensure test fails if invalid concrete query runner given", func(t *testing.T) {
		invalidQueryRunner := struct{}{}

		tester := &SpyTester{T: t}
		Test(tester)(
			Given("InvalidRunner", invalidQueryRunner),
			When(context.Background(), mock.TestQuery{}, &mock.TestReport{}),
			Then(nil),
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
