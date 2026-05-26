package aabb2

import (
	"fmt"

	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/mathex"
)

// A is always less or equal to B
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

func NanFloat64() Float64 {
	return New(
		vector2.NanFloat64(),
		vector2.NanFloat64(),
	)
}

func NanFloat32() Float32 {
	return New(
		vector2.NanFloat32(),
		vector2.NanFloat32(),
	)
}

func (r Of[T]) String() string {
	return fmt.Sprintf("A: %v; B: %v;", r.A, r.B)
}

func (v Of[T]) IsZero() bool {
	return v.A.Sub(v.B).IsZero()
}

func (v Of[T]) IsNan() bool {
	return v.A.IsNan() || v.B.IsNan()
}

func (v Of[T]) Width() T {
	return v.A.X - v.B.X
}

func (v Of[T]) Height() T {
	return v.A.Y - v.B.Y
}

func (v Of[T]) Size() vector2.Of[T] {
	return vector2.NewT[T](v.Width(), v.Height())
}

// Does a fully contain b
func (a Of[T]) Contains(b Of[T]) bool {
	return a.A.X <= b.A.X &&
		a.A.Y <= b.A.Y &&
		b.B.X <= a.B.X &&
		b.B.Y <= a.B.Y
}

// Get the center of the AABB.
func (a Of[T]) Center() vector2.Of[T] {
	return vector2.NewT[T](
		0.5*float64(a.A.X+a.B.X),
		0.5*float64(a.A.Y+a.B.Y),
	)
}

// Get the extents of the AABB (half-widths).
func (a Of[T]) Extents() vector2.Of[T] {
	return vector2.NewT[T](
		0.5*float64(a.B.X-a.A.X),
		0.5*float64(a.B.Y-a.A.Y),
	)
}

// Union of two AABBs
func (a Of[T]) Union(b Of[T]) Of[T] {
	if a.IsNan() {
		return b
	}
	if b.IsNan() {
		return a
	}

	return Of[T]{
		A: vector2.Min(a.A, b.A),
		B: vector2.Max(a.B, b.B),
	}
}

// Do a and b overlap
func (a Of[T]) Overlaps(b Of[T]) bool {
	return !(b.A.X > a.B.X ||
		b.A.Y > a.B.Y ||
		a.A.X > b.B.X ||
		a.A.Y > b.B.Y)
}
