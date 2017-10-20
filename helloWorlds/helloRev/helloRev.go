package main

import (
	"de.budde/helloWorlds/stringutil"
	"fmt"
)

func main() {
	var s = "!oG ,olleH"
	stringutil.Double("qay")
	fmt.Printf(stringutil.Reverse(s) + "\n")
	fmt.Printf(stringutil.Double(s) + "\n")
}
