package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ToggleFloor struct {
	Obj
}

func (t ToggleFloor) Update() error { return nil }

func (t ToggleFloor) Draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, t.x+t.w/2, t.y+t.h/2, t.w/2, image.White)
}