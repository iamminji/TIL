# JDK14

자바 14 new feature  정리

## Records
jdk 14 에 새롭게 등장한 예약어이다. 코딩을 하다 보면 객체 값이 변하지 않기를 원하는 경우가 있다.


```
public class OrgMessage {
    private final int id;
    private final String message;

    public RequestDataOrg(int id, String message) {
        this.id = id;
        this.message = message;
    }

    public int getId() {
        return id;
    }

    public String getMessage() {
        return message;
    }
}
```

이 처럼 클래스 변수들을 `final` 로 선언하고 `setter` 를 만들지 않는 경우가 대부분일 것이다. jdk14 부터는 `records` 라는 개념이 존재 하며 기본적으로 `immutable` 한 객체를 만들고 싶을 때 사용한다. 즉 위의 코드를 `records` 로 변경하면 이렇게 바꿀 수 있다.


```java
public record NewMessage(int id, String message) {}
```

여기엔 `getter` 도 `setter` 도 없다.

[참고](https://openjdk.java.net/jeps/359)
