package state

import "fmt"

type hasItemState struct {
	vendingMachine *VendingMachine
}

func (i *hasItemState) RequestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item requestded\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *hasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *hasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (i *hasItemState) DispenseItem() error {
	return fmt.Errorf("Please select item first")
}
