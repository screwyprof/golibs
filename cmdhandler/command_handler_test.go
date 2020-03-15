package cmdhandler_test

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v4"

	"github.com/screwyprof/golibs/assert"
	"github.com/screwyprof/golibs/cmdhandler/mock"
)

func TestCommandHandler(t *testing.T) {
	t.Run("Valid command provided, valid report returned", func(t *testing.T) {
		// arrange
		commandID := gofakeit.UUID()

		concreteCommandHandler := mock.ConcreteCommandHandlerSpy{}

		// act
		err := concreteCommandHandler.Handle(context.Background(), mock.TestCommand{ID: commandID})

		// assert
		assert.NoError(t, err)
		assert.True(t, concreteCommandHandler.WasCalled)
	})

	t.Run("Invalid command provided, an error returned", func(t *testing.T) {
		// arrange
		concreteCommandHandler := mock.ConcreteCommandHandlerStub{Err: errors.New("an error")}

		// act
		err := concreteCommandHandler.Handle(context.Background(), mock.TestCommand{})

		// assert
		assert.NotNil(t, err)
	})
}
