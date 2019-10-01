# Redis sentinel

#### 원인
Redis에서 같은 키를 사용하는 서로 다른 App 이 존재하였다. 이를 위해 Redis 가 제공하는 database select 를 찾아보았었다.

그러나 불행히도 현재 구성되어있는 Redis cluster 에선 __당연하게도__ database select 를 제공하지 않았다.

결국 database select 를 사용하기 위해선 standalone 으로 띄워서 마스터-슬레이브 구조로 가야 하는데 그럴 경우 failover 가 걱정 되었다.

그렇게 해서 찾아본 것이 sentinel 이였다

> 결론부터 말하자면 이 방법은 꼼수 of 꼼수였고 [Redis SELECT](https://redis.io/commands/select) 도큐먼테이션에도 추천하지 않는 방법이다. (즉, 서로 다른 API 에서 database select 를 사용하는 방식은 권하지 않았다.)

#### Sentinel 이란 무엇인가
Sentinel 은 단순하게 말하면 (사실 정말 간단히 리서치 해서 자세히는 모름) master-slave 구조에서 master 가 다운 되었을 시 다른 slave 들 중에 master 를 선출하는 또 다른 앱? 레디스? 인 것이다.

그러니까 결국 서버에서 redis 를 띄우면, redis-sentinel 이란 것도 띄워야 한다는 것이다. (물리적으로 같은 서버에서 띄워도 되고, 다른 서버에서 띄워도 된다.)

다시 말하자면 auto failover 를 제공해주는 것이라고 보면 될 것 같다.

#### Sentinel 은 어떻게 사용할까?
우선 레디스를 master-slave 구조로 띄워야 한다. 마스터는 그냥 띄우면 되고 슬레이브를 띄울때 config 를 조정한다. 중간에 `replicaof master-ip` 를 추가 하면 된다. 다른 예제에서는 `slaveof` 로 추가하라고 하는데 아마 버젼 차이 때문인 것 같다. 참고로 내가 테스트했을 시 레디스 버젼은  `5.0.3` 이었다.

이렇게 각 슬레이브에 마스터 ip 를 적어주면 우선은 master-slave 구성은 완성 되었다. _참 쉽죠?_

다음은 Sentinel 을 띄운다. 나는 물리적으로 같은 서버에 띄웠다.
sentinel 은 `sentinel.conf` 라는 파일이 있고, 여기에 같은 방법으로 마스터 ip 를 적어주면 된다. (정확히 어떤 문구였는지는 까먹음)

레디스를 띄울때 `redis-server` 였다면 sentinel을 띄울땐 `redis-sentinel` 로 띄우면 된다. 그리고 뒤에 각각의 config 를 바로 넣어주면 알아서 그 config로 띄워진다.

1대의 master, 2대의 slave 그리고 3개의 sentinel 로 테스트해보았고 마스터를 다운 시켰을 때 다른 slave가 master로 선출되고 해당 config 들에서도 기존 마스터가 아닌 새로운 마스터의 ip 로 오버라이트 되는 것을 볼 수 있었다.

만약 제대로 안된다고 하면 bind 로 해당 서버에 접근 가능한지 우선 확인해보고 그 다음에 ip 를 제대로 적었는지 보면 될 것 같다.

### 언제 cluster를 사용하고 언제 sentinel 을 사용할까?
우선 cluster 는 하나가 맛이 가면, 다른 것들도 맛이 가는데 sentinel 로 구성하면 하나가 맛이 가도 다른 애들은 정상적으로 서비스가 가능하다.

이렇게 보면 cluster 가 더 안좋은 것 같지만 대신 cluster 는 샤딩을 제공한다. sentinel은 단독 모드에다가 master-slave라 단순히 복제라서 데이터 사이즈가 엄~청 커지면 스케일업을 해야하는데 cluster 는 스케일아웃을 하면 된다. (맞나??)

그리고 sentinel 로 구성해도 마스터 선출만 자동이지, 노드를 다시 띄우는건 수동이다. (클러스터는 내가 구성안해서 잘 모르겠음)

### 결론
cluster 를 사용하자. 어차피 sentinel 은 꼼수였다.
서로 다른 앱이라면 redis 도 분리해야하는게 맞다.
