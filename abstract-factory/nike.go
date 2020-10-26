package abstract_factory

//具体工厂
type nike struct {
}

func (a *nike) MakeShoe() IShoe {
	return &NikeShoe{
		logo: "adidas",
		size: 14,
	}
}

func (a *nike) MakeShirt() IShirt {
	return &NikeShirt{
		logo: "nike",
		size: 100,
	}
}
