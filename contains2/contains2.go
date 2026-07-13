package contains2

import (
	"github.com/Mishka-Squat/gamemath/line2"
	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/mathex"
)

type Result int

const (
	Exclude Result = iota
	Partial
	Contains
)

type Quadrant int

const (
	Inside  Quadrant = 0
	Left    Quadrant = 0b0001
	Right   Quadrant = 0b0010
	Top     Quadrant = 0b0100
	Bottom  Quadrant = 0b1000
	Invalid Quadrant = Left | Right | Top | Bottom
	Outside Quadrant = Left | Right | Top | Bottom
)

//type Result

func Wrap(Result, Quadrant) {

}

func VectorVector[T mathex.SignedNumber](a vector2.Of[T], b vector2.Of[T]) (Result, Quadrant) {
	if a.X == b.X && a.Y == b.Y {
		return Contains, Inside
	}

	q := Inside
	if a.X < b.X {
		q |= Left
	} else if a.X > b.X {
		q |= Right
	}

	if a.Y < b.Y {
		q |= Top
	} else if a.Y > b.Y {
		q |= Bottom
	}

	return Exclude, q
}

func LineVector[T mathex.SignedNumber](line line2.Of[T], vector vector2.Of[T]) (Result, Quadrant) {
	a, b, c := line.ABC()

	d := a*vector.X + b*vector.Y + c
	if d == 0 {
		return Contains, Inside
	}

	// TODO(iga): Quadrant randmoly chosen, need to check
	if d > 0 {
		return Exclude, Left
	} else {
		return Exclude, Right
	}
}

func RectVector[T mathex.SignedNumber](rect rect2.Of[T], vector vector2.Of[T]) (Result, Quadrant) {
	_, qva := VectorVector(vector, rect.A())
	_, qvb := VectorVector(vector, rect.B())

	if qva&qvb != 0 {
		return Exclude, qva & qvb
	}

	if (qva == 0 || (qva&(Right|Bottom)) != 0) && (qvb == 0 || (qvb&(Top|Left) != 0)) {
		return Contains, Inside
	}

	return Exclude, Invalid
}

func RectLine[T mathex.SignedNumber](rect rect2.Of[T], line line2.Of[T]) (Result, Quadrant, Quadrant) {
	ac, qac := RectVector(rect, line.A)
	bc, qbc := RectVector(rect, line.B)

	// both points inside rect
	if ac == Contains && bc == Contains {
		return Contains, qac, qbc
	}

	// A is inside rect
	if ac == Contains {
		return Partial, qac, qbc
	}

	// B is inside rect
	if bc == Contains {
		return Partial, qac, qbc
	}

	// Both point share common side
	if qac&qbc != 0 {
		return Exclude, qac, qbc
	}

	// Point lie exectly on opposite sides
	switch qac {
	case Left:
		if qbc == Right {
			return Partial, qac, qbc
		}
	case Right:
		if qbc == Left {
			return Partial, qac, qbc
		}
	case Top:
		if qbc == Bottom {
			return Partial, qac, qbc
		}
	case Bottom:
		if qbc == Top {
			return Partial, qac, qbc
		}
	}

	// Check all four corners relative to the line
	_, cq := LineVector(line, rect.A())
	_, ncq := LineVector(line, rect.B())
	if cq != ncq {
		return Partial, qac, qbc
	}
	_, ncq = LineVector(line, rect.AB())
	if cq != ncq {
		return Partial, qac, qbc
	}
	_, ncq = LineVector(line, rect.BA())
	if cq != ncq {
		return Partial, qac, qbc
	}

	// All four corners are on one side from the line
	return Exclude, qac, qbc
}

func RectRect[T mathex.SignedNumber](a rect2.Of[T], b rect2.Of[T]) (Result, Quadrant) {
	ac, aq := RectVector(a, b.A())
	abc, abq := RectVector(a, b.AB())
	bc, bq := RectVector(a, b.B())
	bac, baq := RectVector(a, b.BA())

	if ac == Contains {
		if abc == Contains {
			if bc == Contains {
				if bac == Contains {
					return Contains, Inside
				}

				return Partial, baq
			}

			return Partial, bq
		}

		return Partial, abq
	}

	fq := aq | abq | bq | baq
	concat_q := aq & abq & bq & baq

	if fq == Outside {
		// concat_q is zero here
		return Contains, Outside
	}

	if (concat_q&Left)|(concat_q&Right)|(concat_q&Top)|(concat_q&Bottom) != 0 {
		return Exclude, fq
	}

	return Partial, fq
}
