package main

import "fmt"

func hello2(n int) {
	n = 2
}

func hello3(n *int) {
	*n = 2
}

func main() {
	var numPtr *int = new(int)
	fmt.Println(numPtr)

	*numPtr = 1
	fmt.Println(*numPtr)

	var num int = 2
	var num2Ptr *int = &num

	fmt.Println(num2Ptr)
	fmt.Println(&num)
	fmt.Println(*num2Ptr)

	var n int = 1

	// 값 복사 - Call By Value
	hello2(n)
	fmt.Println(n)

	var n2 int = 1

	// 참조 값 - Call By Ref
	hello3(&n2)
	fmt.Println(n2)

}
