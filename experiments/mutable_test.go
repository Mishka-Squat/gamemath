package experiments_test

import "github.com/Mishka-Squat/goex/mathex"

type MVector[T mathex.Number] struct {
	X T
	Y T
	Z T
}

func (v MVector[T]) Reciprocal() MVector[float64] {
	return MVector[float64]{
		X: 1.0 / float64(v.X),
		Y: 1.0 / float64(v.Y),
		Z: 1.0 / float64(v.Z),
	}
}

func (v MVector[T]) Add(a MVector[T]) MVector[T] {
	return MVector[T]{
		X: v.X + a.X,
		Y: v.Y + a.Y,
		Z: v.Z + a.Z,
	}
}

func (v MVector[T]) MulByVector(a MVector[T]) MVector[T] {
	return MVector[T]{
		X: v.X * a.X,
		Y: v.Y * a.Y,
		Z: v.Z * a.Z,
	}
}

func (v *MVector[T]) AddInPlace(a MVector[T]) {
	v.X += a.X
	v.Y += a.Y
	v.Z += a.Z
}

func (v *MVector[T]) MultInPlace(a MVector[T]) {
	v.X *= a.X
	v.Y *= a.Y
	v.Z *= a.Z
}

func (v *MVector[T]) AddInPlaceAndReturn(a MVector[T]) MVector[T] {
	v.X += a.X
	v.Y += a.Y
	v.Z += a.Z
	return *v
}

func (v *MVector[T]) MultInPlaceAndReturn(a MVector[T]) MVector[T] {
	v.X *= a.X
	v.Y *= a.Y
	v.Z *= a.Z
	return *v
}

func (v *MVector[T]) AddInPlaceTakingPointer(a *MVector[T]) {
	v.X += a.X
	v.Y += a.Y
	v.Z += a.Z
}

func (v *MVector[T]) MultInPlaceTakingPointer(a *MVector[T]) {
	v.X *= a.X
	v.Y *= a.Y
	v.Z *= a.Z
}
