package vector2

import (
	"cmp"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"

	. "github.com/Mishka-Squat/gamemath"
	"github.com/Mishka-Squat/goex/mathex"
)

type Of[T mathex.SignedNumber] struct {
	X T
	Y T
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

func Make[T mathex.SignedNumber](x T, y T) Of[T] {
	return Of[T]{
		X: x,
		Y: y,
	}
}

func MakeT[T mathex.SignedNumber, XT, YT mathex.Number](x XT, y YT) Of[T] {
	return Of[T]{
		X: T(x),
		Y: T(y),
	}
}

func MakeFloat64[XT, YT mathex.Number](x XT, y YT) Float64 {
	return MakeT[float64](x, y)
}

func MakeFloat32[XT, YT mathex.Number](x XT, y YT) Float32 {
	return MakeT[float32](x, y)
}

func MakeInt[XT, YT mathex.Number](x XT, y YT) Int {
	return MakeT[int](x, y)
}

func MakeInt64[XT, YT mathex.Number](x XT, y YT) Int64 {
	return MakeT[int64](x, y)
}

func MakeInt32[XT, YT mathex.Number](x XT, y YT) Int32 {
	return MakeT[int32](x, y)
}

func MakeInt16[XT, YT mathex.Number](x XT, y YT) Int16 {
	return MakeT[int16](x, y)
}

func MakeInt8[XT, YT mathex.Number](x XT, y YT) Int8 {
	return MakeT[int8](x, y)
}

func MakePolarT[T mathex.SignedNumber, RT, PHIT mathex.Number](r RT, phi PHIT) Of[T] {
	return Of[T]{
		X: T(float64(r) * math.Cos(float64(phi))),
		Y: T(float64(r) * math.Sin(float64(phi))),
	}
}

func MakePolarFloat64[RT, PHIT mathex.Number](r RT, phi PHIT) Float64 {
	return MakePolarT[float64](r, phi)
}

func MakePolarFloat32[RT, PHIT mathex.Number](r RT, phi PHIT) Float32 {
	return MakePolarT[float32](r, phi)
}

func MakePolarInt[RT, PHIT mathex.Number](r RT, phi PHIT) Int {
	return MakePolarT[int](r, phi)
}

func MakePolarInt64[RT, PHIT mathex.Number](r RT, phi PHIT) Int64 {
	return MakePolarT[int64](r, phi)
}

func MakePolarInt32[RT, PHIT mathex.Number](r RT, phi PHIT) Int32 {
	return MakePolarT[int32](r, phi)
}

func MakePolarInt16[RT, PHIT mathex.Number](r RT, phi PHIT) Int16 {
	return MakePolarT[int16](r, phi)
}

func MakePolarInt8[RT, PHIT mathex.Number](r RT, phi PHIT) Int8 {
	return MakePolarT[int8](r, phi)
}

func MinFloat64() Float64 {
	return Make[float64](-math.MaxFloat64, -math.MaxFloat64)
}

func MinFloat32() Float32 {
	return Make[float32](-math.MaxFloat32, -math.MaxFloat32)
}

func MinInt() Int {
	return Make[int](math.MinInt, math.MinInt)
}

func MinInt64() Int64 {
	return Make[int64](math.MinInt64, math.MinInt64)
}

func MinInt32() Int32 {
	return Make[int32](math.MinInt32, math.MinInt32)
}

func MinInt16() Int16 {
	return Make[int16](math.MinInt16, math.MinInt16)
}

func MinInt8() Int8 {
	return Make[int8](math.MinInt8, math.MinInt8)
}

func MaxFloat64() Float64 {
	return Make[float64](math.MaxFloat64, math.MaxFloat64)
}

func MaxFloat32() Float32 {
	return Make[float32](math.MaxFloat32, math.MaxFloat32)
}

func MaxInt() Int {
	return Make[int](math.MaxInt, math.MaxInt)
}

func MaxInt64() Int64 {
	return Make[int64](math.MaxInt64, math.MaxInt64)
}

func MaxInt32() Int32 {
	return Make[int32](math.MaxInt32, math.MaxInt32)
}

func MaxInt16() Int16 {
	return Make[int16](math.MaxInt16, math.MaxInt16)
}

func MaxInt8() Int8 {
	return Make[int8](math.MaxInt8, math.MaxInt8)
}

func NanFloat64() Float64 {
	return Make(math.NaN(), math.NaN())
}

func NanFloat32() Float32 {
	return Make(float32(math.NaN()), float32(math.NaN()))
}

// Fill creates a vector where each component is equal to v
func Fill[T mathex.SignedNumber](v T) Of[T] {
	return Of[T]{
		X: v,
		Y: v,
	}
}

func Zero[T mathex.SignedNumber]() Of[T] {
	return Of[T]{
		X: 0,
		Y: 0,
	}
}

func Up[T mathex.SignedNumber]() Of[T] {
	return Of[T]{
		X: 0,
		Y: 1,
	}
}

func Down[T mathex.SignedNumber]() Of[T] {
	return Of[T]{
		X: 0,
		Y: -1,
	}
}

func Left[T mathex.SignedNumber]() Of[T] {
	return Of[T]{
		X: -1,
		Y: 0,
	}
}

func Right[T mathex.SignedNumber]() Of[T] {
	return Of[T]{
		X: 1,
		Y: 0,
	}
}

func One[T mathex.SignedNumber]() Of[T] {
	return Of[T]{
		X: 1,
		Y: 1,
	}
}

func Compare[T mathex.SignedNumber](a, b Of[T]) int {
	yc := cmp.Compare(a.Y, b.Y)
	if yc != 0 {
		return yc
	}
	return cmp.Compare(a.X, b.X)
}

func (v Of[T]) SignI() Int {
	sv := Zero[int]()
	if v.X < 0 {
		sv.X = -1
	} else if v.X > 0 {
		sv.X = 1
	}
	if v.Y < 0 {
		sv.Y = -1
	} else if v.Y > 0 {
		sv.Y = 1
	}
	return sv
}

func (v Of[T]) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v Of[T]) NearZeroF(epsilon float32) bool {
	fv := v.ToFloat32()
	return fv.X > -epsilon && fv.X < epsilon &&
		fv.Y > -epsilon && fv.Y < epsilon
}

