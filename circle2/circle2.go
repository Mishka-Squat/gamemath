package circle2

import (
	"iter"

	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/mathex"
)

type Of[T mathex.SignedNumber] struct {
	Center vector2.Of[T]
	Radius T
}

type (
	Float64 = Of[float64]
	Float32 = Of[float32]
	Int     = Of[int]
	Int64   = Of[int64]
	Int32   = Of[int32]
	Int16   = Of[int16]
	Int8    = Of[int8]
)

func Make[T mathex.SignedNumber](center vector2.Of[T], radius T) Of[T] {
	return Of[T]{
		Center: center,
		Radius: radius,
	}
}

// BresenhamCircle calculates all grid points along the perimeter of a circle.
// (xc, yc) is the center point, and r is the radius.
// The yield callback processes each point. Return false to stop early.
func Enum[T mathex.SignedNumber](c vector2.Of[T], r T) iter.Seq[vector2.Of[T]] {
	return func(yield func(vector2.Of[T]) bool) {
		p := vector2.Make(0, r)
		d := 3 - (2 * r) // Initial decision parameter

		// Plot the initial points on the main axes
		if !plot8Points(c, p, yield) {
			return
		}

		for p.Y >= p.X {
			p.X++

			// Check decision parameter to update error margin
			if d > 0 {
				p.Y--
				d = d + 4*(p.X-p.Y) + 10
			} else {
				d = d + 4*p.X + 6
			}

			// Mirror the newly calculated point across all 8 octants
			if !plot8Points(c, p, yield) {
				return
			}
		}
	}
}

// plot8Points mirrors a point (x, y) relative to center (xc, yc) into all 8 octants.
func plot8Points[T mathex.SignedNumber](c, p vector2.Of[T], yield func(vector2.Of[T]) bool) bool {
	if !yield(vector2.Make(c.X+p.X, c.Y+p.Y)) {
		return false
	}
	if !yield(vector2.Make(c.X-p.X, c.Y+p.Y)) {
		return false
	}
	if !yield(vector2.Make(c.X+p.X, c.Y-p.Y)) {
		return false
	}
	if !yield(vector2.Make(c.X-p.X, c.Y-p.Y)) {
		return false
	}
	if !yield(vector2.Make(c.X+p.Y, c.Y+p.X)) {
		return false
	}
	if !yield(vector2.Make(c.X-p.Y, c.Y+p.X)) {
		return false
	}
	if !yield(vector2.Make(c.X+p.Y, c.Y-p.X)) {
		return false
	}
	if !yield(vector2.Make(c.X-p.Y, c.Y-p.X)) {
		return false
	}

	return true
}
