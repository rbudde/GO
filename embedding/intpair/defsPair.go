package intpair

type Pair struct {
	X, Y int
}

type IPair interface {
	Show()
	Sum()
	Add(x int, y int) Pair
}
