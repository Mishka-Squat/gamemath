package circle2_test

import (
	"sort"
	"testing"

	"github.com/Mishka-Squat/gamemath/circle2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/stretchr/testify/assert"
)

func TestEnumScoreCheckRadiusZeroPanics(t *testing.T) {
	assert.NotPanics(t, func() {
		for range circle2.Make(vector2.Make(5, 5), 0).EnumScoreCheck() {
		}
	})

	assert.NotPanics(t, func() {
		for range circle2.Make(vector2.MakeFloat32(5.5, 4.5), 0.5).EnumScoreCheck() {
		}
	})
}

func TestEnumScoreCheckNeverYieldsCenter(t *testing.T) {
	center := vector2.Make(4, 4)
	c := circle2.Make(center, 6)

	for xy := range c.EnumScoreCheck() {
		assert.NotEqual(t, center, xy)
	}
}

func TestEnumScoreCheckStopsOnYieldFalse(t *testing.T) {
	c := circle2.Make(vector2.Make(0, 0), 8)

	count := 0
	for range c.EnumScoreCheck() {
		count++
		if count == 3 {
			break
		}
	}

	assert.Equal(t, 3, count)
}

// A negative score written into the yielded pointer is the caller's way of
// marking a cell opaque/blocking. Points strictly further out than the
// blocker, along the exact same ray, must not be visited afterwards.
func TestEnumScoreCheckBlockedCellStopsPropagationAlongItsRay(t *testing.T) {
	centers := []vector2.Int{
		vector2.Make(0, 0),
		vector2.Make(37, 21),
	}

	for _, center := range centers {
		radius := 8
		wallX := center.X + 3

		c := circle2.Make(center, radius)
		beyondWall := 0
		sawWall := false
		for xy, score := range c.EnumScoreCheck() {
			if xy.X == wallX {
				sawWall = true
				*score = -1
				continue
			}
			if xy.X > wallX && xy.Y == center.Y {
				beyondWall++
			}
			*score = 1
		}

		assert.True(t, sawWall, "center=%v: the wall cell itself should still be visited", center)
		assert.Equal(t, 0, beyondWall, "center=%v: no cell beyond the wall on the blocked ray should be reached", center)
	}
}

func TestEnumCircleRadiusBelowOneYieldsNothing(t *testing.T) {
	count := 0
	for range circle2.Make(vector2.Make(5, 5), 0).EnumCircle() {
		count++
	}
	assert.Equal(t, 0, count)

	count = 0
	assert.NotPanics(t, func() {
		for range circle2.Make(vector2.MakeFloat32(5.5, 4.5), 0.5).EnumCircle() {
			count++
		}
	})
	assert.Equal(t, 0, count)
}

func TestEnumCircleNeverYieldsCenter(t *testing.T) {
	center := vector2.Make(4, 4)
	c := circle2.Make(center, 6)

	for xy := range c.EnumCircle() {
		assert.NotEqual(t, center, xy)
	}
}

func TestEnumCircleStopsOnYieldFalse(t *testing.T) {
	c := circle2.Make(vector2.Make(0, 0), 8)

	count := 0
	for range c.EnumCircle() {
		count++
		if count == 3 {
			break
		}
	}

	assert.Equal(t, 3, count)
}

// EnumCircle rasterizes a filled disk by walking rays from the center out to
// each Bresenham circle boundary point. Because adjacent octant rays share
// pixels near the center, the same point can legitimately be yielded more
// than once.
func TestEnumCirclePointsCanBeYieldedMultipleTimes(t *testing.T) {
	c := circle2.Make(vector2.Make(0, 0), 2)

	total := 0
	seen := map[vector2.Int]int{}
	for xy := range c.EnumCircle() {
		total++
		seen[xy]++
	}

	assert.Greater(t, total, len(seen), "expected at least one point to be yielded more than once")
}

// All yielded points stay close to the requested radius; the Bresenham
// circle boundary used to drive the rays is only an approximation of a
// true circle, so a small slack is allowed.
func TestEnumCirclePointsStayWithinRadius(t *testing.T) {
	center := vector2.Make(10, -3)
	for _, radius := range []int{1, 2, 3, 5, 8, 13, 21} {
		c := circle2.Make(center, radius)
		for xy := range c.EnumCircle() {
			dist := center.ToFloat64().Distance(xy.ToFloat64())
			assert.LessOrEqual(t, dist, float64(radius)+1.5,
				"radius=%d: point %v is too far from center", radius, xy)
		}
	}
}

// EnumCircle only depends on the offset from the center, so translating the
// circle should translate every yielded point by the same amount.
func TestEnumCircleIsTranslationInvariant(t *testing.T) {
	radius := 5
	offset := vector2.Make(37, -13)

	origin := map[vector2.Int]bool{}
	for xy := range circle2.Make(vector2.Make(0, 0), radius).EnumCircle() {
		origin[xy] = true
	}

	translated := map[vector2.Int]bool{}
	for xy := range circle2.Make(offset, radius).EnumCircle() {
		translated[xy.Sub(offset)] = true
	}

	assert.Equal(t, origin, translated)
}

// Golden-value regression test locking down the exact filled shape produced
// for a small radius, since EnumCircle has no simpler specification to test
// against.
func TestEnumCircleRadius2ExactShape(t *testing.T) {
	expected := []vector2.Int{
		vector2.Make(-1, -2), vector2.Make(0, -2), vector2.Make(1, -2),
		vector2.Make(-2, -1), vector2.Make(0, -1), vector2.Make(2, -1),
		vector2.Make(-2, 0), vector2.Make(-1, 0), vector2.Make(1, 0), vector2.Make(2, 0),
		vector2.Make(-2, 1), vector2.Make(0, 1), vector2.Make(2, 1),
		vector2.Make(-1, 2), vector2.Make(0, 2), vector2.Make(1, 2),
	}

	seen := map[vector2.Int]bool{}
	for xy := range circle2.Make(vector2.Make(0, 0), 2).EnumCircle() {
		seen[xy] = true
	}

	var actual []vector2.Int
	for p := range seen {
		actual = append(actual, p)
	}
	sort.Slice(actual, func(i, j int) bool {
		if actual[i].Y != actual[j].Y {
			return actual[i].Y < actual[j].Y
		}
		return actual[i].X < actual[j].X
	})

	assert.Equal(t, expected, actual)
}

func TestEnumCircleRadius1IsCardinalNeighbors(t *testing.T) {
	expected := map[vector2.Int]bool{
		vector2.Make(1, 0):  true,
		vector2.Make(-1, 0): true,
		vector2.Make(0, 1):  true,
		vector2.Make(0, -1): true,
	}

	actual := map[vector2.Int]bool{}
	for xy := range circle2.Make(vector2.Make(0, 0), 1).EnumCircle() {
		actual[xy] = true
	}

	assert.Equal(t, expected, actual)
}
