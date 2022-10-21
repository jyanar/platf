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
}

func (l *Level) init(width int, data []int, c *Collisions, player *Player) {
	l.width = width
	l.data = data
	l.levers = []Lever{}
	l.toggleFloors = []ToggleFloor{}
	l.spikes = []Spikes{}
	l.tiles = []Tile{}
	for i := 0; i < len(l.data); i++ {
		x, y := float64((i%l.width)*TILESIZE), math.Floor(float64(i)/float64(l.width))*TILESIZE
		obj := Obj{x: x, y: y, w: float64(TILESIZE), h: float64(TILESIZE), isSolid: true}
		switch TILETYPES[l.data[i]] {
		case "Tile":
			l.tiles = append(l.tiles, Tile{obj})
			// l.tiles = append(l.tiles, *NewTile(obj))
			c.add(&l.tiles[len(l.tiles)-1].Obj)

		case "ToggleFloor":
			l.toggleFloors = append(l.toggleFloors, ToggleFloor{obj})
			// l.toggleFloors = append(l.toggleFloors, *NewToggleFloor(obj))
			// objpointer := &l.toggleFloors[len(l.toggleFloors)-1].Obj
			c.add(&l.toggleFloors[len(l.toggleFloors)-1].Obj)

		case "Spikes":
			l.spikes = append(l.spikes, Spikes{obj})
			c.add(&l.spikes[len(l.spikes)-1].Obj)

		case "Lever":
			obj.isSolid = false
			l.levers = append(l.levers, Lever{obj, false})
			c.add(&l.levers[len(l.levers)-1].Obj)

		case "Player":
			player.Obj = obj
			player.Collisions = c
			player.vy = 0
			player.speed = 220
			player.alive = true
			c.add(&player.Obj)
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
