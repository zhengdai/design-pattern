package flyweight

type countTerroristDress struct {
	color string
}

func (t *countTerroristDress) getColor() string {
	return t.color
}

func newCountTerroristDress() *countTerroristDress {
	return &countTerroristDress{color: "green"}
}
