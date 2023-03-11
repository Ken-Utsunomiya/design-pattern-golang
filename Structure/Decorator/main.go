package main

import (
	"fmt"

	"github.com/projects/design-pattern-golang/Structure/Decorator/component"
	"github.com/projects/design-pattern-golang/Structure/Decorator/decorator"
)

func main() {

	pizza := &component.VeggieMania{}

	//Add cheese topping
	pizzaWithCheese := &decorator.CheeseTopping{
		Pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &decorator.TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}
