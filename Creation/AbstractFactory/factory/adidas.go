package factory

import "github.com/projects/design-pattern-golang/Creation/AbstractFactory/product"

// concrete factory

// Adidas implements ISportsFactory
type Adidas struct {
}

func (a *Adidas) MakeShoe() product.IShoe {
	return &product.AdidasShoe{
		Shoe: product.Shoe{
			Logo: "adidas",
			Size: 14,
		},
	}
}

func (a *Adidas) MakeShirt() product.IShirt {
	return &product.AdidasShirt{
		Shirt: product.Shirt{
			Logo: "adidas",
			Size: 14,
		},
	}
}
