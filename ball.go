package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)
type Ball struct {
	x, y, height, width, velX, velY float64
	clr                 color.Color
}

func (ball *Ball) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, ball.x, ball.y, ball.width, ball.height, ball.clr)
}

//TODO replace magic values with constants
func (ball *Ball) Update() {
	ball.x += ball.velX
	ball.y += ball.velY

	if (ball.x < 0) {
		ball.x = 0
		ball.velX = -ball.velX
	} else if (ball.x + ball.width > 320) {
		ball.x = 320 - ball.width
		ball.velX = -ball.velX
	}
	if (ball.y < 0) {
		ball.y = 0
		ball.velY = -ball.velY
	} else if (ball.y + ball.height > 240) {
		ball.y = 240 - ball.height
		ball.velY = -ball.velY
	}
}

func (ball Ball) CollidesWith(player Player) bool {
	// left edge of either is futher to the right than right edge of other
	if ball.x > player.x + player.width || player.x > ball.x + ball.width {
		return false;
	}

	if ball.y > player.y + player.height || player.y > ball.y + ball.height {
		return false;
	}

	return true;
}