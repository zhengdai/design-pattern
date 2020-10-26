package test

import (
	"design-pattern/decorator"
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	pizza := &decorator.VeggieMania{}

	pizzWithCheese := &decorator.CheeseTopping{
		Pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &decorator.TomatoTopping{
		Pizza: pizzWithCheese,
	}
	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}
