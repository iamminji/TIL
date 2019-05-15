# Scala

마틴 오더스키가 만든 [scala by example](https://www.scala-lang.org/docu/files/ScalaByExample.pdf) 읽어볼 것

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

#### 인자 값으로 함수 사용 (인자 목록 추론)

```
scala> def m1[A](a: A, f: A => String) = f(a)
m1: [A](a: A, f: A => String)String

scala> def m2[A](a: A)(f: A => String) = f(a)
m2: [A](a: A)(f: A => String)String

scala> m1(100, i => s"$i + $i")
<console>:13: error: missing parameter type
       m1(100, i => s"$i + $i")
               ^

scala> m2(100)(i => s"$i + $i")
res1: String = 100 + 100
```

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

### 스칼라의 타입추론
다음은 스칼라에서 명시적으로 타입을 표기해야 하는 경우다.

- `var` 나 `val` 선언에 값을 대입하지 않는 경우.
- 모든 메서드 매개변수
- 메서드 안에서 `return` 을 명시적으로 호출하는 경우
- 메서드가 재귀적인 경우
- 오버로딩한 둘 이상의 메서드가 있고, 그 중 한 메서드가 다른 메서드를 호출하는 경우.(호출하는 메서드에는 반환 타입을 표기해야 함)
- 컴파일러가 추론한 타입이 일반적인 경우(`Any`)

### 스칼라 연산자

스칼라는 기본적으로 마침표와 괄호를 생략할 수 있지만, 아래와 같이 연산자 우선순위에 의해서 예상치 못한 값이 나올 수도 있다.

```
scala> 1.+(2)
res6: Int = 3

scala> 1 + 2 * 3
res10: Int = 7

scala> 1.+(2) * 3
res11: Int = 9
```

아래와 같이 괄호를 생략하는 과정도 가능하다

```
scala> List(1, 2, 3, 4).filter((i: Int) => isEven(i)).foreach((i: Int) => println(i))
scala> List(1, 2, 3, 4).filter(i => isEven(i)).foreach(i => println(i))
scala> List(1, 2, 3, 4).filter(isEven(_)).foreach(println(_))
scala> List(1, 2, 3, 4).filter(isEven).foreach(println)
scala> List(1, 2, 3, 4) filter isEven foreach println
```

### 암시 (implicit)
암시를 사용하는 이유
- 준비를 위한 코드를 없애준다.
- 매개변수화한 타입을 받는 메서드에 사용해서 버그를 줄이거나 허용되는 타입을 제한하기 위한 제약사항으로 사용한다.

#### 암시적 인자 처리 규칙
- 마지막 인자 목록에만 (인자가 단 하나뿐인 메서드의 유일한 인자를 포함해서) 암시적 인자가 들어갈 수 있다.
- `implicit` 키워드는 인자 목록의 맨 처음에 와야하며, 오직 한 번만 나타날 수 있다. 인자 목록 안에서 암시적 인자 다음에 '비암시적' 인자가 따라올 수 없다.
- 인자 목록이 `implicit` 키워드로 시작하면, 그 인자 목록 아느이 모든 인자가 암시적 인자가 된다.

그러나 해당 규칙을 깨는 예외들이 있다.

```
scala> class Bad {
     | def m(i: Int, implicit s: String) = "boo"
<console>:2: error: identifier expected but 'implicit' found.
       def m(i: Int, implicit s: String) = "boo"
                     ^

scala> }
<console>:1: error: eof expected but '}' found.
       }
       ^

scala> class Bad2 {
     | def m(i: Int)(implicit s: String)(implicit d:Double) = "boo"
<console>:2: error: an implicit parameter section must be last
       def m(i: Int)(implicit s: String)(implicit d:Double) = "boo"
                     ^
<console>:2: error: multiple implicit parameter sections are not allowed
       def m(i: Int)(implicit s: String)(implicit d:Double) = "boo"
                                         ^

scala> }
<console>:1: error: eof expected but '}' found.
       }
       ^

scala> class Good1 {
     | def m(i: Int)(implicit s: String, d:Double) = "boo"
     | }
defined class Good1

scala> class Good2 {
     | def m(implicit i:Int, s:String, d:Double) = "boo"
     | }
defined class Good2

scala>
```

### 지연 (lazy)
메서드의 경우 메서드를 호출할 때 마다 본문을 실행한다. `lazy` 값에서는 초기화 본문은 해당 값을 처음으로 사용하는 순간 오직 한번만 평가된다.

### 기타

#### 꼬리 재귀
꼬리 재귀란 마지막 연산이 자기 자신을 호출하는 재귀 연산인 경우를 의미하고 꼬리재귀 최적화(tail call optimization) 는 이 꼬리 재귀 함수를 일반 루프로 변환하는 것이다.
스칼라 자체는 꼬리재귀 최적화를(JVM 언어가 그렇듯) 수행하지 않지만 `tailrec` 어노테이션을 이용하면 가능하다.

다음과 같은 경우에서 num이 0보다 큰 else statement 같은 경우는 꼬리재귀가 아니다. 연산이 있기 때문이다.
```
def recursive(num: Int) : Int = {
    if (num == 0) {
        return 0
    } else {
        return recursive(num-1) + 1
    }
}
```

루프를 사용하면 스택 오버 플로우의 걱정도 없고, 재귀 호출의 비용보다 저렴하다.

## 스칼라 공부 시 참고
- [scala tutorial](https://www.tutorialspoint.com/scala/index.htm)
- [스칼라 도큐먼트 한국어 번역](https://docs.scala-lang.org/ko/tutorials/tour/tour-of-scala.html)
