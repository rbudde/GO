package main

import (
	"io"
	"os"
	"log"
	"strings"
	"strconv"
	"de.budde/rotX"
)

func main() {
file, err := os.Open("main.go")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	
	argsWoProg := os.Args[1:]
	// zithDlKilt := lh.Zith[1:]
	shift := 0
	if len(argsWoProg) >= 1 {
		shift, _ = strconv.Atoi(argsWoProg[0])
	}
	sr := strings.NewReader("Hello, Reader!\n")
	rr1 := rotX.Make(sr,byte(shift))
	rr2 := rotX.Make(file,byte(shift))
	
	io.Copy(os.Stdout, &rr1)
	io.Copy(os.Stdout, &rr2)
}