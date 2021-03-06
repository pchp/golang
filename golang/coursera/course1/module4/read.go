package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname, lname string
}

type Contacts struct {
	persons []Person
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("File path:")
	path, _ := reader.ReadString('\n')
	path = strings.Trim(path, "\n")

	if len(path) == 0 {
		fmt.Println("File path can't be empty")
		fmt.Println("")
		os.Exit(0)
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Couldn't open file path ", path)
		fmt.Println("")
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)

	contacts := Contacts{}

	for scanner.Scan() {
		str := scanner.Text()

		if strings.Contains(str, " ") {
			line := strings.Split(scanner.Text(), " ")
			person := Person{line[0], line[1]}
			contacts.persons = append(contacts.persons, person)
		} else {
			fmt.Print("line ", str, " don't contains ' ' ")
			fmt.Print("")
		}
	}
	fmt.Println("Contacts:", contacts)

	for i, _ := range contacts.persons {
		p := contacts.persons[i]
		fmt.Println("firstName: ", p.fname, " lastName: ", p.lname)
	}

}
