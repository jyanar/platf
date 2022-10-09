package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Alright, so this is how it works. We have a:
// - SceneManager (struct)
// - Scene (interface)
// - GameState (struct)
//
// The SceneManager stores the current scene, and calls the update and draw
// functions on it. In turn, any given scene that we implement (which satisfies
// the Scene interface) needs to implement those Draw() and Update() methods.
// Note that the Scene interface is defined as such:
//
// 		type Scene interface {
//			Update(state *GameState) error
//			Draw(screen *ebiten.Image)
//      }
//
// GameState, which is also a struct, merely carries a pointer to the SceneManager.
// In this way, individual Scenes (such as PlayScene, GameScene, etc) can reach the
// SceneManager.
//
// For instance, in ebiten's examples/blocks/gamepadscene.go, we have the following
// Update() method:
//
// 		func (s *GamePadScene) Update(state *GameState) error {
//			...
//			if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
//				...
//				state.SceneManager.GoTo(&TitleScene{})
//				return nil
//			}
//			...
//		}
//
// In this way, code within GamePadScene can call functions in SceneManager.
// However, do we even need GameState? Can't we just do something like
//
//		type Scene interface {
//			Update(sm *SceneManager) error
//			Draw(screen *ebiten.Image)
//		}
//
//		func (sm *SceneManager) Update() error {
//			return s.scenes[curidx].Update(sm)
//		}
//

type GameState struct {
	SceneManager *SceneManager
}

type Scene interface {
	Update(state *GameState) error
	Draw(screen *ebiten.Image)
}

type SceneManager struct {
	scenes []Scene
	curidx int
}

func (sm *SceneManager) Update() error {
	return sm.scenes[sm.curidx].Update(&GameState{sm})
	// if res := sm.scenes[sm.curidx].Update(); res != nil {
	// 	return res
	// }
	// return nil
}

func (sm *SceneManager) Draw(screen *ebiten.Image) {
	// sm.states[sm.curidx].Draw(screen)
	sm.scenes[sm.curidx].Draw(screen)
}

func (sm *SceneManager) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOWSIZE, WINDOWSIZE
}

func (sm SceneManager) getCurrent() Scene {
	return sm.scenes[sm.curidx]
}

func (sm *SceneManager) push(scene Scene) {
	sm.scenes = append(sm.scenes, scene)
	sm.curidx = len(sm.scenes) - 1
	// sm.states[sm.curidx].initState(sm)
}

func (sm *SceneManager) pop() {
	sm.curidx = len(sm.scenes) - 2
	sm.scenes = sm.scenes[:len(sm.scenes)-1] // discard last state
}

// func (sm *SceneManager) setCurrent(idx int) {
// 	sm.fromidx = sm.curidx
// 	sm.curidx = idx
// }

// func (sm *SceneManager) returnToPrevious() {
// 	tmp := sm.curidx
// 	sm.curidx = sm.fromidx
// 	sm.fromidx = tmp
// }

// type GameState interface {
// 	initState(sm *SceneManager)
// 	processMsg(msg string)
// 	Update() error
// 	Draw(screen *ebiten.Image)
// }