상태 없는 객체는 항상 스레드 안전하다.
=======
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

## 쓰레드

### 작업마다 쓰레드를 만들기
어플리케이션에서 반응 속도를 높이려면 요청이 들어올 떄 마다 쓰레드를 생성해서 실행 시키는 방법이 있다. 예를 들면 아래 처럼 서버 쪽에서 connection 을 쓰레드가 처리 하는 것이다.

```java
public class SimpleWebServer {

    public static void main(String[] args) throws IOException {
        ServerSocket socket = new ServerSocket(80);
        while (true) {
            final Socket connection = socket.accept();
            Runnable task = () -> System.out.println("Do something with connetion!");
            new Thread(task).start();
        }
    }
}
```

이렇게 처리하면 3가지 이점을 얻을 수 있다.

1. 작업을 처리하는 기능과 클라이언트의 접속을 기다리는 부분이 분리 되어 있어서 서버의 응답 속도를 높여 줄 수 있다.
2. 동시에 여러 작업을 병렬로 처리할 수 있어, 두 개 이상의 요청을 받아 동시에 처리할 수 있다.
3. 작업별로 쓰레드를 생성해 처리하는 방법으로 웬만한 부하까지는 견딜 수 있으며, 순차적인 실행 방법에 비하면 속도가 크게 향상된다.

그러나 사실 여기엔 문제점이 있다. 예를 들면 요청이 (엄청) 많이 들어오는 경우 쓰레드가 대량으로 생성될 수 있다는 것이다. 쓰레드가 대량으로 생성되면 문제점이 발생한다.

#### 쓰레드 라이프 사이클 문제
쓰레드를 생성하고 제거하는 작업에도 자원이 소모된다.
#### 자원 낭비
실행 중인 쓰레드는 시스템의 자원, 특히 메모리를 소모한다. 프로세서 보다 많은 쓰레드가 만들어진다면 대부분의 쓰레드는 idle 상태에 빠지게 된다.
idle 상태에 빠진 쓰레드는 메모리를 갖고(점유) 하고 있으며 이와 같은 많은 쓰레드가 CPU 를 선점하기 위해 경쟁까지 하면서 메모리 외에 다른 자원들도 소모하게 된다.
#### 안정성 문제
OOM 이 날 수 있다.


이러한 이유 때문에 일정한 수준 까지는 쓰레드를 위해 성능의 이점을 얻을 수 있지만, 어느 수준을 넘어서면 성능이 떨어질 수 있다. 

### Executor
Executor 는 작업 등록(task submissions) 과 작업 실행(task execution) 을 분리하는 표준적인 방법이며, 각 작업은 `Runnable` 형태로 정의한다.
Executor 는 producer-consumer 패턴에 기반하고 있어,  producer-consumer 을 구현하기가 편하다.

```java
public class TaskExecutionWebServer {
    private static final int THREADS_COUNT = 10;
    private static final Executor exec = Executors.newFixedThreadPool(THREADS_COUNT);

    public static void main(String[] args) throws IOException {
        ServerSocket socket = new ServerSocket(80);
        while (true) {
            final Socket connection = socket.accept();
            Runnable task = () -> System.out.println("Do something with connection!");
            exec.execute(task);
        }
    }
}
```

### 쓰레드 풀
작업 쓰레드는 먼저 작업 큐에서 실행할 다음 작업을 사져오고, 작업을 실행하고, 가져와 실행할 다음 작업이 나타날 때 까지 대기하는 일을 반복한다.
쓰레드 풀을 사용하는데 장점은 다음과 같다.

1. 매번 쓰레드를 생성할 필요가 없다(쓰레드를 재사용하기 때문에). 따라서 시스템 자원이 줄어들지 않는다.
2. 이미 쓰레드가 만들어져 있어서, 쓰레드 생성에 따른 딜레이가 없다. 이 때문에 전체적인 반응 속도가 향상된다.
3. 적절한 쓰레드 풀의 사이즈는, 프로세서가 쉬지 않고 동작하게 할 수 있게 한다.

쓰레드 풀의 사용 방법 중에는 `Executors` 클래스에 만들어져 있는 다음과 같은 메서드를 주로 사용하곤 한다.

- `newFixedThreadPool` 생성할 수 있는 쓰레드 수가 정해져있고, 제한된 개수까지만 생성된다.
- `newCachedThreadPool` 쉬는 쓰레드가 있을 경우 쉬는 쓰레드를 종료 시킨다.
- `newSingleThreadExecutor` 단일 쓰레드만 생성되고, 해당 쓰레드 작업이 exception 이 나 비 정상 종료 되면 새로운 스레드가 생성된다. (single 만 갖고 있는다.)
- `newScheduledThreadPool` 주기적으로 작업을 실행할 수 있는 쓰레드 풀이다.


### 쓰레드 중단 및 종료
자바에는 쓰레드가 작업을 실행하고 있을 때 강제로 멈추도록 하는 방법이 없다. 대신 __인터럽트__ 라는 방법을 사용한다. 인터럽트는 특정 쓰레드에게 작업을 멈춰 달라고 요청하는 행위이다.

가장 기본적인 형태로는 __취소 요청__ 에 대한 플래그를 설정하는 것이다. 예를 들면 큐를 두고 프로듀서-컨슈머 패턴으로 개발한다고 하였을 때 __Graceful__ 하게 종료하려면 큐에 사전에 정의된 특정 객체를 넣어, 해당 객체를 큐에서 `take` 할 때 발견하게 된다면 종료 시키는 방법이 있다. (이 때 해당 객체를 `poison` 이라고 명명하기도 한다.)


> 
특정 쓰레드의 interrupt 메서드를 호출한다 해도 해당 스레드가 처리하던 작업을 멈추지는 않는다. 단지 해당 쓰레드에게 인터럽트 요청이 있었다는 메시지를 전달할 뿐이다. 
#### InterruptedException 처리 방법

- 발생한 예외를 호출 스택의 상위 메소드로 전달한다. 이 방법을 사용하는 메서드 역시 인터럽트를 걸 수 있는 블로킹 메서드가 된다.
- 호출 스택의 상ㅇ단에 위치한 메소드가 직접 처리할 수 있도록 인터럽트 상태를 유지한다.

만약 Runnable 을 직접 구현했다면 상위에 전달하는 방식을 사용할 수 없다. 이 경우 가장 일반적인 방법은 `interrupt` 를 한 번 더 호출 하는 것이다.

작업 중단 기능을 지원하지 않으면서 인터럽트를 걸 수 있는 블로킹 메서드를 호출하는 작업은 인터럽트가 걸렸을 떄 블로킹 메서드의 기능을 자동으로 재시도하도록 반복문 내부에서 블로킹 메서드를 호출하도록 구성하는 것이 좋다.


