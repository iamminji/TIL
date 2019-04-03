# case class

###### 최초 작성일: 2018-09-27


케이스 클래스 / case class 는 아래와 같은 특징을 가지는 클래스이다.
- 기본적으로 불변 (immutable)
    - 생성자 매개변수가 자동으로 변경 불가한 필드로 바꾸어준다.
- 패턴 매칭을 통해 분해 가능
- 레퍼런스가 아닌 구조적인 동등성으로 비교됨
- 초기화와 운영이 간결함

또한 case 키워드가 붙으면 컴파일러가 이름이 같은 싱글턴 객체인 동반 객체 (Companion Object)를 자동으로 만들어낸다.

동반 객체에 자동으로 추가 되는 메서드도 몇가지 있는데 그 중에 하나가 apply 이다.

<pre><code>val p1 = Point.apply(1 .0, 2@0)
val p2 = Point(1.0, 2.0)
</code></pre>

위 두 가지는 동일하다.

즉 아래 처럼 케이스 클래스(1) 를 생성하면, (2) 와 같은 동반 객체가 생성된다.

(1)
<pre><code>case class Point(x: Double = 0.0, y: Double = 0.0)
</code></pre>
(2)
<pre><code>object Point {
    def apply(x: Double = 0.0, y: Double = 0.0) = new Point(x, y)
}
</code></pre>


참고
* https://docs.scala-lang.org/ko/tutorials/tour/case-classes.html.html
* 프로그래밍 스칼라 2nd Edition
