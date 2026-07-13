package contains2

import (
	"testing"

	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/stretchr/testify/assert"
)

func TestRectRect(t *testing.T) {
	a := rect2.MakeFloat32(
		vector2.MakeFloat32(5, 5),
		vector2.MakeFloat32(20, 40),
	)

	b := rect2.MakeFloat32(
		vector2.MakeFloat32(55, 95),
		vector2.MakeFloat32(20, 40),
	)

	r, q := RectRect(a, b)
	assert.Equal(t, r, Exclude)
	assert.Equal(t, q, Right|Bottom)
}
