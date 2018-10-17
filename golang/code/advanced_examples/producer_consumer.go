package advanced_examples

import "fmt"

func producer(c chan<- int) {

	/* 보내기 전용 */

	for i := 0; i < 5; i++ {
		c <- i
	}

	c <- 100
}

func consumer(c <-chan int) {

	/* 받기 전용 */

	for i := range c {
		fmt.Println(i)
	}

	fmt.Println(<-c)
}
