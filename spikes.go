package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Spikes struct {
	Obj
	assets Assets
}

func NewSpikes(obj Obj) *Spikes {
	return &Spikes{obj, Assets{}}
}

func (s Spikes) Update(state *GameState) error { return nil }

func (s Spikes) Draw(screen *ebiten.Image) {
	s.assets.qdraw(screen, 1, s.x, s.y)
	s.assets.qdraw(screen, 2, s.x, s.y-4)
	// ebitenutil.DrawCircle(screen, s.x+s.w/2, s.y+s.h/2, s.w/2, image.White)
}

func (s Spikes) Solid() bool {
	return true
}

// func (s Spikes) onTouch() {}
