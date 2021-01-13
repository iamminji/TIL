# Architecture

## NoSQL?
HBase 는 NoSQL 의 한 종류입니다. HBase 가 갖는 특징은 다음과 같습니다.

- 일관된 읽기/쓰기에 강합니다. HBase is not an "eventually consitent" DataStore. 고속 카운터 같은 task 에 잘 맞습니다.
- sharding 이 자동으로 이루어진다. HBase table 은 region 을 통해 클러스터에 분산되며 각 region 은 데이터 증가에 따라 자동으로 split 되고 재 분산됩니다.
- Automatic RegionServer failover
- Hadoop/HDFS Integration. HBase 는 HDFS 를 분산 파일 시스템으로 지원합니다.
- MapReduce 를 사용할 수 있습니다. HBase supports massively paralleized processing via MapReduce for using HBase as both source and sink.
- Java Client API 를 지원합니다.
- Thrift/REST API 를 사용할 수 있습니다.
- Block Cache and Bloom Filter 를 지원합니다. (쿼리 최적화를 위해)
- Operational Management.

### When Should I Use HBase?
HBase 는 모든 문제에 맞지는 않습니다.


우선적으로 데이터 사이즈에 영향을 받습니다. 만약 데이터 row 수가 수억, 수십억 정도 된다면 HBase 는 좋은 대안이 될 것입니다. 그러나 고작 수천/수백만개라면 전통적인 RDBMS 가 더 좋은 선택이 될 것입니다.


두번째로 RDBMS 가 제공하는 추가적인 기능(예를 들면 typed column, secondary index, transaction, advanced query languages...)이 없어도 되는지 확인해야 합니다. RDBMS 에서 HBase 로 이전하려면 설계를 다시 해야 합니다.


세번째로, 하드웨어가 충분한지 살펴보십시오. HDFS 가 5개 미만의 DataNode와 NamdeNode 에서는 제대로 작동하지 않습니다. (HBase 가 stand-alone 모드에서 잘 동작하지만 그건 오직 개발을 위한 것이기 때문입니다.)

### What is The Difference Between HBase and Hadoop/HDFS?
HDFS 는 대용량 파일의 저장에 잘 맞는 분산 파일 시스템입니다. HBase 는 HDFS 를 기반으로 구축되며, 대형 테이블에서 빠른 조회(및 업데이트) 를 제공합니다. 

HBase 는 빠른 조회를 위해 HDFS 에 존재하는 색인화된 `StoreFile` 에 데이터를 저장합니다.

## Catalog Tables
HBase 에는 catalog table 인 `hbase:meta` 라는 테이블이 존재한다.

### hbase:meta
`hbase:meta` (예전에는 `.META.` 라고 불리웠던) 테이블은 모든 region 정보를 가지고 있다. `hbase:meta`  의 위치는 Zookeeper 에 저장되어 있다.

### Startup Sequencing
먼저 주키퍼에서 `hbase:meta` 정보를 찾는다. 그리고 `hbase:meta` 의 server 와 startcode 정보를 업데이트 한다. (region 과 regionServer 의 할당 방식은 [여기](https://hbase.apache.org/book.html#regions.arch.assignment) 를 참고하도록 하자.)


## Client

### Master Registry (new as of 2.3.0)
Client 는 내부적으로 커넥션에 필요로 한 메타 데이터 정보를 connection registry 를 사용한다.
connection registry 구현에는 아래의 메타데이터 정보를 가져오도록 되어 있다.

- Acitve master address
- Current meta region(s) location
- Cluster ID (Unique to thie cluster)

이 정보들은 set up, scan, get 과 같은 클라이언트 측의 일을 수행하기 위해 필요로 한다.

전통적으로 connection registry 는 Zookeeper quroum 에 다이렉트로 접근하여 (위의) 메타데이터 정보를 가져오지만 HBase 2.3.0 부터는 Master 와 통신하여 이 정보를 가져올 수 있다. 클라이언트는 이제 Zookeeper 와의 통신 대신 master RPC  를 필요로 한다.

이렇게 한 이유는 다음과 같다.

- Zookeeper 부하를 줄이기 위해서
- Holistic client timeout and retry configurations since the new registry brings all the client operations under HBase rpc framework. 
- HBase client library 에서 Zookeeper 의존성을 줄이기 위해서

이 의미는

- 클러스터 연결에 최소한의 single active 또는 stand by master 가 필요하다. (자세한 내용은 [Runtime impact](https://hbase.apache.org/book.html#master.runtime) 를 참고)
- master 가 client metadata 캐시가 비어있거나 오래된 경우 critical 한 경로에서 읽기/쓰기 수행을 하게된다. 
- 클라이언트가 Zookeeper 앙상블 대신에 HMaster 과 직접 통신하기 때문에 master 에 커넥션 부하가 더 생긴다.

single master 에 생기는 hot-spotting 을 줄이기 위해서는 모든 마스터 (active 와 stand-by) 는 커넥션 메타데이터를 fetch 하기위해서 필요한 서비스가 노출되어야 한다. 
이렇게 하면 클라이언트는 (active 뿐만 아니라) 모든 마스터와 통신할 수 있다.

Zookeeper 그리고 Master 중심의 connection registry 는 2.3 이상에서 사용할 수 있다. (3.0.0 이후부터는 default 가 된다.)

connection registry 를 주키퍼로 사용하려면

```
property>
    <name>hbase.client.registry.impl</name>
    <value>org.apache.hadoop.hbase.client.ZKConnectionRegistry</value>
 </property>
```
master 로 사용하려면 아래 처럼 설정해준다.
```
<property>
    <name>hbase.client.registry.impl</name>
    <value>org.apache.hadoop.hbase.client.MasterRegistry</value>
 </property>
 ```

#### MasterRegistory RPC heading
MasterRegistry

##### Additinal Nodes

- 단일 master 의 hot-spotting 을 피하기 위해 클라이언트가 무작위 순서로 요청하는 것을 차단(hedge) 한다..
- 클라이언트의 내부 커넥션 (master <-> regionserver) 은 여전히 Zookeeper 기반의 connection registry 를 사용한다.
- 클러스터의 내부 상태는 Zookeeper 에 의해 추적된다. 그래서 ZK 가용성은 이전과 똑같이 요구된다.
- 클러스터간 복제는 여전이 Zookeeper 기반 connection registry 를 사용한다.

더 많은 정보는 [design-docs](https://github.com/apache/hbase/blob/master/dev-support/design-docs/HBASE-18095-Zookeeper-less-client-connection-design.adoc) 와 [HBASE-18095](https://issues.apache.org/jira/browse/HBASE-18095) 를 참고하도록 하자.

## Client Request Filters


## Master
`HMaster` 는 Master Server 로 구현된다. Master Server 는 클러스터의  모든 Region Server 인스턴스와 모든 메타데이터 변화를 모니터링한다.

분산된 클러스터에 전통적으로 Master 는 NameNode 위에서 동작한다.

### Startup Behavior
멀티 마스터 환경에서 동작할때, 모든 마스터는 클러스터 위에서 실행하기 위해 경쟁한다. 만약 active master 가 zookeeper 에서 lease를 잃는 경우 (또는 마스터가 shutdown 될때) 남은 마스터들은 서로 마스터가 되기 위해 경쟁하게 된다.
