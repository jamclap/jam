package jam

import "github.com/hajimehoshi/ebiten/v2"

type Window interface {
	SetTitle(s string) // TODO Should title be more general than window?
}

type EbitenWindow struct{}

func (e *EbitenWindow) SetTitle(s string) {
	ebiten.SetWindowTitle(s)
}
