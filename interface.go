package service

type IService interface {
	Start()
	Stop()
	Errors()
}
