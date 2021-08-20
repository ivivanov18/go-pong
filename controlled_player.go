package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type ControlledPlayer struct {
	object Object 
}

//TODO replace magic number 240 with adequate const
func (player *ControlledPlayer) Update() {
	if (ebiten.IsKeyPressed(ebiten.KeyArrowDown)) {
		player.object.y = math.Min(player.object.y +2, 240 - player.object.height )
	} else if (ebiten.IsKeyPressed(ebiten.KeyArrowUp)) {
		player.object.y = math.Max(player.object.y - 2, 0)
	}
}
