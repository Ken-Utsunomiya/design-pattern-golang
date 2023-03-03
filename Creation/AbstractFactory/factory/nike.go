package factory

import "github.com/projects/design-pattern-golang/Creation/AbstractFactory/product"

// concrete factory

// Nike implements ISportsFactory
type Nike struct {
}

func (n *Nike) MakeShoe() product.IShoe {
	return &product.NikeShoe{
		Shoe: product.Shoe{
			Logo: "nike",
			Size: 14,
		},
	}
}

func (n *Nike) MakeShirt() product.IShirt {
	return &product.NikeShirt{
		Shirt: product.Shirt{
			Logo: "nike",
			Size: 14,
		},
	}
}
