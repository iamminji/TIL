package main

import "fmt"

func main() {

	// 맵 만들기
	var a map[string]int = make(map[string]int)
	fmt.Println(a)

	// 맵 초기화 하기
	b := map[string]int{"Hello": 10, "world": 20}

	c := map[string]int{
		"Hello": 20,
		"World": 30,
	}

	fmt.Println(b)
	fmt.Println(c["Hello"])

	// 맵에 데이터 있는지 여부 조회하기
	value, ok := c["No"]
	fmt.Println(value)
	fmt.Println(ok)

	// if 문과 함께 키 조회하기
	if _, exists := c["You"]; exists {
		fmt.Println("Exist!")
	} else {
		fmt.Println("No way...")
	}

	// 맵에서 데이터 삭제하기
	d := map[string]int{"Hello": 10, "World": 20}
	delete(d, "World")
	fmt.Println(d)
}
