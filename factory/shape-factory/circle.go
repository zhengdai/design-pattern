package shape_factory

import "fmt"

type Circle struct {
}

func (r *Circle) Draw() {
	fmt.Println("circle draw")
}
