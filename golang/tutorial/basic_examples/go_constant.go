package basic_examples

import "fmt"

func main() {

	// 선언하고 사용하지 않아도 컴파일 에러가 안나네?
	const age int = 10
	const name string = "minji"

	fmt.Println(age)
	fmt.Println(name)

	const (
		firstName string = "민지"
		lastName  string = "김"
	)
	fmt.Println(lastName + firstName)
}
