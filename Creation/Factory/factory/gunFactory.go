package factory

import (
	"fmt"

	"github.com/projects/design-pattern-golang/Creation/Factory/product"
)

func GetGun(gunType string) (product.IGun, error) {
	switch gunType {
	case "ak47":
		return product.NewAk47(), nil
	case "musket":
		return product.NewMusket(), nil
	default:
		return nil, fmt.Errorf("wrong gun type passed")
	}
}
