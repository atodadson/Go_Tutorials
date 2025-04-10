package readinputs

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// ReadStringInput prompts the user to enter an input via the terminal,
// The prompts has a message which is passed as an argument to the func call
// reads the input, and returns it as a string. If an error occurs while
// reading the input, the program logs the error and exits.
func ReadStringInput(message string) string {
	fmt.Println(message)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return input
}

// func ReadTxt(filepath string) string {
// 	return os.FileInfo(filepath)
// }
