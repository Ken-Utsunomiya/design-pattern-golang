package main

import (
	"fmt"

	"github.com/projects/design-pattern-golang/Structure/Bridge/abstraction"
	"github.com/projects/design-pattern-golang/Structure/Bridge/implementation"
)

func main() {

	hpPrinter := &implementation.Hp{}
	epsonPrinter := &implementation.Epson{}

	macComputer := &abstraction.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &abstraction.Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
