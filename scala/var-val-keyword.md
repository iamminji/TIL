# var & val keyword in scala

###### 최초 작성일: 2018-09-28


스칼라 var / val 키워드

*  스칼라에서 불변(immutable) 변수 선언하기

<code>val book = "Programming Scala"</code>

* 스칼라에서 가변 변수 선언하기

<code>var book = "Programming Scala"</code>


아래와 같을 때 값을 바꿀 수 있을까 없을까?

<pre><code>val array: Array[String] = new Array(5)
array(0) = "Hello World"
</code></pre>

정답은 바꿀 수 있다.

array 객체 자체는 변경하지 못하지만 (immutable) 값은 변경할 수 있다.
