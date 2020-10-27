package visitor

import "fmt"

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(c *Circle) {
	fmt.Println("Calculating area for circle")
}

func (a *AreaCalculator) visitForRectangle(r *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}
