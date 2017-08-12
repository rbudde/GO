package intpair

import (
	"strconv"
)

func (p *Pair) Show() string {
	return "(" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + ")"
}

func (v *Pair) Sum() int {
	return v.X + v.Y
}

func (v *Pair) Add(x, y int) *Pair {
	v.X += x
	v.Y += y
	return v
}
