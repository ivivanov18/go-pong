package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	x, y, height, width float64
	clr                 color.Color
}

func (ball *Ball) Draw(screen *ebiten.Image) {

}

func (ball *Ball) Update(key ebiten.Key) {

}