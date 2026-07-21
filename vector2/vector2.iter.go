package vector2

import (
	"iter"
)

func (v Of[T]) EnumRegion(w, h T, step ...T) iter.Seq[Of[T]] {
	return func(yield func(Of[T]) bool) {
		sx := T(1)
		sy := T(1)
		if len(step) > 0 {
			sx = step[0]
		}
		if len(step) > 1 {
			sy = step[1]
		}

		for y := v.Y; y <= v.Y+h; y += sy {
			for x := v.X; x <= v.X+w; x += sx {
				if !yield(Of[T]{X: x, Y: y}) {
					return
				}
			}
		}
	}
}

// Enumerate coordinates around vector position
// rw and rh are width and height radiuses of enumerated region
func (v Of[T]) EnumRegionAround(wr, hr T, step ...T) iter.Seq[Of[T]] {
	return func(yield func(Of[T]) bool) {
		sx := T(1)
		sy := T(1)
		if len(step) > 0 {
			sx = step[0]
		}
		if len(step) > 1 {
			sy = step[1]
		}

		for y := v.Y - hr; y <= v.Y+hr; y += sy {
			for x := v.X - wr; x <= v.X+wr; x += sx {
				if !yield(Of[T]{X: x, Y: y}) {
					return
				}
			}
		}
	}
}

// Enumerate coordinates around vector position
// rw and rh are width and height radiuses of enumerated region
func (v Of[T]) EnumCircleAround(r T, step ...T) iter.Seq[Of[T]] {
	return func(yield func(Of[T]) bool) {
		sx := T(1)
		sy := T(1)
		if len(step) > 0 {
			sx = step[0]
		}
		if len(step) > 1 {
			sy = step[1]
		}

		for y := v.Y - r; y <= v.Y+r; y += sy {
			for x := v.X - r; x <= v.X+r; x += sx {
				if v.SubXY(x, y).LengthSquared() > r*r {
					continue
				}

				if !yield(Of[T]{X: x, Y: y}) {
					return
				}
			}
		}
	}
}
