package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var mplusNormalFont font.Face

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
	if (g.ball.CollidesWith(*g.player)) {
		g.ball.velX = -g.ball.velX
		g.ball.x = g.player.x + g.player.width
	}
	if (g.ball.CollidesWith(*g.ai)) {
		g.ball.velX = -g.ball.velX
		g.ball.x = g.ai.x - g.ball.width
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var currentTPS string = fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS())
	ebitenutil.DebugPrint(screen, currentTPS)
	text.Draw(screen, "GO PONG", mplusNormalFont, 140, 10, color.White )
	g.player.Draw(screen)
	g.ai.Draw(screen)
	g.ball.Draw(screen)
}

func (g *Game) Layout (outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	const dpi = 20
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(86)

	g := &Game {
		player: &Player {10, 40, 8, 40, color.White},
		ai: &Player {300, 40, 8, 40, color.White },
		ball: &Ball {160, 120, 5, 5, float64(rand.Intn(3)), float64(rand.Intn(3)), color.White},
	}

	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}