package main

import (
	"image/color"
	"log"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player *Player
	ai *Player
	ball *Ball
}

func (g *Game) ExitGame() {
	os.Exit(0)
}

func (g *Game) Update() error {
	if (ebiten.IsKeyPressed(ebiten.KeyEscape)) {
		g.ExitGame()
	}
	g.player.Update()
	g.ball.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "GO PONG")
	g.player.Draw(screen)
	g.ai.Draw(screen)
	g.ball.Draw(screen)
}

func (g *Game) Layout (outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	rand.Seed(86)
	g := &Game {
		player: &Player {10, 40, 8, 40, color.White},
		ai: &Player {300, 40, 8, 40, color.White },
		ball: &Ball {160, 120, 5, 5, float64(rand.Intn(3)), float64(rand.Intn(3)), color.White},
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}