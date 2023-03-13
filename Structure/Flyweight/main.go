package main

import (
	"fmt"

	"github.com/projects/design-pattern-golang/Structure/Flyweight/factory"
)

func main() {
	game := newGame()

	//Add Terrorist
	game.addTerrorist(factory.TerroristDressType)
	game.addTerrorist(factory.TerroristDressType)
	game.addTerrorist(factory.TerroristDressType)
	game.addTerrorist(factory.TerroristDressType)

	//Add CounterTerrorist
	game.addCounterTerrorist(factory.CounterTerrroristDressType)
	game.addCounterTerrorist(factory.CounterTerrroristDressType)
	game.addCounterTerrorist(factory.CounterTerrroristDressType)

	dressFactoryInstance := factory.GetDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}
