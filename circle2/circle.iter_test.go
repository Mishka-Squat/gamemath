package circle2_test

import (
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
