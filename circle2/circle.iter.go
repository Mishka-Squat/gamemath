package circle2

import (
	"iter"
	"math"

	"deedles.dev/xiter"
	"github.com/Mishka-Squat/gamemath/line2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/mathex"
)

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

func (c Of[T]) Enum() iter.Seq[vector2.Of[T]] {
	return Enum(c.Center, c.Radius)
}

func (c Of[T]) EnumCircle() iter.Seq[vector2.Of[T]] {
	if c.Radius < 1 {
		return func(yield func(vector2.Of[T]) bool) {
		}
	}

	return func(yield func(vector2.Of[T]) bool) {
		type ray_t struct {
			xy vector2.Of[T]
		}
		sector_rays := make([][]ray_t, 8)
		for i := range 8 {
			sector_rays[i] = make([]ray_t, int(c.Radius)+1)
			for ri := range sector_rays[i] {
				sector_rays[i][ri] = ray_t{c.Center}
			}
		}

		first_q := 4
		sector_i := -1
		for circle_xy := range c.Enum() {
			sector_i = (sector_i + 1) % 8
			ray := sector_rays[sector_i]
			if first_q > 0 {
				ray = sector_rays[sector_i*2]
			}

			full_throttle := false
			//line_iter:
			for line_i, line_xy := range xiter.Enumerate(line2.Enum(c.Center, circle_xy)) {
				if line_i == 0 {
					continue
				}

				ray_i := &ray[line_i]
				if ray_i.xy != line_xy || full_throttle {
					if !full_throttle {
						full_throttle = true
					}

					//delta_xy := line_xy.Sub(ray_i.xy)
					//if delta_xy.Y < -1 {
					//	break
					//}

					ray_i.xy = line_xy
					if first_q > 0 {
						sector_rays[sector_i*2+1][line_i].xy = line_xy
					}

					if !yield(line_xy) {
						return
					}
				}
			}

			if first_q > 0 {
				first_q--
			}
		}
	}
}

func (c Of[T]) EnumScoreCheck() iter.Seq2[vector2.Of[T], *float32] {
	if c.Radius < 1 {
		return func(yield func(vector2.Of[T], *float32) bool) {
		}
	}

	return func(yield func(vector2.Of[T], *float32) bool) {
		type ray_t struct {
			xy    vector2.Of[T]
			score float32
		}
		sector_rays := make([][]ray_t, 8)
		for i := range 8 {
			sector_rays[i] = make([]ray_t, int(c.Radius)+1)
			for ri := range sector_rays[i] {
				sector_rays[i][ri] = ray_t{vector2.MakeT[T](-1, -1), 0}
			}
		}

		first_q := 4
		sector_i := -1
		for circle_xy := range c.Enum() {
			sector_i = (sector_i + 1) % 8
			ray := sector_rays[sector_i]
			if first_q > 0 {
				ray = sector_rays[sector_i*2]
			}

			full_throttle := false
			//line_iter:
			for line_i, line_xy := range xiter.Enumerate(line2.Enum(c.Center, circle_xy)) {
				if line_i == 0 {
					continue
				}
				ray_pi := ray[line_i-1]
				if ray_pi.score < 0 {
					break
				}

				ray_i := &ray[line_i]
				if ray_i.xy != line_xy || full_throttle {
					if !full_throttle {
						for i := range ray[line_i:] {
							ray[line_i:][i].score = 0
						}
						full_throttle = true
					}

					//delta_xy := line_xy.Sub(ray_i.xy)
					//if delta_xy.Y < -1 {
					//	break
					//}
					if ray_i.score < ray_pi.score {
						ray_i.score = ray_pi.score - 1 // guarantee lest later
					}

					score := ray_pi.score
					ray_i.xy = line_xy
					if first_q > 0 {
						sector_rays[sector_i*2+1][line_i].xy = line_xy
					}

					if !yield(line_xy, &score) {
						return
					}
					// returns delta score in score
					if score < 0 {
						ray_i.score = -0 - ray_i.score + score
					} else {
						var length_coef float32 = 1
						d_xy := line_xy.Sub(ray_pi.xy)
						if d_xy.X != 0 && d_xy.Y != 0 {
							length_coef = math.Sqrt2
						}

						score = ray_pi.score + length_coef*score
						if score < ray_i.score {
							break
						}
						ray_i.score = score
					}
					if first_q > 0 {
						sector_rays[sector_i*2+1][line_i].score = ray_i.score
					}
				}
			}

			if first_q > 0 {
				first_q--
			}
		}
	}
}

// Enumerate coordinates around vector position
// rw and rh are width and height radiuses of enumerated region
func (v Of[T]) EnumCircleAround() iter.Seq[vector2.Of[T]] {
	return func(yield func(vector2.Of[T]) bool) {
		for p, score := range v.EnumScoreCheck() {
			if !yield(p) {
				return
			}

			*score = 0.1
		}
	}
}
