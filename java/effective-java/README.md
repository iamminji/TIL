# 이펙티브 자바

### Item4 인스턴스화를 막으려거든 private 생성자를 사용하라
생성자를 명시하지 않으면 컴파일러가 자동으로 기본 생성자를 만들어준다. 이 때 매개변수를 받지 않는 public 생성자가 만들어지며 사용자는 자동 생성된 것인지 알지 못한다.

인스턴스화를 막으려면 __private 생성자를 추가하면 된다.__

```java
public class UtilityClass {
    // 기본 생성자가 만들어지는 것을 막는다(인스턴스화 방지용).
    private UtilityClass() {
        throw new AssertrionError();
    }
}
```

이 방식은 상속도 불가능하게 한다.

### Item5 자원을 직접 명시하지 말고 의존 객체 주입을 사용하라
많은 클래스가 하나 이상의 자원에 의존한다.

다음은 정적 유틸리티 클래스를 잘못 사용한 예이다.
아래 코드는 유연하지 않고 테스트 하기가 어렵다.
```java
    public class SpellChecker {
        private static final Lexicon dictionary = ...;

        private SpellChecker() {} // 객체 생성 방지

        public static boolean isValid(String word) { ... }
        public static List<String> suggestions(String typo) { ... }
    }
```

싱글턴을 잘못사용한 예. 역시 유연하지 않고 테스트 하기가 어렵다.
```java
public class SpellChecker {
    private final Lexicon dictionary = ...;

    private SpellChecker(...) {}
    public static SpellChecker INSTANCE = new SpellChecker(...);

    public boolean isValid(String word) { ... }
    public List<String> suggestions(String typo) { ... }
}
```

사용하는 자원 (위의 코드에서는 딕셔너리) 에 따라 동작이 달라지는 클래스에는 정적 유틸리티 클래스나 싱글턴 방식이 적합하지 않다. 이 조건을 만족하려면 인스턴스를 생성할 때 생성자에 필요한 자원을 넘겨주는 방식이 있다.

의존 객체 주입
```java
public class SpellChecker {
    private final Lexicon dictionary;

    public SpellChecker(Lexicon dictionary) {
        this.dictionary = Objects.requireNonNull(dictionary);
    }

    public boolean isValid(String word) { ... }
    public List<String> suggestions(String typo) { ... }
}
```

이 패턴에서 생성자에 자원 팩터리를 넘겨주는 방식으로 변형할 수 있다.

```java
Mosaic create(Supplier<? extends Tile> tileFactory) { ... }
```

### item6 불필요한 객체 생성을 피하라

아래와 같은 코드는 절대 만들지 말자. 이 코드가 반복문 안에 들어가게 된다면 쓸데 없는 String 인스턴스가 엄청 많이 생성될 수도 있다.

```java
String s = new String("어쩌고");
```

대신 아래 코드를 쓴다면 같은 가상 머신 안에서 똑같은 문자열 리터럴을 사용하는 모든 코드가 같은 객체를 재사용함이 보장된다.
```java
String s = "어쩌고";
```

또한 박싱된 기본 타입 보다는 기본 타입을 사용하고, 의도치 않은 오토박싱이 숨어들지 않도록 주의하자.

### item7 다 쓴 객체 참조를 해제하라

자기 메모리를 직접 관리하는 클래스라면 메모리 누수에 주의해야 한다.

### item8 finalizer 와 cleaner 사용을 피하라

`finalizer` 와 `cleaner` 는 즉시 수행된다는 보장이 없다. 자바 언어 명세에서는 수행 시점 뿐만 아니라 수행 여부조차 보장하지 않는다.

### 새 코드에는 무인자 제네릭 자료형을 사용하지 마라
rawtype(무인자) 대신에 제네릭을 사용하자. 형 변환에 있어서 안정적이기 때문이다. (자바 1.5부터 가능하며, 기존에 호환성 때문에 남아있는 것들도 있긴 하다.)

`List` 같은 경우는 `List<E>` 나 `List<Object>` 로 사용하자.

`List` 와 `List<Object>` 의 차이는 `List` 는 형 검사 절차를 완전히 생략한 것이고, `List<Object>` 는 아무 객체나 넣을 수 있다는 것을 컴파일러에게 알리는 것이다.

만약 자료형을 모르는 상태에서 제네릭 자료형을 쓰고 싶거나, 자료형을 신경 쓰고 싶지 않다면 와일드 카드를 사용하자.

### 추상 클래스 대신 인터페이스를 사용하라

- 인터페이스를 사용하면 wrapper class idiom 을 통해 안전하면서도 강력한 기능 개선이 가능하다.
- 추상 클래스를 사용해 자료형을 정의하면 계승 이외의 수단을 사용할 수 없다.

