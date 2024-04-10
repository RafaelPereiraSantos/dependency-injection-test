package main

import (
	"fmt"
	"time"

	lib "github.com/RafaelPereiraSantos/injection-test"
	"github.com/RafaelPereiraSantos/injection-test/example/interfaces"
	"github.com/RafaelPereiraSantos/injection-test/example/services"
	"github.com/RafaelPereiraSantos/injection-test/example/workers"
)

func main() {
	regularWayTime := time.Now()
	initializeWorkersInTheRegularWay()
	elapsedRegularWay := time.Since(regularWayTime)

	dependencyInjectionTime := time.Now()
	initializeWorkersWithDependencyInjection()
	elapsedDependencyInjectionTime := time.Since(dependencyInjectionTime)

	fmt.Println("Elapsed time regular way in nanoseconnds:", elapsedRegularWay.Nanoseconds())
	fmt.Println("Elapsed time dependency injection way in nanoseconnds:", elapsedDependencyInjectionTime.Nanoseconds())
}

func initializeWorkersInTheRegularWay() {
	// Regular way of creating a new instance of a fake worker, it creates each parameter separately and then pass them
	// in the contruction function.

	srv01 := services.NewServiceImplementation01()
	srv02 := services.NewServiceImplementation02()
	srv03 := services.NewServiceImplementation03()

	regularWrk1, _ := createWorker01(srv01, srv02, srv03)
	regularWrk1.DoSomething()

	regularWrk2, _ := createWorker02(srv02, srv03)
	regularWrk2.DoSomething()

	regularWrk3, _ := createWorker03(srv01, srv02)
	regularWrk3.DoSomething()
}

func initializeWorkersWithDependencyInjection() {
	// injecting parameters way:
	// it firstly register each available parameter to be used later with their respective interface names. Then it
	// calls the function that will be injected with the registered parameters as many times as it is needed without
	// the need of passing each paramter separately again.

	inj := lib.Injector{
		AvailableParamebers: make(map[string]any),
	}
	inj.RegisterInterface("ServiceInterface01", services.NewServiceImplementation01())
	inj.RegisterInterface("ServiceInterface02", services.NewServiceImplementation02())
	inj.RegisterInterface("ServiceInterface03", services.NewServiceImplementation03())

	injectedWorker01, _ := inj.FillAndCall(createWorker01)
	injectedWorker01.(*workers.Worker01).DoSomething()

	injectedWorker02, _ := inj.FillAndCall(createWorker02)
	injectedWorker02.(*workers.Worker02).DoSomething()

	injectedWorker03, _ := inj.FillAndCall(createWorker03)
	injectedWorker03.(*workers.Worker03).DoSomething()
}

// each build function below has its own parameters needed, each combination is diferently.

func createWorker01(
	service01 interfaces.ServiceInterface01,
	service02 interfaces.ServiceInterface02,
	service03 interfaces.ServiceInterface03,
) (*workers.Worker01, error) {
	return &workers.Worker01{service01, service02, service03}, nil
}

func createWorker02(
	service02 interfaces.ServiceInterface02,
	service03 interfaces.ServiceInterface03,
) (*workers.Worker02, error) {
	return &workers.Worker02{service02, service03}, nil
}

func createWorker03(
	service01 interfaces.ServiceInterface01,
	service02 interfaces.ServiceInterface02,
) (*workers.Worker03, error) {
	return &workers.Worker03{service01, service02}, nil
}
