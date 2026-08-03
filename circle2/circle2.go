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
	if r < 1 {
		return func(yield func(vector2.Of[T]) bool) {
		}
	}

	return func(yield func(vector2.Of[T]) bool) {
		p := vector2.Make(0, r)
		d := 3 - (2 * r) // Initial decision parameter

		// Plot the initial points on the main axes
		for p := range enum4Points(c, p) {
			if !yield(p) {
				return
			}
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

			// The update above can overshoot past the diagonal (p.X > p.Y),
			// which would re-emit points already yielded by a prior
			// iteration (mirrored with X/Y swapped). Stop before that.
			if p.X > p.Y {
				break
			}

			// Mirror the newly calculated point across all 8 octants
			for p := range enum8Points(c, p) {
				if !yield(p) {
					return
				}
			}
		}
	}
}

func enum4Points[T mathex.SignedNumber](c, p vector2.Of[T]) iter.Seq[vector2.Of[T]] {
	return func(yield func(vector2.Of[T]) bool) {
		if !yield(vector2.Make(c.X+p.X, c.Y+p.Y)) { // 0
			return
		}
		if !yield(vector2.Make(c.X+p.X, c.Y-p.Y)) { // 2
			return
		}
		if !yield(vector2.Make(c.X+p.Y, c.Y+p.X)) { // 4
			return
		}
		if !yield(vector2.Make(c.X-p.Y, c.Y+p.X)) { // 6
			return
		}
	}
}

func enum8Points[T mathex.SignedNumber](c, p vector2.Of[T]) iter.Seq[vector2.Of[T]] {
	return func(yield func(vector2.Of[T]) bool) {
		if !yield(vector2.Make(c.X+p.X, c.Y+p.Y)) { // 0
			return
		}
		if !yield(vector2.Make(c.X-p.X, c.Y+p.Y)) { // 1
			return
		}
		if !yield(vector2.Make(c.X+p.X, c.Y-p.Y)) { // 2
			return
		}
		if !yield(vector2.Make(c.X-p.X, c.Y-p.Y)) { // 3
			return
		}
		if !yield(vector2.Make(c.X+p.Y, c.Y+p.X)) { // 4
			return
		}
		if !yield(vector2.Make(c.X+p.Y, c.Y-p.X)) { // 5
			return
		}
		if !yield(vector2.Make(c.X-p.Y, c.Y+p.X)) { // 6
			return
		}
		if !yield(vector2.Make(c.X-p.Y, c.Y-p.X)) { // 7
			return
		}
	}
}