func (v Of[T]) ZeroF(epsilon float32) Float32 {
	fv := v.ToFloat32()
	if fv.X < epsilon && fv.X > -epsilon {
		fv = fv.SetX(0)
	}
	if fv.Y < epsilon && fv.Y > -epsilon {
		fv = fv.SetY(0)
	}

	return fv
}

func (v Of[T]) IsNan() bool {
	return v.X != v.X || v.Y != v.Y
}

func (v Of[T]) Inv() Float64 {
	return Float64{
		X: 1.0 / float64(v.X),
		Y: 1.0 / float64(v.Y),
	}
}

func (v Of[T]) InvF() Float32 {
	return Float32{
		X: 1.0 / float32(v.X),
		Y: 1.0 / float32(v.Y),
	}
}

// Lerp linearly interpolates between a and b by t
func Lerp[T mathex.SignedNumber](t float32, a, b Of[T]) Of[T] {
	return Of[T]{
		X: mathex.Lerp(t, a.X, b.X),
		Y: mathex.Lerp(t, a.Y, b.Y),
	}
}

func Min[T mathex.SignedNumber](a, b Of[T]) Of[T] {
	return Make(
		min(a.X, b.X),
		min(a.Y, b.Y),
	)
}

func Max[T mathex.SignedNumber](a, b Of[T]) Of[T] {
	return Make(
		max(a.X, b.X),
		max(a.Y, b.Y),
	)
}

func MaxX[T mathex.SignedNumber](a, b Of[T]) T {
	return max(a.X, b.X)
}

func MaxY[T mathex.SignedNumber](a, b Of[T]) T {
	return max(a.Y, b.Y)
}

func MinX[T mathex.SignedNumber](a, b Of[T]) T {
	return min(a.X, b.X)
}

func MinY[T mathex.SignedNumber](a, b Of[T]) T {
	return min(a.Y, b.Y)
}

func Less[T mathex.SignedNumber](a, b Of[T]) bool {
	return a.X < b.X && a.Y < b.Y
}

func LessEq[T mathex.SignedNumber](a, b Of[T]) bool {
	return a.X <= b.X && a.Y <= b.Y
}

func Greater[T mathex.SignedNumber](a, b Of[T]) bool {
	return a.X > b.X && a.Y > b.Y
}

func GreaterEq[T mathex.SignedNumber](a, b Of[T]) bool {
	return a.X >= b.X && a.Y >= b.Y
}

func Midpoint[T mathex.SignedNumber](a, b Of[T]) Of[T] {
	// center = (b - a)0.5 + a
	// center = b0.5 - a0.5 + a
	// center = b0.5 + a0.5
	// center = 0.5(b + a)
	return Of[T]{
		X: T(float64(a.X+b.X) * 0.5),
		Y: T(float64(a.Y+b.Y) * 0.5),
	}
}

