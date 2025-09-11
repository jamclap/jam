package jam

import "image"

type Vec2 struct {
	X, Y float64
}

func XY(x, y float64) Vec2 {
	return Vec2{x, y}
}

func (v Vec2) AddAll(xy float64) Vec2 {
	v.X += xy
	v.Y += xy
	return v
}

func (v Vec2) MulPoint(point image.Point) Vec2 {
	v.X *= float64(point.X)
	v.Y *= float64(point.Y)
	return v
}
