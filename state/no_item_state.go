package state

import "fmt"

type noItemState struct {
	vendingMachine *VendingMachine
}

func (i *noItemState) RequestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (i *noItemState) AddItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *noItemState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (i *noItemState) DispenseItem() error {
	return fmt.Errorf("Item out of stock")
}
