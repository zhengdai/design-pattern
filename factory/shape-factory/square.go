package shape_factory

import "fmt"

type Square struct {
}

func (s *Square) Draw() {
	fmt.Println("square draw")
}
