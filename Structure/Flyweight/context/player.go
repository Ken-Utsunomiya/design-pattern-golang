package context

import (
	"github.com/projects/design-pattern-golang/Structure/Flyweight/factory"
	"github.com/projects/design-pattern-golang/Structure/Flyweight/object"
)

type Player struct {
	dress      object.Dress
	playerType string
	lat        int
	long       int
}

func NewPlayer(playerType, dressType string) *Player {
	dress, _ := factory.GetDressFactorySingleInstance().GetDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) NewLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
