# 주키퍼를 다루면서 생겼던 트러블 슈팅

## TTL Node 생성 불가능
주키퍼 클라이언트에서 (CLI 든 자바 API 를 호출하든) `KeeperErrorCode = Unimplemented for /node` 가 떴다.
TTL 노드는 주키퍼 3.5 이후이 생긴 기능으로, 환경만 구성하면 사용할 수 있다.


환경은 아래와 같이 구성해야 한다.

- 우선 서버 아이디 (`zoo.cfg` 에 작성하는 아이디) 는 255 이하여야만 한다. (만약 주키퍼 3.5.3 을 사용하고 있다면 127 이하여야만 한다.)
- `zookeeper.extendTypesEnabled` 를 `true` 로 맞춰준다. (기본은 false 이기 때문이다.)

이러면 주키퍼에서 TTL 노드를 사용할 수 있다.

나의 원인은 해당 옵션 `zookeeper.extendTypesEnabled` 을 `zoo.cfg` 에 적용한게 문제였다.
이 옵션은 주키퍼를 실행할 때 주어야 하는 프로퍼티이다. 따라서 `zkServer.sh` 에 해당 옵션을 주자.

#### 참고
- [https://zookeeper.apache.org/doc/r3.5.5/zookeeperAdmin.html#sc_advancedConfiguration](https://zookeeper.apache.org/doc/r3.5.5/zookeeperAdmin.html#sc_advancedConfiguration)
