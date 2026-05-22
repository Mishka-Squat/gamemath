package transform2

import (
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/mathex"
)

type Of[T mathex.SignedNumber] struct {
	T vector2.Of[T]
	R vector2.Of[T]
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

func New[T mathex.SignedNumber](t vector2.Of[T], r vector2.Of[T]) Of[T] {
	return Of[T]{
		T: t,
		R: r,
	}
}

func Identity[T mathex.SignedNumber]() Of[T] {
	return Of[T]{
		T: vector2.Zero[T](),
		R: vector2.Right[T](),
	}
}

// Transform a point (e.g. local space to world space)
func (t Of[T]) TransformPoint(p vector2.Of[T]) vector2.Of[T] {
	return vector2.Of[T]{
		X: (t.R.X*p.X - t.R.Y*p.Y) + t.T.X,
		Y: (t.R.Y*p.X + t.R.X*p.Y) + t.T.Y,
	}
}

// / Inverse transform a point (e.g. world space to local space)
func (t Of[T]) InvTransformPoint(p vector2.Of[T]) vector2.Of[T] {
	vx := p.X - t.T.X
	vy := p.Y - t.T.Y
	return vector2.Of[T]{
		X: t.R.X*vx + t.R.Y*vy,
		Y: -t.R.Y*vx + t.R.X*vy,
	}
}

// Multiply two transforms. If the result is applied to a point p local to frame B,
// the transform would first convert p to a point local to frame A, then into a point
// in the world frame.
// v2 = A.q.Rot(B.q.Rot(v1) + B.p) + A.p
//
//	= (A.q * B.q).Rot(v1) + A.q.Rot(B.p) + A.p
func (t Of[T]) MulTransforms(B Of[T]) Of[T] {
	return Of[T]{
		T: B.T.RotateVector(t.R).Add(t.T),
		R: t.R.MulRot(B.R),
	}
}

// Creates a transform that converts a local point in frame B to a local point in frame A.
// v2 = A.q' * (B.q * v1 + B.p - A.p)
//
//	= A.q' * B.q * v1 + A.q' * (B.p - A.p)
func (t Of[T]) InvMulTransforms(B Of[T]) Of[T] {
	return Of[T]{
		T: B.T.Sub(t.T).InvRotateVector(t.R),
		R: t.R.InvMulRot(B.R),
	}
}
