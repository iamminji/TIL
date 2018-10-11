package main

import "fmt"

func main() {

	/* 배열 초기화 하기 */
	var a [5]int
	a[2] = 7
	fmt.Println(a)

	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 초기화로 지정한 값은 앞에서 차례대로 넣고 나머지 값은 0 으로 채워진다.
	var c = [3]int{1}
	fmt.Println(c)

	d := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(d)

	/* 배열 순회하기 */

	// 배열의 길이를 통한 인덱스 접근 순회
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// iterate
	for i, value := range a {
		fmt.Println(i, value)
	}

	// 인덱스 제외하고 순회하기
	// value 만 가지고 돌면 인덱스가 나온다.
	for _, value := range a {
		fmt.Println(value)
	}

	/* 배열 복사하기 */
	k := b
	fmt.Println(k)
	fmt.Println(b)
}
