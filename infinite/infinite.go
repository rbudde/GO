package main

import "fmt"

func main() {
	fmt.Printf("Infinite starts.\n")
	count := 0
	for ; ;  {
		count = count + 1;
		count = count - 1;
	}
}