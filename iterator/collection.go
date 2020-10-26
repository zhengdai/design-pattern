package iterator

type collection interface {
	CreateInterator() Iterator
}
