package line2

import (
	"iter"
	"math"

	"deedles.dev/xiter"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/mathex"
)

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
		for p := a; ; {
			// Send current coordinate to the caller
			if !yield(p) {
				return
			}
			if p == b {
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

func (c Of[T]) EnumScoreCheck() iter.Seq2[vector2.Of[T], *float32] {
	return func(yield func(vector2.Of[T], *float32) bool) {
		var score float32
		var prev_xy vector2.Of[T]
		for line_i, line_xy := range xiter.Enumerate(Enum(c.A, c.B)) {
			if line_i == 0 {
				prev_xy = line_xy
				continue
			}

			delta_score := score
			if !yield(line_xy, &delta_score) {
				return
			}
			// returns delta score in score
			if delta_score < 0 {
				return
			} else {
				var length_coef float32 = 1
				d_xy := line_xy.Sub(prev_xy)
				if d_xy.X != 0 && d_xy.Y != 0 {
					length_coef = math.Sqrt2
				}

				score = score + length_coef*delta_score
			}

			prev_xy = line_xy
		}
	}
}
