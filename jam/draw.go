package jam

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Draw struct {
	// Local ebiten.GeoM
	// Local  ebiten.GeoM
	Scale1 Vec2
	Target *ebiten.Image
}

type Op struct {
	pos    Vec2
	scale1 Vec2
}

func Pos(pos Vec2) Op {
	return Op{}.Pos(pos)
}

func PosXY(x, y float64) Op {
	return Op{}.PosXY(x, y)
}

func ScaleX(x float64) Op {
	return Op{}.ScaleX(x)
}

func (o Op) Pos(pos Vec2) Op {
	o.pos = pos
	return o
}

func (o Op) PosXY(x, y float64) Op {
	return o.Pos(XY(x, y))
}

func (o Op) ScaleX(x float64) Op {
	// TODO Actually int scale?
	o.scale1.X = (o.scale1.X+1)*x - 1
	return o
}

func (d *Draw) Fill(color color.Color) {
	d.Target.Fill(color)
}

func (d *Draw) ScaleX(x float64) *Draw {
	// TODO Actually int scale?
	d.Scale1.X = (d.Scale1.X+1)*x - 1
	return d
}

func (d *Draw) Sprite(image *ebiten.Image, op Op) {
	// TODO Keep this cached in draw.
	options := ebiten.DrawImageOptions{}
	// options.GeoM.Concat(d.Global)
	// // TODO Round or floor xy based on option scale to nearest virtual pixel.
	// if op.FlipX {
	// 	options.GeoM.Scale(-1, 1)
	// }
	// if op.FlipY {
	// 	options.GeoM.Scale(1, -1)
	// }
	scale := op.scale1.AddAll(1)
	options.GeoM.Scale(scale.X, scale.Y)
	options.GeoM.Translate(op.pos.X, op.pos.Y)
	offset := Vec2{}
	if scale.X < 0 {
		offset.X = -scale.X
	}
	if scale.Y < 0 {
		offset.Y = -scale.Y
	}
	offset = offset.MulPoint(image.Bounds().Size())
	options.GeoM.Translate(offset.X, offset.Y)
	// The point of this is make things like flip work well.
	d.Target.DrawImage(image, &options)
}
