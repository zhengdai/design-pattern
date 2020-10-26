package shape_factory

import "fmt"

type Shape interface {
	Draw()
}

func GetShape(shape string) (Shape, error) {
	switch shape {
	case "rectangle":
		return &Rectangle{}, nil
	case "circle":
		return &Circle{}, nil
	case "square":
		return &Square{}, nil
	default:
		return nil, fmt.Errorf("Wrong shape type passed!")
	}
}
