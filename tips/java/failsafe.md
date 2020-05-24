# FailSafe

Delay 시간을 주면서 Retry 를 하려고 하다가, FailSafe 라는 오픈 소스를 발견했다. Star 도 2k 개가 넘고,
최근 (8월) 까지도 커밋하길래 사용하기로 하였다.

> 사실 Retry 관련해서 스택오버플로우 검색을 했는데, FailSafe 개발자가 열심히 홍보하더라

## Usage
사용하기 전에 `RetryPolicy` 라는 클래스로 정책을 생성해야 한다.

### Example

#### RetryPolicy 생성
아래의 예제는 ConnectionException 이 날 경우 1초의 Delay 로 3번 Try 를 하게 만드는 코드다.
```
RetryPolicy<Object> retryPolicy = new RetryPolicy<>()
  .handle(ConnectException.class)
  .withDelay(Duration.ofSeconds(1))
  .withMaxRetries(3);
```

만약 예외가 아니라 결과에 대한 특정 조건을 주고 싶다면 `handleResultIf` 를 사용하면 된다.
```
RetryPolicy<Integer> retryPolicy = new RetryPolicy<Integer>()
  .handleResultIf(result -> result > 100)
  .withDelay(Duration.ofSeconds(1))
  .withMaxRetries(3);
```

어떤 조건에서 Retry 를 진행하고 싶지 않다면 `abort` 를 사용한다. 아래처럼 RetryPolicy 를 만들면 result 는 100을 넘을 때 Retry 를 진행하고,
만약 ConnectionException 예외가 발생한다면 Retry 를 진행하지 않는다는 의미다.
```
RetryPolicy<Integer> retryPolicy = new RetryPolicy<Integer>()
  .abortOn(ConnectionException.class)
  .handleResultIf(result -> result > 100)
  .withDelay(Duration.ofSeconds(1))
  .withMaxRetries(3);
```

주어진 예제들은 단순히 Delay 만큼의 갭을 주고 시도를 해보는 정책이다. 
만약 backoff 를 하고 싶다면 (시도를 할 때 갭이 아니라 Delay 를 점차 증가시키고 싶다면) withBackoff 를 빌더에 넣으면 된다.

생김새는 아래와 같다. (기본은 2 지수 승이다.)

```
public RetryPolicy<R> withBackoff(long delay, long maxDelay, ChronoUnit chronoUnit) { 
    return withBackoff(delay, maxDelay, chronoUnit, 2); 
}
```

> 이를 [exponential Backoff](https://en.wikipedia.org/wiki/Exponential_backoff) 라 부른다.


`RetryPolicy` 객체를 만들었다면 이제 메서드를 실행한다.

#### 실행
`RetryPolicy` 를 아래처럼 넘기고, 실행시키고자 하는 함수 (예제에선 connect) 를 실행한다.
```
Connection connection = Failsafe.with(retryPolicy).get(() -> connect());
```

Async 하게 실행할 수도 있다.
```
CompletableFuture<Connection> future = Failsafe.with(retryPolicy).getAsync(() -> connect());
```

그 밖에 다양한 사용법, 예제들은 [도큐먼트](https://jodah.net/failsafe/) 혹은 [테스트 코드](https://github.com/jhalterman/failsafe/tree/master/src/test/java/net/jodah/failsafe)를 살펴보면 된다.

## 참고
- [https://jodah.net/failsafe/](https://jodah.net/failsafe/)
- [https://github.com/jhalterman/failsafe](https://github.com/jhalterman/failsafe)
