package basic_examples

import (
	"fmt"
	"runtime"
)

func hello4() {
	fmt.Println("Hello World!")
}

func f2(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func closureWithGoroutine() {
	runtime.GOMAXPROCS(1)

	s := "Hello, World!"

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()
}

func main() {
	//go f2(0)
	//var input string
	//fmt.Scanln(&input)

	closureWithGoroutine()
}
