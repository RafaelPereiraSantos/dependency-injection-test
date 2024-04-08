package main

import (
	"fmt"
	"reflect"

	"github.com/RafaelPereiraSantos/injection-test/interfaces"
	"github.com/RafaelPereiraSantos/injection-test/services"
)

type (
	Worker struct {
		service01 interfaces.ServiceInterface01
		service02 interfaces.ServiceInterface02
		service03 interfaces.ServiceInterface03
	}

	Injector struct {
		availableParamebers map[string]any
	}
)

func main() {
	// usual way
	srv01 := services.NewServiceImplementation01()
	srv02 := services.NewServiceImplementation02()
	srv03 := services.NewServiceImplementation03()

	wrk := createWorker(srv01, srv02, srv03)
	wrk.DoSomething()

	// injecting parameters
	inj := &Injector{
		availableParamebers: make(map[string]any),
	}
	inj.RegisterInterface("ServiceInterface01", services.NewServiceImplementation01())
	inj.RegisterInterface("ServiceInterface02", services.NewServiceImplementation02())
	inj.RegisterInterface("ServiceInterface03", services.NewServiceImplementation03())

	wrk2 := inj.FillAndCall(createWorker).(*Worker)
	wrk2.DoSomething()
}

func createWorker(
	service01 interfaces.ServiceInterface01,
	service02 interfaces.ServiceInterface02,
	service03 interfaces.ServiceInterface03,
) *Worker {
	return &Worker{service01, service02, service03}
}

func (w *Worker) DoSomething() {
	w.service01.MethodA()
	w.service02.MethodB()
	w.service03.MethodC()

	fmt.Println("Called")
}

func (i *Injector) RegisterInterface(name string, value any) {
	i.availableParamebers[name] = value
}

func (inj *Injector) FillAndCall(fn any) any {
	amountOfParams := reflect.TypeOf(fn).NumIn()

	parameters := []reflect.Value{}

	for i := 0; i < amountOfParams; i++ {
		paramName := reflect.TypeOf(fn).In(i).Name()

		if value, ok := inj.availableParamebers[paramName]; ok {
			parameters = append(parameters, reflect.ValueOf(value))
		} else {
			fmt.Println("Parameter not found")
		}

	}

	return reflect.ValueOf(fn).Call(parameters)[0].Interface()
}
