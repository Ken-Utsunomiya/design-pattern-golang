package adapter

import (
	"fmt"

	"github.com/projects/design-pattern-golang/Structure/Adapter/service"
)

type WindowsAdapter struct {
	WindowMachine *service.Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.WindowMachine.InsertIntoUSBPort()
}
