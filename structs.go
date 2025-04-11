package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Defining structs
// starting a field in a struct with a lowercase means the field
// is not to be exported
type Person struct {
	Name   string  `json:"name"`       // Tag line for json format
	Age    int     `db:"age"`          // Tag line for database operations
	Height float32 `validate:"max=20"` // Tag line for validation. Validates when triggered
}

type Group struct {
	Name    string
	Persons []Person
}

func main() {
	person1 := Person{
		Name:   "Ato Kofi",
		Height: 172,
		Age:    19,
	}
	person2 := Person{
		Name:   "Jon Snow",
		Age:    23,
		Height: 191,
	}

	group1 := Group{
		Name: "Group1",
		Persons: []Person{
			person1,
			person2,
		},
	}

	// // Prints the field as well as the values
	// fmt.Printf("%+v", person)

	// // Prints the values of structs
	// fmt.Printf("%+v", person)

	// Json marshalling
	// Converting Go data (struct) to Json format
	byteList, _ := json.Marshal(group1)
	fmt.Printf("Value: %v\n", string(byteList))

	indentByteList, _ := json.MarshalIndent(group1, "", "\t")
	_ = os.WriteFile("myjson.json", indentByteList, 0644)

	var group2 Group

	filename := "myjson.json"
	jsonByte, _ := os.ReadFile(filename)

	json.Unmarshal(jsonByte, &group2)
	group2.Name = "Group II"
	fmt.Printf("Group2: %+v\n", group2)

	fmt.Println(string(indentByteList))
}
