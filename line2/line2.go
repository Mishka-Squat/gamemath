package line2

import (
	"iter"

	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/mathex"
)

type Of[T mathex.SignedNumber] struct {
	A vector2.Of[T]
	B vector2.Of[T]
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

func Make[T mathex.SignedNumber](a vector2.Of[T], b vector2.Of[T]) Of[T] {
	return Of[T]{
		A: a,
		B: b,
	}
}

// ABC - Calculate line's parametruc A,B,C coeffecients
func (l Of[T]) ABC() (a T, b T, c T) {
	return (l.B.Y - l.A.Y), (l.A.X - l.B.X), (l.B.X*l.A.Y - l.A.X*l.B.Y)
}

// Bresenham calculates all grid points along a line segment.
// The yield callback function processes each point. Return false to stop early.
func Enum[T mathex.SignedNumber](a, b vector2.Of[T]) iter.Seq[vector2.Of[T]] {
	d := a.Sub(b).Abs()

	sx := -1
	if a.X < b.X {
		sx = 1
	}
	sy := -1
	if a.Y < b.Y {
		sy = 1
	}

	err := d.X - d.Y

	return func(yield func(vector2.Of[T]) bool) {
		for p := a; p != b; {
			// Send current coordinate to the caller
			if !yield(p) {
				return
			}

			e2 := 2 * err
			if e2 > -d.Y {
				err -= d.Y
				p.X += T(sx)
			}
			if e2 < d.X {
				err += d.X
				p.Y += T(sy)
			}
		}
	}
}

func (l Of[T]) Enum() iter.Seq[vector2.Of[T]] {
	return Enum(l.A, l.B)
}
