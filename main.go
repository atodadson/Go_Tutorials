package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// name := readinputs.ReadLine("What is your name")
	// fmt.Printf("Hello %v", name)

	filename := "textfile.txt"
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("> %s", line)
		if err != nil {
			return
		}

	}
}
