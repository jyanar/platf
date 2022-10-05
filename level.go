package main

import (
	"math"
)

type Level struct {
	width int
	data  []int
	world *World
}

func NewLevel(width int, data []int, world *World, player *Player) *Level {
	l := &Level{}
	l.width = width
	l.data = data
	l.world = world
	for i := 0; i < len(l.data); i++ {
		x, y := float64((i%l.width)*TILESIZE), math.Floor(float64(i)/float64(l.width))*TILESIZE
		obj := Obj{x: x, y: y, w: float64(TILESIZE), h: float64(TILESIZE)}
		switch TILETYPES[l.data[i]] {
		case "Tile":
			l.world.add(NewTile(obj))

		case "ToggleFloor":
			l.world.add(NewToggleFloor(obj))

		case "Spikes":
			l.world.add(NewSpikes(obj))

		case "Lever":
			l.world.add(NewLever(obj))

		case "Player":
			l.world.add(NewPlayer(obj, world, 0, 220, true))
		}
	}
	return l
}

func (l *Level) init(width int, data []int, world *World, player *Player) {
	l.width = width
	l.data = data
	l.world = world
	for i := 0; i < len(l.data); i++ {
		x, y := float64((i%l.width)*TILESIZE), math.Floor(float64(i)/float64(l.width))*TILESIZE
		obj := Obj{x: x, y: y, w: float64(TILESIZE), h: float64(TILESIZE)}
		switch TILETYPES[l.data[i]] {
		case "Tile":
			l.world.add(NewTile(obj))

		case "ToggleFloor":
			l.world.add(NewToggleFloor(obj))

		case "Spikes":
			l.world.add(NewSpikes(obj))

		case "Lever":
			l.world.add(NewLever(obj))

		case "Player":
			player.Obj = obj
			player.world = world
			player.vy = 0
			player.speed = 220
			player.alive = true
			l.world.add(player)
		}
	}
}