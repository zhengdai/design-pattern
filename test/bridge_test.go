package test

import (
	"design-pattern/bridge"
	"testing"
)

func TestBridge(t *testing.T) {
	hpPrinter := &bridge.Hp{}
	epsonPrinter := &bridge.Epson{}

	mac := &bridge.Mac{}
	windows := &bridge.Windows{}

	mac.SetPrinter(hpPrinter)
	mac.Print()

	mac.SetPrinter(epsonPrinter)
	mac.Print()

	windows.SetPrinter(hpPrinter)
	windows.Print()

	windows.SetPrinter(epsonPrinter)
	windows.Print()
}
