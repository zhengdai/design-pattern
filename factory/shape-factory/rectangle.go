package shape_factory

import "fmt"

type Rectangle struct {
}

func (r *Rectangle) Draw() {
	fmt.Println("rectangle draw")
}
