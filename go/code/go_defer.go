package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

func HelloWorld() {
	defer func() {
		fmt.Println("world")
	}()

	func() {
		fmt.Println("Hello")
	}()

}

func main() {
	//defer world() // 현재 함수(main)가 끝나기 직전에 호출
	//hello()
	//hello()
	//hello()

	HelloWorld()
}
