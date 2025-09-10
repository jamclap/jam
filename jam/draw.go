package jam

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Draw struct {
	// TODO Option stack
	Target *ebiten.Image
}

func (d *Draw) Image(image *ebiten.Image, x float64, y float64) {
	// TODO Keep this cached in draw.
	options := ebiten.DrawImageOptions{}
	// TODO Round or floor xy based on option scale to nearest virtual pixel.
	options.GeoM.Translate(x, y)
	d.Target.DrawImage(image, &options)
}
