package jam

// Dense 2D grid that pretends to be larger than its official size.
// Stores content in row-major order.
type Grid[T any] struct {
	items []T
	size  Vec2i
}

func NewGrid[T any](size Vec2i) Grid[T] {
	return Grid[T]{items: make([]T, size.X*size.Y), size: size}
}

// Panics on negative. Provides default value if beyond current size.
func (g *Grid[T]) At(xy Vec2i) T {
	g.checkMin(xy)
	if xy.X >= g.size.X || xy.Y >= g.size.Y {
		var item T
		return item
	}
	return g.items[g.Index(xy)]
}

func (g *Grid[T]) Index(xy Vec2i) int {
	return xy.Y*g.size.X + xy.X
}

// Raw item access for fast iteration.
func (g *Grid[T]) Items() []T {
	return g.items
}

// Panics on negative. Grows if beyond current size, but only as far as needed,
// so don't resize by only one at a time in a tight loop.
func (g *Grid[T]) SetAt(xy Vec2i, item T) {
	g.checkMin(xy)
	wantedSize := g.size
	if xy.X >= g.size.X {
		wantedSize.X = xy.X + 1
	}
	if xy.Y >= g.size.Y {
		wantedSize.Y = xy.Y + 1
	}
	if wantedSize != g.size {
		g.SetSize(wantedSize)
	}
	g.items[g.Index(xy)] = item
}

// Copies contents if able. Clips if made smaller than current size.
func (g *Grid[T]) SetSize(newSize Vec2i) {
	// Presume we don't do this super often, so just allocate new storage.
	oldX, oldY := g.size.X, g.size.Y
	newItems := make([]T, newSize.X*newSize.Y)
	minX := min(oldX, newSize.X)
	minY := min(oldY, newSize.Y)
	for y := 0; y < minY; y++ {
		copy(
			newItems[y*newSize.X:y*newSize.X+minX],
			g.items[y*oldX:y*oldX+minX],
		)
	}
	g.items = newItems
	g.size = newSize
}

// The actual allocated size.
func (g *Grid[T]) Size() Vec2i {
	return g.size
}

// Trims off x and y edges that have only default values in them, leaving a size
// that precisely meets needs. Worst case loops the entire grid.
func TrimGrid[T comparable](g *Grid[T]) {
	var empty T
	// For row-major, count up the rows first for better data locality.
	// Maybe we trim some rows before counting back the columns.
	// Not sure if it matters, but maybe has some value to it.
	needed := g.size
Rows:
	for y := g.size.Y - 1; y >= 0; y-- {
		index := y * g.size.X
		for x := 0; x < needed.X; x++ {
			if g.items[index] != empty {
				break Rows
			}
			index++
		}
		needed.Y--
	}
Cols:
	for x := g.size.X - 1; x >= 0; x-- {
		index := x
		for y := 0; y < needed.Y; y++ {
			if g.items[index] != empty {
				break Cols
			}
			index += g.size.Y
		}
		needed.X--
	}
	if needed != g.size {
		g.SetSize(needed)
	}
}

func (g *Grid[T]) checkMin(xy Vec2i) {
	if xy.X < 0 || xy.Y < 0 {
		panic("negative index")
	}
}
