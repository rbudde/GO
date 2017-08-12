package main

import (
	"de.budde/stringutil"
	"fmt"
)

func main() {
	var s = "!oG ,olleH"
	fmt.Printf(stringutil.Reverse(s) + "\n")
	fmt.Printf(stringutil.Double(s) + "\n")
}
