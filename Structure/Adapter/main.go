package main

import (
	"github.com/projects/design-pattern-golang/Structure/Adapter/adapter"
	"github.com/projects/design-pattern-golang/Structure/Adapter/client"
	"github.com/projects/design-pattern-golang/Structure/Adapter/service"
)

func main() {

	client := &client.Client{}
	mac := &service.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &service.Windows{}
	windowsMachineAdapter := &adapter.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
