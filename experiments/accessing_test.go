package experiments_test

import (
	"runtime"
	"testing"

	"github.com/Mishka-Squat/gamemath/vector3"
)

var accessResult int

func BenchmarkAccessVector(b *testing.B) {
	var r int

	for n := 0; n < b.N; n++ {
		a := vector3.Make(n, n, n)
		r += a.X + a.Y + a.Z
	}

	accessResult = r
	runtime.KeepAlive(r)
}

func BenchmarkAccessMutableVector(b *testing.B) {
	var r int

	for n := 0; n < b.N; n++ {
		a := MVector[int]{n, n, n}
		r += a.X + a.Y + a.Z
	}

	accessResult = r
	runtime.KeepAlive(r)
}
