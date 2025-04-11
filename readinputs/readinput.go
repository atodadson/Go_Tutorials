package readinputs

import (
	"bufio"
	"fmt"
	"os"
)

// ReadLine prompts the user to enter an input via the terminal \n
// The prompt message is passed as an argument to the func call
// Returns prompts and error if it occurs
// reading the input, the program logs the error and exits.
func ReadLine(message string) (string, error) {
	fmt.Println(message)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return input, nil
}

// ReadTxt reads the content of a textfile and returns the whole content as a string
// It takes the file directory as a string and returns
func ReadTxt(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)
	output := ""
	for scanner.Scan() {
		output = output + scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		return output, err
	}
	return output, nil

}

// func ReadTxt(filepath string) string {
// 	return os.FileInfo(filepath)
// }
