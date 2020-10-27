package observer

type Subject interface {
	Register(observer observer)
	Deregister(observer observer)
	notifyAll()
}
