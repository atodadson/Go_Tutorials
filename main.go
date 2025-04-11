package main

import (
	"fmt"
	"go_tuts/readinputs"
)

func main() {
	name, _ := readinputs.ReadLine("What is your name")
	fmt.Printf("Hello %v", name)

	filename := "textfile.txt"
	output, err := readinputs.ReadTxt(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Type: %T, Value: %v", output, output)
	}

}
