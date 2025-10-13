package jam

import "log"

type TileLayer struct {
	hidden   bool
	offset   Vec2i
	tiles    Grid[Tile]
	tileSize Vec2i
}

type Tile struct {
	sheet uint16 // TODO Store as 8 bits + 2 * 12 bits for 32 bits as images?
	pos   Vec2[uint16]
}

type TileMap struct {
	layers []*TileLayer
	sheets []*Sheet
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

func (l *TileLayer) TileSize() Vec2i {
	return l.tileSize
}
