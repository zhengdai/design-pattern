package observer

import "fmt"

type item struct {
	observerList []observer
	name         string
	inStock      bool
}

func NewItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) UpdateAvailability() {
	fmt.Printf("Item %s is now is stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *item) Register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) Deregister(o observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func removeFromSlice(observerList []observer, o observer) []observer {
	length := len(observerList)
	for i, observer := range observerList {
		if o.getID() == observer.getID() {
			observerList[length-1], observerList[i] = observerList[i], observerList[length-1]
			return observerList[:length-1]
		}
	}
	return observerList
}
