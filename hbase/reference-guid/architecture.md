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
