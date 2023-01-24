package code

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// receiver
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d.", p.FirstName, p.LastName, p.Age)
}

func ClosureExample() {
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

}

func MakeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
