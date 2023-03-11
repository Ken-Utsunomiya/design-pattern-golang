package decorator

import "github.com/projects/design-pattern-golang/Structure/Decorator/component"

type TomatoTopping struct {
	Pizza component.IPizza
}

func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 7
}
