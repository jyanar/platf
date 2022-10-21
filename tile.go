package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jyanar/platf/graphics"
)

type Tile struct {
	Obj
}

func NewTile(obj Obj) *Tile {
	return &Tile{obj}
}

func (t Tile) getPosition() (float64, float64) {
	return t.x, t.y
}

func (t *Tile) setPosition(x, y float64) {
	t.x = x
	t.y = y
}

func (t Tile) Solid() bool {
	return t.Obj.Solid()
}

func (t Tile) getPosAndSize() (float64, float64, float64, float64) {
	return t.x, t.y, t.w, t.h
}

func (t Tile) Update(state *GameState) error { return nil }

func (t Tile) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.x, t.y)
	screen.DrawImage(graphics.Tile, op)
}
