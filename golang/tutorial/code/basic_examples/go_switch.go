package basic_examples

import "fmt"

func main() {

	i := 3
	switch i {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	default:
		fmt.Println(-1)
	}

	// break 를 써서 문장 실행 중단 시키기
	s := "Hello"
	j := 2

	switch j {
	case 1:
		fmt.Println(1)
	case 2:
		if s == "Hello" {
			fmt.Println("Hello 2")
			break
		}
		fmt.Println(2)
	}

	k := 3
	switch k {
	case 4:
		fmt.Println("4 이상")
		fallthrough
	case 3:
		fmt.Println("3 이상")
		fallthrough
	case 2:
		fmt.Println("2 이상")
		fallthrough
	case 1:
		fmt.Println("1 이상")
		fallthrough
	case 0:
		// 마지막 case 전에는 fallthrough를 사용할 수 없다.
		fmt.Println("0 이상")
	}

	// 여러 조건 함께 처리하기
	m := 3

	switch m {
	case 2, 4, 6:
		fmt.Println("짝수")
	case 1, 3, 5:
		fmt.Println("홀수")
	}

	// case 구문에 조건식 사용하기
	n := 7
	switch { // case에 조건식을 지정했으므로 판단할 변수는 생략
	case n >= 5 && n < 10:
		fmt.Println("5 이상 10 미만")
	case n >= 0 && n < 5:
		fmt.Println("0 이상 5 미만")
	}
}
