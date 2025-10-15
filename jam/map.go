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
	sheet uint16 // TODO Store as 8 bits + 2 * 12 bits for 32 bits as images?
	pos   Vec2[uint16]
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
	if sheet < 0 || sheet > maxSheet {
		log.Panicf("bad sheet index: %d", sheet)
	}
	if pos.X < 0 || pos.X > maxPos || pos.Y < 0 || pos.Y > maxPos {
		log.Panicf("bad pos: %v", pos)
	}
	return Tile{sheet: uint16(sheet), pos: Vec2Of[uint16](pos)}
}

func (t *Tile) Pos() Vec2i {
	return Vec2Of[int](t.pos)
}

func (t *Tile) Sheet() int {
	return int(t.sheet)
}

// Constrain extra in case we want to save in 32-bit values.
const maxPos = (1 << 12) - 1
const maxSheet = (1 << 8) - 1
