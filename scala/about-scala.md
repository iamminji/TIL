# Scala

## 스칼라의 특징
- 정적 타입
- 객체지향 프로그래밍
- 함수형 프로그래밍
- 복잡한 타입 시스템
- 간결하고 우아하며 유연한 문법
- 규모 확장성

## 문법

### 변수 정의
스칼라에서는 변수가 불변인지 아닌지 선언 시 지정할 수 있다.
변경이 불가능한 변수는 `val` 이라는 키워드로 선언할 수 있다. 이를 __값 객체__ 라 한다.

<pre><code>val array: Array[String] = new Array(5)
</code></pre>

위와 같은 코드에서 변수 array를 `val` 로 지정하면, 변수 array는 다른 값은 참조하지 못한다. (즉 재할당이 불가능하다는 것이다.)
__배열의 원소는 변경 가능하다__

>`var` 나 `val` 키워드는 어떤 참조가 다른 객체를 참조하도록 변경될 수 있는지(var), 또는 없는지(val) 여부만 지정한다. 참조가 가리키는 대상 객체의 내부 상태를 변경가능한지 여부는 지정하지 않는다.

#### 변수 스코프
스칼라에서 변수는 사용하고 있는 위치에 따라 세 가지 다른 스코프를 가진다.

##### Fields
필드는 객체에 속해 있는 변수다. 필드는 객체의 모든 메서드 안에서 접근이 가능할 뿐만 아니라,
객체 밖에서도 접근 지시자를 선언 했다면 접근이 가능하다.
객체 필드는 mutable과 immutable 타입으로 지정할 수 있고 이는 앞서 말했듯 `var` 과 `val` 로 선언하면 된다.

##### Method Parameters
메서드 파라미터는 메서드 안에서 사용되는 변수다. 메서드 파라미터는 오직 메서드 안에서만 접근이 가능하지만, 메서드 밖에서 객체를 참조한다면 객체를 통해서는 접근 가능할 수 있다(?)

메서드 파라미터는 항상 immutable 하다. (`val` 키워드로 선언해야 한다.)

##### Local Variables
로컬 변수는 메서드 안에서 선언된 것이다. 로컬 변수는 오직 메서드 안에서만 접근 가능하지만 메서드에서 객체를 생성하고 리턴한다면 벗어날 수 있다(?) 로컬 변수는 `var` 과 `val` 둘 다 사용가능하다.

### range

<pre><code>1 to 10 // 1~10
1 until 10 // 1~9
1 to 10 by 3 // 1, 4, 7, 10
</code></pre>

### 함수 선언
<pre><code>def greet() { println("hi") }
</code></pre>

>등호가 없으면 결과 타입이 `Unit` 인 `프로시저` 라고 부른다. (사실 똑같음)

### 예외처리
스칼라는 자바와는 다르게 `finally` 구문안에서 무조건 값을 덮어 씌우지 않는다. 명시적으로 `return` 문을 사용하거나 예외를 발생시켜야 원래의 결과를 덮어 씌운다.

### 상위 타입 경계 Upper type Bounds
상위 타입 경계는 `T <: A` 형식으로 사용하며 이는 타입 `A`의 서브 타입 `T` 란 의미이다.

```
abstract class Animal {
 def name: String
}

abstract class Pet extends Animal {}

class Cat extends Pet {
  override def name: String = "Cat"
}

class Dog extends Pet {
  override def name: String = "Dog"
}

class Lion extends Animal {
  override def name: String = "Lion"
}

class PetContainer[P <: Pet](p: P) {
  def pet: P = p
}

// Pet 의 서브타입인 Dog, Cat 은 PetContainer 에 담기지만...
val dogContainer = new PetContainer[Dog](new Dog)
val catContainer = new PetContainer[Cat](new Cat)

// Lion 은 실패한다! Lion 은 Pet의 subtype 이 아니고 Animal 이기 때문이다.
val lionContainer = new PetContainer[Lion](new Lion)
```

클래스 `PetContainer` 는 파라미터 `P` 를 가지는데 이는 `Pet` 의 서브 타입인 `Dog`, `Cat` 이 된다.

### 하위 타입 경계 Lower Type Bounds


## 스칼라 공부 시 참고
- [scala tutorial](https://www.tutorialspoint.com/scala/index.htm)
- [스칼라 도큐먼트 한국어 번역](https://docs.scala-lang.org/ko/tutorials/tour/tour-of-scala.html)
