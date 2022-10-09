package main

import (
	"math"
)

type Level struct {
	sm       *SceneManager
	cs       *CollisionSystem
	width    int
	data     []int
	entities []Entity
}

// func NewLevel(sm *SceneManager, cs *CollisionSystem, width int, data []int, player *Player) *Level {
// 	l := &Level{}
// 	l.width = width
// 	l.data = data
// 	l.world = world
// 	for i := 0; i < len(l.data); i++ {
// 		x, y := float64((i%l.width)*TILESIZE), math.Floor(float64(i)/float64(l.width))*TILESIZE
// 		obj := Obj{x: x, y: y, w: float64(TILESIZE), h: float64(TILESIZE)}
// 		var e Entity
// 		switch TILETYPES[l.data[i]] {
// 		case "Tile":
// 			e = NewTile(sm, obj)
// 		case "ToggleFloor":
// 			e = NewToggleFloor(obj)
// 		case "Spikes":
// 			e = NewSpikes(sm, obj)
// 		case "Lever":
// 			e = NewLever(obj)
// 		case "Player":
// 			e = player
// 		}
// 		l.cs.add(e)
// 		l.entities = append(l.entities, e)
// 	}
// 	return l
// }

func (l *Level) init(sm *SceneManager, cs *CollisionSystem, width int, data []int, player *Player) {
	l.sm = sm
	l.cs = cs
	l.width = width
	l.data = data
	for i := 0; i < len(l.data); i++ {
		x, y := float64((i%l.width)*TILESIZE), math.Floor(float64(i)/float64(l.width))*TILESIZE
		obj := Obj{x: x, y: y, w: float64(TILESIZE), h: float64(TILESIZE)}
		var e Entity
		switch TILETYPES[l.data[i]] {
		case "Tile":
			e = NewTile(sm, obj)
		case "ToggleFloor":
			e = NewToggleFloor(sm, obj)
		case "Spikes":
			e = NewSpikes(sm, obj)
		case "Lever":
			e = NewLever(sm, obj)
		case "Player":
			player.Obj = obj
			player.cs = cs
			player.vy = 0
			player.speed = 220
			player.alive = true
			e = player
		}
		l.entities = append(l.entities, e)
		l.cs.add(&e)
	}
}