func Index(xy Int, i int) Int {
	return Int{
		i % xy.X,
		i / xy.X,
	}
}

func (v Of[T]) String() string {
	return fmt.Sprintf("X: %v; Y: %v;", v.X, v.Y)
}

// Builds a vector from the data found from the passed in array to the best of
// it's ability. If the length of the array is smaller than the vector itself,
// only those values will be used to build the vector, and the remaining vector
// components will remain the default value of the vector's data type (some
// version of 0).
func FromArray[T mathex.SignedNumber](data []T) Of[T] {
	v := Of[T]{}

	if len(data) > 0 {
		v.X = data[0]
	}

	if len(data) > 1 {
		v.Y = data[1]
	}

	return v
}

func Rand(r *rand.Rand) Of[float64] {
	return Of[float64]{
		X: r.Float64(),
		Y: r.Float64(),
	}
}

func (v Of[T]) MinComponent() T {
	return min(v.X, v.Y)
}

func (v Of[T]) MaxComponent() T {
	return max(v.X, v.Y)
}

func (v Of[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}{
		X: float64(v.X),
		Y: float64(v.Y),
	})
}

func (v *Of[T]) UnmarshalJSON(data []byte) error {
	aux := &struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}{
		X: 0,
		Y: 0,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	v.X = T(aux.X)
	v.Y = T(aux.Y)
	return nil
}

func (v Of[T]) Format(format string) string {
	return fmt.Sprintf(format, v.X, v.Y)
}

// Sqrt applies the Sqrt to each component of the vector
func (v Of[T]) Sqrt() Of[T] {
	return Make(
		mathex.Sqrt(v.X),
		mathex.Sqrt(v.Y),
	)
}

func (v Of[T]) Clamp(vmin, vmax T) Of[T] {
	return Of[T]{
		X: mathex.Clamp(v.X, vmin, vmax),
		Y: mathex.Clamp(v.Y, vmin, vmax),
	}
}

func (v Of[T]) ClampV(vmin, vmax Of[T]) Of[T] {
	return Of[T]{
		X: mathex.Clamp(v.X, vmin.X, vmax.X),
		Y: mathex.Clamp(v.Y, vmin.Y, vmax.Y),
	}
}

func (v Of[T]) Clamp0V(vmax Of[T]) Of[T] {
	return Of[T]{
		X: mathex.Clamp(v.X, 0, vmax.X),
		Y: mathex.Clamp(v.Y, 0, vmax.Y),
	}
}

func (v Of[T]) ToNpot() Of[T] {
	return Of[T]{
		X: mathex.Npot(v.X),
		Y: mathex.Npot(v.Y),
	}
}

func To[T, OT mathex.SignedNumber](v Of[OT]) Of[T] {
	return Of[T]{
		X: T(v.X),
		Y: T(v.Y),
	}
}

func (v Of[T]) ToFloat64() Of[float64] {
	return To[float64](v)
}

func (v Of[T]) ToFloat32() Of[float32] {
	return To[float32](v)
}

func (v Of[T]) ToInt() Of[int] {
	return To[int](v)
}

func (v Of[T]) ToInt64() Of[int64] {
	return To[int64](v)
}

func (v Of[T]) ToInt32() Of[int32] {
	return To[int32](v)
}

func (v Of[T]) ToInt16() Of[int16] {
	return To[int16](v)
}

func (v Of[T]) ToInt8() Of[int8] {
	return To[int8](v)
}

// SetX changes the x component of the vector
func (v Of[T]) SetX(X T) Of[T] {
	return Of[T]{
		X: X,
		Y: v.Y,
	}
}

// AddX adds to the x component of the vector
func (v Of[T]) AddX(dX T) Of[T] {
	return Of[T]{
		X: v.X + dX,
		Y: v.Y,
	}
}

// SetY changes the y component of the vector
func (v Of[T]) SetY(Y T) Of[T] {
	return Of[T]{
		X: v.X,
		Y: Y,
	}
}

// AddY adds to the y component of the vector
func (v Of[T]) AddY(dY T) Of[T] {
	return Of[T]{
		X: v.X,
		Y: v.Y + dY,
	}
}

func (v Of[T]) Axis(i Axis) T {
	switch i {
	case AxisX:
		return v.X
	case AxisY:
		return v.Y
	}

	return 0
}

