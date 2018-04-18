### 첫번째 스칼라

<pre>
<code>object HelloWorld {
  def main(args: Array[String]) {
    var myVar: Int = 10
    val myVal: String = "Hello Scala with datatype declaration."
    var myVar1 = 20
    val myVal1 = "Hello Scala new without datatype declaration."

    println(myVar)
    println(myVal)
    println(myVar1)
    println(myVal1)

    /* without datatypes */
    println(myVal1 + myVar)
    println(myVar + myVar1)
  }
}
</code></pre>

실행 시키면 아래 처럼 나온다.

<pre>
10
Hello Scala with datatype declaration.
20
Hello Scala new without datatype declaration.
Hello Scala new without datatype declaration.10
30
</pre>

세미콜론이 없어도 되고, 타입을 지정하지 않아도 된다는게 신기하다.
String 에서 덧셈 을 하면 문자열로 합쳐지고, Int는 덧셈 결과가 나온다.
