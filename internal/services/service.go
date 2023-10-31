package services

type Service struct {
	Model interface{}
}

func (s Service) Index() string {
	return "Index"
}
