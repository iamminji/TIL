package examples

import "fmt"

func CountCharacterNum(s1 string) map[string]int {

	var dict = make(map[string]int)

	for _, char := range s1 {
		dict[string(char)] += 1
	}

	return dict
}

func main() {
	r := CountCharacterNum("Hello World!")
	fmt.Println(r)
}
