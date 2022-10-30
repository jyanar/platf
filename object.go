package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Object struct {
	position Vector
	velocity Vector
	w, h     float64
	isSolid  bool
}

func (o Object) getPosAndSize() (float64, float64, float64, float64) {
	return o.position.x, o.position.y, o.w, o.h
}

func (o Object) Solid() bool {
	return o.isSolid
}

func (o Object) Update(state *GameState) error { return nil }

func (o Object) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, o.position.x, o.position.y, o.w, o.h, color.White)
}
