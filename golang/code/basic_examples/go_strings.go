package basic_examples

import "fmt"

func main() {
	var s1 string = "한글"
	var s2 string = "Hello"

	fmt.Println(s1)
	fmt.Println(len(s2))

	var s3 string = `문자열
여러줄
쓰기`

	fmt.Println(s3)

	var s4 string = "한글"
	var s5 string = "한글"
	var s6 string = "Go"

	fmt.Println(s4 == s5)
	fmt.Println(s4 + s6 + s5)
	fmt.Println("Hello " + s4)

}
