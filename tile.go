package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jyanar/platf/graphics"
)

type Tile struct {
	Object
}

func NewTile(Object Object) *Tile {
	return &Tile{Object}
}

func (t Tile) Solid() bool {
	return t.Object.Solid()
}

func (t Tile) Update(state *GameState) error { return nil }

func (t Tile) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.position.x, t.position.y)
	screen.DrawImage(graphics.Tile, op)
}
