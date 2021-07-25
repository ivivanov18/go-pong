package main

import (
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player *Player
	ai *Player
}

func (g *Game) ExitGame() {
	os.Exit(0)
}

func (g *Game) Update() error {
	if (ebiten.IsKeyPressed(ebiten.KeyArrowDown)) {
		g.player.Update(ebiten.KeyArrowDown)
	} else if (ebiten.IsKeyPressed(ebiten.KeyArrowUp)) {
		g.player.Update(ebiten.KeyArrowUp)
	} else if (ebiten.IsKeyPressed(ebiten.KeyEscape)) {
		g.ExitGame()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "GO PONG")
	g.player.Draw(screen)
	g.ai.Draw(screen)
}

func (g *Game) Layout (outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	g := &Game {
		player: &Player {10, 40, 8, 40, color.White},
		ai: &Player {300, 40, 8, 40, color.White },
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}