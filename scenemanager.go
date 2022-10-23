package main

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	init()
	trigger(msg string)
	Update(state *GameState) error
	Draw(screen *ebiten.Image)
}

type SceneManager struct {
	scenes []Scene
	curidx int
}

type GameState struct {
	*SceneManager
}

func (sm *SceneManager) Update() error {
	return sm.scenes[sm.curidx].Update(&GameState{sm})
}

func (sm *SceneManager) Draw(screen *ebiten.Image) {
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
	sm.scenes[sm.curidx].init()
}

func (sm *SceneManager) pop() {
	sm.curidx = sm.curidx - 1
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
