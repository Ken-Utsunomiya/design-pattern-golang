package main

import (
	"fmt"

	"github.com/projects/design-pattern-golang/Creation/Factory/factory"
	"github.com/projects/design-pattern-golang/Creation/Factory/product"
)

func main() {
	ak47, _ := factory.GetGun("ak47")
	musket, _ := factory.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g product.IGun) {
	fmt.Printf("Gun: %s\n", g.GetName())
	fmt.Printf("Power: %d\n", g.GetPower())
}
