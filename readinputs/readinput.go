package readinputs

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadStringInput() string {
	fmt.Println("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return input
}
