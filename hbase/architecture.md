# HBase 아키텍쳐

HBase 는 3 가지 타입으로 구성되어 있다.

- Region Server
- HMaseter
- Zookeeper

## 데이터 읽고 쓰기
HBase 에는 Catalog table (aka Meta table) 이라는 특별한 테이블이 있다.
이 테이블의 위치 정보는 주키퍼에 있으며, 이 테이블은 클러스터의 모든 region 들에 대한 정보를 갖고 있다.

데이터를 처음 읽고 쓸 때 HBase 는 다음과 같이 동작한다.

1. 클라이언트는 주키퍼로 부터 Region Server 정보를 얻는다. (Region Server 정보는 메타 Meta 테이블에 있다.)
2. 클라이언트는 Region Server 에게 쿼리 (질의) 해서 데이터의 rowKey 를 갖고 있는 Region Server 를 알게 된다. 그러면 클라이언트는 Meta table 의 위치도 캐싱하고 Region Server 정보도 캐싱한다.
3. 클라이언트는 Region Server 로 부터 데이터를 가져온다.

이후 요청은 (처음 읽기/쓰기가 끝나고) 클라이언트 캐시가 만료되지 않는 이상 (혹은 실패) 캐시한 메타 테이블 정보와 Region Server 의 정보로 데이터를 읽고 쓴다.

만약 Region migration 으로 캐시가 실패하게 된다면 다시 1~3 을 반복하게 된다.


## Region Server

요청에 대한 데이터 쓰기/읽기등을 담당하고 있다. 클라이언트는 Region Server 와 다이렉트 통신한다. (근데 이것도 3.x.x 버젼 부터는 HMaster 랑 통신하게 바뀌었다고 한다.)

Region Server는 HDFS DataNode 에서 돌아가고 다음과 같은 컴포넌트들을 갖는다.

- WAL
- BlockCache
- MemStore
- HFile

데이터 쓰기는 아래 처럼 동작한다.

1. 새로운 데이터는 WAL 파일에 추가 작성된다.
2. WAL 은 데이터를 recover 할 때 사용된다. (장애 복구시에는 유지되지 않는다. ??)


## HBase Master (aka HMaster)

Region 할당을 담당하고 있다. 또한 테이블의 Create, Delete 같은 연산, 그리고 기타 다른 운영들을 담당한다.

## Zookeeper

클러스터 상태 유지에 대한 책임을 담당한다.


당연하겠지만 Hadoop HDFS 를 기반으로 되어 있다.

### DataNode

리젼서버에 의해서 다루어지는 데이터 저장에 대한 책임을 갖고 있다. 모든 HBase 데이터는 HDFS 파일로 저장된다.
Region Server 와 HDFS DataNode 는 종종 같이 구성하는데 locality 문제 때문이다. 

데이터를 쓸 때는 local 에다가 저장 되지만 region 이 migration 되면 (compaction 되기 전 까지는) 로컬에 없을 수도 있다.

## NameNode
모든 HDFS 의 물리적인 데이터 블록에 대한 메타 정보를 담당하고 있다.


## 참고
- https://developpaper.com/deep-understanding-of-hbase-architecture-translation/
