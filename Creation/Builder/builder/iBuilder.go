package builder

import "github.com/projects/design-pattern-golang/Creation/Builder/product"

// builder interface

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() product.House
}

func GetBuilder(builderType string) IBuilder {
	switch builderType {
	case "normal":
		return NewNormalBuilder()
	case "igloo":
		return NewIglooBuilder()
	default:
		return nil
	}
}
