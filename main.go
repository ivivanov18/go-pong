package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const scoreToWin = 3

var mplusNormalFont font.Face
var menuFont font.Face

type State int

const (
	Scored = iota  // the beginning of the game
	Play				// the ball is being played
	Win 				// the game is finished, one of players has won
	NewGame
	Winner
)

const (
	Player = iota
	Ai
)

type Score struct {
	player int
	ai int
}

type Game struct {
	player *ControlledPlayer
	ai *AiPlayer
	ball *Ball
	audioContext *audio.Context
	ballHitWall *audio.Player
	ballHitPaddle *audio.Player
	playerScored *audio.Player
	state State
	score Score
}

func (g *Game) ExitGame() {
	os.Exit(0)
}

func (g *Game) init() {
	g.initPlayers()
	g.score = Score{0, 0}
	g.state = NewGame 
}

func (g *Game) initPlayers() {
	rand.Seed(86)
	g.player = &ControlledPlayer{Object{10, 40, 8, 40, color.White}}
	g.ai = &AiPlayer{Object{300, 40, 8, 40, color.White} }
	g.ball = &Ball {160, 120, 5, 5, float64(rand.Intn(3)), float64(rand.Intn(3)), color.White}
}

func DisplayMenu(screen *ebiten.Image) {
	text.Draw(screen, "NEW GAME - PRESS SPACE", menuFont, 70, 130, color.White )
	text.Draw(screen, "QUIT GAME - PRESS ESC", menuFont, 70, 150, color.White )
}

func DisplayWinner(screen *ebiten.Image, winner int) {
	if (winner == Player) {
		text.Draw(screen, "YOU WON!!!", menuFont, 100, 110, color.White )
	}else {
		text.Draw(screen, "THE OTHER PLAYER WON!", menuFont, 70, 110, color.White )
	}
	DisplayMenu(screen)
}

func (g *Game) Update() error {
	if (ebiten.IsKeyPressed(ebiten.KeyEscape)) {
		g.ExitGame()
	}
	if (ebiten.IsKeyPressed(ebiten.KeySpace) && g.state == NewGame) {
		g.state = Play
	} else if (ebiten.IsKeyPressed(ebiten.KeySpace) && g.state == Winner) {
		g.state = Play
		g.score = Score{0,0}
	}
	if (g.state == Scored) {
			g.ball.x = 160
			g.ball.y = 120
		if (ebiten.IsKeyPressed(ebiten.KeySpace)) {
			g.state = Play
			rand.Seed(86)
			g.ball.velX = float64(rand.Intn(3))
			g.ball.velY = float64(rand.Intn(3))
		}
	} else if (g.state == Play) {
		g.player.Update()
		g.ball.Update()
		if (g.ball.CollidesWith(g.player.object)) {
			g.ball.velX = -g.ball.velX
			g.ball.x = g.player.object.x + g.player.object.width
		}
		if (g.ball.CollidesWith(g.ai.object)) {
			g.ball.velX = -g.ball.velX
			g.ball.x = g.ai.object.x - g.ball.width
		}
		if (g.ball.x <= 0) {
			g.score.ai += 1	
			g.state = Scored
		}
		if (g.ball.x >=  315) {
			g.score.player += 1
			g.state = Scored
		}
		//TODO: change with constant
		if (g.score.player == scoreToWin || g.score.ai == scoreToWin) {
			g.state = Winner
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if ( g.state == NewGame) {
		DisplayMenu(screen)
	} else if (g.state == Winner) {
		var winner int
		if (g.score.player == scoreToWin) {
			winner = Player
		} else if (g.score.ai == scoreToWin) {
			winner = Ai
		}
		DisplayWinner(screen, winner)
	} else {
		var currentTPS string = fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS())
		ebitenutil.DebugPrint(screen, currentTPS)
		text.Draw(screen, "GO PONG", mplusNormalFont, 140, 10, color.White )
		text.Draw(screen, strconv.Itoa(g.score.player), mplusNormalFont, 150, 25, color.White)
		text.Draw(screen, "-", mplusNormalFont, 155, 25, color.White)
		text.Draw(screen, strconv.Itoa(g.score.ai), mplusNormalFont, 160, 25, color.White)
		text.Draw(screen, strconv.FormatFloat(g.ball.x,'f',-1, 64), mplusNormalFont,10, 25, color.White)
		g.player.object.Draw(screen)
		g.ai.object.Draw(screen)
		g.ball.Draw(screen)
	}
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
		Size:    34,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	menuFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    60,
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
