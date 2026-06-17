package rect2

import (
	"fmt"

	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/mathex"
)

type Of[T mathex.SignedNumber] struct {
	Position vector2.Of[T]
	Size     vector2.Of[T]
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

func Make[T mathex.SignedNumber](position vector2.Of[T], size vector2.Of[T]) Of[T] {
	return Of[T]{
		Position: position,
		Size:     size,
	}
}

func MakeSize[T mathex.SignedNumber](size vector2.Of[T]) Of[T] {
	return Of[T]{
		Position: vector2.Zero[T](),
		Size:     size,
	}
}

func MakeXYWH[T mathex.SignedNumber](x, y, w, h T) Of[T] {
	return Of[T]{
		Position: vector2.Make(x, y),
		Size:     vector2.Make(w, h),
	}
}

func MakeT[T mathex.SignedNumber, PT, ST mathex.SignedNumber](position vector2.Of[PT], size vector2.Of[ST]) Of[T] {
	return Of[T]{
		Position: vector2.To[T](position),
		Size:     vector2.To[T](size),
	}
}

func MakeXYWHT[T, XT, YT, WT, HT mathex.SignedNumber](x XT, y YT, w WT, h HT) Of[T] {
	return Of[T]{
		Position: vector2.MakeT[T](x, y),
		Size:     vector2.MakeT[T](w, h),
	}
}

func MakeFloat64[PT, ST mathex.SignedNumber](position vector2.Of[PT], size vector2.Of[ST]) Float64 {
	return MakeT[float64](position, size)
}

func MakeFloat32[PT, ST mathex.SignedNumber](position vector2.Of[PT], size vector2.Of[ST]) Float32 {
	return MakeT[float32](position, size)
}

func MakeInt[PT, ST mathex.SignedNumber](position vector2.Of[PT], size vector2.Of[ST]) Int {
	return MakeT[int](position, size)
}

func MakeInt64[PT, ST mathex.SignedNumber](position vector2.Of[PT], size vector2.Of[ST]) Int64 {
	return MakeT[int64](position, size)
}

func MakeInt32[PT, ST mathex.SignedNumber](position vector2.Of[PT], size vector2.Of[ST]) Int32 {
	return MakeT[int32](position, size)
}

func MakeInt16[PT, ST mathex.SignedNumber](position vector2.Of[PT], size vector2.Of[ST]) Int16 {
	return MakeT[int16](position, size)
}

func MakeInt8[PT, ST mathex.SignedNumber](position vector2.Of[PT], size vector2.Of[ST]) Int8 {
	return MakeT[int8](position, size)
}

func Zero[T mathex.SignedNumber]() Of[T] {
	return Of[T]{
		Position: vector2.Zero[T](),
		Size:     vector2.Zero[T](),
	}
}

func One[T mathex.SignedNumber]() Of[T] {
	return Of[T]{
		Position: vector2.Zero[T](),
		Size:     vector2.One[T](),
	}
}

func (r Of[T]) String() string {
	return fmt.Sprintf("Position: %v; Size: %v;", r.Position, r.Size)
}

func (v Of[T]) IsZero() bool {
	return v.Position.IsZero() && v.Size.IsZero()
}

func (r Of[T]) A() vector2.Of[T] {
	return r.Position
}

func (r Of[T]) AB() vector2.Of[T] {
	return vector2.Of[T]{
		X: r.Position.X,
		Y: r.Position.Y + r.Size.Y,
	}
}

func (r Of[T]) SetA(a vector2.Of[T]) Of[T] {
	dxy := a.Sub(r.Position)
	return Of[T]{
		Position: a,
		Size:     r.Size.Sub(dxy),
	}
}

func (r Of[T]) B() vector2.Of[T] {
	return r.Position.Add(r.Size)
}

func (r Of[T]) BA() vector2.Of[T] {
	return vector2.Of[T]{
		X: r.Position.X + r.Size.X,
		Y: r.Position.Y,
	}
}

func (r Of[T]) SetB(b vector2.Of[T]) Of[T] {
	return Of[T]{
		Position: r.Position,
		Size:     b,
	}
}

func (r Of[T]) HorizontalLine(y T) (vector2.Of[T], vector2.Of[T]) {
	return vector2.Make(r.A().X, y), vector2.Make(r.B().X, y)
}

func (r Of[T]) VerticalLine(x T) (vector2.Of[T], vector2.Of[T]) {
	return vector2.Make(x, r.A().Y), vector2.Make(x, r.B().Y)
}

func (r Of[T]) Center() vector2.Of[T] {
	return r.Position.Add(r.Size.ScaleF(0.5))
}

func (v Of[T]) ToFloat64() Of[float64] {
	return Of[float64]{
		Position: v.Position.ToFloat64(),
		Size:     v.Size.ToFloat64(),
	}
}

func (v Of[T]) ToFloat32() Of[float32] {
	return Of[float32]{
		Position: v.Position.ToFloat32(),
		Size:     v.Size.ToFloat32(),
	}
}

func (v Of[T]) ToInt() Of[int] {
	return Of[int]{
		Position: v.Position.ToInt(),
		Size:     v.Size.ToInt(),
	}
}

func (v Of[T]) ToInt32() Of[int32] {
	return Of[int32]{
		Position: v.Position.ToInt32(),
		Size:     v.Size.ToInt32(),
	}
}

func (v Of[T]) ToInt64() Of[int64] {
	return Of[int64]{
		Position: v.Position.ToInt64(),
		Size:     v.Size.ToInt64(),
	}
}

// X returns the x of the xy component
func (r Of[T]) X() T {
	return r.Position.X
}

// SetX changes the x of the xy component of the rectangle
func (r Of[T]) SetX(newX T) Of[T] {
	return Of[T]{
		Position: r.Position.SetX(newX),
		Size:     r.Size,
	}
}

func (r Of[T]) AddX(dX T) Of[T] {
	return Of[T]{
		Position: r.Position.AddX(dX),
		Size:     r.Size,
	}
}

// Y returns the y of the xy component
func (r Of[T]) Y() T {
	return r.Position.Y
}

// SetY changes the y of the xy component of the rectangle
func (r Of[T]) SetY(newY T) Of[T] {
	return Of[T]{
		Position: r.Position.SetY(newY),
		Size:     r.Size,
	}
}

func (r Of[T]) AddY(dY T) Of[T] {
	return Of[T]{
		Position: r.Position.AddY(dY),
		Size:     r.Size,
	}
}

// Width returns the x of the wh component
func (r Of[T]) Width() T {
	return r.Size.X
}

// SetWidth changes the x of the wh component of the rectangle
func (r Of[T]) SetWidth(newW T) Of[T] {
	return Of[T]{
		Position: r.Position,
		Size:     r.Size.SetX(newW),
	}
}

func (r Of[T]) AddWidth(dW T) Of[T] {
	return Of[T]{
		Position: r.Position,
		Size:     r.Size.AddX(dW),
	}
}

// Y returns the y of the wh component
func (r Of[T]) Height() T {
	return r.Size.Y
}

// SetHeight changes the y of the wh component of the rectangle
func (r Of[T]) SetHeight(newH T) Of[T] {
	return Of[T]{
		Position: r.Position,
		Size:     r.Size.SetY(newH),
	}
}

func (r Of[T]) AddHeight(dH T) Of[T] {
	return Of[T]{
		Position: r.Position,
		Size:     r.Size.AddY(dH),
	}
}

// ResetPosition zero the xy component of the rectangle
func (r Of[T]) ResetPosition() Of[T] {
	return Of[T]{
		Position: vector2.Zero[T](),
		Size:     r.Size,
	}
}

// SetPosition changes the xy component of the rectangle
func (r Of[T]) SetPosition(newXY vector2.Of[T]) Of[T] {
	return Of[T]{
		Position: newXY,
		Size:     r.Size,
	}
}

// SetPosition changes the xy component of the rectangle
func (r Of[T]) SetPositionXY(x, y T) Of[T] {
	return Of[T]{
		Position: vector2.Make(x, y),
		Size:     r.Size,
	}
}

func (r Of[T]) AddPosition(dXY vector2.Of[T]) Of[T] {
	return Of[T]{
		Position: r.Position.Add(dXY),
		Size:     r.Size,
	}
}

func (r Of[T]) AddPositionXY(x, y T) Of[T] {
	return Of[T]{
		Position: r.Position.AddXY(x, y),
		Size:     r.Size,
	}
}

// SetSize changes the wh component of the rectangle
func (r Of[T]) SetSize(newWH vector2.Of[T]) Of[T] {
	return Of[T]{
		Position: r.Position,
		Size:     newWH,
	}
}

// SetSizeXY changes the wh component of the rectangle
func (r Of[T]) SetSizeXY(width, height T) Of[T] {
	return Of[T]{
		Position: r.Position,
		Size:     vector2.Make(width, height),
	}
}

func (r Of[T]) AddSize(dWH vector2.Of[T]) Of[T] {
	return Of[T]{
		Position: r.Position,
		Size:     r.Size.Add(dWH),
	}
}

func (r Of[T]) AddSizeXY(width, height T) Of[T] {
	return Of[T]{
		Position: r.Position,
		Size:     r.Size.AddXY(width, height),
	}
}

// Round takes each component of the rectangle and rounds it to the nearest whole
// number
func (v Of[T]) Round() Of[T] {
	return Make(
		v.Position.Round(),
		v.Size.Round(),
	)
}

// RoundToInt takes each component of the rectangle and rounds it to the nearest
// whole number, and then casts it to a int
func (v Of[T]) RoundToInt() Of[int] {
	return Make(
		v.Position.RoundToInt(),
		v.Size.RoundToInt(),
	)
}

// Ceil applies the ceil math operation to each component of the rectangle
func (v Of[T]) Ceil() Of[T] {
	return Make(
		v.Position.Ceil(),
		v.Size.Ceil(),
	)
}

// CeilToInt applies the ceil math operation to each component of the rectangle,
// and then casts it to a int
func (v Of[T]) CeilToInt() Of[int] {
	return Make(
		v.Position.CeilToInt(),
		v.Size.CeilToInt(),
	)
}

// Floor applies the floor math operation to each component of the rectangle
func (v Of[T]) Floor() Of[T] {
	return Make(
		v.Position.Floor(),
		v.Size.Floor(),
	)
}

// FloorToInt applies the floor math operation to each component of the rectangle,
// and then casts it to a int
func (v Of[T]) FloorToInt() Of[int] {
	return Make(
		v.Position.FloorToInt(),
		v.Size.FloorToInt(),
	)
}

func (r Of[T]) Add(xy vector2.Of[T], wh vector2.Of[T]) Of[T] {
	return Of[T]{
		Position: r.Position.Add(xy),
		Size:     r.Size.Add(wh),
	}
}

func (r Of[T]) AddXYWH(x, y, w, h T) Of[T] {
	return Of[T]{
		Position: r.Position.AddXY(x, y),
		Size:     r.Size.AddXY(w, h),
	}
}

func (r Of[T]) Grow(v T) Of[T] {
	return Of[T]{
		Position: r.Position.AddXY(-v, -v),
		Size:     r.Size.AddXY(v+v, v+v),
	}
}

func (r Of[T]) GrowXYWH(left, top, right, bottom T) Of[T] {
	return Of[T]{
		Position: r.Position.AddXY(-left, -top),
		Size:     r.Size.AddXY(left+right, top+bottom),
	}
}

func (r Of[T]) ShrinkXYWH(left, top, right, bottom T) Of[T] {
	return Of[T]{
		Position: r.Position.AddXY(left, top),
		Size:     r.Size.AddXY(-left-right, -top-bottom),
	}
}

func (r Of[T]) Centered() Of[T] {
	return r.AddPosition(r.Center())
}

func (r Of[T]) Aspect() vector2.Of[float64] {
	return vector2.Of[float64]{
		X: float64(r.Size.X) / float64(r.Size.Y),
		Y: r.Size.Length(),
	}
}

func (r Of[T]) AspectF() vector2.Of[float32] {
	return vector2.Of[float32]{
		X: float32(r.Size.X) / float32(r.Size.Y),
		Y: r.Size.LengthF(),
	}
}

func (r Of[T]) Scale(f float64) Of[T] {
	return Of[T]{
		Position: r.Position,
		Size:     r.Size.Scale(f),
	}
}

func (r Of[T]) ScaleF(f float32) Of[T] {
	return Of[T]{
		Position: r.Position,
		Size:     r.Size.ScaleF(f),
	}
}

func (r Of[T]) ScaleByVector(f vector2.Float64) Of[T] {
	return Of[T]{
		Position: r.Position,
		Size:     r.Size.ScaleByVector(f),
	}
}

func (r Of[T]) ScaleByVectorF(f vector2.Float32) Of[T] {
	return Of[T]{
		Position: r.Position,
		Size:     r.Size.ScaleByVectorF(f),
	}
}

func (r Of[T]) ScaleByXY(x, y float64) Of[T] {
	return Of[T]{
		Position: r.Position,
		Size:     r.Size.ScaleByXY(x, y),
	}
}

func (r Of[T]) ScaleByXYF(x, y float32) Of[T] {
	return Of[T]{
		Position: r.Position,
		Size:     r.Size.ScaleByXYF(x, y),
	}
}

func (r Of[T]) Zoom(f float64) Of[T] {
	return Of[T]{
		Position: r.Position.Scale(f),
		Size:     r.Size.Scale(f),
	}
}

func (r Of[T]) ZoomF(f float32) Of[T] {
	return Of[T]{
		Position: r.Position.ScaleF(f),
		Size:     r.Size.ScaleF(f),
	}
}

func (r Of[T]) ZoomByVector(f vector2.Float64) Of[T] {
	return Of[T]{
		Position: r.Position.ScaleByVector(f),
		Size:     r.Size.ScaleByVector(f),
	}
}

func (r Of[T]) ZoomByVectorF(f vector2.Float32) Of[T] {
	return Of[T]{
		Position: r.Position.ScaleByVectorF(f),
		Size:     r.Size.ScaleByVectorF(f),
	}
}

func (r Of[T]) ZoomByXY(x, y float64) Of[T] {
	return Of[T]{
		Position: r.Position.ScaleByXY(x, y),
		Size:     r.Size.ScaleByXY(x, y),
	}
}

func (r Of[T]) ZoomByXYF(x, y float32) Of[T] {
	return Of[T]{
		Position: r.Position.ScaleByXYF(x, y),
		Size:     r.Size.ScaleByXYF(x, y),
	}
}

func (r Of[T]) Inverse(v vector2.Float64) vector2.Float64 {
	return r.InverseLerp(v).SubXY(0.5, 0.5)
}

func (r Of[T]) InverseF(v vector2.Float32) vector2.Float32 {
	return r.InverseLerpF(v).SubXY(0.5, 0.5)
}

// InverseLerp calculates the inverse lerp of a point within the rectangle, returning a normalized vector2.Vector[T].
func (r Of[T]) InverseLerp(v vector2.Float64) vector2.Float64 {
	return v.Sub(r.Position.ToFloat64()).ToFloat64().ScaleByVector(r.Size.Inv())
}

// InverseLerpF calculates the inverse lerp of a point within the rectangle, returning a normalized vector2.Float32.
func (r Of[T]) InverseLerpF(v vector2.Float32) vector2.Float32 {
	return v.Sub(r.Position.ToFloat32()).ToFloat32().ScaleByVectorF(r.Size.InvF())
}

// InverseLerpXYF calculates the inverse lerp of a point within the rectangle using float32 x and y, returning a normalized vector2.Float32.
//func (r Rectangle[T]) InverseLerpXYF(x, y float32) vector2.Float32 {
//	return vector2.New(float32(x)-float32(r.Position.X), float32(y)-float32(r.Position.Y)).Div(r.Size.ToFloat32())
//}

func (r Of[T]) Lerp(t vector2.Float64) vector2.Of[T] {
	return r.Position.Add(r.Size.ScaleByVector(t))
}

func (r Of[T]) LerpF(t vector2.Float32) vector2.Of[T] {
	return r.Position.Add(r.Size.ScaleByVectorF(t))
}

func (r Of[T]) LerpXYF(x, y float32) vector2.Of[T] {
	return r.Position.Add(r.Size.ScaleByXYF(x, y))
}

func (r Of[T]) Contains(v vector2.Of[T]) bool {
	return vector2.GreaterEq(v, r.A()) && vector2.LessEq(v, r.B())
}

func (r Of[T]) OverlappedBy(or Of[T]) bool {
	return vector2.LessEq(or.A(), r.B()) && vector2.GreaterEq(or.B(), r.A())
}

func (r Of[T]) Pivot(anchor vector2.Of[T], xy vector2.Of[T]) Of[T] {
	return Of[T]{
		Position: xy.Sub(anchor.MulByVector(r.Size)),
		Size:     r.Size,
	}
}
