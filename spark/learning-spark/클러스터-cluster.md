클러스터 cluster
==

## 드라이버 Driver
스파크에서 `driver` 란 `SparkContext` 를 생성하고 `main()` 메서드가 실행되는 프로세스를 말한다.

## 익스큐터 Executor
`executor` 는 개별 task 를 실행하는 프로세스다. (쓰레드가 아니다!)
크게 두 가지 역할을 한다.

- 어플리케이션을 구성하는 작업을 실행하고 `driver` 에게 그 결과를 되돌려준다.
- 각 `executor` 안에 존재하는 `Block Manager` 라는 서비스를 통해 사용자 프로그램에서 캐시하는 __RDD__ 를 저장하기 위한 메모리 저장소를 제공한다.

## 클러스터 매니저 Cluster Manager
`cluster manager` 는 스파크와 붙이거나 뗄 수 있는 컴포넌트다.

### standalone
하나의 마스터와 여러 개의 워커로 구성되며, 각각은 설정에 따른 용량의 메모리와 CPU 코어 개수만큼을 사용한다. 사용자는 어플리케이션을 제출할 때 `executor` 가 메모리를 얼마나 쓸 지뿐만 아니라 모든 `executor` 가 사용할 총 코어 개수도 지정할 수 있다.

스파크 실행시에는 아래 처럼 명령어를 입력한다.
```
spark-submit --master spark://masternode:7077 myapp
```

### yarn

### apache mesos

## 스파크 실행 단계

1. 사용자는 `spark-submit` 을 사용하여 어플리케이션을 제출한다.
2. `spark-submit` 은 `driver` 를 실행하고 사용자가 정의한 `main()` 메서드를 호출한다.
3. `driver` 는 `cluster manager` (스탠드얼론, yarn, mesos 등등) 에게 `executor` 실행을 위한 리소스를 요청한다.
4. `cluster manager` 는 `executor` 를 실행한다.
5. `driver` 가 어플리케이션을 실 행한다. 여기서 작성된 __RDD__ 의 트랜스포메이션과 액션에 기반하여 `driver` 는 작업 내역을 단위 작업(보통 `task` 라 부름) 형태로 나누어 `executor` 에게 보낸다.
6. `executor` 는 이러한 `task` 들을 실행한다.
7. `driver` 의 `main()` 이 끝나거나 `SparkContext.stop()` 이 호출된다면 `executor` 들은 중지되고 `cluster manager` 에게 사용했던 리소스를 돌려준다.
