# Chapter06. 함수형 객체
함수형 객체 _functional object_ 란 변경 가능한 상태를 전혀 갖지 않는 것을 의미한다.

## 자기 참조
현재 실행 중인 메소드의 호출 대상 인스턴스에 대한 참조를 자기 참조 _self reference_ 라고 한다.
생성자 내부에서는 자기 참조가 생성 중인 객체의 인스턴스를 가리킨다.

## 보조 생성자
하나의 클래스에 여러 생성자가 필요한 경우도 있다. 스칼라에서 주 생성자가 아닌 다른 생성자는 보조 생성자 _auxiliary constructor_ 라고 부른다.

<pre><code>class Rational(n: Int, d:Int) {
  require(d != 0)
  val number: Int = n
  val denom: Int = d

  def this(n: Int) = this(n, 1) // 보조 생성자
  override def toString = number + "/" + denom
  def add(that: Rational): Rational =
    new Rational (
          number * that.denom + that.number * denom;
          denom * that.denom
    )
}
</code></pre>

스칼라에서, 모든 보조 생성자는 반드시 같은 클래스에 속한 다른 생성자를 호출하는 코드로 시작해야 한다.
다시 말해, 모든 보조 생성자의 첫 구문은 `this(...)` 여야 한다.
