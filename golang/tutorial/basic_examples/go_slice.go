package basic_examples

import "fmt"

func main() {
	var a []int // int 형 슬라이스 선언
	fmt.Println(a)

	/* 슬라이스 할당하기 */
	var b []int = make([]int, 4)
	fmt.Println(b)

	var c = make([]int, 5)
	fmt.Println(c)

	var s = make([]int, 5, 10)
	fmt.Println(s)

	/* append 로 값 추가하기 */

	// 값 추가
	s = append(s, 1, 2, 3)

	fmt.Println(s)

	// 슬라이스 붙이기
	b = append(b, c...)
	fmt.Println(b)

	// 슬라이스 복사하기
	e := []int{1, 2, 3, 4, 5}
	f := make([]int, 3)

	copy(f, e)

	fmt.Println(e)
	fmt.Println(f)

	// 슬라이스의 길이와 용량
	g := []int{1, 2, 3, 4, 5}
	fmt.Println(len(g), cap(g))

	g = append(g, 6, 7)
	fmt.Println(len(g), cap(g))

	// 부분 슬라이스
	h := []int{1, 2, 3, 4, 5}
	i := h[0:2]
	fmt.Println(i)

	i[0] = 99
	fmt.Println(i)
	fmt.Println(h)
}
