package abstract_factory

import "fmt"

//抽象工厂多个抽象方法
type ISportFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportFactory, error) {
	switch brand {
	case "adidas":
		return &adidas{}, nil
	case "nike":
		return &nike{}, nil
	default:
		return nil, fmt.Errorf("wrong brand type passed")
	}
}
