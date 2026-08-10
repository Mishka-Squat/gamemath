package bvh2

import (
	"slices"
	"testing"

	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/stretchr/testify/assert"
)

func TestBvh2(t *testing.T) {
	bvh := Of[float32, int]{}

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(0, 0),
		vector2.MakeFloat32(320, 200),
	), 0)

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(100, 50),
		vector2.MakeFloat32(200, 50),
	), 1)

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(50, 50),
		vector2.MakeFloat32(200, 50),
	), 1)

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(60, 60),
		vector2.MakeFloat32(100, 30),
	), 3)

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(55, 55),
		vector2.MakeFloat32(130, 40),
	), 2)

	v := slices.Collect(bvh.Query(vector2.MakeFloat32(70, 65)))
	assert.ElementsMatch(t, v, []int{0, 1, 2, 3})
}

func TestBvh2Empty(t *testing.T) {
	bvh := Of[float32, int]{}

	v := slices.Collect(bvh.Query(vector2.MakeFloat32(0, 0)))
	assert.Empty(t, v)
}

func TestBvh2Single(t *testing.T) {
	bvh := Of[float32, int]{}

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(0, 0),
		vector2.MakeFloat32(10, 10),
	), 1)

	inside := slices.Collect(bvh.Query(vector2.MakeFloat32(5, 5)))
	assert.ElementsMatch(t, inside, []int{1})

	outside := slices.Collect(bvh.Query(vector2.MakeFloat32(50, 50)))
	assert.Empty(t, outside)
}

// TestBvh2Disjoint covers two rectangles that never touch. Neither can be
// placed as a descendant of the other, so appending the second one has to
// promote a new (virtual) root without losing track of the first.
func TestBvh2Disjoint(t *testing.T) {
	bvh := Of[float32, int]{}

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(0, 0),
		vector2.MakeFloat32(10, 10),
	), 1)

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(100, 100),
		vector2.MakeFloat32(10, 10),
	), 2)

	inFirst := slices.Collect(bvh.Query(vector2.MakeFloat32(5, 5)))
	assert.ElementsMatch(t, inFirst, []int{1})

	inSecond := slices.Collect(bvh.Query(vector2.MakeFloat32(105, 105)))
	assert.ElementsMatch(t, inSecond, []int{2})

	inNeither := slices.Collect(bvh.Query(vector2.MakeFloat32(50, 50)))
	assert.Empty(t, inNeither)
}

// TestBvh2PartialOverlap covers two rectangles that overlap but neither
// contains the other. A query inside the shared region must return both
// values, not just whichever one the tree happens to visit first.
func TestBvh2PartialOverlap(t *testing.T) {
	bvh := Of[float32, int]{}

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(0, 0),
		vector2.MakeFloat32(10, 10),
	), 1)

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(5, 5),
		vector2.MakeFloat32(10, 10),
	), 2)

	overlap := slices.Collect(bvh.Query(vector2.MakeFloat32(7, 7)))
	assert.ElementsMatch(t, overlap, []int{1, 2})

	onlyFirst := slices.Collect(bvh.Query(vector2.MakeFloat32(2, 2)))
	assert.ElementsMatch(t, onlyFirst, []int{1})

	onlySecond := slices.Collect(bvh.Query(vector2.MakeFloat32(12, 12)))
	assert.ElementsMatch(t, onlySecond, []int{2})

	inNeither := slices.Collect(bvh.Query(vector2.MakeFloat32(50, 50)))
	assert.Empty(t, inNeither)
}

// TestBvh2ThreeWayOverlap chains three mutually overlapping rectangles
// (A-B overlap, B-C overlap, A-C don't) so the tree ends up nesting two
// virtual union nodes. Each query point should only pick up the rectangles
// it actually falls inside.
func TestBvh2ThreeWayOverlap(t *testing.T) {
	bvh := Of[float32, int]{}

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(0, 0),
		vector2.MakeFloat32(10, 10),
	), 1) // x:[0,10]

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(5, 0),
		vector2.MakeFloat32(10, 10),
	), 2) // x:[5,15]

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(12, 0),
		vector2.MakeFloat32(10, 10),
	), 3) // x:[12,22]

	aAndB := slices.Collect(bvh.Query(vector2.MakeFloat32(7, 5)))
	assert.ElementsMatch(t, aAndB, []int{1, 2})

	bAndC := slices.Collect(bvh.Query(vector2.MakeFloat32(13, 5)))
	assert.ElementsMatch(t, bAndC, []int{2, 3})

	onlyA := slices.Collect(bvh.Query(vector2.MakeFloat32(2, 5)))
	assert.ElementsMatch(t, onlyA, []int{1})

	onlyC := slices.Collect(bvh.Query(vector2.MakeFloat32(20, 5)))
	assert.ElementsMatch(t, onlyC, []int{3})
}

// TestBvh2WrapMultipleTimes appends progressively larger wrapping
// rectangles and checks that each one correctly becomes the new root while
// everything appended before it stays reachable.
func TestBvh2WrapMultipleTimes(t *testing.T) {
	bvh := Of[float32, int]{}

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(100, 100),
		vector2.MakeFloat32(10, 10),
	), 1) // x:[100,110] y:[100,110]

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(50, 50),
		vector2.MakeFloat32(120, 120),
	), 2) // x:[50,170] y:[50,170], wraps the first rect

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(0, 0),
		vector2.MakeFloat32(300, 300),
	), 3) // x:[0,300] y:[0,300], wraps both previous rects

	all := slices.Collect(bvh.Query(vector2.MakeFloat32(105, 105)))
	assert.ElementsMatch(t, all, []int{1, 2, 3})

	outerTwo := slices.Collect(bvh.Query(vector2.MakeFloat32(60, 60)))
	assert.ElementsMatch(t, outerTwo, []int{2, 3})

	outerOnly := slices.Collect(bvh.Query(vector2.MakeFloat32(200, 200)))
	assert.ElementsMatch(t, outerOnly, []int{3})

	none := slices.Collect(bvh.Query(vector2.MakeFloat32(350, 350)))
	assert.Empty(t, none)
}

func TestBvh2Random(t *testing.T) {
	bvh := Of[float32, int]{}

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(100, 50),
		vector2.MakeFloat32(200, 50),
	), 1)

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(50, 50),
		vector2.MakeFloat32(200, 50),
	), 1)

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(60, 60),
		vector2.MakeFloat32(100, 30),
	), 3)

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(55, 55),
		vector2.MakeFloat32(130, 40),
	), 2)

	bvh = bvh.Append(rect2.MakeFloat32(
		vector2.MakeFloat32(0, 0),
		vector2.MakeFloat32(320, 200),
	), 0)

	v := slices.Collect(bvh.Query(vector2.MakeFloat32(70, 65)))
	assert.ElementsMatch(t, v, []int{0, 1, 2, 3})
}
