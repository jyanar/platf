package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PlayScene struct {
	World
	Player
	Level
}

func NewPlayScene(w World, p Player, l Level) *PlayScene {
	s := &PlayScene{}
	s.World = w
	s.Player = p
	s.Level = l
	return s
}

func (s *PlayScene) Update(state *GameState) error {
	// Handle input
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state.SceneManager.push(&PauseScene{})
	}
	// Update world
	for _, item := range s.World.items {
		item.Update()
	}
	// Check player status
	if !s.Player.alive {
		state.SceneManager.push(&DeadScene{})
	}
	return nil
}

func (s *PlayScene) Draw(screen *ebiten.Image) {
	for _, item := range s.World.items {
		item.Draw(screen)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (s *PlayScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOWSIZE, WINDOWSIZE
}

func (s *PlayScene) init() {
	s.Player = Player{}
	s.World.init()
	s.Level.init(16, map1, &s.World, &s.Player)
}
