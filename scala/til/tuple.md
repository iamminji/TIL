# 스칼라 튜플

###### 최초 작성일: 2018-09-29


스칼라에서 튜플 선언하는 방법

<pre><code>val tuple = (1, false, "Scala")

val tuple2 = "title" -> “Hey"
</code></pre>

튜플에서 특정 값을 가져오려면 ._(index)를 하면 된다.

<code>println(tuple._3)</code>

스칼라에선 dot 을 생략할 수 있어서 아래도 가능하다.

<code>println(tuple _3)</code>
