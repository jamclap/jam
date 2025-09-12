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
	sprites    []*ebiten.Image
}

// Excludes any partial space at edges.
func NewSheet(image *ebiten.Image, spriteSize Vec2i) *Sheet {
	imageSize := image.Bounds().Size()
	gridSize := Vec2i{}.AddPoint(imageSize).Div(spriteSize)
	return &Sheet{
		image:      image,
		gridSize:   gridSize,
		spriteSize: spriteSize,
		sprites:    make([]*ebiten.Image, gridSize.X*gridSize.Y),
	}
}

func (s *Sheet) At(xy Vec2i) *ebiten.Image {
	if xy.X < 0 || xy.X >= s.gridSize.X || xy.Y < 0 || xy.Y >= s.gridSize.Y {
		return nil
	}
	idx := xy.X*s.gridSize.Y + xy.Y
	if s.sprites[idx] == nil {
		pixelStart := xy.Mul(s.spriteSize)
		pixelEnd := pixelStart.Add(s.spriteSize)
		r := image.Rect(pixelStart.X, pixelStart.Y, pixelEnd.X, pixelEnd.Y)
		s.sprites[idx] = s.image.SubImage(r).(*ebiten.Image)
	}
	return s.sprites[idx]
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
