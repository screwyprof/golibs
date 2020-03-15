package cmdhandler_test

import (
	"context"
	"errors"
	"testing"

	"github.com/screwyprof/golibs/assert"

	"github.com/screwyprof/golibs/cmdhandler"
	"github.com/screwyprof/golibs/cmdhandler/mock"
)

func TestDispatcher(t *testing.T) {
	t.Run("dispatcher implements CommandHandler interface", func(t *testing.T) {
		dispatcher := cmdhandler.NewDispatcher()
		var _ cmdhandler.CommandHandler = dispatcher
	})

	t.Run("when command not found, an error returned", func(t *testing.T) {
		// arrange
		dispatcher := cmdhandler.NewDispatcher()

		// act
		err := dispatcher.Handle(context.Background(), mock.TestCommand{})

		// assert
		assertCommandHandlerNotFound(t, err)
	})

	t.Run("when command registered, it is being run", func(t *testing.T) {
		// arrange
		commandHandler := &mock.GenericCommandHandlerSpy{}

		dispatcher := cmdhandler.NewDispatcher()
		dispatcher.RegisterCommandHandler("TestCommand", commandHandler.Handle)

		// act
		err := dispatcher.Handle(context.Background(), mock.TestCommand{})

		// assert
		assertCommandWasRunSuccessfully(t, err, commandHandler)
	})
}

func assertCommandWasRunSuccessfully(t *testing.T, err error, commandHandler *mock.GenericCommandHandlerSpy) {
	assert.NoError(t, err)
	assert.True(t, commandHandler.WasCalled)
}

func assertCommandHandlerNotFound(t *testing.T, err error) {
	t.Helper()

	assert.True(t, errors.Is(err, cmdhandler.ErrNotFound))
}
