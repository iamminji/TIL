package basic_examples

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func mutexExample01() {
	//fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	//runtime.GOMAXPROCS(1)

	var data = []int{}
	var mutex = new(sync.Mutex)

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1)
			runtime.Gosched()
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1)
			// 양보
			runtime.Gosched()
			mutex.Unlock()
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(len(data))
}

func mutexExample02() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex)

	c := make(chan bool, 3)

	for i := 0; i < 3; i++ {
		go func(n int) {
			mutex.Lock()
			c <- true
			fmt.Println("wait begin :", n)
			cond.Wait()
			fmt.Println("wait end :", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c
	}

	//for i := 0; i < 3; i++ {
	//    mutex.Lock()
	//    fmt.Println("signal :", i)
	//    cond.Signal()
	//    mutex.Unlock()
	//}

	/* 고루틴 다 깨우기 */
	mutex.Lock()
	fmt.Println("broadcast")
	cond.Broadcast()
	mutex.Unlock()

	fmt.Scanln()
}

func main() {
	mutexExample02()
}
