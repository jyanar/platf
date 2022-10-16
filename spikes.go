package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jyanar/platf/graphics"
)

type Spikes struct {
	Obj
}

func NewSpikes(obj Obj) *Spikes {
	return &Spikes{obj}
}

func (s Spikes) Update(state *GameState) error { return nil }

func (s Spikes) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.x, s.y)
	screen.DrawImage(graphics.Tile, op)
	op.GeoM.Translate(0, -4)
	screen.DrawImage(graphics.Spikes, op)
}

func (s Spikes) Solid() bool {
	return true
}

// func (s Spikes) onTouch() {}
