package basic_examples

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func diff(x int, y int) int {
	return x - y
}

func SumAndDiff(a int, b int) (int, int) {
	return a + b, a - b
}

func sum(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}
	return total
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(42, 13))
	fmt.Println(SumAndDiff(1, 2))

	fmt.Println(sum(1, 2, 3, 4, 5))

	n := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(n...))

	// 함수를 변수에 저장한다.
	var hello func(a int, b int) int = add
	world := add

	fmt.Println(hello(1, 2))
	fmt.Println(world(1, 2))

	// 함수를 슬라이스에 저장한다.
	hey := []func(int, int) int{add, diff}
	fmt.Println(hey[0](1, 2))
	fmt.Println(hey[1](1, 2))

	// 함수를 맵에 저장한다.
	you := map[string]func(int, int) int{
		"add":  add,
		"diff": diff,
	}

	fmt.Println(you["add"](1, 2))
	fmt.Println(you["diff"](1, 2))

}
