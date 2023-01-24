package code

import "fmt"

type Person3 struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func ptrEx1() {
	p := Person3{
		FirstName: "Pat",
		// 컴파일 안됨
		// 상수는 컴파일 과정에 생성되기 때문에 메모리 주소를 가질 수 없음
		// MiddleName: "Perry",
		LastName: "Peterson",
	}
	fmt.Println(p)

	// 변수에 상수의 값을 들고 있게 하거나
	// 헬퍼함수를 작성해야 함
	p2 := Person3{
		FirstName:  "Pat",
		MiddleName: stringp("Perry"),
		LastName:   "Peterson",
	}
	fmt.Println(p2)
}

func stringp(s string) *string {
	// 상수를 함수로 전달했을 때 상수는 파라미터로 복사된다.
	// 파라미터로 복사된 값은 변수이기 때문에 메모리에서 주소를 가진다.
	// 그래서 함수는 변수의 멤노리 주소를 반환하게 된다.
	return &s
}
