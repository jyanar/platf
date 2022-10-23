package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	width        int
	data         []int
	levers       []*Lever
	toggleFloors []*ToggleFloor
	spikes       []*Spikes
	tiles        []*Tile
}

func (l *Level) init(width int, data []int, c *Collisions, player *Player) {
	l.width = width
	l.data = data
	l.levers = []*Lever{}
	l.toggleFloors = []*ToggleFloor{}
	l.spikes = []*Spikes{}
	l.tiles = []*Tile{}
	for i := 0; i < len(l.data); i++ {
		x, y := float64((i%l.width)*TILESIZE), math.Floor(float64(i)/float64(l.width))*TILESIZE
		obj := Obj{x: x, y: y, w: float64(TILESIZE), h: float64(TILESIZE), isSolid: true}
		switch TILETYPES[l.data[i]] {
		case "Tile":
			l.tiles = append(l.tiles, NewTile(obj))
			c.add(&l.tiles[len(l.tiles)-1].Obj)

		case "ToggleFloor":
			l.toggleFloors = append(l.toggleFloors, NewToggleFloor(obj))
			c.add(&l.toggleFloors[len(l.toggleFloors)-1].Obj)

		case "Spikes":
			l.spikes = append(l.spikes, NewSpikes(obj))
			c.add(&l.spikes[len(l.spikes)-1].Obj)

		case "Lever":
			obj.isSolid = false
			l.levers = append(l.levers, NewLever(obj))
			c.add(&l.levers[len(l.levers)-1].Obj)

		case "Player":
			obj.w = 8
			obj.x += 4
			// obj.x = obj.x
			player.Obj = obj
			player.Collisions = c
			player.velocity = Vector{0, 0}
			player.speed = 220
			player.alive = true
			player.anim.Init(9, 2, 0.5)
			// player.anim.t = 0
			// player.anim.start = 9
			// player.anim.len = 3
			// player.anim.duration = 0.5
			// 	t:        0,
			// 	start:    9,
			// 	len:      3,
			// 	duration: 0.5,
			// }
			// player.anim = *graphics.NewAnimation()
			// player.anim.t = 0
			// player.anim.start
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
