
unit17. for 반복문 사용하기
==
  
  
Go에선 반복문이 __for__ 만 있다. 우리가 아는 for 구문이랑 구조가 똑같다. 다만 표현식 안에 괄호가 없고 여는 괄호가 다음 줄로 넘어가면 안될 뿐이다. (if 구문과 같음)
  
while loop가 없는 대신에 for 로 비슷하게 만들 수 있다.
  
+ for 조건식 { 변화식 }

<pre><code>i := 0
for i < 5 {
    fmt.Println(i)
    i = i + 1
}
</code></pre>
  
17.1 break 사용하기
--
Go엔 label _레이블_  이란게 있는데 중첩된 반복문에서 좋다.
  
<pre><code>Loop:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if j == 2 {
            break Loop
        }
        fmt.Println(i, j)
    }
}
</code></pre>
  
이러면 한 번에 빠져나온다. 단! Loop 레이블과 for 구문 사이에 다른 코드가 있으면 안된다!
  
17.2 continue 사용하기
--
continue도 break 를 위에서 썼던 것과 똑같다. 레이블을 지정해서 할 수 있는데 결과적으로 레이블을 지정한 continue는 레이블이 없는 break 문이랑 같더라.
  
17.3 반복문에서 변수 여러 개 사용하기
--
Go 언어는 반복문의 변화식에서 여러 변수를 처리하려면 병렬 할당을 사용해야 한다.
<pre><code>for i, j :=0, 0; i < 10; i, j = j+1, j+2 {
    fmt.Println(i, j)
}
</code></pre>
  
병렬 할당 하지 않고 i++, j+2 처럼 하면 컴파일 에러가 생긴다.