package abstraction

import (
	"fmt"

	"github.com/projects/design-pattern-golang/Structure/Bridge/implementation"
)

type Mac struct {
	Printer implementation.Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.Printer.PrintFile()
}

func (m *Mac) SetPrinter(p implementation.Printer) {
	m.Printer = p
}
