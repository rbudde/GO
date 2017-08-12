package main

import (
	"de.budde/aspects/intpair"
	"de.budde/aspects/inttriple"
	"fmt"
)

func main() {
	fmt.Printf("Hello!!!\n")

	sp := intpair.Pair{2, 4}
	spp := &sp

	fmt.Printf(sp.Show() + "\n")
	spp.Add(1, 1)
	fmt.Printf(sp.Show() + "\n")

	st := inttriple.Triple{intpair.Pair{6, 6}, 6}
	st.X = 4
	st.Z = 4
	stp := &st
	fmt.Printf(st.Show() + "\n")
	stp.Add(1, 1)
	fmt.Printf(st.Show() + "\n")
}