func (v Of[T]) SetAxis(i Axis, a T) Of[T] {
	switch i {
	case AxisX:
		return Of[T]{
			X: a,
			Y: v.Y,
		}
	case AxisY:
		return Of[T]{
			X: v.X,
			Y: a,
		}
	}

	return v
}

func (v Of[T]) YX() Of[T] {
	return Of[T]{
		X: v.Y,
		Y: v.X,
	}
}

// Angle return angle in radians between vector and other vector [float64]
func (v Of[T]) Angle(other Of[T]) float64 {
	denominator := mathex.Sqrt(float64(v.LengthSquared() * other.LengthSquared()))
	if denominator < 1e-15 {
		return 0.
	}
	return mathex.Acos(mathex.Clamp(float64(v.Dot(other))/denominator, -1., 1.))
}

// AngleF return angle in radians between vector and other vector [float32]
func (v Of[T]) AngleF(other Of[T]) float32 {
	denominator := mathex.Sqrt(float32(v.LengthSquared() * other.LengthSquared()))
	if denominator < 1e-15 {
		return 0.
	}
	return mathex.Acos(mathex.Clamp(float32(v.Dot(other))/denominator, -1., 1.))
}

// Midpoint returns the midpoint between this vector and the vector passed in.
func (v Of[T]) Midpoint(o Of[T]) Of[T] {
	return o.Add(v).Scale(0.5)
}

// Dot return dot product between vector and other vector
func (v Of[T]) Dot(other Of[T]) T {
	return v.X*other.X + v.Y*other.Y
}

// Cross return cross product between vector and value
func (v Of[T]) CrossV(value T) Of[T] {
	return Of[T]{
		-value * v.Y,
		value * v.X,
	}
}

// Cross return cross product between vector and other vector
func (v Of[T]) Cross(other Of[T]) T {
	return v.X*other.Y - v.Y*other.X
}

// Perpendicular creates a vector perpendicular to the one passed in with the
// same magnitude
func (v Of[T]) Perpendicular() Of[T] {
	return Of[T]{
		X: v.Y,
		Y: -v.X,
	}
}

