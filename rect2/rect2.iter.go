package rect2

import (
	"iter"

	"github.com/Mishka-Squat/gamemath/vector2"
)

func (r Of[T]) EachUnitCell() iter.Seq[vector2.Of[T]] {
	return func(yield func(vector2.Of[T]) bool) {
		for x := T(0); x < r.Size.X; x++ {
			for y := T(0); y < r.Size.Y; y++ {
				if !yield(vector2.Make(r.Position.X+x, r.Position.X+y)) {
					return
				}
			}
		}
	}
}
