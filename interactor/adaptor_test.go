package interactor_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/screwyprof/golibs/assert"
	"github.com/screwyprof/golibs/interactor"
)

func TestAdaptor(t *testing.T) {
	t.Run("given use case runner must be a function", func(t *testing.T) {
		// act
		_, err := interactor.Adapt(struct{}{})

		// assert
		assertUseCaseRunnerIsAFunction(t, err)
	})

	t.Run("given use case runner must have 3 arguments", func(t *testing.T) {
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
		// act
		got, err := interactor.Adapt(ConcreteInteractorStub{}.RunUseCase)

		// assert
		assert.NoError(t, err)
		assertReturnedUseCaseRunnerIsNotNil(t, got)
	})

	t.Run("ensure that the given valid concrete use case runner can return valid result", func(t *testing.T) {
		// arrange
		want := TestResponse{result: 123}

		// act
		runner, err := interactor.Adapt(ConcreteInteractorStub{res: 123}.RunUseCase)
		assert.NoError(t, err)

		var res TestResponse
		err = runner(context.Background(), TestRequest{}, &res)

		// assert
		assert.NoError(t, err)
		assert.Equals(t, want, res)
	})

	t.Run("ensure that the given valid concrete use case runner can return error result", func(t *testing.T) {
		// arrange
		want := errors.New("some error")

		// act
		runner, err := interactor.Adapt(ConcreteInteractorStub{err: want}.RunUseCase)
		assert.NoError(t, err)

		err = runner(context.Background(), TestRequest{}, &TestResponse{})

		// assert
		assert.Equals(t, want, err)
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
