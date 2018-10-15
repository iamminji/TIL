package main

import "fmt"

func CountCharacterNum(s1 string) {

	var dict = make(map[string]int)

	for _, char := range s1 {
		dict[string(char)] += 1
	}

	fmt.Println(dict)
}

func main() {
	CountCharacterNum("Hello World!")
}
