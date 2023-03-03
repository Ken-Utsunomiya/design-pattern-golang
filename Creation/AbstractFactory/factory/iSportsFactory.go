package factory

import (
	"fmt"

	"github.com/projects/design-pattern-golang/Creation/AbstractFactory/product"
)

// Abstract factory interface

type ISportsFactory interface {
	MakeShoe() product.IShoe
	MakeShirt() product.IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	switch brand {
	case "adidas":
		return &Adidas{}, nil
	case "nike":
		return &Nike{}, nil
	default:
		return nil, fmt.Errorf("%s is unknown brand type", brand)
	}
}
