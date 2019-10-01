# Zookeeper
주키퍼는 기본적으로 분산 시스템을 위한 코디네이터다.

## 주키퍼
주키퍼는 데이터를 디렉토리 구조로 관리하며 key-value 스토리지 처럼 key 로 접근 가능하다.

> 물론 value 를 사용하지 않아도 된다. 그럴 땐 임의로 "" 와 같이 빈 문자열을 넣어준다.

이러한 디렉토리 구조는 (당연하겠지만 tree 형태) 계층을 가지고 있고 각 데이터 노드를 `znode` 라고 부른다. 이러한 `znode`를 다루는 주키퍼 작업을 `recipe` 라고 부른다.

### 주키퍼의 쿼럼
주키퍼의 쿼럼은 적어도 하나의 앙상블 (서버군) 과 겹쳐야 하며 과반수가 넘어야 한다. 만약 서버가 5대라면 쿼럼은 3개여야 한다.

만약 5대의 앙상블 서버 들 중 2대만 쿼럼으로 지정한다면 아래와 같은 시나리오가 발생한다.

1. 서버 s1, s2 는 `/z` 라는 znode 를 생성하는 요청이 복제되었다고 리더에게 응답한다.
2. 주키퍼 서비스는 `/z` 라는 znode 가 생성되었다고 클라이언트에게 알린다.
3. 서버 s1, s2가 다른 서버로 새로운 `/z` 를 복제하기 전에 어떤 이유로 긴 시간 동안 다른 서버와 클라이언트와 통신이 불가능해졌다.
4. 남은 3대는 `/z` 를 모름에도 서비스는 계속된다. (쿼럼은 2대이기 때문에 나머지 3대로도 서비스가 가능하기 때문이다.)

## 주키퍼의 장애 복구
워커(슬레이브) 와 마스터의 연결이 끊어졌을 때, 작업을 재할당하게 된다면 같은 작업(Task)가 동시에 두 개의 Worker(워커) 에게 할당 될 수 있다. 이런 경우를 방지 하기 위해 주키퍼는 아래와 같은 프로세스를 따른다.

1. 클라이언트(워커)가 데이터를 임시(Ephemeral) 상태라고 정의할 수 있게 한다.
2. 클라이언트는 주키퍼 앙상블에게 자신의 상태를 주기적으로 알려준다.
3. 만약 상태 알람에 실패한다면 주키퍼 앙상블은 해당 클라이언트에 속한 모든 임시(Ephemeral) 상태를 제거한다.


## 주키퍼 API
주키퍼의 `znode` 의 데이터는 바이트 배열로 저장된다.

주키퍼에는 다음과 같은 API 가 있다.

data 를 담고 있는 /path 라는 이름의 znode 생성
```
create /path data
```

/path 라는 znode 를 제거
```
delete /path
```

> 자바에서 API 를 호출해 삭제를 원한다면, 해당 znode 에 대한 버젼 정보가 반드시 포함되어야 한다.

/path 라는 znode 가 존재하는지 확인 (Stat)
```
exists /path
```

/path 라는 znode 에 data를 저장
```
setData /path data
```

/path 라는 znode의 데이터를 가져옴
```
getData /path
```

/path 라는 znode의 모든 자식 목록을 반환
```
getChildren /path
```

### 주키퍼의 노드 종류

주키퍼에는 총 4가지 종류의 노드가 있다. (3.5 기준으로 TTL, container 가 생겼다.)

- Persistent

영구적인 노드로 `delete` 를 이용해 제거가 가능하다.
- Ephemeral

임시 노드로 주키퍼 클라이언트와 세션이 끊기기 전까지 살아있다가, 끊기면 사라지는 노드이다. 혹은 세션을 맺은 클라이언트가 삭제해야 제거되는 노드이다.
- TTL

만약 znode가 주어진 TTL 시간 동안 변화가 없고 어떠한 자식 노드도 생성되지 않는다면 삭제가 된다. 다만 기본적으로 TTL 은 비활성화 되어 있으므로 이를 활성화 시키는 사전 작업이 필요하다.
- Container

컨테이너 노드는 자식 노드가 없을떄 삭제가 된다. 바로 삭제는 되지 않고 TTL 과 마찬가지로 삭제 노드의 후보가 되는 것이다. 즉 자식 노드가 있다가, 모두 삭제가 된다면 일정 시간 이후에 Container 노드로 설정한 노드는 삭제가 된다.

각 노드에는 순차적 (sequential) 노드도 존재한다. 단순히 노드가 생성되면 뒤에 sequential number 가 붙는다고 생각하면 된다. 즉 tasks 라는 zNode가 있을 때 해당 옵션(-s)으로 생성하면 tasks-(숫자) 가 붙는 것이다.

## 주키퍼 세션 생성
주키퍼 객체를 생성하는 생성자는 다음과 같다.

```
Zookeeper(String connectString, int sessionTimeout, Watcher watcher)
```

