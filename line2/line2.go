package line2

import (
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

func New[T mathex.SignedNumber](a vector2.Of[T], b vector2.Of[T]) Of[T] {
	return Of[T]{
		A: a,
		B: b,
	}
}

func (l Of[T]) ABC() (a T, b T, c T) {
	return (l.B.Y - l.A.Y), (l.A.X - l.B.X), (l.B.X*l.A.Y - l.A.X*l.B.Y)
}
