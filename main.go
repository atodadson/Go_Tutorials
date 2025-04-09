package main

import (
	"fmt"
	"go_tuts/readinputs"
)

func main() {
	name := readinputs.ReadStringInput()
	fmt.Printf("Hello %v", name)
}
