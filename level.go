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
	toggleFloors []*ToggleFloor
}

func (l *Level) init(width int, data []int, col *Collisions, player *Player) {
	l.width = width
	l.data = data
	l.tiles = []*Tile{}
	l.levers = []*Lever{}
	l.spikes = []*Spikes{}
	l.toggleFloors = []*ToggleFloor{}
	l.portals = []*Portal{}
	for i := 0; i < len(l.data); i++ {
		x, y := float64((i%l.width)*TILESIZE), math.Floor(float64(i)/float64(l.width))*TILESIZE
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

		case "Player":
			// Since player is already intialized (in playscene.go), we just shift
			// his position to whatever's specified in the map
			obj.w = 8
			obj.x += 4
			player.Obj = obj
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
	for i := range l.levers {
		l.portals[i].Draw(screen)
	}
}
