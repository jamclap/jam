package jam

import (
	"image"
	"image/color"
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
	sheet uint8
	pos   Vec2[uint8]
}

func LoadTiles(r io.Reader) (grid Grid[Tile], err error) {
	image, err := png.Decode(r)
	if err != nil {
		return
	}
	size := Vec2i(image.Bounds().Size())
	grid.SetSize(size)
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			xy := XY(x, y)
			sheet, sourceX, sourceY, _ := image.At(x, y).RGBA()
			sourceXY := XY(int(sourceX>>8), int(sourceY>>8))
			tile := NewTile(int(sheet>>8), sourceXY)
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

func NewTile(sheet int, pos Vec2i) Tile {
	if sheet < 0 || sheet > maxTile {
		log.Panicf("bad sheet index: %d", sheet)
	}
	if pos.X < 0 || pos.X > maxTile || pos.Y < 0 || pos.Y > maxTile {
		log.Panicf("bad pos: %v", pos)
	}
	return Tile{sheet: uint8(sheet), pos: Vec2Of[uint8](pos)}
}

func SaveTilesToPath(path string, grid Grid[Tile]) error {
	img := image.NewRGBA(image.Rect(0, 0, grid.Size().X, grid.Size().Y))
	index := 0
	tiles := grid.Items()
	for y := 0; y < grid.Size().Y; y++ {
		for x := 0; x < grid.Size().X; x++ {
			tile := tiles[index]
			c := color.RGBA{
				uint8(tile.Sheet()),
				uint8(tile.Pos().X),
				uint8(tile.Pos().Y),
				255,
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

func (t *Tile) Pos() Vec2i {
	return Vec2Of[int](t.pos)
}

func (t *Tile) Sheet() int {
	return int(t.sheet)
}

const maxTile = (1 << 8) - 1
