package readinputs

import (
	"bufio"
	"encoding/json"
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
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	// Scanner for scanning the file content
	scanner := bufio.NewScanner(file)
	output := ""

	// Scan the file token by token
	for scanner.Scan() {
		output = output + scanner.Text() + "\n" // Seperates tokens by end of line
	}
	if err := scanner.Err(); err != nil {
		return output, err
	}
	return output, nil // return string and nil when successful

}

// ReadJson reads the content of a JSON file.
// Parameters: filename (string), struct_ (pointer to a struct for unmarshalling).
// Returns: error if it occurs.
func ReadJson(filename string, struct_ any) error {
	// Read the file's content.
	byteList, err := os.ReadFile(filename)
	if err != nil {
		return err // Return file read error.
	}

	// Unmarshal JSON into the provided struct pointer.
	unmarshalErr := json.Unmarshal(byteList, struct_)
	if unmarshalErr != nil {
		return unmarshalErr // Return unmarshalling error.
	}

	return nil // Success, no errors.
}
