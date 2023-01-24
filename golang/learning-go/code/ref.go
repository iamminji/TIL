package code

import "fmt"

type person2 struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person2) {
	i = i * 2
	s = "GoodBye"
	p.name = "Bob"
}

func RefExample() {
	p := person2{}
	i := 2
	s := "Hello"
	// 값이 변경되지 않음
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
}

func modMap(m map[int]string) {
	m[2] = "Hello"
	m[3] = "GoodBye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func RefExample2() {
	m := map[int]string{
		1: "first",
		2: "second",
	}
	// 값 변경 됨
	modMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	// 슬라이스의 요소는 변경이 가능하지만 길이를 늘리는 것은 안된다.
	modSlice(s)
	fmt.Println(s)

	// 맵과 슬라이스는 포인터로 구현되어 있어서 그런 것임
}
