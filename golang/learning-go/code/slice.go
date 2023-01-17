package code

import "fmt"

func Slice() {
	var x = []int{10, 20, 30}
	fmt.Println(x)

	var x1 = []int{1, 2, 3}
	x1 = append(x1, 4)
	fmt.Println(x1)
	x1 = append(x1, 5, 6, 7)
	fmt.Println(x1)

	y := []int{20, 30, 40}
	x1 = append(x1, y...)
	fmt.Println(x1)
}
