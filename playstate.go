package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PlayState struct {
	sm *StateManager
	World
	Player
	Level
}

func NewPlayState(sm *StateManager, w World, p Player, l Level) *PlayState {
	s := &PlayState{}
	s.sm = sm
	s.World = w
	s.Player = p
	s.Level = l
	return s
}

func (s *PlayState) Update() error {
	for _, item := range s.World.items {
		item.Update()
	}
	if !s.Player.alive {
		s.sm.push(&PauseState{})
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		s.sm.push(&PauseState{})
	}
	return nil
}

func (s *PlayState) Draw(screen *ebiten.Image) {
	for _, item := range s.World.items {
		item.Draw(screen)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (s *PlayState) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOWSIZE, WINDOWSIZE
}

func (s *PlayState) initState(sm *StateManager) {
	s.Player = Player{}
	s.World.init()
	s.Level.init(16, map1, &s.World, &s.Player)
	s.sm = sm
}
