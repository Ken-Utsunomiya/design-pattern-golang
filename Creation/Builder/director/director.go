package director

import (
	"github.com/projects/design-pattern-golang/Creation/Builder/builder"
	"github.com/projects/design-pattern-golang/Creation/Builder/product"
)

type Director struct {
	Builder builder.IBuilder
}

func NewDirector(b builder.IBuilder) *Director {
	return &Director{
		Builder: b,
	}
}

func (d *Director) SetBuilder(b builder.IBuilder) {
	d.Builder = b
}

func (d *Director) BuildHouse() product.House {
	d.Builder.SetDoorType()
	d.Builder.SetWindowType()
	d.Builder.SetNumFloor()
	return d.Builder.GetHouse()
}
