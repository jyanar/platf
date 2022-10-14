package main

import (
	// "fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PauseScene struct {
	assets Assets
}

func (s *PauseScene) Update(state *GameState) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state.SceneManager.pop()
	}
	return nil
}

func (s *PauseScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "=== PAUSE SCREEN ===")
	for i := 0; i < 20; i++ {
		s.assets.qdraw(screen, i, float64(i)*16, 16)
	}
}

func (s *PauseScene) init() {
	s.assets.init()
}

func (s *PauseScene) trigger(msg string) {}
