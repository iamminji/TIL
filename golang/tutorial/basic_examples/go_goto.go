package basic_examples

import "fmt"

func main() {
	var a int = 1
	if a == 1 {
		goto ERROR
		b := 1
		fmt.Println(b)
	}

ERROR:
	fmt.Println("ERROR")
}
