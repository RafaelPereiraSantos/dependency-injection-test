package services

type (
	ServiceImplementation01 struct{}
)

func NewServiceImplementation01() *ServiceImplementation01 {
	return &ServiceImplementation01{}
}

func (srv *ServiceImplementation01) MethodA() string {
	return "ServiceImplementation01"
}
