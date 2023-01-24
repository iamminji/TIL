package code

import (
	"fmt"
	"time"
)

// receiver
// 메서드가 리시버를 수정한다면 반드시 포인터 리시버를 사용해야 한다.
// 메서드가 nil 인스턴스를 처리할 필요가 있다면, 반드시 포인터 리시버를 사용해야 한다.
// 메서드가 리시버를 수정하지 않는다면, 값 리시버를 사용할 수 있다.

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func CounterExample() {
	var c Counter
	fmt.Println(c)
	// 이거랑 같음 (&c).Increment()
	c.Increment()
	fmt.Println(c)
}
