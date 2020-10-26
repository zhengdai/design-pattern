package abstract_factory

//具体工厂
type adidas struct {
}

func (a *adidas) MakeShoe() IShoe {
	return &AdidasShoe{
		logo: "adidas",
		size: 14,
	}
}

func (a *adidas) MakeShirt() IShirt {
	return &AdidasShirt{
		logo: "adidas",
		size: 12,
	}
}
