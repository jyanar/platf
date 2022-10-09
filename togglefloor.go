package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ToggleFloor struct {
	sm *SceneManager
	Obj
	isSolid bool
}

func NewToggleFloor(s *SceneManager, obj Obj) *ToggleFloor {
	return &ToggleFloor{s, obj, true}
}

func (t ToggleFloor) notify(msg string) {
	t.sm.getCurrent().processMsg(msg)
}

func (t ToggleFloor) Update() error { return nil }

func (t ToggleFloor) Draw(screen *ebiten.Image) {
	if t.isSolid {
		ebitenutil.DrawCircle(screen, t.x+t.w/2, t.y+t.h/2, t.w/2, image.White)
	} else {
		ebitenutil.DrawCircle(screen, t.x+t.w/2, t.y+t.h/2, t.w/2, image.White)
	}
}
