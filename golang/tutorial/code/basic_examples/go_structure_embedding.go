package basic_examples

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello")
}

type Student struct {
	p      Person /* has a */
	school string
	grade  int
}

type Student2 struct {
	Person /* is a */
	school string
	grade  int
}

func (p *Student2) greeting() {
	/* Person의 greeting을 오버라이드함 */
	fmt.Println("Hello Students")
}

func main() {
	var s Student

	s.p.greeting()

	fmt.Println("#######################")

	var s2 Student2
	s2.Person.greeting()
	s2.greeting()

}
