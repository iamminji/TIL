package code

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func Struct() {

	var fred person
	bob := person{}

	fmt.Println(fred)
	fmt.Println(bob)

	julia := person{
		"Julia",
		40,
		"cat",
	}
	fmt.Println(julia)

	// anonymous struct
	var person2 struct {
		name string
		age  int
		pet  string
	}

	person2.name = "bob"
	person2.age = 50
	person2.pet = "dog"
	fmt.Println(person2)

	pet := struct {
		name string
		kind string
	}{
		name: "Bori",
		kind: "dog",
	}

	fmt.Println(pet)
}
