package test

import (
	"design-pattern/adapter"
	"testing"
)

func TestAdapter(t *testing.T) {
	client := &adapter.Client{}
	mac := &adapter.Mac{}
	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &adapter.Windows{}
	windowsMachineAdapter := &adapter.WindowsAdapter{
		WindowsMachine: windowsMachine,
	}
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
