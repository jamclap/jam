package jam

import (
	"bytes"
	"image"
	"log"

	"github.com/HugoSmits86/nativewebp"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sheet struct {
	image      *ebiten.Image
	spriteSize Vec2i
	sprites    Grid[*ebiten.Image]
	tags       Grid[uint32]
}

type SheetInfo struct {
	Image      *ebiten.Image
	SpriteSize Vec2i
	Tags       Grid[uint32]
}

// Excludes any partial space at edges.
func NewSheet(info SheetInfo) *Sheet {
	imageSize := info.Image.Bounds().Size()
	gridSize := Vec2i{}.AddPoint(imageSize).Div(info.SpriteSize)
	return &Sheet{
		image:      info.Image,
		spriteSize: info.SpriteSize,
		sprites:    NewGrid[*ebiten.Image](gridSize),
		tags:       info.Tags,
	}
}

func (s *Sheet) At(xy Vec2i) *ebiten.Image {
	gridSize := s.sprites.size
	if xy.X < 0 || xy.X >= gridSize.X || xy.Y < 0 || xy.Y >= gridSize.Y {
		return nil
	}
	sprite := s.sprites.At(xy)
	if sprite == nil {
		pixelStart := xy.Mul(s.spriteSize)
		pixelEnd := pixelStart.Add(s.spriteSize)
		r := image.Rect(pixelStart.X, pixelStart.Y, pixelEnd.X, pixelEnd.Y)
		sprite = s.image.SubImage(r).(*ebiten.Image)
		s.sprites.SetAt(xy, sprite)
	}
	return sprite
}

func (s *Sheet) AtXY(x, y int) *ebiten.Image {
	return s.At(XY(x, y))
}

func (s *Sheet) SpriteSize() Vec2i {
	return s.spriteSize
}

// Panics on fail, presuming the bytes are known good.
// TODO Pass in object with both sprite size and tags.
func LoadSheet(b []byte, spriteSize Vec2i) *Sheet {
	// TODO Guess format from bytes.
	sheetRaw, err := nativewebp.Decode(bytes.NewReader(b))
	if err != nil {
		log.Panicln(err)
	}
	sheet := ebiten.NewImageFromImage(sheetRaw)
	return NewSheet(SheetInfo{Image: sheet, SpriteSize: spriteSize})
}
