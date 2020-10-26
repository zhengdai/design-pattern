package abstract_factory

//具体产品
type NikeShoe struct {
	logo string
	size int
}

func (s *NikeShoe) SetLogo(logo string) {
	s.logo = logo
}

func (s *NikeShoe) SetSize(size int) {
	s.size = size
}

func (s *NikeShoe) GetLogo() string {
	return s.logo
}

func (s *NikeShoe) GetSize() int {
	return s.size
}
