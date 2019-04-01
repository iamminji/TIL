# var & val keyword in scala

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


자바에서 기본 타입 (primitive type) 이라고 부르는 `char`, `byte`, `short`, `int`, `long`, `float`, `double`,
`boolean` 은 근본적으로 객체(참조 타입) 과는 다르다. 실제로 이런 기본 타입에는 객체도 참조도 존재하지 않는다. 단지 __원래의 값__ 만 존재한다.

그러나 스칼라에서는 이런 타입도 참조 타입과 마찬가지로 실제로 메서드가 있는 객체다.
