package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PauseScene struct {
	sm *SceneManager
}

func (s *PauseScene) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		s.sm.pop()
	}
	return nil
}

func (s *PauseScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "=== PAUSE SCREEN ===")
}

func (s *PauseScene) initState(sm *SceneManager) {
	s.sm = sm
}
