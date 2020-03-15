package interactor

import (
	"context"
	"errors"
	"reflect"
)

// Guard errors
var (
	ErrInvalidUseCaseRunnerSignature = errors.New("useCaseRunner must have 3 input params")
	ErrUseCaseRunnerIsNotAFunction   = errors.New("useCaseRunner is not a function")
	ErrFirstArgHasInvalidType        = errors.New("first input argument must have context.Context type")
	ErrSecondArgHasInvalidType       = errors.New("second input argument must be a struct")
	ErrThirdArgHasInvalidType        = errors.New("third input argument must be a pointer to a struct")
)

// Adapt transforms a concrete use case runner into a generic one.
// A concrete runner function should have 3 arguments:
// - ctx context.Context,
// - req - a request struct,
// - res - a pointer to a response struct.
//
// The returned param must have error type.
// An example signature may look like:
//   func(ctx context.Context, req TestRequest, res *TestResponse) error
//
func Adapt(useCaseRunner interface{}) (UseCaseRunnerFn, error) {
	useCaseRunnerType := reflect.TypeOf(useCaseRunner)
	err := ensureSignatureIsValid(useCaseRunnerType)
	if err != nil {
		return nil, err
	}

	fn := func(ctx context.Context, req interface{}, res interface{}) error {
		return invokeUseCaseRunner(useCaseRunner, ctx, req, res)
	}

	return fn, nil
}

func ensureSignatureIsValid(useCaseRunnerType reflect.Type) error {
	if useCaseRunnerType.Kind() != reflect.Func {
		return ErrUseCaseRunnerIsNotAFunction
	}

	if useCaseRunnerType.NumIn() != 3 {
		return ErrInvalidUseCaseRunnerSignature
	}

	return ensureParamsHaveValidTypes(useCaseRunnerType)
}

func ensureParamsHaveValidTypes(useCaseRunnerType reflect.Type) error {
	if !firstArgIsContext(useCaseRunnerType) {
		return ErrFirstArgHasInvalidType
	}

	if !secondArgIsStructure(useCaseRunnerType) {
		return ErrSecondArgHasInvalidType
	}

	if !thirdArgIsAPointerToAStruct(useCaseRunnerType) {
		return ErrThirdArgHasInvalidType
	}

	return nil
}

func firstArgIsContext(useCaseRunnerType reflect.Type) bool {
	ctxtInterface := reflect.TypeOf((*context.Context)(nil)).Elem()
	ctx := useCaseRunnerType.In(0)
	return ctx.Implements(ctxtInterface)
}

func secondArgIsStructure(useCaseRunnerType reflect.Type) bool {
	return useCaseRunnerType.In(1).Kind() == reflect.Struct
}

func thirdArgIsAPointerToAStruct(useCaseRunnerType reflect.Type) bool {
	thirdArg := useCaseRunnerType.In(2)
	return thirdArg.Kind() == reflect.Ptr && thirdArg.Elem().Kind() == reflect.Struct

}

func invokeUseCaseRunner(useCaseRunner interface{}, args ...interface{}) error {
	result := invoke(useCaseRunner, args...)
	resErr := result[0].Interface()
	if resErr != nil {
		return resErr.(error)
	}
	return nil
}

func invoke(fn interface{}, args ...interface{}) []reflect.Value {
	v := reflect.ValueOf(fn)
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}
	return v.Call(in)
}
