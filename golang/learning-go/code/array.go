package code

import "fmt"

func Array() {
	var x = [3]int{10, 20, 30}
	fmt.Println(x)

	var x2 = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(x2)

	var x3 = [...]int{10, 20, 30}
	fmt.Println(x3)

	var x4 = [3]int{10, 20, 30}
	// true
	fmt.Println(x3 == x4)
}
