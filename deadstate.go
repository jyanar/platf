package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type DeadState struct {
	sm *StateManager
}

func (s *DeadState) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.sm.pop()
	}
	return nil
}

func (s *DeadState) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "=== YOU ARE DEAD ===\n")
	ebitenutil.DebugPrint(screen, "\n=== Press [enter] to restart ===")
}

func (s *DeadState) initState(sm *StateManager) {
	s.sm = sm
}
