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
