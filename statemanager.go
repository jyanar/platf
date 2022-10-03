package main

import "github.com/hajimehoshi/ebiten/v2"

type State interface {
	initState(sm *StateManager)
	Update() error
	Draw(screen *ebiten.Image)
}

type StateManager struct {
	current State
}

func (sm *StateManager) Update() error {
	if res := sm.current.Update(); res != nil {
		return res
	}
	return nil
}

func (sm *StateManager) Draw(screen *ebiten.Image) {
	sm.current.Draw(screen)
}

func (sm *StateManager) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOWSIZE, WINDOWSIZE
}

func (sm StateManager) getCurrent() State {
	return sm.current
}

func (sm *StateManager) setCurrent(state State) {
	sm.current = state
	sm.current.initState(sm)
}
