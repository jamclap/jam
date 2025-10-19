package jam

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"os"
)

type TileLayers struct {
	Layers []*TileMap
	Sheets []*Sheet
}

type TileMap struct {
	Offset   Vec2i
	Tiles    Grid[Tile]
	TileSize Vec2i
}

type Tile struct {
	opacity uint8
	sheet   uint8
	pos     Vec2[uint8]
}

func LoadTiles(r io.Reader) (grid Grid[Tile], err error) {
	img, err := png.Decode(r)
	nrgba, ok := img.(*image.NRGBA)
	if !ok {
		// Expect this is uncommon, so do some overhead when it happens.
		nrgba = image.NewNRGBA(img.Bounds())
		draw.Draw(nrgba, nrgba.Bounds(), img, image.Point{}, draw.Src)
	}
	if err != nil {
		return
	}
	size := Vec2i(nrgba.Bounds().Size())
	grid.SetSize(size)
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			xy := XY(x, y)
			c := nrgba.NRGBAAt(x, y)
			sourceXY := XY(int(c.G), int(c.B))
			tile := NewTile(int(c.R), sourceXY, uint8(c.A))
			grid.SetAt(xy, tile)
		}
	}
	return
}

func LoadTilesFromPath(path string) (grid Grid[Tile], err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	grid, err = LoadTiles(f)
	return
}

func NewTile(sheet int, pos Vec2i, opacity uint8) Tile {
	if sheet < 0 || sheet > maxTile {
		log.Panicf("bad sheet index: %d", sheet)
	}
	if pos.X < 0 || pos.X > maxTile || pos.Y < 0 || pos.Y > maxTile {
		log.Panicf("bad pos: %v", pos)
	}
	return Tile{sheet: uint8(sheet), pos: Vec2Of[uint8](pos), opacity: opacity}
}

func SaveTilesToPath(path string, grid Grid[Tile]) error {
	img := image.NewNRGBA(image.Rect(0, 0, grid.Size().X, grid.Size().Y))
	index := 0
	tiles := grid.Items()
	for y := 0; y < grid.Size().Y; y++ {
		for x := 0; x < grid.Size().X; x++ {
			tile := tiles[index]
			c := color.NRGBA{
				uint8(tile.Sheet()),
				uint8(tile.Pos().X),
				uint8(tile.Pos().Y),
				tile.opacity,
			}
			img.Set(x, y, c)
			index++
		}
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	// Saving a one-pixel transparent image with nativewebp crashed.
	// PMG also is better known and always lossless.
	return png.Encode(f, img)
}

func (t *Tile) Opacity() uint8 {
	return t.opacity
}

func (t *Tile) Pos() Vec2i {
	return Vec2Of[int](t.pos)
}

func (t *Tile) Sheet() int {
	return int(t.sheet)
}

const maxTile = (1 << 8) - 1
