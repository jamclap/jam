package main

import (
	_ "embed"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jamclap/jam/jam"
)

func main() {
	jam.Run(InitState)
}

type Game struct {
	faceX   float64
	floored bool
	frame   float64
	move    jam.Vec2f
	pos     jam.Vec2f
	scale   float64
	sprites *jam.Sheet
}

func InitState(hub *jam.Hub) jam.Game {
	return &Game{
		faceX:   1.0,
		floored: false,
		move:    jam.XY(0.0, 0.0),
		pos:     jam.XY(8, 8.0),
		sprites: jam.LoadSheet(spriteBytes, jam.XY(8, 8)),
		scale:   2,
	}
}

func (g *Game) Update(hub *jam.Hub) {
	g.handleInput()
	g.applyPhysics()
	g.updateFrame()
}

func (g *Game) Draw(draw *jam.Draw) {
	// TODO Include some standard palettes.
	draw.Fill(color.RGBA{0x2a, 0x3d, 0x74, 0xff})
	draw.Sprite(
		g.sprites.AtXY(0, int(g.frame)),
		jam.Pos(g.pos).Scale(g.scale).ScaleX(g.faceX),
	)
}

func (g *Game) applyPhysics() {
	size := jam.VecAs[float64](g.sprites.SpriteSize()).MulAll(g.scale)
	floor := 135.0
	// Fall if in the air.
	bottomLeft := g.pos.Add(size)
	if bottomLeft.Y < floor {
		g.move.Y += 0.2
	}
	g.pos = g.pos.Add(g.move)
	// Go up if through the floor.
	bottomLeft = g.pos.Add(size)
	excess := bottomLeft.Y - floor
	g.floored = excess >= 0
	if g.floored {
		g.move.Y = 0
		g.pos.Y -= excess
	}
	// Check walls.
	g.pos.X = max(g.pos.X, 0)
	g.pos.X = min(g.pos.X, 240-size.X)
}

func (g *Game) handleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && g.floored {
			g.move.Y = -6
		}
	} else {
		if g.move.Y < 0 {
			g.move.Y += 0.3
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.faceX = -1
		g.move.X = -3
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.faceX = 1
		g.move.X = 3
	} else {
		g.move.X = 0
	}
}

func (g *Game) updateFrame() {
	if g.move.Y < 0 {
		g.frame = 1
	} else if g.move.Y > 0 {
		g.frame = 2
	} else if g.move.X == 0 {
		g.frame = 0
	} else {
		g.frame = math.Mod((g.frame + 0.2), 3.0)
	}
}

//go:embed sprite.webp
var spriteBytes []byte
