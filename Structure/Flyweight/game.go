package main

import "github.com/projects/design-pattern-golang/Structure/Flyweight/context"

type game struct {
	terrorists        []*context.Player
	counterTerrorists []*context.Player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*context.Player, 1),
		counterTerrorists: make([]*context.Player, 1),
	}
}

func (c *game) addTerrorist(dressType string) {
	player := context.NewPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
}

func (c *game) addCounterTerrorist(dressType string) {
	player := context.NewPlayer("CT", dressType)
	c.counterTerrorists = append(c.counterTerrorists, player)
}
