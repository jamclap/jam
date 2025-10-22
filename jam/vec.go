package jam

import "image"

type Number interface {
	~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Vec2f = Vec2[float64]
type Vec2i = Vec2[int]

type Vec2[T Number] struct {
	X, Y T
}

func XY[T Number](x, y T) Vec2[T] {
	return Vec2[T]{x, y}
}

func (v Vec2[T]) Add(v2 Vec2[T]) Vec2[T] {
	v.X += v2.X
	v.Y += v2.Y
	return v
}

func (v Vec2[T]) AddAll(xy T) Vec2[T] {
	v.X += xy
	v.Y += xy
	return v
}

func (v Vec2[T]) AddPoint(p image.Point) Vec2[T] {
	v.X += T(p.X)
	v.Y += T(p.Y)
	return v
}

func (v Vec2[T]) AddX(x T) Vec2[T] {
	v.X += x
	return v
}

func (v Vec2[T]) AddY(y T) Vec2[T] {
	v.Y += y
	return v
}

func (v Vec2[T]) Div(v2 Vec2[T]) Vec2[T] {
	v.X /= v2.X
	v.Y /= v2.Y
	return v
}

func (v Vec2[T]) DivAll(xy T) Vec2[T] {
	v.X /= xy
	v.Y /= xy
	return v
}

func (v Vec2[T]) DivPoint(p image.Point) Vec2[T] {
	v.X /= T(p.X)
	v.Y /= T(p.Y)
	return v
}

func (v Vec2[T]) Float64() Vec2f {
	return Vec2Of[float64](v)
}

func (v Vec2[T]) Int() Vec2i {
	return Vec2Of[int](v)
}

func (v Vec2[T]) Max(v2 Vec2[T]) Vec2[T] {
	v.X = max(v.X, v2.X)
	v.Y = max(v.Y, v2.Y)
	return v
}

func (v Vec2[T]) Min(v2 Vec2[T]) Vec2[T] {
	v.X = min(v.X, v2.X)
	v.Y = min(v.Y, v2.Y)
	return v
}

func (v Vec2[T]) Mul(v2 Vec2[T]) Vec2[T] {
	v.X *= v2.X
	v.Y *= v2.Y
	return v
}

func (v Vec2[T]) MulAll(xy T) Vec2[T] {
	v.X *= xy
	v.Y *= xy
	return v
}

func (v Vec2[T]) MulPoint(p image.Point) Vec2[T] {
	v.X *= T(p.X)
	v.Y *= T(p.Y)
	return v
}

func (v Vec2[T]) Point() image.Point {
	return image.Pt(int(v.X), int(v.Y))
}

func (v Vec2[T]) Sub(v2 Vec2[T]) Vec2[T] {
	v.X -= v2.X
	v.Y -= v2.Y
	return v
}

func (v Vec2[T]) SubAll(xy T) Vec2[T] {
	v.X -= xy
	v.Y -= xy
	return v
}

func (v Vec2[T]) SubPoint(v2 image.Point) Vec2[T] {
	v.X -= T(v2.X)
	v.Y -= T(v2.Y)
	return v
}

func Vec2FromPoint(p image.Point) Vec2i {
	return Vec2i{}.AddPoint(p)
}

func Vec2Of[T, U Number](v Vec2[U]) Vec2[T] {
	w := Vec2[T]{}
	w.X = T(v.X)
	w.Y = T(v.Y)
	return w
}
