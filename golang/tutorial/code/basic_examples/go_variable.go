package basic_examples

import "fmt"

func main() {
	/* 변수를 선언하는 방법 */

	// 초기값 생략하기
	var i int
	var s string

	fmt.Println(i)
	fmt.Println(s)

	// 자료형 생략하기
	var a = 10
	var b = "Hello"

	fmt.Println(a)
	fmt.Println(b)

	// m = 10 Error
	n := 20

	//fmt.Println(m)
	fmt.Println(n)

	// 변수 여러개 선언하기
	var x, y int = 30, 50
	var p, q = 10, "World"

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(p)
	fmt.Println(q)

	// 선언 뒤에 할당
	var v, t int
	v, t = 3, 5
	fmt.Println(v)
	fmt.Println(t)

	// var와 () 사용하기
	var (
		age, name = 10, "Maria"
	)

	fmt.Println(age)
	fmt.Println(name)

}
