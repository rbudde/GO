package main

import (
	"de.budde/helloWorlds/stringutil"
	"fmt"
)

func main() {
	var s = "!oG ,olleH"
	fmt.Printf(stringutil.Reverse(s) + "\n")
	fmt.Printf(stringutil.Double(s) + "\n")
}
