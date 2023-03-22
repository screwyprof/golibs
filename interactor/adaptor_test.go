package interactor_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/screwyprof/golibs/interactor"
)

func TestAdapt(t *testing.T) {
	t.Parallel()

	t.Run("given use case runner must be a function", func(t *testing.T) {
		t.Parallel()

		// act
		_, err := interactor.Adapt(struct{}{})

		// assert
		assertUseCaseRunnerIsAFunction(t, err)
	})

	t.Run("given use case runner must have 3 arguments", func(t *testing.T) {
		t.Parallel()

		// arrange
		invalidRunner := func(ctx context.Context, req TestRequest) error {
			return nil
		}

		// act
		_, err := interactor.Adapt(invalidRunner)

		// assert
		assertUseCaseRunnerHasInvalidSignature(t, err)
	})

	t.Run("first input param must be context.Context", func(t *testing.T) {
		t.Parallel()

		// arrange
		invalidRunner := func(ctx struct{}, req TestRequest, resp *TestResponse) error {
			return nil
		}

		// act
		_, err := interactor.Adapt(invalidRunner)

		// assert
		assertFirstArgHasContextType(t, err)
	})

	t.Run("second input param must be a structure", func(t *testing.T) {
		t.Parallel()

		// arrange
		invalidRunner := func(ctx context.Context, req interface{}, resp *TestResponse) error {
			return nil
		}

		// act
		_, err := interactor.Adapt(invalidRunner)

		// assert
		assertSecondArgHasStructType(t, err)
	})

	t.Run("third input param must be a pointer to a structure", func(t *testing.T) {
		t.Parallel()

		// arrange
		invalidRunner := func(ctx context.Context, req TestRequest, resp TestResponse) error {
			return nil
		}

		// act
		_, err := interactor.Adapt(invalidRunner)

		// assert
		assertThirdArgHasPointerToAStructType(t, err)
	})

	t.Run("valid concrete use case runner given, valid generic use case runner returned", func(t *testing.T) {
		t.Parallel()

		// act
		got, err := interactor.Adapt(ConcreteUseCase{}.RunUseCase)

		// assert
		assert.NoError(t, err)
		assertReturnedUseCaseRunnerIsNotNil(t, got)
	})

	t.Run("ensure that the given valid concrete use case runner can return valid result", func(t *testing.T) {
		t.Parallel()

		// arrange
		want := TestResponse{result: 123}

		// act
		runner, err := interactor.Adapt(ConcreteUseCase{res: 123}.RunUseCase)
		assert.NoError(t, err)

		var res TestResponse
		err = runner(context.Background(), TestRequest{id: 123}, &res)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, want, res)
	})

	t.Run("ensure that the given valid concrete use case runner can return error result", func(t *testing.T) {
		t.Parallel()

		// arrange
		want := errSomeErr

		// act
		runner, err := interactor.Adapt(ConcreteUseCase{err: want}.RunUseCase)
		assert.NoError(t, err)

		err = runner(context.Background(), TestRequest{}, &TestResponse{})

		// assert
		assert.ErrorIs(t, err, want)
	})
}

func TestMustAdapt(t *testing.T) {
	t.Parallel()

	t.Run("it panics if it cannot adapt a use case runner", func(t *testing.T) {
		t.Parallel()

		assert.Panics(t, func() {
			interactor.MustAdapt(struct{}{})
		})
	})

	t.Run("it adapts a use case runner", func(t *testing.T) {
		t.Parallel()

		// act
		runner := interactor.MustAdapt(ConcreteUseCase{res: 123}.RunUseCase)

		// assert
		assert.NotNil(t, runner)
	})
}

func assertUseCaseRunnerHasInvalidSignature(t *testing.T, err error) {
	t.Helper()
	assert.True(t, errors.Is(err, interactor.ErrInvalidUseCaseRunnerSignature))
}

func assertUseCaseRunnerIsAFunction(t *testing.T, err error) {
	t.Helper()
	assert.True(t, errors.Is(err, interactor.ErrUseCaseRunnerIsNotAFunction))
}

func assertFirstArgHasContextType(t *testing.T, err error) {
	t.Helper()
	assert.True(t, errors.Is(err, interactor.ErrFirstArgHasInvalidType))
}

func assertSecondArgHasStructType(t *testing.T, err error) {
	t.Helper()
	assert.True(t, errors.Is(err, interactor.ErrSecondArgHasInvalidType))
}

func assertThirdArgHasPointerToAStructType(t *testing.T, err error) {
	t.Helper()
	assert.True(t, errors.Is(err, interactor.ErrThirdArgHasInvalidType))
}

func assertReturnedUseCaseRunnerIsNotNil(t *testing.T, got interactor.UseCaseRunnerFn) {
	t.Helper()
	assert.True(t, !reflect.ValueOf(got).IsNil())
}
