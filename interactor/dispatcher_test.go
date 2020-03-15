package interactor_test

import (
	"context"
	"errors"
	"testing"

	"github.com/screwyprof/golibs/assert"

	"github.com/screwyprof/golibs/interactor"
)

func TestDispatcher(t *testing.T) {
	t.Run("dispatcher implements UseCaseRunner interface", func(t *testing.T) {
		dispatcher := interactor.NewDispatcher()
		var _ interactor.UseCaseRunner = dispatcher
	})

	t.Run("when use case not found, an error returned", func(t *testing.T) {
		// arrange
		dispatcher := interactor.NewDispatcher()

		// act
		var res TestResponse
		err := dispatcher.RunUseCase(context.Background(), TestRequest{}, &res)

		// assert
		assertUseCaseRunnerNotFound(t, err)
	})

	t.Run("when use case registered, it is being run", func(t *testing.T) {
		// arrange
		useCaseRunner := &GeneralInteractorSpy{}

		dispatcher := interactor.NewDispatcher()
		dispatcher.RegisterUseCaseRunner("TestRequest", useCaseRunner.RunUseCase)

		// act
		var res TestResponse
		err := dispatcher.RunUseCase(context.Background(), TestRequest{}, &res)

		// assert
		assertUseCaseWasRunSuccessfully(t, err, useCaseRunner)
	})
}

func assertUseCaseWasRunSuccessfully(t *testing.T, err error, useCaseRunner *GeneralInteractorSpy) {
	assert.NoError(t, err)
	assert.True(t, useCaseRunner.wasCalled)
}

func assertUseCaseRunnerNotFound(t *testing.T, err error) {
	t.Helper()

	assert.True(t, errors.Is(err, interactor.ErrNotFound))
}
