package generic

type Config struct {
}

type Service struct {
}

func NewService(config Config) (*Service, error) {
	return &Service{}, nil
}

func (svc *Service) Start() {

}

func (svc *Service) Stop() {

}
