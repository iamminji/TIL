
unit16. if 조건문 안에서 함수 사용하기
==
  

if 조건문 안에서 함수를 실행한 뒤 조건을 판단하는 방식으로 바꿀 수 있다.  
<pre><code>if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
    fmt.Printf("%s", b)
}
</code></pre>
  
하지만 if 조건문 바깥에서는 변수를 사용할 수 없다.
  
<pre><code>if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
    fmt.Printf("%s", b) // 변수 b를 사용할 수 있음
} else {
    fmt.Println(err) // 변수 err을 사용할 수 있음
}

fmt.Println(b) // 변수 b를 사용할 수 없음, 컴파일 에러
fmt.Println(err) // 변수 err을 사용할 수 없음, 컴파일 에러
</code></pre>
