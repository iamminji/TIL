package main

import "fmt"

func calc() func(x int) int {
	a, b := 3, 5

	return func(x int) int {
		return a*x + b // 함수 바깥의 변수 a, b 사용
	}

}

func main() {
	sum := func(a, b int) int {
		return a + b
	}

	r := sum(1, 2)
	fmt.Println(r)

	f := calc()
	fmt.Println(f(1))
	fmt.Println(f(2))

}
