package jam

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Draw struct {
	// Local ebiten.GeoM?
	// Local Op?
	Target *ebiten.Image
}

type Op struct {
	pos    Vec2f
	scale1 Vec2f
}

func Pos(pos Vec2f) Op {
	return Op{}.Pos(pos)
}

func PosXY(x, y float64) Op {
	return Op{}.PosXY(x, y)
}

func ScaleX(x float64) Op {
	return Op{}.ScaleX(x)
}

func (o Op) Pos(pos Vec2f) Op {
	o.pos = pos
	return o
}

func (o Op) PosXY(x, y float64) Op {
	return o.Pos(XY(x, y))
}

func (o Op) Scale(xy float64) Op {
	return o.ScaleXY(xy, xy)
}

func (o Op) ScaleVec(xy Vec2f) Op {
	return o.ScaleXY(xy.X, xy.Y)
}

func (o Op) ScaleX(x float64) Op {
	// TODO Actually int scale?
	o.scale1.X = (o.scale1.X+1)*x - 1
	return o
}

func (o Op) ScaleXY(x, y float64) Op {
	return o.ScaleX(x).ScaleY(y)
}

func (o Op) ScaleY(y float64) Op {
	// TODO Actually int scale?
	o.scale1.Y = (o.scale1.Y+1)*y - 1
	return o
}

func (d *Draw) Fill(color color.Color) {
	d.Target.Fill(color)
}

func (d *Draw) Sprite(image *ebiten.Image, op Op) {
	// TODO Keep this cached in draw.
	eop := ebiten.DrawImageOptions{}
	// options.GeoM.Concat(d.Global)
	// // TODO Round or floor xy based on option scale to nearest virtual pixel.
	// if op.FlipX {
	// 	options.GeoM.Scale(-1, 1)
	// }
	// if op.FlipY {
	// 	options.GeoM.Scale(1, -1)
	// }
	scale := op.scale1.AddAll(1)
	eop.GeoM.Scale(scale.X, scale.Y)
	eop.GeoM.Translate(op.pos.X, op.pos.Y)
	offset := Vec2f{}
	if scale.X < 0 {
		offset.X = -scale.X
	}
	if scale.Y < 0 {
		offset.Y = -scale.Y
	}
	offset = offset.MulPoint(image.Bounds().Size())
	eop.GeoM.Translate(offset.X, offset.Y)
	// The point of this is make things like flip work well.
	d.Target.DrawImage(image, &eop)
}

func (d *Draw) TileLayer(
	m *TileMap, l *TileLayer, bounds image.Rectangle, op Op,
) {
	//
}

func (d *Draw) TileMap(m *TileMap, pixelBounds image.Rectangle, op Op) {
	//
}
