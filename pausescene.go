package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PauseScene struct{}

func (s *PauseScene) Update(state *GameState) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state.SceneManager.pop()
	}
	return nil
}

func (s *PauseScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "=== PAUSE SCREEN ===")
}

func (s *PauseScene) init() {}
