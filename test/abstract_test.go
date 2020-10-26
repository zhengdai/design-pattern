package test

import (
	abstract_factory "design-pattern/abstract-factory"
	"fmt"
	"testing"
)

func TestAbstract(t *testing.T) {
	var factory abstract_factory.ISportFactory
	factory, _ = abstract_factory.GetSportsFactory("adidas")
	printShoeDetails(factory.MakeShoe())
	printShirtDetails(factory.MakeShirt())
	factory, _ = abstract_factory.GetSportsFactory("nike")
	printShoeDetails(factory.MakeShoe())
	printShirtDetails(factory.MakeShirt())
}

func printShoeDetails(s abstract_factory.IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShirtDetails(s abstract_factory.IShirt) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
