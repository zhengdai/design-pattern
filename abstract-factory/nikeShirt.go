package abstract_factory

//具体产品
type NikeShirt struct {
	logo string
	size int
}

func (s *NikeShirt) SetLogo(logo string) {
	s.logo = logo
}

func (s *NikeShirt) SetSize(size int) {
	s.size = size
}

func (s *NikeShirt) GetLogo() string {
	return s.logo
}

func (s *NikeShirt) GetSize() int {
	return s.size
}
