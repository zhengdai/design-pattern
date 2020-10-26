package chain

type department interface {
	Execute(*Patient)
	SetNext(department)
}
