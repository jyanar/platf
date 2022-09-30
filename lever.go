package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Lever struct {
	Obj
}

func (l Lever) Update() error { return nil }

func (l Lever) Draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, l.x+l.w/2, l.y+l.h/2, l.w/2, image.White)
}