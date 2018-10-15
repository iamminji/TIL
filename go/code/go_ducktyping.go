package main

import "fmt"

type Duck struct {
}

func (d Duck) quack() {
	fmt.Println("꽥꽥!!")
}

func (d Duck) feathers() {
	fmt.Println("오리는 흰색과 회색 털을 가지고 있다.")
}

type Person2 struct {
}

func (p Person2) quack() {
	fmt.Println("사람은 오리를 흉내낸다.")
}

func (p Person2) feathers() {
	fmt.Println("사람은 땅에서 깃털을 주워서 보여준다.")
}

type Quacker interface {
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()
	q.feathers()
}

func main() {
	var donald Duck
	var john Person2

	// 어떤 구조체인지와는 상관 없이 메서드만 가지고 있다면 그 메서드가 실행됨
	inTheForest(donald)
	inTheForest(john)

	// 인터페이스를 구현했는지 검사
	if v, ok := interface{}(donald).(Quacker); ok {
		fmt.Println(v, ok)
	}

	if v, ok := interface{}(john).(Quacker); ok {
		fmt.Println(v, ok)
	}
}
