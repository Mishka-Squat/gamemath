package experiments_test

import (
	"math/rand"
	"runtime"
	"testing"

	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/mathex"
)

var r = rand.New(rand.NewSource(99))

func ClampGenericMinMax[T mathex.Number](f, vmin, vmax T) T {
	return max(min(f, vmax), vmin)
}

func ClampGenericCompare[T mathex.Number](f, vmin, vmax T) T {
	if f <= vmin {
		return vmin
	}
	if f >= vmax {
		return vmax
	}
	return f
}

func ClampMinMax(f, vmin, vmax float64) float64 {
	return max(min(f, vmax), vmin)
}

func ClampCompare(f, vmin, vmax float64) float64 {
	if f <= vmin {
		return vmin
	}
	if f >= vmax {
		return vmax
	}
	return f
}

func BenchmarkClampMathEx(b *testing.B) {
	var res int
	i := r.Int()
	clamp := vector2.Make(r.Int(), r.Int())
	imin := clamp.MinComponent()
	imax := clamp.MaxComponent()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		res += mathex.Clamp(i, imin, imax)
	}
	runtime.KeepAlive(res)
}

func BenchmarkClampGenericMinMax(b *testing.B) {
	var res int
	i := r.Int()
	clamp := vector2.Make(r.Int(), r.Int())
	imin := clamp.MinComponent()
	imax := clamp.MaxComponent()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		res += ClampGenericMinMax(i, imin, imax)
	}
	runtime.KeepAlive(res)
}

func BenchmarkClampGenericCompare(b *testing.B) {
	var res int
	i := r.Int()
	clamp := vector2.Make(r.Int(), r.Int())
	imin := clamp.MinComponent()
	imax := clamp.MaxComponent()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		res += ClampGenericCompare(i, imin, imax)
	}
	runtime.KeepAlive(res)
}

func BenchmarkClampMinMax(b *testing.B) {
	var res int
	i := r.Int()
	clamp := vector2.Make(r.Int(), r.Int())
	imin := clamp.MinComponent()
	imax := clamp.MaxComponent()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		res += (int)(ClampMinMax(float64(i), float64(imin), float64(imax)))
	}

	runtime.KeepAlive(res)
}

func BenchmarkClampCompare(b *testing.B) {
	var res int
	i := r.Int()
	clamp := vector2.Make(r.Int(), r.Int())
	imin := clamp.MinComponent()
	imax := clamp.MaxComponent()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		res += (int)(ClampCompare(float64(i), float64(imin), float64(imax)))
	}

	runtime.KeepAlive(res)
}
