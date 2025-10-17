package jam

import (
	"log"
)

type TileMap struct {
	layers []*TileLayer
	sheets []*Sheet
}

type TileLayer struct {
	Hidden   bool
	Offset   Vec2i
	Tiles    Grid[Tile]
	TileSize Vec2i
}

type Tile struct {
	sheet uint8
	pos   Vec2[uint8]
}

func NewTileMap() *TileMap {
	return &TileMap{
		layers: []*TileLayer{&TileLayer{}},
	}
}

func (m *TileMap) Layers() []*TileLayer {
	return m.layers
}

// For extending by only one at a time at most.
func (m *TileMap) SetSheet(index int, sheet *Sheet) {
	if index > len(m.sheets) {
		log.Panicf("index too high: %d vs %d", index, len(m.sheets))
	}
	if index == len(m.sheets) {
		m.sheets = append(m.sheets, sheet)
	} else {
		m.sheets[index] = sheet
	}
}

func (m *TileMap) Sheets() []*Sheet {
	return m.sheets
}

func NewTile(sheet int, pos Vec2i) Tile {
	if sheet < 0 || sheet > maxTile {
		log.Panicf("bad sheet index: %d", sheet)
	}
	if pos.X < 0 || pos.X > maxTile || pos.Y < 0 || pos.Y > maxTile {
		log.Panicf("bad pos: %v", pos)
	}
	return Tile{sheet: uint8(sheet), pos: Vec2Of[uint8](pos)}
}

func (t *Tile) Pos() Vec2i {
	return Vec2Of[int](t.pos)
}

func (t *Tile) Sheet() int {
	return int(t.sheet)
}

const maxTile = (1 << 8) - 1
