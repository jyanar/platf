package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	width        int
	data         []int
	tiles        []*Tile
	levers       []*Lever
	spikes       []*Spikes
	portals      []*Portal
	enemies      []*Enemy
	toggleFloors []*ToggleFloor
}

func (l *Level) init(width int, data []int, col *Collisions, player *Player) {
	l.width = width
	l.data = data
	l.tiles = []*Tile{}
	l.levers = []*Lever{}
	l.spikes = []*Spikes{}
	l.enemies = []*Enemy{}
	l.portals = []*Portal{}
	l.toggleFloors = []*ToggleFloor{}
	for i := 0; i < len(l.data); i++ {
		x := float64((i % l.width) * TILESIZE)
		y := math.Floor(float64(i)/float64(l.width)) * TILESIZE
		obj := Obj{x: x, y: y, w: float64(TILESIZE), h: float64(TILESIZE), isSolid: true}
		switch TILETYPES[l.data[i]] {
		case "Tile":
			l.tiles = append(l.tiles, NewTile(obj))
			col.add(&l.tiles[len(l.tiles)-1].Obj)

		case "Lever":
			obj.isSolid = false
			l.levers = append(l.levers, NewLever(obj))
			col.add(&l.levers[len(l.levers)-1].Obj)

		case "Spikes":
			l.spikes = append(l.spikes, NewSpikes(obj))
			col.add(&l.spikes[len(l.spikes)-1].Obj)

		case "ToggleFloor":
			l.toggleFloors = append(l.toggleFloors, NewToggleFloor(obj))
			col.add(&l.toggleFloors[len(l.toggleFloors)-1].Obj)

		case "Portal":
			l.portals = append(l.portals, NewPortal(obj))

		case "Enemy":
			l.enemies = append(l.enemies, NewEnemy(obj, col))
			col.add(&l.enemies[len(l.enemies)-1].Obj)

		case "Player":
			// Since player is already intialized (in playscene.go), we just shift
			// his position to whatever's specified in the map
			obj.w = 8
			obj.x += 4
			player.Obj = obj
		}
	}
}

func (l *Level) Update(state *GameState) error {
	for i := range l.enemies {
		err := l.enemies[i].Update(state)
		if err != nil {
			return err
		}
	}
	return nil
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
	for i := range l.portals {
		l.portals[i].Draw(screen)
	}
	for i := range l.enemies {
		l.enemies[i].Draw(screen)
	}
}
