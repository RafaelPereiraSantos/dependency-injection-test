package lib

import (
	"log"
	"reflect"
)

type Injector struct {
	AvailableParamebers map[string]any
}

// This methods registers the implementation of interfaces to be used later to inject into functions.
// Params:
// - name: the name of the interface to be registered.
// - value: the implementation of the interface.
func (i *Injector) RegisterInterface(name string, value any) {
	i.AvailableParamebers[name] = value
}

// This method injects the registered interfaces into the function and calls it.
// Params:
// - fn: the function to be called.
// Returns: the product of the function that was given.
func (inj *Injector) FillAndCall(fn any) any {
	amountOfParams := reflect.TypeOf(fn).NumIn()

	parameters := []reflect.Value{}

	for i := 0; i < amountOfParams; i++ {
		paramName := reflect.TypeOf(fn).In(i).Name()

		if value, ok := inj.AvailableParamebers[paramName]; ok {
			parameters = append(parameters, reflect.ValueOf(value))
		} else {
			log.Fatalf("Parameter %s not found in the available parameters", paramName)
		}

	}

	return reflect.ValueOf(fn).Call(parameters)[0].Interface()
}
