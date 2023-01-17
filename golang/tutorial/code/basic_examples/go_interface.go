package basic_examples

import "fmt"

// 기존 자료형을 새 자료형으로 정의
type MyInt int

func (i MyInt) Print() {
	fmt.Println(i)
}

type Rectangle3 struct {
	width, height int
}

func (r Rectangle3) Print() {
	fmt.Println(r.width, r.height)
}

type Printer interface {
	Print()
}

func main() {
	var i MyInt = 5
	r := Rectangle3{10, 20}
	var p Printer

	p = i
	p.Print()

	p = r
	p.Print()

	fmt.Println("#########################")

	var j MyInt = 10
	p1 := Printer(j)

	r1 := Rectangle3{30, 20}
	p2 := Printer(r1)

	p1.Print()
	p2.Print()

	fmt.Println("##########################")

	pArr := []Printer{j, r1}
	for index, _ := range pArr {
		pArr[index].Print()
	}

	for _, value := range pArr {
		value.Print()
	}
}
