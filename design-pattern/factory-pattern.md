# 팩토리 패턴

팩토리 클래스는 서로 다른 구상 클래스 객체를 if-else 나 switch 문으로 분기하여 리턴하는 클래스다.

## 팩토리 패턴의 장점
객체 생성 코드를 한 군데에 넣어 둠으로써, 중복 코드를 제거할 수 있다.

클라이언트 입장에서는 객체 인스턴스를 만들 때, Concrete class 가 아닌 인터페이스만 필요로 하게 된다. (인터페이스를 사용함으로써 유연성과 확장성이 높아진다.)

의존성 뒤집기 원칙 (Dependency Inversion Principle) 추상화된 것에 의존하도록 만들어라. Concrete class (구상 클래스) 에 의존하도록 만들지 않도록 한다.

## 추상 팩토리 디자인 패턴
if-else 와 같은 분기문을 없애고, 각각의 서브 클래스를 위한 팩토리 클래스 자체를 가지고 있는 것이 추상 팩토리 클래스 패턴이다.

### 추상 팩토리 디자인 패턴의 이점
- 추상 팩토리 디자인 패턴은 (Concrete class) 구현 보다는 인터페이스를 위한 코드 접근을 제공한다.
- 추상 팩토리 패턴은 팩토리들을 위한 팩토리 이다. (확장에 열려 있다.)
- 추상 팩토리 패턴은 팩토리 패턴의 분기 로직을 피할 수 있다.

## 팩토리 메서드 패턴
팩토리 메서드 패턴에서는 객체를 생성하기 위한 인터페이스를 정의하는데, 어떤 클래스의 인스턴스를 만들지는 서브 클래스에서 결정하게 만든다.

팩토리 메서드 패턴을 이용하는 것은, 클래스의 인스턴스를 만드는 일을 서브 클래스에게 맡기는 것과 같다.

## 팩토리 메서드 vs 추상 팩토리 패턴
- 팩토리 메서드 패턴은 상속을 통해서 객채를 만든다. (클래스를 확장 / 팩토리 메서드를 오버라이드 한다.)
- 추상 팩토리 패턴은 구성을 통해서 객체를 만든다. (추상 형식을 제공한다.)