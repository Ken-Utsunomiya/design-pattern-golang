package decorator

import "github.com/projects/design-pattern-golang/Structure/Decorator/component"

type CheeseTopping struct {
	Pizza component.IPizza
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 7
}
