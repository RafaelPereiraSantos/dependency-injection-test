package services

type (
	ServiceImplementation03 struct{}
)

func NewServiceImplementation03() *ServiceImplementation03 {
	return &ServiceImplementation03{}
}

func (srv *ServiceImplementation03) MethodC() string {
	return "ServiceImplementation03"
}