### 인터페이스는 자료형을 정의할 때만 사용하라

### 태그 달린 클래스 대신 클래스 계층을 활용하라
아래와 같이 클래스 안에 `enum` 처럼 기능을 제공하는 __태그__ 가 달린 클래스를 만날 때가 있다.
```
class Figure {
  enum Shape { RECTANGLE, CIRCLE };

  final Shape shape;

  // 태그가 RECTANGLE 일 때만 쓰는 필드
  double length;
  double width;

  // 원을 만드는 생성자
  Figure(double radius) {
    shape = Shape.CIRCLE;
    this.radius = radius;
  }
  ...
}
```

이 클래스에는 다양한 문제가 있다. 서로 다른 기능을 가진 코드가 한 클래스에 모여 있어 가독성도 떨어지고 필요 하지 않은 필드도 생성되므로 메모리 요구량도 늘어난다.

생성자에서 관련 없는 필드를 초기화 하지 않는 한, 필드들을 final 로 선언할 수 없으므로 boilerplate 도 늘어난다.

즉, _태그 기반 클래스 는 너저분한데다 오류 발생 가능성이 높고 효율적이지도 않는다._

태그 달린 클래스는 아래 처럼 변경할 수 있다.

```
abstract class Figure {
  abstract double area();
}

class Circle extends Figure {
  final double radius;

  Circle(double radius) {
    this.radius = radius;
  }

  double area() {
    return Math.PI * (radius * radius);
  }
}
```

### 전략을 표현하고 싶을 때는 함수 객체를 사용하라

### 멤버 클래스는 가능하면 static 으로 선언하라
중첩 클래스에는  __static member class__, __nonstatic member class__,
__anonymous class__, __local class__ 4 가지 종류가 있다.

#### 정적 멤버 클래스 static member class
정적 멤버 클래스는 바깥 클래스의 모든 멤버에 (private 까지) 접근할 수 있다. 주로 헬퍼 클래스 를 정의할 때 사용한다.

#### 비 정적 멤버 클래스 nonstatic member class
비 정적 멤버 클래스의 객체는 바깥 클래스 객체 없이는 존재할 수 없다. 주로 Adapter 를 정의할 때 많이 쓰인다. 즉,
바깥 클래스 객체를 다른 클래스 객체인 것 처럼 보이게 하는 용도이며 아래와 같이 작성한다.

```
public class MySet<E> extends AbstractSet<E> {
  ...

  public Iterator<E> iterator() {
    return new MyIterator();
  }

  private class MyIterator implements Iterator<E> {

  }
}
```

바깥 클래스 객체에 접근할 필요가 없는 멤버 클래스를 정의할 때는 항상 선언문 앞에 __static__ 을 붙여서
비 정적 멤버 클래스 대신 정적 멤버 클래스를 만들자. static 을 생략하면 모든 객체는 내부적으로 바깥 객체에 대한
참조를 유지하게 된다. 그 덕분에 시간과 공간 요구량이 늘어나며, 바깥 객체에 대한 참조를 유지하게 된다.

#### 익명 클래스 anonymous member class
사용하는 순간에 선언하고 객체를 만든다.

#### 지역 클래스 local member class

#### int 상수 대신 enum을 사용하라

#### ordinal 대신 객체 필드를 사용하라
`enum` 클래스에는 `ordinal` 이라는 메서드가 있다. `ordinal` 메서드는 `enum` 클래스에서 구현한 상수값에 매핑되는 정수 값을 리턴하는 메서드이다.
따라서 아래와 같이 사용하는 경우가 있다.

```
public enum Ensemble {
  SOLO, DUET, TRIO;

  public int numberOfMusicians() { return ordinal() + 1; }
}
```

이는 끔찍한 코드다. 만약에 상수 순서가 바뀐다면 해당 메서드의 결과도 달라지게 되기 때문이다. (쉽게 깨진다.) 따라서 `enum` 상수 값에 매핑되는 정수 값을 사용하고 싶다면 아래와 같이 사용하도록 하자.

```
public enum Ensemble {
  SOLO(1), DUET(2), TRIO(3);

  private final int numberOfMusicians;
  Ensemble(int size) { this.numberOfMusicians = size; }
  public int numberOfMusicians() { return numberOfMusicians; }
}
```
#### 비트 필드 대신 EnumSet 을 사용하자

#### ordinal 을 배열 첨자로 사용하는 대신 EnumMap 을 이용하라

#### 확장 가능한 enum 을 만들어야 한다면 인터페이스를 이용하라
