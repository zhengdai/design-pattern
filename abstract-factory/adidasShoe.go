package abstract_factory

//具体产品
type AdidasShoe struct {
	logo string
	size int
}

func (s *AdidasShoe) SetLogo(logo string) {
	s.logo = logo
}

func (s *AdidasShoe) SetSize(size int) {
	s.size = size
}

func (s *AdidasShoe) GetLogo() string {
	return s.logo
}

func (s *AdidasShoe) GetSize() int {
	return s.size
}
