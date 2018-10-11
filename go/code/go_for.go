package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 조건식이 없어서 무한 루프와 같음
	// 다른 언어 처럼 break를 써서 빠져나올 수 있다.
	for {
		fmt.Println("Hey!")
		break
	}

Loop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				break Loop
			}
			fmt.Println(i, j)
		}
	}

	fmt.Println("breakLoop!")

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				break
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("Just Break!")

continueLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				continue continueLoop
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("continueLoop!")

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				continue
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("just continue!")

	for i, j := 0, 0; i < 10; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
}
