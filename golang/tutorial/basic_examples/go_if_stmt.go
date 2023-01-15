package basic_examples

import "fmt"

func main() {

	i := 5
	if i < 3 {
		fmt.Println("under 3")
	} else if i == 3 {
		fmt.Println("same")
	} else {
		fmt.Println("over 3")
	}
}
