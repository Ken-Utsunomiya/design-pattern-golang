package abstraction

import "github.com/projects/design-pattern-golang/Structure/Bridge/implementation"

type Computer interface {
	Print()
	SetPrinter(implementation.Printer)
}
