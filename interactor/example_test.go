package interactor_test

import (
	"context"
	"fmt"
	"log"

	"github.com/screwyprof/golibs/interactor"
)

func ExampleDispatcher() {
	// arrange
	useCaseRunner := &ConcreteUseCase{res: 42}

	dispatcher := interactor.NewDispatcher()
	dispatcher.Register(TestRequest{}, interactor.MustAdapt(useCaseRunner.RunUseCase))

	// act
	var res TestResponse
	if err := dispatcher.Run(context.Background(), TestRequest{}, &res); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer to life the universe and everything: %d\n", res.result)

	// Output:
	// The answer to life the universe and everything: 42
}
