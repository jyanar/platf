package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PlayScene struct {
	sm *SceneManager
	World
	Player
	Level
}

func NewPlayState(sm *SceneManager, w World, p Player, l Level) *PlayScene {
	s := &PlayScene{}
	s.sm = sm
	s.World = w
	s.Player = p
	s.Level = l
	return s
}

func (s *PlayScene) Update() error {
	for _, item := range s.World.items {
		item.Update()
	}
	if !s.Player.alive {
		s.sm.push(&PauseScene{})
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		s.sm.push(&PauseScene{})
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

func (s *PlayScene) initState(sm *SceneManager) {
	s.Player = Player{}
	s.World.init()
	s.Level.init(16, map1, &s.World, &s.Player)
	s.sm = sm
}
