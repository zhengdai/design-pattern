package visitor

type Rectangle struct {
	L int
	B int
}

func (r *Rectangle) Accept(v visitor) {
	v.visitForRectangle(r)
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}
