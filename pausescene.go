package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jyanar/platf/graphics"
)

type PauseScene struct{}

func (s *PauseScene) Update(state *GameState) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state.SceneManager.pop()
	}
	return nil
}

func (s *PauseScene) Draw(screen *ebiten.Image) {
	for i := 0; i < len(graphics.Quads); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(i)*16, 16)
		screen.DrawImage(graphics.Quads[i], op)
	}

	ebitenutil.DebugPrint(screen, "=== PAUSE SCREEN ===")
}

func (s *PauseScene) init() {}

func (s *PauseScene) trigger(msg string) {}
