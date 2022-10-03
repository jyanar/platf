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

func (s *PlayState) Update() error {
	s.World.Update()
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		s.sm.push(&PauseState{})
	}
	return nil
}

func (s *PlayState) Draw(screen *ebiten.Image) {
	s.World.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (s *PlayState) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOWSIZE, WINDOWSIZE
}

func (s *PlayState) initState(sm *StateManager) {
	s.World.init()
	s.Level.init(&s.World)
	s.sm = sm
}

// func (s *PlayState) trigger(event, actor, data)
