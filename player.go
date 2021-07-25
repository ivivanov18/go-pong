package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	x, y, width, height float64
	clr color.Color
}

func (player *Player) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, player.x, player.y, player.width, player.height, player.clr)
}

//TODO replace magic number 240 with adequate const
func (player *Player) Update(key ebiten.Key) {
	switch(key) {
	case ebiten.KeyArrowUp:
		player.y = math.Max(player.y - 2, 0)
	case ebiten.KeyArrowDown:
		player.y = math.Min(player.y +2, 240 - player.height )
	}
}