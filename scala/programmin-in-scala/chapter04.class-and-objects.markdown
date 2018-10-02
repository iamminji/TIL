
Chapter04. 클래스와 객체
==
  
클래스, 필드, 메서드
--
  
클래스는 객체에 대한 청사진이다. 클래스를 정의하고 나면, 그 클래스 청사진으로부터 new를 사용해 객체를 만들 수 있다.

<pre><code>class ChecksumAccumulator {
    var sum = 0
}
</code></pre>

위와 같은 클래스가 있다고 해보자. 변수 sum을 <code>var</code>로 정의하였다.
_var와 val의 차이는 [이곳](https://github.com/iamminji/TIL/blob/master/scala/etc/var-val-keyword.md) 에서 확인 가능하다_

이를 두 번 인스턴스화 한다.

<pre><code>val acc = new ChecksumAccumulator
val csa = ChecksumAccumulator
</code></pre>

클래스 안에서 정의된 변수 sum이 var 이기 때문에 해당 인스턴스 변수에 재할당이 가능하다. 즉, <code>acc.sum = 10</code>이 가능하다.
acc가 비록 val 로 선언되었지만 객체의 내부는 변경이 가능하다. 단지 acc에 다른 객체를 할당하지 못할뿐이다.
  

스칼라도 자바와 같이 접근 제어자를 사용한다. 외부에 접근이 불가능하게 하고 싶으면 <code>private</code> 키워드를 붙이기만 하면 된다.
스칼라에서 <code>public</code> 으로 지정하고 싶다면 아무것도 하지 않으면 된다.


위의 클래스를 아래와 같이 재정의해본다.
<pre><code>class ChecksumAccumulator {
    private var sum = 0

    /*함수 명 뒤에 오는 것은 반환형으로 Unit은 자바의 void와 같다.*/
    def add(b: Byte): Unit = {
        b = 1 // error!
        sum += b
    }

    def checksum(): Int = {
        ~(sum & 0xFF) + 1
    }
}

</code></pre>

  
<code>add</code>함수를 보자.
__스칼라 메서드 파라미터 에 있어서 중요한 것은 이들은 <code>val</code>이지 <code>var</code>가 아니기 때문에__ 파라미터 b를 바꾸려고 하면 컴파일 에러가 난다.

_파이썬에선 파라미터 값을 변경하여도 전혀 상관이 없었는데, 대신 파이썬은 파라미터가 immutable인지 mutable인지에 따라 결과가 달라지긴 하다._

  
<code>checksum</code> 함수의 경우 반환형이 Int이지만 <code>return</code> 구문이 없다. 스칼라에서는 명시적으로 사용하지 않는것을 권장한다. 그러니 메서드를 작성할 때
메서드가 한 값을 계산하는 표현식이라고 생각해야 한다. 물론 분기에 따라 다른 결과를 내려면 사용하긴 해야 한다.

  
메서드를 더 간결하게 만들 수 있는 방법 하나는, 어떤 메서드가 오직 하나의 표현식만 계산하는 경우 중괄호를 없애는 것이다. 그러면 아래처럼 바뀔 수 있다.

<pre><code>class ChecksumAccumulator {
    private var sum = 0

    def add(b:Byte) = sum += b
    def checksum() = ~(sum&0xFF) + 1
}
</code></pre>


세미 콜론 추론
--
  
스칼라 프로그램에서는 보통 문장 끝의 세미콜론을 생략할 수 있다.


싱글톤 객체
--
  
잠깐 [여기](https://github.com/iamminji/TIL/blob/master/scala/etc/difference-between-object-and-class.md) 에서 언급하고 지나갔던 부분이다.
  
스칼라 클래스에는 _정적(static) 멤버_ 가 없다.
대신에, 스칼라는 _싱글톤(singleton) 객체_ 를 제공한다. 싱글톤 객체 정의는 클래스 정의와 같아 보이지만, <code>class</code> 대신에 <code>object</code>를 사용한다.

  
<pre><code>import scala.collection.mutable

object ChecksumAccumulator {
   private val cache = mutable.Map.empty[String, Int]

   def calculate(s: String): Int =
       if (cache.contains(s))
           cache(s)
       else {
           val acc = new ChecksumAccumulator
           for (c <- s)
               acc.add(c.toByte)
           val cs = acc.checksum()
           cache += (s -> cs)
           cs
       }
}
</code></pre>

위에서 정의한 클래스와 동명의 싱글톤 객체를 정의하였.
이 때 클래스는 동반 클래스(Companion class) 라 하고 싱글톤 객체는 동반 객체(Companion object) 라 한다.
  
싱글톤 객체와 클래스의 차이점 중에 하나는 싱글톤 객체에는 파라미터를 사용할 수 없다. 싱글톤을 인스턴스화 할 수 없기 때문이다. 따라서 이들의 초기화는 자바의 정적 요소를 초기화하는 것과 의미가 동일하다.
어떤 싱글톤 객체의 초기화는 어떤 코드가 그 객체에 처음 접근할 때 일어난다.
  
동반 클래스가 없는 싱글톤 객체를 독립 객체(Standalone object) 라 한다.

  
스칼라 애플리케이션
--
  
스칼라와 자바가 다른 것 중 하나는 스칼라는 그 클래스 이름과 같은 이름의 파일일 필요는 없다. 다만, 클래스와 파일 이름이 같으면 관리 측면에서 편하니 네이밍 규칙을 클래스 이름으로 정하는 것을 권장한다.

스칼라 프로그램을 실행하려면 독립 싱글톤 객체에 <code>main</code> 메서드가 있으며 반환형은 Unit으로, 인자 값은 Array[String] 만 있으면 된다. 바로 아래와 같이 말이다.
  
<pre><code>object Summer{
    def main(args: Array[String]) = {
        /* something */
    }
}
</code></pre>
