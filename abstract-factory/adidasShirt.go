package abstract_factory

//具体产品
type AdidasShirt struct {
	logo string
	size int
}

func (s *AdidasShirt) SetLogo(logo string) {
	s.logo = logo
}

func (s *AdidasShirt) SetSize(size int) {
	s.size = size
}

func (s *AdidasShirt) GetLogo() string {
	return s.logo
}

func (s *AdidasShirt) GetSize() int {
	return s.size
}
