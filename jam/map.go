package jam

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

func NewTileMap(sheets []*Sheet) *TileMap {
	return &TileMap{
		layers: []*TileLayer{&TileLayer{}},
		sheets: sheets,
	}
}
