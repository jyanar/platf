package main

const (
	DT          = 0.1
	WINDOWSIZE  = 256
	WINDOWSCALE = 3
	TILESIZE    = 16
)

const (
	GRAVITY        = 5
	JUMP_SPEED     = -6.5
	MAX_SPEED      = 25
	ACCELERATION   = 15
	MAX_FALL_SPEED = GRAVITY
	FRICTION       = 8
)

var TILETYPES = map[int]string{
	1: "Tile",
	2: "ToggleFloor",
	5: "Spikes",
	6: "Lever",
	7: "Enemy",
	8: "Portal",
	9: "Player",
}
