package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	width        int
	data         []int
	levers       []Lever
	toggleFloors []ToggleFloor
	spikes       []Spikes
	tiles        []Tile
	*Collisions
}

func (l *Level) init(width int, data []int, c *Collisions, player *Player) {
	l.width = width
	l.data = data
	l.Collisions = c
	l.levers = []Lever{}
	l.toggleFloors = []ToggleFloor{}
	l.spikes = []Spikes{}
	l.tiles = []Tile{}
	for i := 0; i < len(l.data); i++ {
		x, y := float64((i%l.width)*TILESIZE), math.Floor(float64(i)/float64(l.width))*TILESIZE
		obj := Obj{x: x, y: y, w: float64(TILESIZE), h: float64(TILESIZE)}
		switch TILETYPES[l.data[i]] {
		case "Tile":
			l.tiles = append(l.tiles, *NewTile(obj))
			ptr := &l.tiles[len(l.tiles)-1]
			l.Collisions.add(ptr)

		case "ToggleFloor":
			t := ToggleFloor{obj, true}
			l.toggleFloors = append(l.toggleFloors, t)
			ptr := &l.toggleFloors[len(l.toggleFloors)-1]
			l.Collisions.add(ptr)

		case "Spikes":
			l.spikes = append(l.spikes, *NewSpikes(obj))
			ptr := &l.spikes[len(l.spikes)-1]
			l.Collisions.add(ptr)

		case "Lever":
			l.levers = append(l.levers, *NewLever(obj))
			ptr := &l.levers[len(l.levers)-1]
			l.Collisions.add(ptr)

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

func (l Level) Draw(screen *ebiten.Image) {
	for i := range l.tiles {
		l.tiles[i].Draw(screen)
	}
	for i := range l.toggleFloors {
		l.toggleFloors[i].Draw(screen)
	}
	for i := range l.spikes {
		l.spikes[i].Draw(screen)
	}
	for i := range l.levers {
		l.levers[i].Draw(screen)
	}
}
