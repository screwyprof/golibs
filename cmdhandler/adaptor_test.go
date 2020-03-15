package cmdhandler_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/screwyprof/golibs/assert"
	"github.com/screwyprof/golibs/cmdhandler"
	"github.com/screwyprof/golibs/cmdhandler/mock"
)

func TestAdaptor(t *testing.T) {
	t.Run("given command handler must be a function", func(t *testing.T) {
		// act
		_, err := cmdhandler.Adapt(struct{}{})

		// assert
		assertCommandHandlerIsAFunction(t, err)
	})

	t.Run("given command handler must have 2 arguments", func(t *testing.T) {
		// arrange
		invalidRunner := func(req mock.TestCommand) error {
			return nil
		}
		// act
		_, err := cmdhandler.Adapt(invalidRunner)

		// assert
		assertCommandHandlerHasInvalidSignature(t, err)
	})

	t.Run("first input param must be context.Context", func(t *testing.T) {
		// arrange
		invalidRunner := func(ctx struct{}, req mock.TestCommand) error {
			return nil
		}
		// act
		_, err := cmdhandler.Adapt(invalidRunner)

		// assert
		assertFirstArgHasContextType(t, err)
	})

	t.Run("second input param must be a structure", func(t *testing.T) {
		// arrange
		invalidRunner := func(ctx context.Context, req interface{}) error {
			return nil
		}
		// act
		_, err := cmdhandler.Adapt(invalidRunner)

		// assert
		assertSecondArgHasStructType(t, err)
	})

	t.Run("valid concrete command handler given, valid generic command handler returned", func(t *testing.T) {
		// act
		got, err := cmdhandler.Adapt((&mock.ConcreteCommandHandlerStub{}).Handle)

		// assert
		assert.NoError(t, err)
		assertReturnedCommandHandlerIsNotNil(t, got)
	})

	t.Run("ensure that the given valid concrete command handler can return valid result", func(t *testing.T) {
		// act
		runner, err := cmdhandler.Adapt((&mock.ConcreteCommandHandlerStub{}).Handle)
		assert.NoError(t, err)

		err = runner(context.Background(), mock.TestCommand{})

		// assert
		assert.NoError(t, err)
	})

	t.Run("ensure that the given valid concrete command handler can return error result", func(t *testing.T) {
		// arrange
		want := errors.New("some error")

		// act
		runner, err := cmdhandler.Adapt((&mock.ConcreteCommandHandlerStub{Err: want}).Handle)
		assert.NoError(t, err)

		err = runner(context.Background(), mock.TestCommand{})

		// assert
		assert.Equals(t, want, err)
	})
}

func TestMustAdaptor(t *testing.T) {
	t.Run("MustAdapt panics on error", func(t *testing.T) {
		assert.Panic(t, func() {
			cmdhandler.MustAdapt(struct{}{})
		})
	})

	t.Run("valid concrete command handler given, valid generic command handler returned", func(t *testing.T) {
		// act
		got := cmdhandler.MustAdapt((&mock.ConcreteCommandHandlerStub{}).Handle)

		// assert
		assertReturnedCommandHandlerIsNotNil(t, got)
	})
}

func assertCommandHandlerHasInvalidSignature(t *testing.T, err error) {
	t.Helper()
	assert.True(t, errors.Is(err, cmdhandler.ErrInvalidCommandHandlerSignature))
}

func assertCommandHandlerIsAFunction(t *testing.T, err error) {
	t.Helper()
	assert.True(t, errors.Is(err, cmdhandler.ErrCommandHandlerIsNotAFunction))
}

func assertFirstArgHasContextType(t *testing.T, err error) {
	t.Helper()
	assert.True(t, errors.Is(err, cmdhandler.ErrFirstArgHasInvalidType))
}

func assertSecondArgHasStructType(t *testing.T, err error) {
	t.Helper()
	assert.True(t, errors.Is(err, cmdhandler.ErrSecondArgHasInvalidType))
}

func assertReturnedCommandHandlerIsNotNil(t *testing.T, got cmdhandler.CommandHandlerFn) {
	t.Helper()
	assert.True(t, !reflect.ValueOf(got).IsNil())
}
