package workers

import (
	"fmt"

	"github.com/RafaelPereiraSantos/injection-test/example/interfaces"
)

// Here we have a sort of "workers" but they could be http APIs, gRPC clients, or any other kind of service that are
// a good candidate for a dependency injection approach. These workers were defined to have different dependencies
// on purpose, so we can see how the injector can handle different scenarios.
type (
	Worker01 struct {
		Service01 interfaces.ServiceInterface01
		Service02 interfaces.ServiceInterface02
		Service03 interfaces.ServiceInterface03
	}

	Worker02 struct {
		Service02 interfaces.ServiceInterface02
		Service03 interfaces.ServiceInterface03
	}

	Worker03 struct {
		Service01 interfaces.ServiceInterface01
		Service02 interfaces.ServiceInterface02
	}
)

func (w *Worker01) DoSomething() {
	w.Service01.MethodA()
	w.Service02.MethodB()
	w.Service03.MethodC()

	fmt.Println("Called worker 01")
}

func (w *Worker02) DoSomething() {
	w.Service02.MethodB()
	w.Service03.MethodC()

	fmt.Println("Called worker 02")
}

func (w *Worker03) DoSomething() {
	w.Service01.MethodA()
	w.Service02.MethodB()

	fmt.Println("Called worker 03")
}
