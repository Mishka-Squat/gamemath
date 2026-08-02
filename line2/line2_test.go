package line2_test

import (
	"testing"

	"github.com/Mishka-Squat/gamemath/line2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/stretchr/testify/assert"
)

func collect[T any](seq func(yield func(T) bool)) []T {
	var out []T
	for v := range seq {
		out = append(out, v)
	}
	return out
}

func TestEnumIncludesEndpoints(t *testing.T) {
	tests := map[string]struct {
		a vector2.Int
		b vector2.Int
	}{
		"horizontal +x":       {a: vector2.Make(0, 0), b: vector2.Make(5, 0)},
		"horizontal -x":       {a: vector2.Make(5, 0), b: vector2.Make(0, 0)},
		"vertical +y":         {a: vector2.Make(0, 0), b: vector2.Make(0, 5)},
		"vertical -y":         {a: vector2.Make(0, 5), b: vector2.Make(0, 0)},
		"diagonal +x+y":       {a: vector2.Make(0, 0), b: vector2.Make(4, 4)},
		"diagonal -x-y":       {a: vector2.Make(4, 4), b: vector2.Make(0, 0)},
		"diagonal +x-y":       {a: vector2.Make(0, 4), b: vector2.Make(4, 0)},
		"shallow slope":       {a: vector2.Make(0, 0), b: vector2.Make(7, 2)},
		"steep slope":         {a: vector2.Make(0, 0), b: vector2.Make(2, 7)},
		"negative quadrant":   {a: vector2.Make(-3, -2), b: vector2.Make(2, 3)},
		"single point (a==b)": {a: vector2.Make(2, 2), b: vector2.Make(2, 2)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			points := collect(line2.Enum(tc.a, tc.b))

			assert.Contains(t, points, tc.a, "a should be included in the enumerated line")
			assert.Contains(t, points, tc.b, "b should be included in the enumerated line")
		})
	}
}

func TestEnumMethodIncludesEndpoints(t *testing.T) {
	l := line2.Make(vector2.Make(1, 1), vector2.Make(6, 3))

	points := collect(l.Enum())

	assert.Contains(t, points, l.A)
	assert.Contains(t, points, l.B)
}
