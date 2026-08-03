package circle2

import (
	"iter"

	"deedles.dev/xiter"
	"github.com/Mishka-Squat/gamemath/line2"
	"github.com/Mishka-Squat/gamemath/vector2"
)

func (c Of[T]) Enum() iter.Seq[vector2.Of[T]] {
	return Enum(c.Center, c.Radius)
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
		//p8s := slices.Collect(enum8Points(c.Center, c.Center.AddX(c.Radius)))
		for i := range 8 {
			sector_rays[i] = make([]ray_t, int(c.Radius)+1)
			for ri := range sector_rays[i] {
				sector_rays[i][ri] = ray_t{vector2.MakeT[T](-1, -1), 0}
			}
		}

		sector_i := -1
		for circle_xy := range c.Enum() {
			sector_i = (sector_i + 1) % 8
			ray := sector_rays[sector_i]

			full_throttle := false
		line_iter:
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

					delta_xy := line_xy.Sub(ray_i.xy)
					if delta_xy.Y < -1 {
						break
					}
					if ray_i.score < ray_pi.score {
						ray_i.score = ray_pi.score - 1 // guarantee lest later
					}

					score := ray_pi.score
					ray_i.xy = line_xy

					if !yield(line_xy, &score) {
						return
					}
					// returns delta score in score
					if score < 0 {
						ray_i.score = -0 - ray_i.score + score
					} else {
						score = ray_pi.score + score
						if score < ray_i.score {
							break line_iter
						}
						ray_i.score = score
					}
				}
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
