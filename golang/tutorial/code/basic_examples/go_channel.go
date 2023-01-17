package basic_examples

import (
	"fmt"
	"runtime"
	"time"
)

func sum2(a int, b int, c chan int) {
	c <- a + b
}

func channel1() {
	c := make(chan int)
	go sum2(1, 2, c)

	n := <-c
	fmt.Println(n)
}

// func pinger(c chan string) // 양 방향
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// func printer(c  chan string) // 양 방
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func channel2() {
	var c chan string = make(chan string)

	/* 두 고루틴이 동기화 된다. */
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

func channelWithSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("TimeOut!")
			default:
				fmt.Println("nothing to ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}

func channel3() {
	done := make(chan bool)
	count := 3

	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("고루틴: ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("메인 함수: ", i)
	}
}

func channel4() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2)
	count := 4

	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("고루틴 : ", i)
		}
	}()

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("메인 함수 :", i)
	}
}

func channel5() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		// 채널 닫음
		close(c)
	}()

	a, ok := <-c
	fmt.Println(a, ok)

	// 채널에서 값을 꺼냄
	for i := range c {
		fmt.Println(i)
	}

	b, ok := <-c
	fmt.Println(b, ok)
}

func numWithChannel(a, b int) <-chan int {
	out := make(chan int)
	go func() {
		out <- a
		out <- b
		close(out)
	}()

	return out
}

func sumWithChannel(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		r := 0
		for i := range c {
			r = r + i
		}
		out <- r
	}()
	return out
}

func channel6() {
	c := numWithChannel(1, 2)
	out := sumWithChannel(c)
	fmt.Println(<-out)
}

func main() {
	channel6()
}