## 주키퍼의 connectString
주키퍼 서버들의 (앙상블) 호스트 네임과 포트다. 3.2(? 3.3?) 부터는 해당 목록들 뒤에 아래와 같이 root node path 를 지정할 수도 있다.

```
localhost:1234,localhost:1235,localhost:1236/root
```

## 주키퍼의 와쳐 (Watcher)
데이터를 주기적으로 가져오기 위해 주키퍼에게 폴링하는 것은 주키퍼의 설계와는 맞지 않다. 이럴 때 사용하는 것이 와쳐 (Watcher) 다. 주키퍼는 어떤 znode 의 변화가 생길 것을 감지 하기 위해 와쳐를 등록할 수 있다.

주키퍼에서 znode 의 와쳐를 걸 수 있는 함수는 `exists`, `getData`, `getChildren` 이 있다.

각각 걸 수 있는 콜백 종류는 API 를 확인해보면 된다. `exists` 의 경우 `StatCallback` 으로 와쳐를 설정하면 노드의 생성, 삭제 이벤트가 와쳐로 들어오게 된다.

와쳐는 1회성이므로 한번 알림이 가면 해당 노드에 대해 다시 등록/설정을 해주어야 된다. 만약 1000개의 클라이언트가 동일한 znode 에 `exist` 함수로 와쳐를 등록했다면 해당 노드가 존재하지 않았다가, 생성된다면 1000개의 클라이언트에게 알람이 가게 된다. 이런 설계는 옳지 않다. 이상적으로는 한 개의 클라이언트만 하나의 znode 에 와쳐를 등록해야 한다.

와쳐로 등록할 `CallBack` 메소드들은 해당 스레드에서 처리하게 된다.

와쳐는 세션이 종료되거나, 만료되는 등 실제 트리거 되기 전까진 삭제가 불가능 했었다. 그러나 주키퍼 3.5 부터는 [removeWatches](https://zookeeper.apache.org/doc/r3.5.0-alpha/api/org/apache/zookeeper/ZooKeeper.html#removeWatches) 를 제공하고 있다.

## 주키퍼의 sessionTimeout
주키퍼는 처음 객체를 생성할 때 `seesionTimeout` 을 지정할 수 있다. 보통은 tickTime 이라 부르는 `zoo.cfg` 에 정의된 값에 2 를 곱한다(고 한다). (이렇게 충분한 만료 시간을 주는 이유로는) 주키퍼 앙상블이 죽으면 시간도 멈춘 상태가 되고, 다시 앙상블이 살아났을 때 세션 타임아웃이 재 시작 된다.

주키퍼를 사용하는 프로세스들이 그 때까지 남아있다면 복구된 앙상블 장애 때문에 긴 타임아웃이 발생했다는 것을 알게 되고 추가적인 지연 없이 바로 마지막 지점부터 다시 시작할 수 있다.

> HBase 는 이 값을 40초를 준다. [참고](https://hbase.apache.org/book.html#important_configurations) The current default maxSessionTimeout that ZK ships with is 40 seconds, which is lower than HBase’s.

주키퍼와의 세션이 지정한 `sessionTimeout` 을 넘어서면 session expired 가 된다. 그러므로 적절한 `sessionTimeout` 을 하는 것은 중요하다.

> 테스트를 할 때 네트워크 이상과 같은 상태일 때만 해당 sessionTimeout 에 영향을 받는 줄 알았지만, 실제로 그냥 주키퍼와 통신을 전혀 하지 않은 상태에서도 expired 가 되는 것을 발견했다. 내부적으로 ping-pong 을 계속 떄릴 줄 알았는데, 그냥 어플리케이션 단에서 통신을 안해도 expired 가 되었던 것이다. 아무래도 더 공부해서 찾아봐야할 것 같다.

## 멀티옵
주키퍼에는 원자적인 블럭에서 여러 개의 명령을 실행할 수 있게 하는 옵션도(?) 제공한다.
이를 멀티옵 multiOp 이라 부른다. (트랜잭션 개념과 비슷)

멀티옵 블럭에서는 모두 성공하거나 모두 실패하거나만 가능하다. 사용하려면 아래와 같이 한다.
(`Transaction` 이라는 래퍼 버젼도 제공한다, 물론 비동기 버젼도 있다.)

```
Op deleteZnode(String z) {
    return Op.delete(z, -1)
}

...

List<OpResult> results = zk.multi(Arrays.asList(deleteZnode("/a/b"),
    deleteZnode("/a")));
```

## 주키퍼 API 의 유용한 함수들

#### deleteRecursive

zkUtil.java 에 있는 함수다. 어떤 `znode`와 그 자식 노드 `znode` 들이 있을 때 모든 자식 노드들을 가져오지 않고도, 부모 노드만 사용하여 부모 노드와 하위 자식 노드들을 삭제할 수 있는 함수다.

나의 경우엔 zookeeper 를 사용하는 클래스의 테스트 코드를 생성할 때 유용했다.(쓰고 보니 멀티옵을 사용해도 되는 부분이였다.)

