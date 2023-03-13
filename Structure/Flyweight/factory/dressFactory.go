package factory

import (
	"fmt"

	"github.com/projects/design-pattern-golang/Structure/Flyweight/object"
)

const (
	//TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	//CounterTerrroristDressType terrorist dress type
	CounterTerrroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		DressMap: make(map[string]object.Dress),
	}
)

type DressFactory struct {
	DressMap map[string]object.Dress
}

func (d *DressFactory) GetDressByType(dressType string) (object.Dress, error) {
	if d.DressMap[dressType] != nil {
		return d.DressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.DressMap[dressType] = object.NewTerroristDress()
		return d.DressMap[dressType], nil
	}
	if dressType == CounterTerrroristDressType {
		d.DressMap[dressType] = object.NewCounterTerroristDress()
		return d.DressMap[dressType], nil
	}

	return nil, fmt.Errorf("wrong dress type passed")
}

func GetDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}
