package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Spikes struct {
	Obj
}

func NewSpikes(obj Obj) *Spikes {
	return &Spikes{obj}
}

func (s Spikes) Update() error { return nil }

func (s Spikes) Draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, s.x+s.w/2, s.y+s.h/2, s.w/2, image.White)
}
