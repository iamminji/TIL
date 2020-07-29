# 동시성

## CountDownLatch vs CyclicBarrier

- CyclicBarrier 는 여러 쓰레드가 서로를 기다리고 CountDownLatch 는 하나 또는 다수의 쓰레드가 작업이 완료될 때 까지 기다린다.
- CyclicBarrier 에 모든 쓰레드가 도달하면 값이 초기화 되고 CountDownLatch 는 초기화 되지 않는다. 따라서 CountDownLatch 는 재사용 불가능하다.

### CountDownLatch
쓰레드를 N개 실행했을 때, 일정 개수의 쓰레드가 모두 끝날 때 까지 기다려야지만 다음으로 진행할 수 있거나 다른 쓰레드를 실행시킬 수 있는 경우 사용한다

예를 들면 메인 쓰레드에서 5개의 쓰레드를 실행 시키고, CountDownLatch 값을 3으로 설정해본다고 하자. 각 쓰레드가 종료되는 시점에 `countDown()` 을 호출한다. 그러면 3개의 쓰레드가 종료 될 때 메인 쓰레드가 다시 시작된다.

```java
CountDownLatch countDownLatch = new CountDownLatch(3);
  IntStream.range(0, 5)
    .mapToObj(i -> new Worker(i, countDownLatch))
    .map(Thread::new)
    .forEach(Thread::start);

countDownLatch.await(); // 실행되는 쓰레드 (여기선 메인이라고 하자) 가 다른 쓰레드에서 countDown 이 3번 호출 될 때 까지 기다린다.
```

### CyclicBarrier
CyclicBarrier 모든 스레드들이 Barrier Point 에 도착하는 것을 보장해야 할 때 (쓰레드가 동시에 어떤 작업을 실행해야 하는 것이 필요할 때, 약간 느낌상 synchronized 한 작업이 있어야 할 때?) 사용한다.

### 참고
- https://www.baeldung.com/java-cyclicbarrier-countdownlatch
- https://imasoftwareengineer.tistory.com/100

## Semaphore vs Mutex

### Semaphore
`acquire()` 를 호출하면 permits 를 -1 을 하고 `release()` 를 호출하면 permits 를 반환한다.
permits 가 음수이면 쓰레드는 대기 상태에 빠진다.(예를 들어 -1로 초기화하면 세마포어는 release()가 올 때까지 acquire() 호출을 블락한다.)

아래 예제는 n 이 몇이건 간에 무조건 `foo -> bar` 순서대로 호출하게 하는 동시성 예제이다.
```java
class FooBar {
    private int n;

    // s1 은 permits 를 0, s2 는 permits 를 1로 둔다.
    private Semaphore s1 = new Semaphore(0);
    private Semaphore s2 = new Semaphore(1);

    public FooBar(int n) {
        this.n = n;
    }

    public void foo(Runnable printFoo) throws InterruptedException {

        for (int i = 0; i < n; i++) {

            // s2 의 permits 를 감소 시킨다. 이 값은 0 이 된다.
            s2.acquire();
        	  printFoo.run();
            // s1 의 permits 를 반환한다. (+1)
            s1.release();
        }
    }

    public void bar(Runnable printBar) throws InterruptedException {

        for (int i = 0; i < n; i++) {
            // s1은 0 이기 때문에 감소 시키면 -1 로 빠진다.
            // 아래는 blocking 된다.
            s1.acquire();
        	  printBar.run();
            // s2 의 permits 를 반환한다. (+1)
            s2.release();
        }
    }
}
```
