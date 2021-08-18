package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var mplusNormalFont font.Face

type State int

const (
	Start State = iota  // the beginning of the game
	Serve				// waiting key press to start game
	Play				// the ball is being played
	Done				// the game is finished, one of players has won
)
type Game struct {
	player *Player
	ai *Player
	ball *Ball
	audioContext *audio.Context
	ballHitWall *audio.Player
	ballHitPaddle *audio.Player
	playerScored *audio.Player
	state State
}

func (g *Game) ExitGame() {
	os.Exit(0)
}

func (g *Game) init() {
	g.initPlayers()
	g.state = Start
}

func (g *Game) initPlayers() {
	rand.Seed(86)
	g.player = &Player {10, 40, 8, 40, color.White}
	g.ai = &Player {300, 40, 8, 40, color.White }
	g.ball = &Ball {160, 120, 5, 5, float64(rand.Intn(3)), float64(rand.Intn(3)), color.White}
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

func initFonts() {
	const dpi = 20
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	g := new(Game)
	g.init()

	initFonts()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}