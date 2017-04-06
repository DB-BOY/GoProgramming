package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	personDB["12345"] = PersonInfo{"12345", "Tom", "Rom 203"}
	personDB["1"] = PersonInfo{"1", "Jack", "Romm 103..."}

	person, ok := personDB["12345"]

	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234")
	} else {
		fmt.Println("Did not Found person with ID 1234")
	}
	delete(personDB, "12345")
	person, ok = personDB["12345"]

	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234")
	} else {
		fmt.Println("Did not Found person with ID 1234")
	}

}
