package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"log"

	"github.com/HugoSmits86/nativewebp"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jamclap/jam/jam"
)

func main() {
	jam.Run(InitState)
}

type State struct {
	Sprites []*ebiten.Image
}

func InitState(hub *jam.Hub) jam.App {
	sprites := loadSprites()
	return &State{Sprites: sprites}
}

func (s *State) Update(hub *jam.Hub) {
	//
}

func (s *State) Draw(draw *jam.Draw) {
	draw.Fill(color.RGBA{0x2a, 0x3d, 0x74, 0xff})
	draw.Sprite(s.Sprites[0], jam.PosXY(8, 8).ScaleX(-1))
}

func loadSprites() []*ebiten.Image {
	sheetRaw, err := nativewebp.Decode(bytes.NewReader(spriteBytes))
	if err != nil {
		log.Fatalln(err)
	}
	sheet := ebiten.NewImageFromImage(sheetRaw)
	sprites := []*ebiten.Image{}
	for y := 0; y < sheet.Bounds().Dy(); y += 8 {
		sprite := sheet.SubImage(image.Rect(0, y, 8, y+8)).(*ebiten.Image)
		sprites = append(sprites, sprite)
	}
	return sprites
}

//go:embed sprite.webp
var spriteBytes []byte
