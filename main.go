package main

import (
	"fmt"
	"go_tuts/readinputs"
	"os"
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

	// System arguments
	// Arguments added terminal command when running file
	programName, args := os.Args[0], os.Args[1:]
	fmt.Printf("Args 1: %v, Args: %v", programName, args)
}
