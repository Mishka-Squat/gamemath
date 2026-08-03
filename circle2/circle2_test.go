package circle2_test

import (
	"testing"

	"github.com/Mishka-Squat/gamemath/circle2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/stretchr/testify/assert"
)

func TestMake(t *testing.T) {
	c := circle2.Make(vector2.Make(3, 4), 5)
	assert.Equal(t, vector2.Make(3, 4), c.Center)
	assert.Equal(t, 5, c.Radius)
}

func TestEnumRadiusBelowOneYieldsNothing(t *testing.T) {
	for _, r := range []int{0, -1, -5} {
		count := 0
		for range circle2.Enum(vector2.Make(0, 0), r) {
			count++
		}
		assert.Equal(t, 0, count, "radius=%d", r)
	}
}

func TestEnumNeverYieldsCenter(t *testing.T) {
	center := vector2.Make(4, 4)

	for r := 1; r <= 20; r++ {
		for xy := range circle2.Enum(center, r) {
			assert.NotEqual(t, center, xy, "radius=%d", r)
		}
	}
}

func TestEnumPointsStayWithinRoundedRadius(t *testing.T) {
	center := vector2.Make(10, -7)

	for r := 1; r <= 20; r++ {
		for xy := range circle2.Enum(center, r) {
			dx := float64(xy.X - center.X)
			dy := float64(xy.Y - center.Y)
			dist := dx*dx + dy*dy
			// Bresenham points approximate the radius; allow generous slack
			// rather than asserting exact distance.
			assert.LessOrEqual(t, dist, float64((r+1)*(r+1)), "radius=%d point=%v", r, xy)
		}
	}
}

func TestEnumStopsOnYieldFalse(t *testing.T) {
	count := 0
	for range circle2.Enum(vector2.Make(0, 0), 8) {
		count++
		if count == 3 {
			break
		}
	}
	assert.Equal(t, 3, count)
}

// Enum rasterizes the circle's perimeter, so every grid cell it visits must
// be reported exactly once. Duplicate yields would cause callers (e.g.
// visibility/FOV propagation) to process the same cell more than once.
func TestEnumYieldsEveryPointExactlyOnce(t *testing.T) {
	r := 20
	centers := []vector2.Int{
		vector2.Make(5, 5),
	}

	for _, center := range centers {
		counts := map[vector2.Int]int{}
		for xy := range circle2.Enum(center, r) {
			counts[xy]++
		}

		for xy, n := range counts {
			assert.Equal(t, 1, n, "center=%v radius=%d point=%v yielded %d times", center, r, xy, n)
		}
	}
}
