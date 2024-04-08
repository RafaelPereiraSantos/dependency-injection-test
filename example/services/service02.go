package services

type (
	ServiceImplementation02 struct{}
)

func NewServiceImplementation02() *ServiceImplementation02 {
	return &ServiceImplementation02{}
}

func (srv *ServiceImplementation02) MethodB() string {
	return "ServiceImplementation02"
}
