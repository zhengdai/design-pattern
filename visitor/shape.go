package visitor

type shape interface {
	getType() string
	Accept(visitor)
}
