package basic_examples

import "fmt"

func f() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	a := [...]int{1, 2}

	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	f()
	fmt.Println("Hello World!")
}
