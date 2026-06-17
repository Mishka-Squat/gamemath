package vector4

import (
	"cmp"
	"encoding/json"
	"fmt"
	"image/color"
	"math"

	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/gamemath/vector3"
	"github.com/Mishka-Squat/goex/mathex"
)

// Of contains 4 components
type Of[T mathex.SignedNumber] struct {
	X T
	Y T
	Z T
	W T
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

// Make creates a new vector with corresponding 3 components
func Make[T mathex.SignedNumber](x, y, z, w T) Of[T] {
	return Of[T]{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

func MakeT[T mathex.SignedNumber, XT, YT, ZT, WT mathex.Number](x XT, y YT, z ZT, w WT) Of[T] {
	return Of[T]{
		X: T(x),
		Y: T(y),
		Z: T(z),
		W: T(w),
	}
}

func MakeFloat64[XT, YT, ZT, WT mathex.Number](x XT, y YT, z ZT, w WT) Float64 {
	return MakeT[float64](x, y, z, w)
}

func MakeFloat32[XT, YT, ZT, WT mathex.Number](x XT, y YT, z ZT, w WT) Float32 {
	return MakeT[float32](x, y, z, w)
}

func MakeInt[XT, YT, ZT, WT mathex.Number](x XT, y YT, z ZT, w WT) Int {
	return MakeT[int](x, y, z, w)
}

func MakeInt64[XT, YT, ZT, WT mathex.Number](x XT, y YT, z ZT, w WT) Int64 {
	return MakeT[int64](x, y, z, w)
}

func MakeInt32[XT, YT, ZT, WT mathex.Number](x XT, y YT, z ZT, w WT) Int32 {
	return MakeT[int32](x, y, z, w)
}

func MakeInt16[XT, YT, ZT, WT mathex.Number](x XT, y YT, z ZT, w WT) Int16 {
	return MakeT[int16](x, y, z, w)
}

func MakeInt8[XT, YT, ZT, WT mathex.Number](x XT, y YT, z ZT, w WT) Int8 {
	return MakeT[int8](x, y, z, w)
}

// Fill creates a vector where each component is equal to v
func Fill[T mathex.SignedNumber](v T) Of[T] {
	return Of[T]{
		X: v,
		Y: v,
		Z: v,
		W: v,
	}
}

func FromColor(c color.Color) Float64 {
	r, g, b, a := c.RGBA()
	return Make(float64(r)/0xffff, float64(g)/0xffff, float64(b)/0xffff, float64(a)/0xffff)
}

// Zero is (0, 0, 0)
func Zero[T mathex.SignedNumber]() Of[T] {
	return Make[T](0, 0, 0, 0)
}

// One is (1, 1, 1)
func One[T mathex.SignedNumber]() Of[T] {
	return Make[T](1, 1, 1, 1)
}

func Compare[T mathex.SignedNumber](a, b Of[T]) int {
	wc := cmp.Compare(a.W, b.W)
	if wc != 0 {
		return wc
	}
	zc := cmp.Compare(a.Z, b.Z)
	if zc != 0 {
		return zc
	}
	yc := cmp.Compare(a.Y, b.Y)
	if yc != 0 {
		return yc
	}
	return cmp.Compare(a.X, b.X)
}

// Average sums all vector4's components together and divides each
// component by the number of vectors added
func Average[T mathex.SignedNumber](vectors []Of[T]) Of[T] {
	var center Of[T]
	for _, v := range vectors {
		center = center.Add(v)
	}
	return center.DivByConstant(float64(len(vectors)))
}

// Lerp linearly interpolates between a and b by t
// func Lerp[T vector.Number](a, b Vector[T], t float64) Vector[T] {

// 	// (b - a) * t + a
// 	// bt - at + a
// 	// bt - a(1 - t)
// 	tm1 := 1. - t
// 	return Vector[T]{
// 		x: T((float64(b.x) * t) - (float64(a.x) * tm1)),
// 		y: T((float64(b.y) * t) - (float64(a.y) * tm1)),
// 		z: T((float64(b.z) * t) - (float64(a.z) * tm1)),
// 		w: T((float64(b.w) * t) - (float64(a.w) * tm1)),
// 	}
// }

// Lerp linearly interpolates between a and b by t
func Lerp[T mathex.SignedNumber](a, b Of[T], t float64) Of[T] {

	// return b.Sub(a).Scale(t).Add(a)
	return Of[T]{
		X: T((float64(b.X-a.X) * t) + float64(a.X)),
		Y: T((float64(b.Y-a.Y) * t) + float64(a.Y)),
		Z: T((float64(b.Z-a.Z) * t) + float64(a.Z)),
		W: T((float64(b.W-a.W) * t) + float64(a.W)),
	}
}

func (v Of[T]) Negated() Of[T] {
	return Of[T]{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
		W: -v.W,
	}
}

func (v Of[T]) Scale(t float64) Of[T] {
	return Of[T]{
		X: T(float64(v.X) * t),
		Y: T(float64(v.Y) * t),
		Z: T(float64(v.Z) * t),
		W: T(float64(v.W) * t),
	}
}

func (v Of[T]) DivByConstant(t float64) Of[T] {
	return Of[T]{
		X: T(float64(v.X) / t),
		Y: T(float64(v.Y) / t),
		Z: T(float64(v.Z) / t),
		W: T(float64(v.W) / t),
	}
}

func Min[T mathex.SignedNumber](a, b Of[T]) Of[T] {
	return Make(
		min(a.X, b.X),
		min(a.Y, b.Y),
		min(a.Z, b.Z),
		min(a.W, b.W),
	)
}

func Max[T mathex.SignedNumber](a, b Of[T]) Of[T] {
	return Make(
		max(a.X, b.X),
		max(a.Y, b.Y),
		max(a.Z, b.Z),
		max(a.W, b.W),
	)
}

func MaxX[T mathex.SignedNumber](a, b Of[T]) T {
	return max(a.X, b.X)
}

func MaxY[T mathex.SignedNumber](a, b Of[T]) T {
	return max(a.Y, b.Y)
}

func MaxZ[T mathex.SignedNumber](a, b Of[T]) T {
	return max(a.Z, b.Z)
}

func MaxW[T mathex.SignedNumber](a, b Of[T]) T {
	return max(a.W, b.W)
}

func MinX[T mathex.SignedNumber](a, b Of[T]) T {
	return min(a.X, b.X)
}

func MinY[T mathex.SignedNumber](a, b Of[T]) T {
	return min(a.Y, b.Y)
}

func MinZ[T mathex.SignedNumber](a, b Of[T]) T {
	return min(a.Z, b.Z)
}

func MinW[T mathex.SignedNumber](a, b Of[T]) T {
	return min(a.W, b.W)
}

func Midpoint[T mathex.SignedNumber](a, b Of[T]) Of[T] {
	// center = (b - a)0.5 + a
	// center = b0.5 - a0.5 + a
	// center = b0.5 + a0.5
	// center = 0.5(b + a)
	return Of[T]{
		X: T(float64(a.X+b.X) * 0.5),
		Y: T(float64(a.Y+b.Y) * 0.5),
		Z: T(float64(a.Z+b.Z) * 0.5),
		W: T(float64(a.W+b.W) * 0.5),
	}
}

func (v Of[T]) String() string {
	return fmt.Sprintf("X: %v; Y: %v; Z: %v; W: %v;", v.X, v.Y, v.Z, v.W)
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

	if len(data) > 2 {
		v.Z = data[2]
	}

	if len(data) > 3 {
		v.W = data[3]
	}

	return v
}

func (v Of[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
		W float64 `json:"w"`
	}{
		X: float64(v.X),
		Y: float64(v.Y),
		Z: float64(v.Z),
		W: float64(v.W),
	})
}

func (v *Of[T]) UnmarshalJSON(data []byte) error {
	aux := &struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
		W float64 `json:"w"`
	}{
		X: 0,
		Y: 0,
		Z: 0,
		W: 0,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	v.X = T(aux.X)
	v.Y = T(aux.Y)
	v.Z = T(aux.Z)
	v.W = T(aux.W)
	return nil
}

func (v Of[T]) Format(format string) string {
	return fmt.Sprintf(format, v.X, v.Y, v.Z, v.W)
}

func (v Of[T]) MinComponent() T {
	return min(v.X, v.Y, v.Z, v.W)
}

func (v Of[T]) MaxComponent() T {
	return max(v.X, v.Y, v.Z, v.W)
}

func To[T, OT mathex.SignedNumber](v Of[OT]) Of[T] {
	return Of[T]{
		X: T(v.X),
		Y: T(v.Y),
		Z: T(v.Z),
		W: T(v.W),
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
func (v Of[T]) SetX(newX T) Of[T] {
	return Of[T]{
		X: newX,
		Y: v.Y,
		Z: v.Z,
		W: v.W,
	}
}

func (v Of[T]) AddX(dX T) Of[T] {
	return Of[T]{
		X: v.X + dX,
		Y: v.Y,
		Z: v.Z,
		W: v.W,
	}
}

// SetY changes the y component of the vector
func (v Of[T]) SetY(newY T) Of[T] {
	return Of[T]{
		X: v.X,
		Y: newY,
		Z: v.Z,
		W: v.W,
	}
}

func (v Of[T]) AddY(dY T) Of[T] {
	return Of[T]{
		X: v.X,
		Y: v.Y + dY,
		Z: v.Z,
		W: v.W,
	}
}

// SetZ changes the z component of the vector
func (v Of[T]) SetZ(newZ T) Of[T] {
	return Of[T]{
		X: v.X,
		Y: v.Y,
		Z: newZ,
		W: v.W,
	}
}

func (v Of[T]) AddZ(dZ T) Of[T] {
	return Of[T]{
		X: v.X,
		Y: v.Y,
		Z: v.Z + dZ,
		W: v.W,
	}
}

// SetW changes the w component of the vector
func (v Of[T]) SetW(newW T) Of[T] {
	return Of[T]{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
		W: newW,
	}
}

func (v Of[T]) AddW(dW T) Of[T] {
	return Of[T]{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
		W: v.W + dW,
	}
}

// Add takes each component of our vector and adds them to the vector passed
// in, returning a resulting vector
func (v Of[T]) Add(other Of[T]) Of[T] {
	return Of[T]{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
		W: v.W + other.W,
	}
}

func (v Of[T]) Sub(other Of[T]) Of[T] {
	return Of[T]{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
		W: v.W - other.W,
	}
}

func (v Of[T]) ReciprocalF() Of[float32] {
	return Of[float32]{
		X: 1.0 / float32(v.X),
		Y: 1.0 / float32(v.Y),
		Z: 1.0 / float32(v.Z),
		W: 1.0 / float32(v.W),
	}
}

func (v Of[T]) Reciprocal() Of[float64] {
	return Of[float64]{
		X: 1.0 / float64(v.X),
		Y: 1.0 / float64(v.Y),
		Z: 1.0 / float64(v.Z),
		W: 1.0 / float64(v.W),
	}
}

func (v Of[T]) Product() T {
	return v.X * v.Y * v.Z * v.W
}

func (v Of[T]) Dot(other Of[T]) float64 {
	return float64((v.X * other.X) + (v.Y * other.Y) + (v.Z * other.Z) + (v.W * other.W))
}

func (v Of[T]) Normalized() Of[T] {
	return v.DivByConstant(v.Length())
}

func (v Of[T]) Length() float64 {
	return mathex.Sqrt(float64(v.LengthSquared()))
}

func (v Of[T]) LengthF() float32 {
	return mathex.Sqrt(float32(v.LengthSquared()))
}

func (v Of[T]) LengthSquared() T {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z) + (v.W * v.W)
}

// Sqrt applies the Sqrt to each component of the vector
func (v Of[T]) Sqrt() Of[T] {
	return Make(
		mathex.Sqrt(v.X),
		mathex.Sqrt(v.Y),
		mathex.Sqrt(v.Z),
		mathex.Sqrt(v.W),
	)
}

// Abs applies the Abs math operation to each component of the vector
func (v Of[T]) Abs() Of[T] {
	return Make(
		mathex.Abs(v.X),
		mathex.Abs(v.Y),
		mathex.Abs(v.Z),
		mathex.Abs(v.W),
	)
}

func (v Of[T]) Clamp(vmin, vmax T) Of[T] {
	return Of[T]{
		X: mathex.Clamp(v.X, vmin, vmax),
		Y: mathex.Clamp(v.Y, vmin, vmax),
		Z: mathex.Clamp(v.Z, vmin, vmax),
		W: mathex.Clamp(v.W, vmin, vmax),
	}
}

// Round takes each component of the vector and rounds it to the nearest whole
// number
func (v Of[T]) Round() Of[T] {
	return Make(
		mathex.Round(v.X),
		mathex.Round(v.Y),
		mathex.Round(v.Z),
		mathex.Round(v.W),
	)
}

// RoundToInt takes each component of the vector and rounds it to the nearest
// whole number, and then casts it to a int
func (v Of[T]) RoundToInt() Of[int] {
	return Make(
		int(mathex.Round(v.X)),
		int(mathex.Round(v.Y)),
		int(mathex.Round(v.Z)),
		int(mathex.Round(v.W)),
	)
}

// Floor applies the floor math operation to each component of the vector
func (v Of[T]) Floor() Of[T] {
	return Make(
		mathex.Floor(v.X),
		mathex.Floor(v.Y),
		mathex.Floor(v.Z),
		mathex.Floor(v.W),
	)
}

// FloorToInt applies the floor math operation to each component of the vector,
// and then casts it to a int
func (v Of[T]) FloorToInt() Of[int] {
	return Make(
		int(mathex.Floor(v.X)),
		int(mathex.Floor(v.Y)),
		int(mathex.Floor(v.Z)),
		int(mathex.Floor(v.W)),
	)
}

// Ceil applies the ceil math operation to each component of the vector
func (v Of[T]) Ceil() Of[T] {
	return Make(
		mathex.Ceil(v.X),
		mathex.Ceil(v.Y),
		mathex.Ceil(v.Z),
		mathex.Ceil(v.W),
	)
}

// CeilToInt applies the ceil math operation to each component of the vector,
// and then casts it to a int
func (v Of[T]) CeilToInt() Of[int] {
	return Make(
		int(mathex.Ceil(v.X)),
		int(mathex.Ceil(v.Y)),
		int(mathex.Ceil(v.Z)),
		int(mathex.Ceil(v.W)),
	)
}

// MulByVector is component wise multiplication, also known as Hadamard product.
func (v Of[T]) MulByVector(o Of[T]) Of[T] {
	return Of[T]{
		X: v.X * o.X,
		Y: v.Y * o.Y,
		Z: v.Z * o.Z,
		W: v.W * o.W,
	}
}

func (v Of[T]) ContainsNaN() bool {
	if math.IsNaN(float64(v.X)) {
		return true
	}

	if math.IsNaN(float64(v.Y)) {
		return true
	}

	if math.IsNaN(float64(v.Z)) {
		return true
	}

	if math.IsNaN(float64(v.W)) {
		return true
	}

	return false
}

func (v Of[T]) NearZero() bool {
	return mathex.NearZero(v.X) && mathex.NearZero(v.Y) && mathex.NearZero(v.Z) && mathex.NearZero(v.W)
}

func (v Of[T]) Flip() Of[T] {
	return Of[T]{
		X: v.X * -1,
		Y: v.Y * -1,
		Z: v.Z * -1,
		W: v.W * -1,
	}
}

func (v Of[T]) FlipX() Of[T] {
	return Of[T]{
		X: v.X * -1,
		Y: v.Y,
		Z: v.Z,
		W: v.W,
	}
}

func (v Of[T]) FlipY() Of[T] {
	return Of[T]{
		X: v.X,
		Y: v.Y * -1,
		Z: v.Z,
		W: v.W,
	}
}

func (v Of[T]) FlipZ() Of[T] {
	return Of[T]{
		X: v.X,
		Y: v.Y,
		Z: v.Z * -1,
		W: v.W,
	}
}

func (v Of[T]) FlipW() Of[T] {
	return Of[T]{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
		W: v.W * -1,
	}
}

func (v Of[T]) XYZ() vector3.Of[T] {
	return vector3.Make(v.X, v.Y, v.Z)
}

// XY returns vector2 with the x and y components
func (v Of[T]) XY() vector2.Of[T] {
	return vector2.Make(v.X, v.Y)
}

// XZ returns vector2 with the x and z components
func (v Of[T]) XZ() vector2.Of[T] {
	return vector2.Make(v.X, v.Z)
}

// YZ returns vector2 with the y and z components
func (v Of[T]) YZ() vector2.Of[T] {
	return vector2.Make(v.Y, v.Z)
}

// YX returns vector2 with the y and x components
func (v Of[T]) YX() vector2.Of[T] {
	return vector2.Make(v.Y, v.X)
}

// ZX returns vector2 with the z and x components
func (v Of[T]) ZX() vector2.Of[T] {
	return vector2.Make(v.Z, v.X)
}

// ZY returns vector2 with the z and y components
func (v Of[T]) ZY() vector2.Of[T] {
	return vector2.Make(v.Z, v.Y)
}

// Log returns the natural logarithm for each component
func (v Of[T]) Log() Of[T] {
	return Of[T]{
		X: mathex.Log(v.X),
		Y: mathex.Log(v.Y),
		Z: mathex.Log(v.Z),
		W: mathex.Log(v.W),
	}
}

// Log10 returns the decimal logarithm for each component.
func (v Of[T]) Log10() Of[T] {
	return Of[T]{
		X: mathex.Log10(v.X),
		Y: mathex.Log10(v.Y),
		Z: mathex.Log10(v.Z),
		W: mathex.Log10(v.W),
	}
}

// Log2 returns the binary logarithm for each component
func (v Of[T]) Log2() Of[T] {
	return Of[T]{
		X: mathex.Log2(v.X),
		Y: mathex.Log2(v.Y),
		Z: mathex.Log2(v.Z),
		W: mathex.Log2(v.W),
	}
}

// Exp2 returns 2**x, the base-2 exponential for each component
func (v Of[T]) Exp2() Of[T] {
	return Of[T]{
		X: mathex.Exp2(v.X),
		Y: mathex.Exp2(v.Y),
		Z: mathex.Exp2(v.Z),
		W: mathex.Exp2(v.W),
	}
}

// Exp returns e**x, the base-e exponential for each component
func (v Of[T]) Exp() Of[T] {
	return Of[T]{
		X: mathex.Exp(v.X),
		Y: mathex.Exp(v.Y),
		Z: mathex.Exp(v.Z),
		W: mathex.Exp(v.W),
	}
}

// Expm1 returns e**x - 1, the base-e exponential for each component minus 1. It is more accurate than Exp(x) - 1 when the component is near zero
func (v Of[T]) Expm1() Of[T] {
	return Of[T]{
		X: mathex.Expm1(v.X),
		Y: mathex.Expm1(v.Y),
		Z: mathex.Expm1(v.Z),
		W: mathex.Expm1(v.W),
	}
}

func (v Of[T]) Values() (T, T, T, T) {
	return v.X, v.Y, v.Z, v.W
}
