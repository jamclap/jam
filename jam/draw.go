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
	scale1 Vec2f // Awkward math to allow zero default.
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

func (o Op) GetScale() Vec2f {
	return o.scale1.AddAll(1)
}

func (o Op) Move(xy Vec2f) Op {
	o.pos = o.pos.Add(xy)
	return o
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
	scale := op.GetScale()
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

// Always draw full tiles but only those touching pixelBounds?
func (d *Draw) TileLayers(
	layers *TileLayers,
	pixelBounds image.Rectangle,
	op Op,
) {
	sheets := layers.Sheets
	for _, layer := range layers.Layers {
		d.TileMap(layer, sheets, pixelBounds, op)
	}
}

func (d *Draw) TileMap(
	m *TileMap,
	sheets []*Sheet,
	pixelBounds image.Rectangle, // TODO Give grid bounds instead?
	op Op,
) {
	scale := op.GetScale()
	tileSize := Vec2Of[float64](m.TileSize)
	// Group by sheet for sprite batching.
	for sheetIndex := 0; sheetIndex < len(sheets); sheetIndex++ {
		sheet := sheets[sheetIndex]
		// TODO Use bounds, apply offsets.
		start := Vec2i{}
		drawSize := m.Tiles.Size()
		semiStride := m.Tiles.Size().X - drawSize.X
		tiles := m.Tiles.Items()
		index := m.Tiles.Index(start)
		offset := XY(0, 0.0)
		for tileY := 0; tileY < drawSize.Y; tileY++ {
			tileOp := op.Move(offset)
			for tileX := 0; tileX < drawSize.X; tileX++ {
				tile := tiles[index]
				if tile.Opacity() > 0 {
					if tile.Sheet() == sheetIndex {
						// TODO Apply opacity.
						d.Sprite(sheet.At(tile.Pos()), tileOp)
					}
				}
				// TODO Should move be autoscaled?
				tileOp = tileOp.Move(XY(tileSize.X*scale.X, 0))
				index++
			}
			offset.Y += tileSize.Y * scale.Y
			index += semiStride
		}
	}
}
