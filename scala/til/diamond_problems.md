# Scala Diamond Problem

스칼라의 <code>trait</code>은 자바의 인터페이스와 유사하지만 좀 더 유연한 구조를 제공한다. 메서드를 구현할 수 있기 때문이다.
(자바8에서 <code>default</code> 키워드를 사용하면 된다 하던데, 자바를 잘 몰라서 패스.)

또한 <code>trait</code>은 여러개 사용할 수 있다. 이를 <code>mixin(믹스인)</code>~파이썬에서도 믹스인이라는 용어를 사용할 때가 있다. 주로 django에서 클래스 명에 붙이는듯~ 이라 한다.
_대신에, <code>trait</code> 이 아닌 <code>class</code> 는 다중 상속을 허용하지 않는다!_


다중 상속을 허용하지 않는 언어들은 대게 __Diamond Problem__ 때문에 그런 것으로 알고 있다. 스칼라에서도 역시 이 얘기가 나온다.

## Diamond Problem

다이아몬드 문제란 결국엔 다중 상속시 부모들에게 동일한 함수가 있다면 어떤 것을 먼저 호출해야 하는가? 에 대한 것이다.
파이썬에선 method resolution order 라고, 예전에 파이썬 클래스 정리할때 [적어놓은게](https://www.pymoon.com/entry/%ED%8C%8C%EC%9D%B4%EC%8D%AC-%ED%81%B4%EB%9E%98%EC%8A%A4-%EC%83%81%EC%86%8D) 있다.

무튼, 스칼라를 공부하면서 <code>trait</code>이 다중 상속이 된다 하니, 다이아몬드 문제를 어떻게 해결했는지 궁금해서 찾아보았다.


## Multiple Inheritance

위키피디아에 아래와 같이 적혀 있었다.
> Scala allows multiple instantiation of traits, 
which allows for multiple inheritance by adding a distinction between the class hierarchy and the trait hierarchy. 
A class can only inherit from a single class, but can mix-in as many traits as desired. 
Scala resolves method names using a right-first depth-first search of extended 'traits', before eliminating all but the last occurrence of each module in the resulting list. 
So, the resolution order is: [D, C, A, B, A], which reduces down to [D, C, B, A].

쉽게 말해서 가장 우측이 먼저 실행된다는 것 같다. 그리고 중복 되는 호출은 하나만 된다는? 것 같다.
예제를 살펴보자.

<pre><code>trait A {
    def name = println("This is A")
}

trait B extends A {
    override def name: Unit = {
        println("I am B")
        super.name
    }
}

trait C extends A {
    override def name: Unit = {
        println("I am C")
        super.name
    }
}

object Test {
    def main(args: Array[String]): Unit = {
        object O extends B with C
        O.name
    }
}
</pre></code>

결과는 아래와 같다.

<pre><code>I am C
I am B
This is A
</code></pre>

딱 봤을 때 <code>super.name</code> 이 각각 트레잇에 있어서 두 번 호출 될 줄 알았는데 결과는 아니여서 놀랐다.
<code>super</code>를 호출했을 때 <code>class</code>나 <code>trait</code>의 메서드를 호출하게 되는데, 이 때 따르는 룰을 __linearization__  이라 한다.
이 룰이 호출 순서를 결정한다.

이 룰에 의하여 중복된 부분의 콜인 A -> AnyRef -> Any 가 한번만 나오게 된다. C 가 먼저 호출 되므로 최종적으로 
```C -> B -> A -> AnyRef -> Any``` 가 된다카더라. 
결과적으로 다이아몬드 문제를 피할 수 있다.


* 참고
- [https://www.artima.com/pins1ed/traits.html#12.6](https://www.artima.com/pins1ed/traits.html#12.6)
- [https://riptutorial.com/scala/example/13106/solving-the-diamond-problem](https://riptutorial.com/scala/example/13106/solving-the-diamond-problem)
- [https://riptutorial.com/scala/example/14607/linearization](https://riptutorial.com/scala/example/14607/linearization)
