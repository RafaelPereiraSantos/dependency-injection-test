package lib

import (
	"log"
	"reflect"
)

type Injector struct {
	AvailableParamebers map[string]any
}

const (
	valueIndex = 0
	errIndex   = 1
)

// This methods registers the implementation of interfaces to be used later to inject into functions.
// Params:
// - name: the name of the interface to be registered.
// - value: the implementation of the interface.
// TODO: remove the "name" parameter and allow the function to identify the name of the interface it is receiving
// by itself.
func (inje *Injector) RegisterInterface(name string, value any) {
	inje.AvailableParamebers[name] = value
}

// This method injects the registered interfaces into the function and calls it.
// Params:
// - buildFunction: the function to be called. This function could receive any arange of parameters, however it must
// return exactly two values, a product and an error, the product being what you are trying to "build" with the
// injection.
// Returns: the product of the function that was given and a possible error.
func (inje *Injector) FillAndCall(buildFunction any) (any, error) {
	amountOfParams := reflect.TypeOf(buildFunction).NumIn() // checks what is the type of the function.

	parameters := []reflect.Value{}

	// in this loop the function tries to identify which type of value is required in each position of the
	// functino parameters and arranges them acordingly.
	for i := 0; i < amountOfParams; i++ {
		paramName := reflect.TypeOf(buildFunction).In(i).Name()

		if value, ok := inje.AvailableParamebers[paramName]; ok {
			parameters = append(parameters, reflect.ValueOf(value))
		} else {
			log.Fatalf("Parameter %s not found in the available parameters", paramName)
		}
	}

	returnedValuesCout := len(reflect.ValueOf(buildFunction).Call(parameters))

	// For now, the logic expects that the function returns exactly two values, a "product" and an error.
	// The product being what the client expect to receive from the given function and the error an error that
	// indicates that something went wrong during the process of creating the "product".
	// TODO: create a more generic function to return N values in any order.
	if returnedValuesCout != 2 {
		log.Fatalf("The function must return two values, the product and an error")
	}

	value := reflect.ValueOf(buildFunction).Call(parameters)[valueIndex].Interface()
	err := reflect.ValueOf(buildFunction).Call(parameters)[errIndex].Interface()

	switch err := err.(type) {
	case error:
		return value, err
	case nil:
		return value, nil
	default:
		log.Fatalf("fn must return an error as the second value.")
	}

	return nil, nil
}
