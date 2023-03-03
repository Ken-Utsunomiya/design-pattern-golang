package main

import (
	"fmt"

	"github.com/projects/design-pattern-golang/Creation/AbstractFactory/factory"
	"github.com/projects/design-pattern-golang/Creation/AbstractFactory/product"
)

func main() {
	adidasFactory, _ := factory.GetSportsFactory("adidas")
	nikeFactory, _ := factory.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s product.IShoe) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}

func printShirtDetails(s product.IShirt) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}
