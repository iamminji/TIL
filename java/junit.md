# Junit

- `setUp()` 매 method 시작전에 실행된다.
- `tearDown()` 매 method 끝난 후에 실행된다. 

## Mockito

### 모킹할 때 리턴과 동시에 어떤 함수가 실행되기를 원할 때
테스트를 진행하다 보면 `thenReturn` 으로 함수의 결과 값을 임의로 지정할 때가 많지만 때로는 함수 실행 시 동작해야 하는 어떤 액션을 지정하고 싶을때도 있다.

즉, 정리해보면 `A()` 함수가 리턴하는 값은 `a` 지만 `B()` 라는 메서드도 실행되고 있다고 가정해보자.
```
A() {
    ...
    B();
    return a;
}
```

주로 모킹은 아래와 같이 할 것이다.

```
Mockito.when(A()).thenReturn(a);
```

`B()` 의 동작은 감춰지게 되는게 보통이다. 그러나 `B()` 가 디비 업데이트라던지, 다른 어떤 외부 동작(?) 일 때 필연적으로 제어해야할 때가 있다. (내가 그랬다.)
그럴 때는 이렇게 사용하면 된다.

```
Mockito.when(A()).thenAnswer(invocation -> {
    B();
    return a;
});
```

그러면 `B()` 제어 뿐만 아니라 리턴해야 하는 값 `a` 도 제대로 동작한다.

##### 참고
- [https://stackoverflow.com/questions/18238709/how-to-make-mockito-do-something-every-time-the-method-on-the-mock-is-called](https://stackoverflow.com/questions/18238709/how-to-make-mockito-do-something-every-time-the-method-on-the-mock-is-called)

### 모킹할 때 함수 호출 순서에 따라 리턴 값을 다르게 주고 싶을 때
무슨 말인가 하면 어떤 함수 `A()` 에 대한 모킹을 할때, 첫번째 호출 될 때는 `a`, 두번째 호출 될 때는 `b` 와 같은 식으로 리턴 받으려는 결과를 호출 순서에 따라 다르게 주고 싶을 때가 있다.

> 예를 들면 처음 호출 때는 실패했지만, 두번째는 성공하는 테스트 코드를 짤 때 같은 경우에

`Mockito` 는 정말 간단하게 이를 지정할 수 있다. 

```
Mockito.when(A()).thenReturn(a).thenReturn(b);
```
또는
```
Mockito.when(A()).thenReturn(a, b);
```
처럼 하면 된다.
