# Zookeeper 레시피

많은 애플리케이션들이 주키퍼를 이용해 지속성과 가용성 (CAP 이론 중 CA) 을 보장받고 있다.

#### 아파치 HBase

HBase 에서는 클러스터 마스터 (리더) 선출과 가용 서버들의 목록을 저장하고, 클러스터
의 메타데이터를 관리하는 데 주키퍼를 사용한다.

#### 아파치 Kafka
카프카는 장애 감지, 토픽 디스커버리, 토픽에 대한 생성과 소비 상태를 관리하는 데 주키퍼를 이용한다.

#### 아파치 Solr
솔라는 검색 엔진 플랫폼으로, 클러스터의 메타데이터를 저장하고 메타데이터에 대한 업데이트들을 코디네이션 하기 위해 주키퍼를 사용한다.

이처럼 메타데이터를 다루는데 주로 주키퍼를 사용하고는 하지만, 리더 선출 시에도 활용할 수 있다.

## Leader Election
주키퍼를 이용해 고 가용성 어플리케이션을 설계할 일이 생겨서, 리더 선출 알고리즘들을 살펴 보았다.

### 리더 선출 준비
우선 임의의 path 를 지정한다. 여기서는 `/election` 이라고 하겠다.

`/election` 하위에, 가용하는 서버 목록의 정보들을 __Ephemeral-sequential__ 모드로 생성하는 것이 핵심이다.

만약 서버가 1번부터 10번까지 총 10개 있다고 해보자. (혹은 10개의 프로세스라고 볼 수도 있다.)

1번 서버(혹은 프로세스) 가 `/election` 의 자식으로 아래와 같이 __Ephemeral-sequential__ 노드를 생성한다.

 ```
 create -es /election server-n
 ```

 결과는 아마 `server-n-00000` 과 같은 형태가 될 것이다. (sequential number 는 다를 수 있다.)

 그리고 그 다음 2번 프로세스가 같은 방식으로 노드를 생성한다고 해 보자. 다른 설정을 하지 않았다면 결과 노드는 `server-n-00001` 처럼 최초에 만들어졌던 노드보다 더 큰 숫자가 할당 될 것이다.

 > 물론 수는 1씩 증가하지 않을 수 있다. 중요한 건 sequential 하다는 것이다.

 여기까지만 보면 어떻게 이렇게 리더 선출을 하나 싶기도 할지 모른다.

 그렇다면 이렇게 생성된 자식 노드들 중에 어떤것이 리더가 되어야 할까?
 보통 가장 작은 수를 갖는 노드가 리더로 한다.

 즉 아래와 같은 노드 들이 있을때

```
/electopn/server-n-00000
/election/server-n-00001
/election/server-n-00009
/election/server-n-00012
...
```

리더는 `server-n-00000` 이 된다. 이를 __primary node__ 라 부른다.

#### 리더 장애 복구

이제는 리더가 죽거나, 장애가 발생 시 나머지 노드들 중 어떤것이 리더가 되어야 하고, 어떻게 리더가 장애가 났는지를 알 수 있냐는 문제가 남았다.

리더 `server-n-00000` 의 대용품 __backup node__ 는 어떤것이 되어야할까?

가장 쉬운 방법은 이미 노드들을 sequential 하게 만들었기 때문에 `server-n-00000` 의 다음 노드 (다음 숫자) `server-n-00001` 이 __backup node__ 가 되면 될 것이다.

리더가 죽을 시, 선출될 다음 노드도 이제 선택했다면 이제 리더가 죽었음을 해당 __backup node__ 가 감지하는 것이다.

이 때, 사용하는 것이 바로 `Watcher` 이다. 주키퍼는 `Watcher` 라는, 특정 노드에 생기는 이벤트를 감지하고, 액션을 취할 수 있는 인터페이스가 존재한다.

`server-n-00001` 은 `server-n-00000` 에게 _exists watcher_ 를 등록하면 된다. 그러면 노드 `server-n-00000` 가 삭제 된다면 (Ephemeral 이기 때문에 프로세스가 중지 되면 노드도 날라간다.) `server-n-00001` 을 등록한 프로세스는 워쳐로 알람을 받을 수 있다.

삭제 알람을 받게되면 `server-n-00001` 를 등록한 프로세스는 `server-n-00001` 를 마스터로 승격 시키면 된다. (마스터가 해야할 역할을 위임하면 된다.)

그러면 `server-n-00001` 가 죽게 된다면?

역시 마찬가지로 `server-n-00001` 의 다음 노드 `server-n-00009` 가 `server-n-00001` 의 __backup node__ 로써 _exists watcher_ 를 등록하면 된다.

워쳐를 등록하는 로직은 최초에 `/election` 하위에 노드를 생성할 때 미리 진행해주면 될 것이다.

이 방법을 사용하면 하나의 노드에 대해선 하나의 워쳐만 등록하게 된다.
(실제로 주키퍼에선 하나의 노드엔 하나의 워쳐를 지향하고 있다.)

### 정리
- 서버 (혹은 프로세스) 정보를 __Ephemeral sequential__ 모드로 zNode 를 생성한다.
- 생성된 zNode 들 중에 리더는 항상 가장 작은 sequential 을 갖는 zNode 이다.
- 각각의 sequential zNode 들은 바로 그 다음 수의 zNode 를 __backup node__ 로 갖고 있는다.
- 각각의 __backup node__ 들은 바로 이전의 sequential zNode 에 `Watcher` 를 등록한다.

### 참고
- [https://zookeeper.apache.org/doc/current/recipes.html#sc_leaderElection](https://zookeeper.apache.org/doc/current/recipes.html#sc_leaderElection)
- [HBase 의 리더선출](https://github.com/apache/hbase/blob/master/hbase-zookeeper/src/main/java/org/apache/hadoop/hbase/zookeeper/ZKLeaderManager.java)
