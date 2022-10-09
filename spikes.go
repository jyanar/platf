package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Spikes struct {
	sm *SceneManager // Could do this? That way interactions can be passed to the PlayState and be handled there.
	Obj
}

// But why not just have it be passed to the playstate itself?

func NewSpikes(sm *SceneManager, obj Obj) *Spikes {
	return &Spikes{sm, obj}
}

func (s Spikes) Update() error {
	return nil
}

func (s Spikes) Draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, s.x+s.w/2, s.y+s.h/2, s.w/2, image.White)
}

func (s Spikes) notify(msg string) {}

func (s Spikes) onTouch(other Entity) {
	// I guess with this way, you could technically send messages to other sections
	// of the world, that are not the current one.
	if typeof(other) == "*main.Player" {
		s.sm.getCurrent().processMsg("Player:kill")
	}
}

// Implemented through something like this
// func (s *Spikes) onTouch(other PosObj) {
// 	if typeof(other) == "Player" {
// 		s.sm.getCurrent().trigger("Player.kill")
// 	}
// }
