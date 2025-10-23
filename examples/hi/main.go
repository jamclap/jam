package main

import (
	_ "embed"
	"math"

	"github.com/jamclap/jam/jam"
	"github.com/jamclap/jam/jam/pal"
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
	tileMap *jam.TileMap
}

func InitState(hub *jam.Hub) jam.Game {
	hub.Window.SetTitle("Jumping Demo for Jam")
	spriteSize := jam.XY(8, 8)
	sprites := jam.LoadSheet(sheetBytes, spriteSize)
	tmap := jam.LoadMap(tilesBytes, spriteSize)
	hub.TileSheets = []*jam.Sheet{sprites}
	return &Game{
		faceX:   1,
		floored: false,
		move:    jam.XY(0, 0.0),
		pos:     extractPlayerPos(tmap),
		sprites: sprites,
		scale:   1,
		tileMap: tmap,
	}
}

func (g *Game) Update(hub *jam.Hub) {
	g.handleInput(hub.Control)
	g.applyPhysics()
	g.updateFrame()
}

func (g *Game) Draw(draw *jam.Draw) {
	draw.Fill(pal.Jam[pal.JamBlue1])
	draw.Map(g.tileMap, jam.MapOp{})
	draw.Sprite(
		// TODO Encourage [0.0, 1.0) indexed frame sequences?
		// TODO Identify by metadata files, ideally from editor.
		g.sprites.AtXY(1, int(g.frame)),
		jam.Pos(g.pos).Scale(g.scale).ScaleX(g.faceX),
	)
}

func (g *Game) applyPhysics() {
	size := jam.Vec2Of[float64](g.sprites.SpriteSize()).MulAll(g.scale)
	floor := 90.0
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
	g.pos.X = min(g.pos.X, 160-size.X)
}

func (g *Game) handleInput(c jam.Control) {
	if c.Active(jam.ActionA) {
		if c.Duration(jam.ActionA) < 20 && g.floored {
			g.move.Y = -5
		}
	} else {
		if g.move.Y < 0 {
			g.move.Y += 0.3
		}
	}
	speed := 2.5
	if c.Active(jam.ActionLeft) {
		g.faceX = -1
		g.move.X = -speed
	} else if c.Active(jam.ActionRight) {
		g.faceX = 1
		g.move.X = speed
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

// Replace the first instance of a player tile with empty, and
// return its position in pixels.
func extractPlayerPos(m *jam.TileMap) jam.Vec2f {
	size := m.Tiles.Size()
	p := jam.Vec2i{}
	for p.X = 0; p.X < size.X; p.X++ {
		for p.Y = 0; p.Y < size.Y; p.Y++ {
			tile := m.Tiles.At(p)
			if tile.Pos().X == 1 {
				m.Tiles.SetAt(p, jam.Tile{})
				return p.Mul(m.TileSize).Float64()
			}
		}
	}
	return jam.Vec2f{}
}

//go:embed sheet.webp
var sheetBytes []byte

//go:embed map.png
var tilesBytes []byte
