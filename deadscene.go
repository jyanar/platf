package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type DeadScene struct {
	sm *SceneManager
}

func (s *DeadScene) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.sm.pop()
	}
	return nil
}

func (s *DeadScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "=== YOU ARE DEAD ===\n")
	ebitenutil.DebugPrint(screen, "\n=== Press [enter] to restart ===")
}

func (s *DeadScene) initState(sm *SceneManager) {
	s.sm = sm
}
