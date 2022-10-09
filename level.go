package main

import (
	"math"
)

type Level struct {
	width int
	data  []int
	*Collisions
}

func NewLevel(width int, data []int, collisions *Collisions, player *Player) *Level {
	l := &Level{}
	l.width = width
	l.data = data
	l.Collisions = collisions
	for i := 0; i < len(l.data); i++ {
		x, y := float64((i%l.width)*TILESIZE), math.Floor(float64(i)/float64(l.width))*TILESIZE
		obj := Obj{x: x, y: y, w: float64(TILESIZE), h: float64(TILESIZE)}
		switch TILETYPES[l.data[i]] {
		case "Tile":
			l.Collisions.add(NewTile(obj))

		case "ToggleFloor":
			l.Collisions.add(NewToggleFloor(obj))

		case "Spikes":
			l.Collisions.add(NewSpikes(obj))

		case "Lever":
			l.Collisions.add(NewLever(obj))

		case "Player":
			l.Collisions.add(NewPlayer(obj, collisions, 0, 220, true))
		}
	}
	return l
}

func (l *Level) init(width int, data []int, c *Collisions, player *Player) {
	l.width = width
	l.data = data
	l.Collisions = c
	for i := 0; i < len(l.data); i++ {
		x, y := float64((i%l.width)*TILESIZE), math.Floor(float64(i)/float64(l.width))*TILESIZE
		obj := Obj{x: x, y: y, w: float64(TILESIZE), h: float64(TILESIZE)}
		switch TILETYPES[l.data[i]] {
		case "Tile":
			l.Collisions.add(NewTile(obj))

		case "ToggleFloor":
			l.Collisions.add(NewToggleFloor(obj))

		case "Spikes":
			l.Collisions.add(NewSpikes(obj))

		case "Lever":
			l.Collisions.add(NewLever(obj))

		case "Player":
			player.Obj = obj
			player.Collisions = c
			player.vy = 0
			player.speed = 220
			player.alive = true
			l.Collisions.add(player)
		}
	}
}
