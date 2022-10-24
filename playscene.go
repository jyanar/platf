package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PlayScene struct {
	Player
	Level
	Collisions
	levelNum int
}

func (s *PlayScene) init() {
	if s.levelNum == 0 {
		s.levelNum = 1
	}
	s.Collisions.init()
	s.Player = *NewPlayer(Obj{}, &s.Collisions)
	s.Level.init(16, maps[s.levelNum], &s.Collisions, &s.Player)
}

func (s *PlayScene) Update(state *GameState) error {
	// Handle input
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state.SceneManager.push(&PauseScene{})
	}
	// Update collisions
	s.Player.Update(state)
	// Update environment
	s.Level.Update(state)
	// Check player status
	groundedObj := s.Player.NewGroundedObj()
	for i := range s.Level.spikes {
		if s.Collisions.areOverlapping(groundedObj, &s.Level.spikes[i].Obj) {
			s.Player.alive = false
		}
	}
	// Check if player has touched portal
	for i := range s.Level.portals {
		if s.Collisions.areOverlapping(&s.Player.Obj, &s.Level.portals[i].Obj) {
			s.levelNum += 1
			s.init()
		}
	}
	// Check if player has touched enemy
	for i := range s.Level.enemies {
		if s.Collisions.areOverlapping(&s.Player.Obj, &s.Level.enemies[i].Obj) {
			s.Player.alive = false
		}
	}
	if !s.Player.alive {
		state.SceneManager.push(&DeadScene{})
	}
	return nil
}

func (s *PlayScene) Draw(screen *ebiten.Image) {
	s.Player.Draw(screen)
	s.Level.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
	// drawGrid(screen)
	// ebitenutil.DrawRect(screen, s.Player.x, s.Player.y, s.Player.w, s.Player.h, color.White)
}

func (s *PlayScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOWSIZE, WINDOWSIZE
}

func (s *PlayScene) trigger(msg string) {
	switch msg {
	case "player:action":
		if s.Collisions.areOverlapping(&s.Player.Obj, &s.Level.levers[0].Obj) {
			for i := range s.Level.toggleFloors {
				s.Level.toggleFloors[i].toggleSolid()
			}
			s.Level.levers[0].toggle = !s.Level.levers[0].toggle
		}
	}
}
