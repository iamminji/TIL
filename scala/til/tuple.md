# 스칼라 튜플

스칼라에서 튜플 선언하는 방법

<pre><code>val tuple = (1, false, "Scala")

val tuple2 = "title" -> “Hey"
</code></pre>

튜플에서 특정 값을 가져오려면 ._(index)를 하면 된다.

<code>println(tuple._3)</code>

스칼라에선 dot 을 생략할 수 있어서 아래도 가능하다.

<code>println(tuple _3)</code>

## Tuple22
스칼라에선 튜플을 만들기 위해 특별히 정의된 클래스를 사용할 수도 있다.

```
scala> Tuple2(1,2)
res19: (Int, Int) = (1,2)

scala> Tuple5(1,2,3,4,5)
res20: (Int, Int, Int, Int, Int) = (1,2,3,4,5)
```

### 기타
바로 Tuple2 ~ Tuple22 인데, 뒤의 숫자는 튜플의 element 개수를 의미한다. 개수와 수가 맞지 않으면 에러가 난다. 찾아보이 `Tuple` 뿐만 아니라 `Function`, `Product` 도 마찬가지였다.

왜 하필 22까지 있는건지 찾아보는데, 나중에 읽어보는걸로... todo...

- [스택오버플로우](https://stackoverflow.com/questions/6241441/why-does-the-scala-library-only-defines-tuples-up-to-tuple22)
- [메일링 리스트](http://scala-programming-language.1934581.n4.nabble.com/Why-tuples-only-to-22-td1945314.html)
