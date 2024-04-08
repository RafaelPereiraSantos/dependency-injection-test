package main

import (
	"github.com/RafaelPereiraSantos/injection-test/interfaces"
	"github.com/RafaelPereiraSantos/injection-test/lib"
	"github.com/RafaelPereiraSantos/injection-test/services"
	"github.com/RafaelPereiraSantos/injection-test/workers"
)

func main() {
	// usual way
	srv01 := services.NewServiceImplementation01()
	srv02 := services.NewServiceImplementation02()
	srv03 := services.NewServiceImplementation03()

	regularWrk := createWorker01(srv01, srv02, srv03)
	regularWrk.DoSomething()

	// injecting parameters
	inj := lib.Injector{
		AvailableParamebers: make(map[string]any),
	}
	inj.RegisterInterface("ServiceInterface01", services.NewServiceImplementation01())
	inj.RegisterInterface("ServiceInterface02", services.NewServiceImplementation02())
	inj.RegisterInterface("ServiceInterface03", services.NewServiceImplementation03())

	injectedWorker01 := inj.FillAndCall(createWorker01).(*workers.Worker01)
	injectedWorker01.DoSomething()

	injectedWorker02 := inj.FillAndCall(createWorker02).(*workers.Worker02)
	injectedWorker02.DoSomething()

	injectedWorker03 := inj.FillAndCall(createWorker03).(*workers.Worker03)
	injectedWorker03.DoSomething()
}

func createWorker01(
	service01 interfaces.ServiceInterface01,
	service02 interfaces.ServiceInterface02,
	service03 interfaces.ServiceInterface03,
) *workers.Worker01 {
	return &workers.Worker01{service01, service02, service03}
}

func createWorker02(
	service02 interfaces.ServiceInterface02,
	service03 interfaces.ServiceInterface03,
) *workers.Worker02 {
	return &workers.Worker02{service02, service03}
}

func createWorker03(
	service01 interfaces.ServiceInterface01,
	service02 interfaces.ServiceInterface02,
) *workers.Worker03 {
	return &workers.Worker03{service01, service02}
}
