
unit24. 함수 사용하기
==
  
  
+ func 함수명() {}
  
<pre><code>package main

import "fmt"

func hello() {
    fmt.Println("Hello World!")
}
</code></pre>
  
함수도 여는 괄호를 밑에 쓰면 컴파일 에러가 발생한다. _Go 그렇게 안 봤는데... 은근히 강제적인데..._
  
24.1 매개 변수와 리턴값 사용하기
--
+ func 함수명(매개변수명 자료형) 리턴값_자료형 {}
  
Go는 리턴 값 자료형에 변수명도 지을 수 있다! 헐!
  
<pre><code>func sum(a int, b int) (r int) {
    r = a + b
    return // 리턴 값 변수를 사용할 때는 return 뒤에 변수를 지정하지 않음
}
</code></pre>
  
24.2 리턴 값 여러개 사용하기
--
Go 언어는 여러 개 리턴할 수도 있다. 파이썬에선 튜플로 받았었고 다른 언어들에서는 레퍼런스도 받고 클래스도 받고 그랬더랬지.
  
+ func 함수명(매개변수명 자료형) (리턴값_자료형1, 리턴값_자료형2) {}
  
값을 받을 때 특정 값을 생략하려면 언더바를 쓰면 된다. _파이썬도 마찬가지_
  
리턴 여러개 일때도 마찬가지로 변수명을 지을 있다.
  
<pre><code>func SumAndDiff(a int, b int) (sum int, diff int) {
    sum = a + b
    diff = a - b
    return
}
</code></pre>
  
24.3 가변인자 사용하기
--
함수의 매개변수 개수가 정해져 있지 않고 유동적으로 변하는 형태를 가변인자라고 한다.
  
+ func 함수명(매개변수명 ...자료형) 리턴값_자료형 {}
  
가변 인자가 슬라이스 타입이므로 슬라이스를 바로 넘겨줄 수도 있다. 하지만 슬라이스 자체는 받을 수 없다는 점! 그러므로 매개 변수에 슬라이스만 넣지 말고 <code>...</code> 도
붙여준다. 이걸 붙이면 슬라이스에 들어있는 요소를 각각 넘겨준다.
  
24.4 재귀 호출 사용하기
--
  
24.5 함수를 변수에 저장하기
--
+ var 변수명 func(매개변수명 자료형) 리턴값_자료형 = 함수명
+ 슬라이스 = []func(매개변수명 자료형) 리턴값_자료형{함수명1, 함수명2}
+ 맵 := map[키_자료형]func(매개변수명 자료형) 리턴값_자료형{"키": 함수명}
  
24.6 익명 함수 사용하기
--
+ func(매개변수명 자료형) 리턴값_자료형 {}()
  
<pre><code>package main

import "fmt"

func main() {
    func() {
        fmt.println("Hello World!")
    }()
    
    func(s string) {
        fmt.Println(s)
    }("Hello, World!")
    
    r := func(a int, b int) int {
        return a + b
    }(1, 2)
    
    fmt.Println(r)
}
</code></pre>
