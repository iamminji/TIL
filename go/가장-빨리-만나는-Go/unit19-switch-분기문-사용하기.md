
unit19. switch 분기문 사용하기
==
  
  
+ switch 변수 { case 값: 코드 }
  
<pre><code>switch 변수 {
    case 값1:
        // 값1일 때 실행할 코드를 작성한다.
    case 값2:
        // 값2일 때 실행할 코드를 작성한다.
    default:
        // 모든 case에 해당하지 않을 때 실행할 코드를 작성한다.
}
</code></pre>
  
각 __case__ 는 처음부터 순서대로 값을 판단하며 값이 일치하면 해당 코드를 실행한 뒤 switch 분기문을 중단한다. 그리고 모든 case에 해당하지 않을 때는 
__default__ 를 실행한다. 한가지 특이한 점은 다른 언어와 달리 <code>break</code> 키워드를 생략한다.
  
19.1 break 사용하기
--
<code>break</code>를 써서 문장 실행을 중단할 수 있다.
  
19.2 fallthrough 사용하기
--
특정 case의 문장을 실행한 뒤 다음 case의 문장을 실행하고 싶을 때는 <code>fallthrough</code> 키워드를 사용한다. 마치 다른 언어의 
switch 분기문에서 break 키워드를 생략한 것 처럼 동작한다.
  
19.3 여러 조건을 함께 처리하기
--
여러 조건을 같은 문장으로 처리하고 싶을 때는 case에서 콤마로 값을 구분해 준다.
  
+ case 값1, 값2, 값3:
  
19.4 조건식으로 분기하기
--
switch 키워드 다음에 판별할 변수를 지정하지 않고 case에서 조건식만으로 문장을 실행할 수도 있다.
  
+ case 조건식:
  
switch 분기문 안에서 함수를 실행하고 결과값으로 분기를 할 수 있다. 이 때는 함수를 호출하고 뒤에 세미콜론을 붙여준다. 또한 case에서는 값으로 분기할 수 없고 
조건식만 사용할 수 있다.
  
<pre><code>package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    switch i := rand.Intn(10); {
    case i >= 3 && i < 6:
        fmt.Println("3이상 6미만")
    case i == 9:
        fmt.Println("9")
    default:
        fmt.Println(i)
    }
}

</code></pre>