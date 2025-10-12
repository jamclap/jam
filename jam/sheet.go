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
	gridSize   Vec2i
	spriteSize Vec2i
	sprites    Grid[*ebiten.Image]
	tags       Grid[uint32]
}

// Excludes any partial space at edges.
func NewSheet(image *ebiten.Image, spriteSize Vec2i) *Sheet {
	imageSize := image.Bounds().Size()
	gridSize := Vec2i{}.AddPoint(imageSize).Div(spriteSize)
	return &Sheet{
		image:      image,
		gridSize:   gridSize,
		spriteSize: spriteSize,
		sprites:    NewGrid[*ebiten.Image](gridSize),
	}
}

func (s *Sheet) At(xy Vec2i) *ebiten.Image {
	if xy.X < 0 || xy.X >= s.gridSize.X || xy.Y < 0 || xy.Y >= s.gridSize.Y {
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
func LoadSheet(b []byte, spriteSize Vec2i) *Sheet {
	// TODO Guess format from bytes.
	sheetRaw, err := nativewebp.Decode(bytes.NewReader(b))
	if err != nil {
		log.Panicln(err)
	}
	sheet := ebiten.NewImageFromImage(sheetRaw)
	return NewSheet(sheet, spriteSize)
}
