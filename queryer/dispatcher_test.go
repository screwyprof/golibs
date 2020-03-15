package queryer_test

import (
	"context"
	"errors"
	"testing"

	"github.com/screwyprof/golibs/assert"
	"github.com/screwyprof/golibs/queryer"
	"github.com/screwyprof/golibs/queryer/mock"
)

func TestDispatcher(t *testing.T) {
	t.Run("dispatcher implements QueryRunner interface", func(t *testing.T) {
		dispatcher := queryer.NewDispatcher()
		var _ queryer.QueryRunner = dispatcher
	})

	t.Run("when query not found, an error returned", func(t *testing.T) {
		// arrange
		dispatcher := queryer.NewDispatcher()

		// act
		var res mock.TestReport
		err := dispatcher.RunQuery(context.Background(), mock.TestQuery{}, &res)

		// assert
		assertQueryRunnerNotFound(t, err)
	})

	t.Run("when query registered, it is being run", func(t *testing.T) {
		// arrange
		queryRunner := &mock.GenericQueryerSpy{}

		dispatcher := queryer.NewDispatcher()
		dispatcher.RegisterQueryRunner("TestQuery", queryRunner.RunQuery)

		// act
		var res mock.TestReport
		err := dispatcher.RunQuery(context.Background(), mock.TestQuery{}, &res)

		// assert
		assertQueryWasRunSuccessfully(t, err, queryRunner)
	})
}

func assertQueryWasRunSuccessfully(t *testing.T, err error, queryRunner *mock.GenericQueryerSpy) {
	assert.NoError(t, err)
	assert.True(t, queryRunner.WasCalled)
}

func assertQueryRunnerNotFound(t *testing.T, err error) {
	t.Helper()

	assert.True(t, errors.Is(err, queryer.ErrNotFound))
}
