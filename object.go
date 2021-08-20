package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Object struct {
	x, y, width, height float64
	clr color.Color
}

func (object *Object) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, object.x, object.y, object.width, object.height, object.clr)
}
