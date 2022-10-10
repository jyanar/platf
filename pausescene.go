package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PauseScene struct {
	ebitenImage *ebiten.Image
}

func (s *PauseScene) Update(state *GameState) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state.SceneManager.pop()
	}
	return nil
}

func (s *PauseScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "=== PAUSE SCREEN ===")

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(20, 20)
	screen.DrawImage(s.ebitenImage, op)
}

func (s *PauseScene) init() {
	ebitenImage, err := getEbitenImage("tex.png")
	if err != nil {
		fmt.Println(err)
	}
	s.ebitenImage = ebitenImage
}

func (s *PauseScene) trigger(msg string) {}
