package test

import (
	"design-pattern/factory/gun-factory"
	"fmt"
	"testing"
)

func TestGun(t *testing.T) {
	var gun gun_factory.IGun
	gun, _ = gun_factory.GetGun("ak47")
	printDetails(gun)
	gun, _ = gun_factory.GetGun("musket")
	printDetails(gun)
}

func printDetails(g gun_factory.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
