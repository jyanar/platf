package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PlayScene struct {
	Collisions
	Player
	Level
}

func NewPlayScene(w Collisions, p Player, l Level) *PlayScene {
	s := &PlayScene{}
	s.Collisions = w
	s.Player = p
	s.Level = l
	return s
}

func (s *PlayScene) init() {
	s.Player = Player{}
	s.Collisions.init()
	s.Level.init(16, map2, &s.Collisions, &s.Player)
}

func (s *PlayScene) Update(state *GameState) error {
	// Handle input
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state.SceneManager.push(&PauseScene{})
	}
	// Update collisions
	s.Player.Update(state)
	// Check player status
	groundedObj := s.Player.NewGroundedObj()
	for i := range s.Level.spikes {
		if s.Collisions.areOverlapping(groundedObj, &s.Level.spikes[i].Obj) {
			s.Player.alive = false
		}
	}
	if !s.Player.alive {
		state.SceneManager.push(&DeadScene{})
	}
	return nil
}

func (s *PlayScene) Draw(screen *ebiten.Image) {
	s.Level.Draw(screen)
	s.Player.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
	drawGrid(screen)
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
