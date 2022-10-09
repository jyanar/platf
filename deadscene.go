package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type DeadScene struct{}

func (s *DeadScene) Update(state *GameState) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		state.SceneManager.pop()
		state.SceneManager.getCurrent().init()
	}
	return nil
}

func (s *DeadScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "=== YOU ARE DEAD ===\n")
	ebitenutil.DebugPrint(screen, "\n=== Press [enter] to restart ===")
}

func (s *DeadScene) init() {}
