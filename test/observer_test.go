package test

import (
	"design-pattern/observer"
	"testing"
)

func TestObserver(t *testing.T) {
	shirtItem := observer.NewItem("Nike shirt")
	observer1 := &observer.Customer{
		Id: "aaa@gmail.com",
	}
	observer2 := &observer.Customer{
		Id: "bbb@gmail.com",
	}

	shirtItem.Register(observer1)
	shirtItem.Register(observer2)
	shirtItem.UpdateAvailability()
}
