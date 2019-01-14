# Chapter11. 스칼라의 계층구조 (Scala's Hierarchy)

스칼라의 모든 클래스는 공통의 슈퍼클래스 `Any` 를 상속한다.

스칼라는 맨 밑바닥에도 클래스가 몇개 있는데 바로 `Null` 과 `Nothing` 이다.
이들은 본질적으로 공통 서브클래스 역할을 하는데 예를 들어서 `Nothing` 은 모든 다른 클래스의 서브 클래스가 되는 것이다.


루트 클래스 `Any` 에는 서브 클래스 두개가 있는데 바로 `AnyVal` 과 `AnyRef` 이다.

###### AnyVal
모든 스칼라의 내장 __값 클래스__ 의 부모 클래스다. ex) Byte, Short, Char, Int, Long, Float, Double, Boolean, Unit

이들 클래스는 `new` 를 사용해 인스턴스화 할 수 없다.

###### AnyRef
이 클래스는 모든 __참조 클래스__ 의 베이스다.
실제 자바에서의 `Object` 를 스칼라에선 `AnyRef` 라고 부르는 것에 지나지 않는다. 따라서 자바나, 스칼라로 작성한 모든 클래스는 `AnyRef` 를 상속한다.
