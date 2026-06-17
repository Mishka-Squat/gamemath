package vector3

import (
	"cmp"
	"encoding/json"
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/mathex"
)

// Of contains 3 components
type Of[T mathex.SignedNumber] struct {
	X T
	Y T
	Z T
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
func Make[T mathex.SignedNumber](x T, y T, z T) Of[T] {
	return Of[T]{
		X: x,
		Y: y,
		Z: z,
	}
}

func MakeT[T mathex.SignedNumber, XT, YT, ZT mathex.Number](x XT, y YT, z ZT) Of[T] {
	return Of[T]{
		X: T(x),
		Y: T(y),
		Z: T(z),
	}
}

func MakeFloat64[XT, YT, ZT mathex.Number](x XT, y YT, z ZT) Float64 {
	return MakeT[float64](x, y, z)
}

func MakeFloat32[XT, YT, ZT mathex.Number](x XT, y YT, z ZT) Float32 {
	return MakeT[float32](x, y, z)
}

func MakeInt[XT, YT, ZT mathex.Number](x XT, y YT, z ZT) Int {
	return MakeT[int](x, y, z)
}

func MakeInt64[XT, YT, ZT mathex.Number](x XT, y YT, z ZT) Int64 {
	return MakeT[int64](x, y, z)
}

func MakeInt32[XT, YT, ZT mathex.Number](x XT, y YT, z ZT) Int32 {
	return MakeT[int32](x, y, z)
}

func MakeInt16[XT, YT, ZT mathex.Number](x XT, y YT, z ZT) Int16 {
	return MakeT[int16](x, y, z)
}

func MakeInt8[XT, YT, ZT mathex.Number](x XT, y YT, z ZT) Int8 {
	return MakeT[int8](x, y, z)
}

// Fill creates a vector where each component is equal to v
func Fill[T mathex.SignedNumber](v T) Of[T] {
	return Of[T]{
		X: v,
		Y: v,
		Z: v,
	}
}

// Right is (1, 0, 0)
func Right[T mathex.SignedNumber]() Of[T] {
	return Make[T](1, 0, 0)
}

// Left is (-1, 0, 0)
func Left[T mathex.SignedNumber]() Of[T] {
	return Make[T](-1, 0, 0)
}

// Forward is (0, 0, 1)
func Forward[T mathex.SignedNumber]() Of[T] {
	return Make[T](0, 0, 1)
}

// Backwards is (0, 0, -1)
func Backwards[T mathex.SignedNumber]() Of[T] {
	return Make[T](0, 0, -1)
}

// Up is (0, 1, 0)
func Up[T mathex.SignedNumber]() Of[T] {
	return Make[T](0, 1, 0)
}

// Down is (0, -1, 0)
func Down[T mathex.SignedNumber]() Of[T] {
	return Make[T](0, -1, 0)
}

// Zero is (0, 0, 0)
func Zero[T mathex.SignedNumber]() Of[T] {
	return Make[T](0, 0, 0)
}

// One is (1, 1, 1)
func One[T mathex.SignedNumber]() Of[T] {
	return Make[T](1, 1, 1)
}

func Compare[T mathex.SignedNumber](a, b Of[T]) int {
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

func FromColor(c color.Color) Float64 {
	r, g, b, _ := c.RGBA()
	return Make(float64(r)/0xffff, float64(g)/0xffff, float64(b)/0xffff)
}

// Average sums all vector3's components together and divides each
// component by the number of vectors added
func Average[T mathex.SignedNumber](vectors []Of[T]) Of[T] {
	var center Of[T]
	for _, v := range vectors {
		center = center.Add(v)
	}
	return center.DivByConstant(float64(len(vectors)))
}

// Lerp linearly interpolates between a and b by t
func Lerp[T mathex.SignedNumber](t float32, a, b Of[T]) Of[T] {
	return Of[T]{
		X: mathex.Lerp(t, a.X, b.X),
		Y: mathex.Lerp(t, a.Y, b.Y),
		Z: mathex.Lerp(t, a.Z, b.Z),
	}
}

func Min[T mathex.SignedNumber](a, b Of[T]) Of[T] {
	return Make(
		min(a.X, b.X),
		min(a.Y, b.Y),
		min(a.Z, b.Z),
	)
}

func Max[T mathex.SignedNumber](a, b Of[T]) Of[T] {
	return Make(
		max(a.X, b.X),
		max(a.Y, b.Y),
		max(a.Z, b.Z),
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

func MinX[T mathex.SignedNumber](a, b Of[T]) T {
	return min(a.X, b.X)
}

func MinY[T mathex.SignedNumber](a, b Of[T]) T {
	return min(a.Y, b.Y)
}

func MinZ[T mathex.SignedNumber](a, b Of[T]) T {
	return min(a.Z, b.Z)
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
	}
}

func (v Of[T]) String() string {
	return fmt.Sprintf("X: %v; Y: %v; Z: %v;", v.X, v.Y, v.Z)
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

	return v
}

func (v Of[T]) ToArr() []T {
	return []T{v.X, v.Y, v.Z}
}

func (v Of[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	}{
		X: float64(v.X),
		Y: float64(v.Y),
		Z: float64(v.Z),
	})
}

func (v *Of[T]) UnmarshalJSON(data []byte) error {
	aux := &struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	}{
		X: 0,
		Y: 0,
		Z: 0,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	v.X = T(aux.X)
	v.Y = T(aux.Y)
	v.Z = T(aux.Z)
	return nil
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

	return false
}

func (v Of[T]) Format(format string) string {
	return fmt.Sprintf(format, v.X, v.Y, v.Z)
}

func (v Of[T]) MinComponent() T {
	return min(v.X, v.Y, v.Z)
}

func (v Of[T]) MaxComponent() T {
	return max(v.X, v.Y, v.Z)
}

func To[T, OT mathex.SignedNumber](v Of[OT]) Of[T] {
	return Of[T]{
		X: T(v.X),
		Y: T(v.Y),
		Z: T(v.Z),
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
	}
}

func (v Of[T]) AddX(dX T) Of[T] {
	return Of[T]{
		X: v.X + dX,
		Y: v.Y,
		Z: v.Z,
	}
}

// SetY changes the y component of the vector
func (v Of[T]) SetY(newY T) Of[T] {
	return Of[T]{
		X: v.X,
		Y: newY,
		Z: v.Z,
	}
}

func (v Of[T]) AddY(dY T) Of[T] {
	return Of[T]{
		X: v.X,
		Y: v.Y + dY,
		Z: v.Z,
	}
}

// SetZ changes the z component of the vector
func (v Of[T]) SetZ(newZ T) Of[T] {
	return Of[T]{
		X: v.X,
		Y: v.Y,
		Z: newZ,
	}
}

func (v Of[T]) AddZ(dZ T) Of[T] {
	return Of[T]{
		X: v.X,
		Y: v.Y,
		Z: v.Z + dZ,
	}
}

func (v Of[T]) XZY() Of[T] {
	return Of[T]{
		X: v.X,
		Y: v.Z,
		Z: v.Y,
	}
}

func (v Of[T]) ZXY() Of[T] {
	return Of[T]{
		X: v.Z,
		Y: v.X,
		Z: v.Y,
	}
}

func (v Of[T]) ZYX() Of[T] {
	return Of[T]{
		X: v.Z,
		Y: v.Y,
		Z: v.X,
	}
}

func (v Of[T]) YXZ() Of[T] {
	return Of[T]{
		X: v.Y,
		Y: v.X,
		Z: v.Z,
	}
}

func (v Of[T]) YZX() Of[T] {
	return Of[T]{
		X: v.Y,
		Y: v.Z,
		Z: v.X,
	}
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

// Midpoint returns the midpoint between this vector and the vector passed in.
func (v Of[T]) Midpoint(o Of[T]) Of[T] {
	return Of[T]{
		X: T(float64(o.X+v.X) * 0.5),
		Y: T(float64(o.Y+v.Y) * 0.5),
		Z: T(float64(o.Z+v.Z) * 0.5),
	}
}

func (v Of[T]) Barycenter(a, b, c Of[T]) Of[T] {
	v0 := b.Sub(a)
	v1 := c.Sub(a)
	v2 := v.Sub(a)
	d00 := v0.Dot(v0)
	d01 := v0.Dot(v1)
	d11 := v1.Dot(v1)
	d20 := v2.Dot(v0)
	d21 := v2.Dot(v1)

	denom := d00*d11 - d01*d01

	result := Of[T]{}

	result.Y = (d11*d20 - d01*d21) / denom
	result.Z = (d00*d21 - d01*d20) / denom
	result.X = 1.0 - (result.Z + result.Y)

	return result
}

// Perpendicular finds a vector that meets this vector at a right angle.
// https://stackoverflow.com/a/11132720/4974261
func (v Of[T]) Perpendicular() Of[T] {
	var c Of[T]
	if v.Y != 0 || v.Z != 0 {
		c = Right[T]()
	} else {
		c = Up[T]()
	}
	return v.Cross(c)
}

// Round takes each component of the vector and rounds it to the nearest whole
// number
func (v Of[T]) Round() Of[T] {
	return Make(
		mathex.Round(v.X),
		mathex.Round(v.Y),
		mathex.Round(v.Z),
	)
}

// RoundToInt takes each component of the vector and rounds it to the nearest
// whole number, and then casts it to a int
func (v Of[T]) RoundToInt() Of[int] {
	return Make(
		int(mathex.Round(v.X)),
		int(mathex.Round(v.Y)),
		int(mathex.Round(v.Z)),
	)
}

// Floor applies the floor math operation to each component of the vector
func (v Of[T]) Floor() Of[T] {
	return Make(
		mathex.Floor(v.X),
		mathex.Floor(v.Y),
		mathex.Floor(v.Z),
	)
}

// FloorToInt applies the floor math operation to each component of the vector,
// and then casts it to a int
func (v Of[T]) FloorToInt() Of[int] {
	return Make(
		int(mathex.Floor(v.X)),
		int(mathex.Floor(v.Y)),
		int(mathex.Floor(v.Z)),
	)
}

// Ceil applies the ceil math operation to each component of the vector
func (v Of[T]) Ceil() Of[T] {
	return Make(
		mathex.Ceil(v.X),
		mathex.Ceil(v.Y),
		mathex.Ceil(v.Z),
	)
}

// CeilToInt applies the ceil math operation to each component of the vector,
// and then casts it to a int
func (v Of[T]) CeilToInt() Of[int] {
	return Make(
		int(mathex.Ceil(v.X)),
		int(mathex.Ceil(v.Y)),
		int(mathex.Ceil(v.Z)),
	)
}

// Sqrt applies the Sqrt to each component of the vector
func (v Of[T]) Sqrt() Of[T] {
	return Make(
		mathex.Sqrt(v.X),
		mathex.Sqrt(v.Y),
		mathex.Sqrt(v.Z),
	)
}

// Abs applies the Abs math operation to each component of the vector
func (v Of[T]) Abs() Of[T] {
	return Make(
		T(mathex.Abs(v.X)),
		T(mathex.Abs(v.Y)),
		T(mathex.Abs(v.Z)),
	)
}

func (v Of[T]) Clamp(vmin, vmax T) Of[T] {
	return Of[T]{
		X: mathex.Clamp(v.X, vmin, vmax),
		Y: mathex.Clamp(v.Y, vmin, vmax),
		Z: mathex.Clamp(v.Z, vmin, vmax),
	}
}

// Add takes each component of our vector and adds them to the vector passed
// in, returning a resulting vector
func (v Of[T]) Add(other Of[T]) Of[T] {
	return Of[T]{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (v Of[T]) Sub(other Of[T]) Of[T] {
	return Of[T]{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

func (v Of[T]) ReciprocalF() Of[float32] {
	return Of[float32]{
		X: 1.0 / float32(v.X),
		Y: 1.0 / float32(v.Y),
		Z: 1.0 / float32(v.Z),
	}
}

func (v Of[T]) Reciprocal() Of[float64] {
	return Of[float64]{
		X: 1.0 / float64(v.X),
		Y: 1.0 / float64(v.Y),
		Z: 1.0 / float64(v.Z),
	}
}

func (v Of[T]) Product() T {
	return v.X * v.Y * v.Z
}

func (v Of[T]) Dot(other Of[T]) T {
	return (v.X * other.X) + (v.Y * other.Y) + (v.Z * other.Z)
}

func (v Of[T]) Cross(other Of[T]) Of[T] {
	return Of[T]{
		X: (v.Y * other.Z) - (v.Z * other.Y),
		Y: (v.Z * other.X) - (v.X * other.Z),
		Z: (v.X * other.Y) - (v.Y * other.X),
	}
}

func (v Of[T]) Normalized() Of[T] {
	return v.DivByConstant(v.Length())
}

// Rand returns a vector with each component being a random value between [0.0, 1.0)
func Rand(r *rand.Rand) Of[float64] {
	return Of[float64]{
		X: r.Float64(),
		Y: r.Float64(),
		Z: r.Float64(),
	}
}

// RandRange returns a vector where each component is a random value that falls
// within the values of min and max
func RandRange[T mathex.SignedNumber](r *rand.Rand, min, max T) Of[T] {
	dist := float64(max - min)
	return Of[T]{
		X: T(r.Float64()*dist) + min,
		Y: T(r.Float64()*dist) + min,
		Z: T(r.Float64()*dist) + min,
	}
}

// RandInUnitSphere returns a randomly sampled point in or on the unit
func RandInUnitSphere(r *rand.Rand) Of[float64] {
	for {
		p := RandRange(r, -1., 1.)
		if p.LengthSquared() < 1 {
			return p
		}
	}
}

// RandNormal returns a random normal
func RandNormal(r *rand.Rand) Of[float64] {
	return Of[float64]{
		X: -1. + (r.Float64() * 2.),
		Y: -1. + (r.Float64() * 2.),
		Z: -1. + (r.Float64() * 2.),
	}.Normalized()
}

func (v Of[T]) Negated() Of[T] {
	return Of[T]{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

func (v Of[T]) Scale(t float64) Of[T] {
	return Of[T]{
		X: T(float64(v.X) * t),
		Y: T(float64(v.Y) * t),
		Z: T(float64(v.Z) * t),
	}
}

func (v Of[T]) ScaleF(t float32) Of[T] {
	return Of[T]{
		X: T(float32(v.X) * t),
		Y: T(float32(v.Y) * t),
		Z: T(float32(v.Z) * t),
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

func (v Of[T]) Refract(normal Of[T], etaiOverEtat float64) Of[T] {
	cosTheta := min(float64(v.Scale(-1).Dot(normal)), 1.0)
	perpendicular := v.Add(normal.Scale(cosTheta)).Scale(etaiOverEtat)
	parallel := normal.ScaleF(-mathex.Sqrt(mathex.Abs(1.0 - float32(perpendicular.LengthSquared()))))
	return perpendicular.Add(parallel)
}

// MulByVector is component wise multiplication, also known as Hadamard product.
func (v Of[T]) MulByVector(o Of[T]) Of[T] {
	return Of[T]{
		X: v.X * o.X,
		Y: v.Y * o.Y,
		Z: v.Z * o.Z,
	}
}

func (v Of[T]) DivByConstant(t float64) Of[T] {
	return Of[T]{
		X: T(float64(v.X) / t),
		Y: T(float64(v.Y) / t),
		Z: T(float64(v.Z) / t),
	}
}

func (v Of[T]) Length() float64 {
	return mathex.Sqrt(float64(v.LengthSquared()))
}

func (v Of[T]) LengthF() float32 {
	return mathex.Sqrt(float32(v.LengthSquared()))
}

func (v Of[T]) LengthSquared() T {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

func (v Of[T]) DistanceSquared(other Of[T]) T {
	xDist := other.X - v.X
	yDist := other.Y - v.Y
	zDist := other.Z - v.Z
	return T((xDist * xDist) + (yDist * yDist) + (zDist * zDist))
}

func (v Of[T]) Distance(other Of[T]) float64 {
	return mathex.Sqrt(float64(v.DistanceSquared(other)))
}

func (v Of[T]) DistanceF(other Of[T]) float32 {
	return mathex.Sqrt(float32(v.DistanceSquared(other)))
}

func (v Of[T]) Angle(other Of[T]) float64 {
	denominator := mathex.Sqrt(float64(v.LengthSquared() * other.LengthSquared()))
	if denominator < 1e-15 {
		return 0.
	}
	return mathex.Acos(mathex.Clamp(float64(v.Dot(other))/denominator, -1., 1.))
}

func (v Of[T]) AngleF(other Of[T]) float32 {
	denominator := mathex.Sqrt(float32(v.LengthSquared() * other.LengthSquared()))
	if denominator < 1e-15 {
		return 0.
	}
	return mathex.Acos(mathex.Clamp(float32(v.Dot(other))/denominator, -1., 1.))
}

func (v Of[T]) NearZero() bool {
	return mathex.NearZero(v.X) && mathex.NearZero(v.Y) && mathex.NearZero(v.Z)
}

func (v Of[T]) Flip() Of[T] {
	return Of[T]{
		X: v.X * -1,
		Y: v.Y * -1,
		Z: v.Z * -1,
	}
}

func (v Of[T]) FlipX() Of[T] {
	return Of[T]{
		X: v.X * -1,
		Y: v.Y,
		Z: v.Z,
	}
}

func (v Of[T]) FlipY() Of[T] {
	return Of[T]{
		X: v.X,
		Y: v.Y * -1,
		Z: v.Z,
	}
}

func (v Of[T]) FlipZ() Of[T] {
	return Of[T]{
		X: v.X,
		Y: v.Y,
		Z: v.Z * -1,
	}
}

// Log returns the natural logarithm for each component
func (v Of[T]) Log() Of[T] {
	return Of[T]{
		X: mathex.Log(v.X),
		Y: mathex.Log(v.Y),
		Z: mathex.Log(v.Z),
	}
}

// Log10 returns the decimal logarithm for each component.
func (v Of[T]) Log10() Of[T] {
	return Of[T]{
		X: mathex.Log10(v.X),
		Y: mathex.Log10(v.Y),
		Z: mathex.Log10(v.Z),
	}
}

// Log2 returns the binary logarithm for each component
func (v Of[T]) Log2() Of[T] {
	return Of[T]{
		X: mathex.Log2(v.X),
		Y: mathex.Log2(v.Y),
		Z: mathex.Log2(v.Z),
	}
}

// Exp2 returns 2**x, the base-2 exponential for each component
func (v Of[T]) Exp2() Of[T] {
	return Of[T]{
		X: mathex.Exp2(v.X),
		Y: mathex.Exp2(v.Y),
		Z: mathex.Exp2(v.Z),
	}
}

// Exp returns e**x, the base-e exponential for each component
func (v Of[T]) Exp() Of[T] {
	return Of[T]{
		X: mathex.Exp(v.X),
		Y: mathex.Exp(v.Y),
		Z: mathex.Exp(v.Z),
	}
}

// Expm1 returns e**x - 1, the base-e exponential for each component minus 1. It is more accurate than Exp(x) - 1 when the component is near zero
func (v Of[T]) Expm1() Of[T] {
	return Of[T]{
		X: mathex.Expm1(v.X),
		Y: mathex.Expm1(v.Y),
		Z: mathex.Expm1(v.Z),
	}
}

func (v Of[T]) Values() (T, T, T) {
	return v.X, v.Y, v.Z
}
