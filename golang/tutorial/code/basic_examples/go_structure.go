package basic_examples

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func rectangleArea(rect *Rectangle) int {
	return rect.width * rect.height
}

func rectangleScaleA(rect *Rectangle, factor int) {
	rect.width *= factor
	rect.height *= factor
}

func rectangleScaleB(rect Rectangle, factor int) {
	rect.width *= factor
	rect.height *= factor
}

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height}
}

func main() {
	/* 구조체 인스턴스를 생성하면서 값 초기화 */
	var rect Rectangle = Rectangle{10, 20}
	fmt.Println(rect.width, rect.height)

	/* 구조체 포인터 선언, 구조체 포인터에 메모리 할당 */
	var rect1 *Rectangle
	rect1 = new(Rectangle)
	rect1.width = 10
	rect1.height = 20

	fmt.Println(rect1.width, rect1.height)

	/* 구조체 포인터 선언과 동시에 메모리 할당 */
	var rect2 *Rectangle = new(Rectangle)
	rect3 := new(Rectangle)

	fmt.Println("#######################")
	fmt.Println("인스턴스 ->", rect)
	fmt.Println("포인터 ->", rect1)
	fmt.Println("포인터 ->", rect2)
	fmt.Println("포인터 ->", rect3)

	fmt.Println("#########################")

	rect4 := Rectangle{30, 30}
	rect4.width = 20
	area := rectangleArea(&rect4)
	fmt.Println(area)
	fmt.Println(rect4)

	rect1.width = 2000
	area2 := rectangleArea(rect1)
	fmt.Println(area2)
	fmt.Println(rect1)

	fmt.Println("##########################")

	rect5 := Rectangle{20, 20}
	// 포인터로 넘겨주어서 값이 변경 됨
	rectangleScaleA(&rect5, 3)

	rect6 := Rectangle{20, 30}
	// 값이 복사가 되서 원래 값에는 변화가 없음
	rectangleScaleB(rect6, 3)

	fmt.Println(rect5)
	fmt.Println(rect6)
}
