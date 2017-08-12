package inttriple

import (
	"strconv"
)

func (p *Triple) Show11() string {
	return "(" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + ")"
}

func (v *Triple) Sum11() int {
	return v.X + v.Y
}

func (v *Triple) Add11(x, y int) *Triple {
	v.X += x
	v.Y += y
	return v
}
