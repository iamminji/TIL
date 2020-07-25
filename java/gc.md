# GC

자바 어플리케이션에서 GC 가 일어나는 주된 원인은 다음 두 가지 이다.

1. 할당률
2. 객체 수명

## Weak Generational Hypothesis
__단명 객체를 빠르게 쉽고 빠르게 수집하자__ 가 목표다.

- 객체마다 세대 카운트(객체가 지금까지 무사 통과한 가비지 수집 횟수)를 센다.
- 큰 객체를 제외한 나머지 객체는 에덴 공간에 생성한다. 여기서 살아 남는 객체는 다른 곳으로 옮긴다.
- 장수했다고 할 정도로 충분히 오래 살아남은 객체들은 별도의 메모리 영역 (올드) 에 보관한다.

## Stop the world
GC 사이클이 발생하여 가비지를 수집하는 동안에는 모든 애플리케이션 스레드가 중단된다.
> GC 스레드는 애플리케이션 스레드와 동시에 실행 될 수 있는데, 사실은 100% 동시는 아니고 거의 준 동시다.

## 핫스팟

## Parallel Collector
자바 8 이전의 GC 는 Parallel Collector 이다.
Parallel Collector 는 young GC, Full GC 모두 STW 를 일으킨다.

종류도 여러가지다

- Parallel GC
- ParNew GC
- ParallelOld GC

종류는 달라도 여러 스레드를 이용해 가극ㅂ적 빠른 시간 내에 살아 있는 객체를 식별하고 기록 작업을 최소화 하도록
설계된 점은 비슷하다.

### Parallel GC / Young Generation
스레드가 에덴에 객체를 할당하려는데 자신이 할당받은 TLAB 가 부족하고 JVM 은 새 TLAB 를 할당할 수 없을 때 Young Genration Collection 이 발생한다.
>TLAB란 스레드 로컬 할당 버퍼(Thread Loacal Allocation Buffer)로 다른 스레드가 자신의 버퍼에 객체를 할당하지 못하게 하는 스레드 자신의 공간이다.

전체 애플리케이션 스레드가 중단되면 (STW) 핫스팟은 Young Generation 을 뒤져서 가비지가 아닌 객체를 골라내고,
살아남은 객체는 Survivor 공간으로 __방출__(이동)시킨다. (이 때 count 를 증가시킨다.)

### ParalleOld GC / Old Generation
Parallel GC와 비슷하지만, 다르다.
Parallel GC 가 Young 에서 Survivor 로 객체를 방출(이동) 시킨 다면 ParallelOld GC 는 하나의 연속된 메모리 공간에서 객체를 __압착__ 시킨다.

Old 세대에서 더 이상 이동시킬 공간이 없으면 Parallel Collector 는 Old 세대 내부에서 객체들을 재 배치해, (오래된 객체가 없어져 생긴) 공간을 회수한다.
따라서 메모리 사용 면에서 아주 효율적이고, 메모리 단편화 (Fragmentation) 가 일어날 일도 없다.

### 단점
Parallel GC의 경우엔 상관 없지만,
ParallelOld GC 의 경우엔 Old 영역의 객체들은 상당수 살아 남을 것이고, 이 때문에 STW의 시간이 힙 크기에 비례하여서 오래 걸리게 될 것이다.

## 참고
- https://johngrib.github.io/wiki/java-g1gc/
- 자바 최적화 (도서 한빛미디어)
