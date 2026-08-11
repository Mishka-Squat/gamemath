package circle2

import (
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
