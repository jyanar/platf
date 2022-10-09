package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// We could make it so that PlayState has three things. A CollisionHandler which is
// the world, which contains a list of the objects in the world and all the logic to
// move them, check for collisions, etc.
// Then there is a Player object -- this is just another object, but we want to have
// it exposed here because it's nice to have.
// Then there is the Level object. It is here that we read in the data of the level,
// generate all of the objects.

type PlayScene struct {
	Player
	Level
	CollisionSystem
}

func (s *PlayScene) Update(state *GameState) error {
	for _, e := range s.Level.entities {
		e.Update()
		// fmt.Println(typeof(e))
	}

}











type PlayState struct {
	*SceneManager
	CollisionSystem
	Player
	Level
}

func NewPlayState(sm *SceneManager, cs CollisionSystem, p Player, l Level) *PlayState {
	s := &PlayState{}
	s.SceneManager = sm
	s.CollisionSystem = cs
	s.Player = p
	s.Level = l
	return s
}

func (s *PlayState) initState(sm *SceneManager) {
	s.Player = Player{}
	s.CollisionSystem.init()
	s.Level.init(s.SceneManager, &s.CollisionSystem, 16, map1, &s.Player)
	s.SceneManager = sm
}

func (s *PlayState) processMsg(msg string) {
	switch msg {
	case "player:kill":
		s.Player.alive = false
	}
}

func (s *PlayState) Update() error {
	for _, e := range s.Level.entities {
		fmt.Println(typeof(e))
		e.Update()
	}
	if !s.Player.alive {
		s.SceneManager.push(&DeadState{})
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		s.SceneManager.push(&PauseState{})
	}
	return nil
}

func (s *PlayState) Draw(screen *ebiten.Image) {
	for _, e := range s.Level.entities {
		e.Draw(screen)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (s *PlayState) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOWSIZE, WINDOWSIZE
}