// Add returns a vector that is the result of two vectors added together
func (v Of[T]) Add(other Of[T]) Of[T] {
	return Of[T]{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Of[T]) AddXY(x, y T) Of[T] {
	return Of[T]{
		X: v.X + x,
		Y: v.Y + y,
	}
}

func (v Of[T]) Sub(other Of[T]) Of[T] {
	return Of[T]{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Of[T]) SubXY(x, y T) Of[T] {
	return Of[T]{
		X: v.X - x,
		Y: v.Y - y,
	}
}

func (v Of[T]) ReciprocalF() Of[float32] {
	return Of[float32]{
		X: 1.0 / float32(v.X),
		Y: 1.0 / float32(v.Y),
	}
}

func (v Of[T]) Reciprocal() Of[float64] {
	return Of[float64]{
		X: 1.0 / float64(v.X),
		Y: 1.0 / float64(v.Y),
	}
}

func (v Of[T]) Product() T {
	return v.X * v.Y
}

func (v Of[T]) LengthSquared() T {
	return v.X*v.X + v.Y*v.Y
}

func (v Of[T]) Length() float64 {
	return mathex.Sqrt((float64)(v.LengthSquared()))
}

func (v Of[T]) LengthF() float32 {
	return mathex.Sqrt((float32)(v.LengthSquared()))
}

func (v Of[T]) Normalized() Of[T] {
	return v.DivByConstant(v.Length())
}

func (v Of[T]) NormalizeF(a Of[T]) Float32 {
	return Float32{
		X: mathex.NormalizeF(a.X, 0, v.X),
		Y: mathex.NormalizeF(a.Y, 0, v.Y),
	}
}

func (v Of[T]) Normalize(a Of[T]) Float64 {
	return Float64{
		X: float64(mathex.Normalize(a.X, 0, v.X)),
		Y: float64(mathex.Normalize(a.Y, 0, v.Y)),
	}
}

func (v Of[T]) Negated() Of[T] {
	return Of[T]{
		X: -v.X,
		Y: -v.Y,
	}
}

func (v Of[T]) Scale(t float64) Of[T] {
	return Of[T]{
		X: T(float64(v.X) * t),
		Y: T(float64(v.Y) * t),
	}
}

func (v Of[T]) ScaleF(t float32) Of[T] {
	return Of[T]{
		X: T(float32(v.X) * t),
		Y: T(float32(v.Y) * t),
	}
}

func (v Of[T]) ScaleByVector(o Float64) Of[T] {
	return Of[T]{
		X: T(float64(v.X) * o.X),
		Y: T(float64(v.Y) * o.Y),
	}
}

func (v Of[T]) ScaleByVectorF(o Float32) Of[T] {
	return Of[T]{
		X: T(float32(v.X) * o.X),
		Y: T(float32(v.Y) * o.Y),
	}
}

func (v Of[T]) ScaleByVectorI(o Int) Of[T] {
	return Of[T]{
		X: v.X * T(o.X),
		Y: v.Y * T(o.Y),
	}
}

func (v Of[T]) ScaleByXY(x, y float64) Of[T] {
	return Of[T]{
		X: T(float64(v.X) * x),
		Y: T(float64(v.Y) * y),
	}
}

func (v Of[T]) ScaleByXYF(x, y float32) Of[T] {
	return Of[T]{
		X: T(float32(v.X) * x),
		Y: T(float32(v.Y) * y),
	}
}

func (v Of[T]) ScaleByXYI(x, y int) Of[T] {
	return Of[T]{
		X: v.X * T(x),
		Y: v.Y * T(y),
	}
}

func (v Of[T]) MulByVector(o Of[T]) Of[T] {
	return Of[T]{
		X: v.X * o.X,
		Y: v.Y * o.Y,
	}
}

func (v Of[T]) DivByVector(o Of[T]) Of[T] {
	return Of[T]{
		X: v.X / o.X,
		Y: v.Y / o.Y,
	}
}

func (v Of[T]) DivByConstant(t float64) Of[T] {
	return v.Scale(1.0 / t)
}

func (v Of[T]) DivByConstantF(t float32) Of[T] {
	return v.ScaleF(1.0 / t)
}

// Rotate a vector
func (v Of[T]) RotateVector(q Of[T]) Of[T] {
	return Of[T]{
		X: q.X*v.X - q.Y*v.Y,
		Y: q.Y*v.X + q.X*v.Y,
	}
}

// Inverse rotate a vector
func (v Of[T]) InvRotateVector(q Of[T]) Of[T] {
	return Of[T]{
		X: q.X*v.X + q.Y*v.Y,
		Y: -q.Y*v.X + q.X*v.Y,
	}
}

// Multiply two rotations: q * r
func (v Of[T]) MulRot(u Of[T]) Of[T] {
	// [vc -vs] * [uc -us] = [vc*uc-vs*us -vc*us-vs*uc]
	// [vs  vc]   [us  uc]   [vs*uc+vc*us -vs*us+vc*uc]
	// s(v + u) = vs * uc + vc * us
	// c(v + u) = vc * uc - vs * us
	return Of[T]{
		X: v.X*u.X - v.Y*u.Y,
		Y: v.Y*u.X + v.X*u.Y,
	}
}

// Transpose multiply two rotations: inv(a) * b
// This rotates a vector local in frame b into a vector local in frame a
func (v Of[T]) InvMulRot(u Of[T]) Of[T] {
	// [ vc vs] * [uc -us] = [vc*uc+qs*us -vc*us+vs*uc]
	// [-vs vc]   [us  uc]   [-vs*uc+vc*us vs*us+vc*uc]
	// s(v - u) = vc * us - vs * uc
	// c(v - u) = vc * uc + vs * us
	return Of[T]{
		X: v.X*u.X + v.Y*u.Y,
		Y: v.X*u.Y - v.Y*u.X,
	}
}

func (v Of[T]) Project(normal Of[T]) Of[T] {
	vdn := float64(v.Dot(normal))
	ndn := float64(normal.Dot(normal))
	mag := vdn / ndn
	return normal.Scale(mag)
}

func (v Of[T]) Reject(normal Of[T]) Of[T] {
	return v.Sub(v.Project(normal))
}

func (v Of[T]) Reflect(normal Of[T]) Of[T] {
	return v.Sub(normal.Scale(2. * float64(v.Dot(normal))))
}

func (v Of[T]) DistanceSquared(other Of[T]) T {
	xDist := other.X - v.X
	yDist := other.Y - v.Y
	return (xDist * xDist) + (yDist * yDist)
}

// Distance is the euclidean distance between two points
func (v Of[T]) Distance(other Of[T]) float64 {
	return mathex.Sqrt((float64)(v.DistanceSquared(other)))
}

// Round takes each component of the vector and rounds it to the nearest whole
// number
func (v Of[T]) Round() Of[T] {
	return Of[T]{
		X: mathex.Round(v.X),
		Y: mathex.Round(v.Y),
	}
}

// RoundToInt takes each component of the vector and rounds it to the nearest
// whole number, and then casts it to a int
func (v Of[T]) RoundToInt() Of[int] {
	return Make(
		int(mathex.Round(v.X)),
		int(mathex.Round(v.Y)),
	)
}

// RoundToInt takes each component of the vector and rounds it to the nearest
// whole number, and then casts it to a int
func (v Of[T]) RoundToVector(scale Of[T]) Of[T] {
	return Make(
		mathex.Round(v.X*scale.X)/scale.X,
		mathex.Round(v.Y*scale.Y)/scale.Y,
	)
}

// Ceil applies the ceil math operation to each component of the vector
func (v Of[T]) Ceil() Of[T] {
	return Of[T]{
		X: mathex.Ceil(v.X),
		Y: mathex.Ceil(v.Y),
	}
}

// CeilToInt applies the ceil math operation to each component of the vector,
// and then casts it to a int
func (v Of[T]) CeilToInt() Of[int] {
	return Make(
		int(mathex.Ceil(v.X)),
		int(mathex.Ceil(v.Y)),
	)
}

func (v Of[T]) Floor() Of[T] {
	return Of[T]{
		X: mathex.Floor(v.X),
		Y: mathex.Floor(v.Y),
	}
}

// FloorToInt applies the floor math operation to each component of the vector,
// and then casts it to a int
func (v Of[T]) FloorToInt() Of[int] {
	return Make(
		int(mathex.Floor(v.X)),
		int(mathex.Floor(v.Y)),
	)
}

// Abs applies the Abs math operation to each component of the vector
func (v Of[T]) Abs() Of[T] {
	return Of[T]{
		X: mathex.Abs(v.X),
		Y: mathex.Abs(v.Y),
	}
}

func (v Of[T]) NearZero() bool {
	return mathex.NearZero(v.X) && mathex.NearZero(v.Y)
}

func (v Of[T]) ContainsNaN() bool {
	if math.IsNaN(float64(v.X)) {
		return true
	}

	if math.IsNaN(float64(v.Y)) {
		return true
	}

	return false
}

func (v Of[T]) Flip() Of[T] {
	return Of[T]{
		X: v.X * -1,
		Y: v.Y * -1,
	}
}

func (v Of[T]) FlipX() Of[T] {
	return Of[T]{
		X: v.X * -1,
		Y: v.Y,
	}
}

func (v Of[T]) FlipY() Of[T] {
	return Of[T]{
		X: v.X,
		Y: v.Y * -1,
	}
}

func (v Of[T]) Pivot(anchor Of[T], wh Of[T]) Of[T] {
	return Of[T]{
		X: v.X - wh.X*anchor.X,
		Y: v.Y - wh.Y*anchor.Y,
	}
}

// Log returns the natural logarithm for each component
func (v Of[T]) Log() Of[T] {
	return Of[T]{
		X: mathex.Log(v.X),
		Y: mathex.Log(v.Y),
	}
}

// Log10 returns the decimal logarithm for each component.
func (v Of[T]) Log10() Of[T] {
	return Of[T]{
		X: mathex.Log10(v.X),
		Y: mathex.Log10(v.Y),
	}
}

// Log2 returns the binary logarithm for each component
func (v Of[T]) Log2() Of[T] {
	return Of[T]{
		X: mathex.Log2(v.X),
		Y: mathex.Log2(v.Y),
	}
}

// Exp2 returns 2**x, the base-2 exponential for each component
func (v Of[T]) Exp2() Of[T] {
	return Of[T]{
		X: mathex.Exp2(v.X),
		Y: mathex.Exp2(v.Y),
	}
}

// Exp returns e**x, the base-e exponential for each component
func (v Of[T]) Exp() Of[T] {
	return Of[T]{
		X: mathex.Exp(v.X),
		Y: mathex.Exp(v.Y),
	}
}

// Expm1 returns e**x - 1, the base-e exponential for each component minus 1. It is more accurate than Exp(x) - 1 when the component is near zero
func (v Of[T]) Expm1() Of[T] {
	return Of[T]{
		X: mathex.Expm1(v.X),
		Y: mathex.Expm1(v.Y),
	}
}

func (v Of[T]) Values() (T, T) {
	return v.X, v.Y
}
