package test

import (
	"design-pattern/factory/shape-factory"
	"testing"
)

func TestFactory(t *testing.T) {
	var shape shape_factory.Shape
	shape, _ = shape_factory.GetShape("rectangle")
	shape.Draw()
	shape, _ = shape_factory.GetShape("circle")
	shape.Draw()
	shape, _ = shape_factory.GetShape("square")
	shape.Draw()
}
