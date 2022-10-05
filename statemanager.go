package main

import "github.com/hajimehoshi/ebiten/v2"

type State interface {
	initState(sm *StateManager)
	Update() error
	Draw(screen *ebiten.Image)
}

type StateManager struct {
	states []State
	curidx int
}

func (sm *StateManager) Update() error {
	if res := sm.states[sm.curidx].Update(); res != nil {
		return res
	}
	return nil
}

func (sm *StateManager) Draw(screen *ebiten.Image) {
	sm.states[sm.curidx].Draw(screen)
}

func (sm *StateManager) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOWSIZE, WINDOWSIZE
}

func (sm StateManager) getCurrent() State {
	return sm.states[sm.curidx]
}

func (sm *StateManager) push(state State) {
	sm.states = append(sm.states, state)
	sm.curidx = len(sm.states) - 1
	sm.states[sm.curidx].initState(sm)
}

func (sm *StateManager) pop() {
	sm.curidx = len(sm.states) - 2
	sm.states = sm.states[:len(sm.states)-1] // discard last state
}

// func (sm *StateManager) setCurrent(idx int) {
// 	sm.fromidx = sm.curidx
// 	sm.curidx = idx
// }

// func (sm *StateManager) returnToPrevious() {
// 	tmp := sm.curidx
// 	sm.curidx = sm.fromidx
// 	sm.fromidx = tmp
// }
