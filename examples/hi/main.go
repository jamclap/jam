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
	sprites *jam.Sheet
	tileMap *jam.TileMap
}

func InitState(hub *jam.Hub) jam.Game {
	hub.Window.SetTitle("Jumping Demo for Jam")
	spriteSize := jam.XY(8, 8)
	sprites := jam.LoadSheet(sheetBytes, spriteSize)
	tmap := jam.LoadMap(tilesBytes, spriteSize)
	// TODO Uncomment for testing submaps.
	// *tmap = tmap.SliceStart(jam.XY(1, 1)).SliceSize(jam.XY(10, 10))
	tmap.Sheets = []*jam.Sheet{sprites}
	return &Game{
		faceX:   1,
		floored: false,
		move:    jam.XY(0, 0.0),
		pos:     extractPlayerPos(tmap),
		sprites: sprites,
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
	draw.Sprite(g.sprites.AtXY(1, int(g.frame)), jam.Pos(g.pos).ScaleX(g.faceX))
}

func (g *Game) atFloor(pos jam.Vec2f) bool {
	tileSize := g.tileMap.TileSize.Float64()
	tilePos := pos.Div(tileSize).Int()
	tile := g.tileMap.Tiles.At(tilePos)
	return tile == jam.NewTile(0, jam.XY(0, 0), 255) // TODO Tile tags.
}

func (g *Game) applyPhysics() {
	wasFloored := g.floored
	size := jam.Vec2Of[float64](g.sprites.SpriteSize())
	floor := 90.0
	// Fall if in the air.
	bottomLeft := g.pos.Add(size)
	tileSize := g.tileMap.TileSize.Float64()
	tilePos0 := bottomLeft.Div(tileSize).Int()
	aligned := g.move.Y == 0 && float64(tilePos0.Y)*tileSize.Y == bottomLeft.Y
	if !g.floored {
		g.move.Y += 0.2
	}
	g.pos = g.pos.Add(g.move)
	bottomLeft = g.pos.AddY(size.Y)
	// Check entering floor tile.
	tilePos1 := bottomLeft.Div(tileSize).Int()
	fellPast := g.move.Y > 0 && tilePos0.Y != tilePos1.Y
	if g.move.Y < 0 {
		g.floored = false
	} else if bottomLeft.Y > floor {
		g.floored = true
	} else if aligned || fellPast || g.move.X != 0 {
		g.floored = g.atFloor(bottomLeft.AddX(1)) ||
			g.atFloor(bottomLeft.AddX(size.X-2))
		if aligned || fellPast {
			floor = float64(tilePos1.Y * g.tileMap.TileSize.Y)
		} else {
			g.floored = g.floored && wasFloored
		}
	}
	// Go up if through the floor.
	if !wasFloored && g.floored {
		g.move.Y = 0
		excess := bottomLeft.Y - floor
		g.pos.Y -= excess
	}
	// Check walls.
	g.pos.X = max(g.pos.X, 0)
	g.pos.X = min(g.pos.X, 160-size.X)
}

func (g *Game) handleInput(c jam.Control) {
	if c.Active(jam.ActionA) {
		if c.Duration(jam.ActionA) < 20 && g.floored {
			g.move.Y = -2.3
		}
	} else {
		if g.move.Y < 0 {
			g.move.Y += 0.3
		}
	}
	speed := 1.5
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
