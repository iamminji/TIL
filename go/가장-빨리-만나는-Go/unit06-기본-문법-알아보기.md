
unit06. Go 언어란?
==
  
  
Go 언어는 문법의 작성 스타일을 강제하고 있다. 컴파일러와는 별도로 문법 스타일을 자동으로 맞춰주는 <code>gofmt</code> 명령도 제공한다.
<pre><code>func main() {
    if i >= 5 {
        fmt.Println("5 이상")
    }
    
    for i:=0; i<5; i++ {
    fmt.Println(i)
    }
}
</code></pre>
  
6.1 세미콜론
--
Go 언어는 보통 구문 마지막의 세미콜론을 생략하고, 한 줄에 여러 구문을 사용할 경우에만 쓴다.
  
6.2 주석
--
한 줄 주석
<pre><code>// fmt.Println("Hello World")
</code></pre>
  
범위 주석
<pre><code>/* fmt.Println("Hello") */
</code></pre>
  
6.3 중괄호
Go 언어의 중괄호는 구문의 맨 뒤에서 시작되어야 한다.

<pre><code>func main()
{   // 컴파일 에러
    var num int = 1
    if num == 1
    {   // 컴파일 에러
        fmt.Println(num)
    }
}
</code></pre>
  
6.4 들여쓰기
--
