package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PauseState struct {
	sm *SceneManager
}

func (s *PauseState) processMsg(msg string) {}

func (s *PauseState) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		s.sm.pop()
	}
	return nil
}

func (s *PauseState) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "=== PAUSE SCREEN ===")
}

func (s *PauseState) initState(sm *SceneManager) {
	s.sm = sm
}
