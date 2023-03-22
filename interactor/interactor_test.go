package interactor_test

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v4"
	"github.com/stretchr/testify/assert"
)

func TestInteractor(t *testing.T) {
	t.Run("Valid request provided, valid response returned", func(t *testing.T) {
		// arrange
		ID := gofakeit.Number(1, 100)
		want := TestResponse{result: ID}

		sut := ConcreteUseCase{res: ID}

		// act
		var res TestResponse
		err := sut.RunUseCase(context.Background(), TestRequest{id: ID}, &res)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, want, res)
	})

	t.Run("Invalid request provided, an error returned", func(t *testing.T) {
		// arrange
		sut := ConcreteUseCase{err: errors.New("an error")}

		// act
		var res TestResponse
		err := sut.RunUseCase(context.Background(), TestRequest{}, &res)

		// assert
		assert.Error(t, err)
	})
}
