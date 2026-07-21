package projection2

import (
	"github.com/Mishka-Squat/gamemath/vector2"
)

type PanZoom struct {
	Pan  vector2.Float32
	Zoom float32
}

func (c PanZoom) Unproject(point vector2.Float32) vector2.Float32 {
	return point.Sub(c.Pan).ScaleF(1 / c.Zoom)
}

func (c PanZoom) Project(point vector2.Float32) vector2.Float32 {
	return point.ScaleF(c.Zoom).Add(c.Pan)
}

func (c *PanZoom) ZoomAt(xy vector2.Float32, zoom float32) {
	old_xy := c.Project(xy)
	c.Zoom = zoom
	new_xy := c.Project(xy)
	d_xy := old_xy.Sub(new_xy)

	c.Pan = c.Pan.Add(d_xy)
}
