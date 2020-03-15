package queryer_test

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v4"

	"github.com/screwyprof/golibs/assert"
	"github.com/screwyprof/golibs/queryer/mock"
)

func TestQueryer(t *testing.T) {
	t.Run("Valid query provided, valid report returned", func(t *testing.T) {
		// arrange
		ID := gofakeit.Number(1, 100)
		want := mock.TestReport{Value: ID}

		concreteQueryer := mock.ConcreteQueryerStub{Rep: ID}

		// act
		var r mock.TestReport
		err := concreteQueryer.Run(context.Background(), mock.TestQuery{ID: ID}, &r)

		// assert
		assert.NoError(t, err)
		assert.Equals(t, want, r)
	})

	t.Run("Invalid query provided, an error returned", func(t *testing.T) {
		// arrange
		concreteQueryer := mock.ConcreteQueryerStub{Err: errors.New("an error")}

		// act
		var r mock.TestReport
		err := concreteQueryer.Run(context.Background(), mock.TestQuery{}, &r)

		// assert
		assert.NotNil(t, err)
	})
}
