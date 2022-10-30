package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jyanar/platf/graphics"
)

type Spikes struct {
	Object
}

func NewSpikes(Object Object) *Spikes {
	return &Spikes{Object}
}

func (s Spikes) Update(state *GameState) error { return nil }

func (s Spikes) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.position.x, s.position.y)
	screen.DrawImage(graphics.Tile, op)
	op.GeoM.Translate(0, -4)
	screen.DrawImage(graphics.Spikes, op)
}

func (s Spikes) Solid() bool {
	return s.Object.Solid()
}

// func (s Spikes) onTouch() {}
