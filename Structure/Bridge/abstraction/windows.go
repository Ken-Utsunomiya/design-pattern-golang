package abstraction

import (
	"fmt"

	"github.com/projects/design-pattern-golang/Structure/Bridge/implementation"
)

type Windows struct {
	Printer implementation.Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.Printer.PrintFile()
}

func (w *Windows) SetPrinter(p implementation.Printer) {
	w.Printer = p
}
