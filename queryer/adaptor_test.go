package queryer_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/screwyprof/golibs/assert"
	"github.com/screwyprof/golibs/queryer"
	"github.com/screwyprof/golibs/queryer/mock"
)

func TestAdaptor(t *testing.T) {
	t.Run("given query runner must be a function", func(t *testing.T) {
		// act
		_, err := queryer.Adapt(struct{}{})

		// assert
		assertQueryRunnerIsAFunction(t, err)
	})

	t.Run("given query runner must have 3 arguments", func(t *testing.T) {
		// arrange
		invalidRunner := func(ctx context.Context, req mock.TestQuery) error {
			return nil
		}
		// act
		_, err := queryer.Adapt(invalidRunner)

		// assert
		assertQueryRunnerHasInvalidSignature(t, err)
	})

	t.Run("first input param must be context.Context", func(t *testing.T) {
		// arrange
		invalidRunner := func(ctx struct{}, req mock.TestQuery, resp *mock.TestReport) error {
			return nil
		}
		// act
		_, err := queryer.Adapt(invalidRunner)

		// assert
		assertFirstArgHasContextType(t, err)
	})

	t.Run("second input param must be a structure", func(t *testing.T) {
		// arrange
		invalidRunner := func(ctx context.Context, req interface{}, resp *mock.TestReport) error {
			return nil
		}
		// act
		_, err := queryer.Adapt(invalidRunner)

		// assert
		assertSecondArgHasStructType(t, err)
	})

	t.Run("third input param must be a pointer to a structure", func(t *testing.T) {
		// arrange
		invalidRunner := func(ctx context.Context, req mock.TestQuery, resp mock.TestReport) error {
			return nil
		}
		// act
		_, err := queryer.Adapt(invalidRunner)

		// assert
		assertThirdArgHasPointerToAStructType(t, err)
	})

	t.Run("valid concrete query runner given, valid generic query runner returned", func(t *testing.T) {
		// act
		got, err := queryer.Adapt((&mock.ConcreteQueryerStub{}).Run)

		// assert
		assert.NoError(t, err)
		assertReturnedQueryRunnerIsNotNil(t, got)
	})

	t.Run("ensure that the given valid concrete query runner can return valid result", func(t *testing.T) {
		// arrange
		want := mock.TestReport{Value: 123}

		// act
		runner, err := queryer.Adapt((&mock.ConcreteQueryerStub{Rep: 123}).Run)
		assert.NoError(t, err)

		var res mock.TestReport
		err = runner(context.Background(), mock.TestQuery{}, &res)

		// assert
		assert.NoError(t, err)
		assert.Equals(t, want, res)
	})

	t.Run("ensure that the given valid concrete query runner can return error result", func(t *testing.T) {
		// arrange
		want := errors.New("some error")

		// act
		runner, err := queryer.Adapt((&mock.ConcreteQueryerStub{Err: want}).Run)
		assert.NoError(t, err)

		err = runner(context.Background(), mock.TestQuery{}, &mock.TestReport{})

		// assert
		assert.Equals(t, want, err)
	})
}

func TestMustAdaptor(t *testing.T) {
	t.Run("MustAdapt panics on error", func(t *testing.T) {
		assert.Panic(t, func() {
			queryer.MustAdapt(struct{}{})
		})
	})

	t.Run("valid concrete query runner given, valid generic query runner returned", func(t *testing.T) {
		// act
		got := queryer.MustAdapt((&mock.ConcreteQueryerStub{}).Run)

		// assert
		assertReturnedQueryRunnerIsNotNil(t, got)
	})
}

func assertQueryRunnerHasInvalidSignature(t *testing.T, err error) {
	t.Helper()
	assert.True(t, errors.Is(err, queryer.ErrInvalidQueryRunnerSignature))
}

func assertQueryRunnerIsAFunction(t *testing.T, err error) {
	t.Helper()
	assert.True(t, errors.Is(err, queryer.ErrQueryRunnerIsNotAFunction))
}

func assertFirstArgHasContextType(t *testing.T, err error) {
	t.Helper()
	assert.True(t, errors.Is(err, queryer.ErrFirstArgHasInvalidType))
}

func assertSecondArgHasStructType(t *testing.T, err error) {
	t.Helper()
	assert.True(t, errors.Is(err, queryer.ErrSecondArgHasInvalidType))
}

func assertThirdArgHasPointerToAStructType(t *testing.T, err error) {
	t.Helper()
	assert.True(t, errors.Is(err, queryer.ErrThirdArgHasInvalidType))
}

func assertReturnedQueryRunnerIsNotNil(t *testing.T, got queryer.QueryRunnerFn) {
	t.Helper()
	assert.True(t, !reflect.ValueOf(got).IsNil())
}
