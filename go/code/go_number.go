package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 10.0
	for i := 0; i < 10; i++ {
		a = a - 0.1
	}

	/* 실수에선 오차 때문에 원하는 결과를 얻지 못할 수 있다. */
	fmt.Println(a)
	fmt.Println(a == 9.0)

	const epsilon = 1e-14
	fmt.Println(math.Abs(a-9.0) <= epsilon)

	// var name rune = "김민지" 컴파일 에러
	// var name rune = '김민지' 컴파일 에러

	var name rune = '김'
	fmt.Println(name)

}
