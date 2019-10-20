# 이펙티브 자바

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
