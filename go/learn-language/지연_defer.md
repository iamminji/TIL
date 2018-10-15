지연 defer
==
  
  
지연 __defer__ 는 함수 호출이 나중에 프로그램의 실행에서 수행하도록 보장하기 위해 사용된다.(LIFO) 
일반적으로 다른 언어의 <code>finally</code> 블록에서 _cleanup_ 작업을 하는 것 처럼 사용된다.
    
<pre><code>package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.Open("파일명")
    if err ! = nil {
        fmt.Println(err)
    }
    
    // main 마지막에 파일 close 실행
    defer f.Close()
    
    bytes := make([]byte, 1024)
    f.Read(bytes)
    fmt.Println(len(bytes))
}

</code></pre>