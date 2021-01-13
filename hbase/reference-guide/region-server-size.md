# RegionServer Sizing Rules of Thumb
리젼 서버 사이즈를 줄이는 법

## On the number of column families

컬럼 패밀리는 2~3 개 정도가 좋다. 사실 하나의 컬럼 패밀리 데이터를 가져온다고 해도 인접한 컬럼 패밀리의 데이터 까지 같이 가져온다(정확히는 flush 된다). 그 이유는 Region 별로 flush 가 수행되기 때문이다.
그래서 많은 컬럼 패밀리가 존재하게 된다면 flush 작용 때문에 불필요한 I/O 가 발생할 수 있다.
그리고 table/region 레벨에서 컴팩션은 저장소 별로 수행된다.


컬럼 패밀리는 가능하면 1개를 사용하도록 하자. 데이터 접근 시 컬럼 범위로 사용할 경우에만 두번째, 세번째 컬럼 패밀리를 만들도록 하자. (보통 하나의 쿼리에 여러개의 컬럼 패밀리를 사용하지는 않기 때문에)

### Cardinality of ColumnFamilies

하나의 테이블에 여러개의 컬럼 패밀리가 존재하는 경우, 카디널리티를 유의하자. 하나의 컬럼 패밀리A 에 백만개의 row 가 있고, 다른 컬럼 패밀리B에 10억개가 있는 경우 컬럼 패밀리A의 row들은 여러 Region(Region Server) 에 분산될 가능성이 높다. 이로 인해 컬럼 패밀리 A 의 대량 스캔의 효율이 떨어지게 된다.

## Rowkey Design

### HotSpotting

HBase 의 Row는 row key 의 문자열 순으로 정렬 된다. 이런 설계(디자인은)는 스캔을 위해 최적화 되어 관련 row 또는 읽을 row 를 서로 가까이에 저장하게 해 준다.
그러나 잘못 설계된 row key 는 hostpotting 의 원인이 된다.
hotspotting 은 많은 양의 클라이언트 트래픽(읽기,쓰기 또는 기타 다른 작업)들이 클러스터의 하나의 노드 또는 적은 수의 노드로 접근할 때 발생한다.
트래픽은 해당 region 을 서빙하는 단일 시스템의 성능을 저하시키고 region 을 사용하지 못하게 할 수 있다.
또한 해당 호스트(서버/시스템) 이 요청을 처리할 수 없기 때문에, 같은 리젼 서버에서 서빙하는 다른 리젼에도 영향을 끼칠 수도 있다.


클러스터가 완전하고 균등하게 활용되도록 데이터 접근 패턴을 설계하는 것은 중요하다.


데이터를 쓸 때, Hotspotting 을 방지하려면 같은 region 에 있어야 하는 row 들만 모을 수 있게 row key 를 설계해야 한다.(큰 그림으로 보자면 데이터가 클러스터 전체의 여러 region 에 쓰여질 수 있도록 설계 되어야 한다.)

#### salting

salting 은 랜덤한 데이터를 row key 시작 부분에 추가하는 것을 의미한다. (암호화와는 관련이 없다.)
이 경우의 salting 은 랜덤으로 할당된 prefix 를 row key 에 추가함으로써 row key 가 정렬 되도록 만든다. 가능한 prefix 수는 데이터를 분산하려는 region 수에 해당된다.


##### salting example
아래와 같이 row key 가 설계되어 있다면, 테이블이 분리될 때 f 로 시작하기 때문에 같은 region 에 몰릴 수 있다.
```
foo0001
foo0002
foo0003
foo0004
```

salting 을 사용한다면 이렇게 임의의 알파벳을 앞에 추가 하는 것이다. prefix 가 달라서 다른 region 에 위치하게 된다.
```
a-foo0003
b-foo0001
c-foo0004
d-foo0002
```

대신 쓰기 throughput 은 좋아졌지만 읽기에는 비용이 들 것이다. (다른 region 에 있기 때문에)

### Hashing
랜덤한 문자열 할당 대신 row key 에 동일한 prefix 로 salt 되도록 하는 단방향 hash 를 사용할 수도 있다.
이 방식은 Region Server 에 부하를 분산시키지만, read 예측이 가능합니다. (어떤 row 를 읽을 것인지 예측 가능하다는 의미?)

##### hashing example

`foo0003` 의 prefix 로 `a` 가 오는 단방향 hash 를 사용하게 하면, 해당 row key 를 조회할 때 이미 키를 알고 있기 때문에 검색 가능하다.
또한 특정 키 pair 가 동일한 region 에 있도록 최적화도 가능하다.

#### Reversing in key
Hotspotting 을 피하는 세번째 방법으로는 고정 길이 또는숫자로 된 row key 를 역순으로 변경하는 것이다. 그래서 자주 변경 되는 (최하위 숫자와 같은) 부분이 첫 번째가 되도록 한다. 이 경우 row key 를 효율적으로 랜덤화 시킬 수 있지만 순서에 대한 특성은 포기하게 된다.

### Monotonically Increasing Row Keys/Timeseries Data


### Try to minimize row and column sizes


### Relationship Between RowKeys and Region Splits
제대로 이해했나 모르겠는데... row key 를 기준으로 region 분할이 일어날 때, region 이 몰리는 (hot) 가능성이 있을 수 있으니, 데이터를 잘 안다면 split 전략을 잘 짜라 이런 의미인거 같다.

## Number of Versions

### Maximum Number of Versions
maximum version 의 기본 값은 1이다. maximum number of version 을 지나치게 크게 (몇백개 이상) 하면 HFile 이 증가할 수 있다.

### Minimum Number of Versions
minimum version 의 기본 값은 0 이다. TTL 파라미터랑 주로 같이 사용한다.

## Supported Datatypes

## Joins
HBase 에는 RDBMS 처럼 Join 관련해서 제공해주는 그런건 없다. 필요하면 hash 를 하든... 직접 설계해서 사용하면 된다.

## Time To Live (TTL)
ColumnFamilies 는 TTL 을 초 단위로 설정할 수 있다. 설정하면 HBase 에서는 자동으로 expire time 에 다다르면 row 가 삭제ㄷ가 된다. (모든 버젼에 적용됨)

만료된 row 가 포함된 StroeFile 은 minor compaction 때 삭제가 된다.
`hbase.store.delete.expired.storefile` 옵션을 `false` 로 만들어주면 이 기능이 비활성화 된다.
아니면 minimum number of version 을 0 보다 작게 해도 비활성화 된다.


최근 버젼 HBase 는 cell(row, column family, column qualifier, timestamp, type 다 합친거) 에도 TTL 을 적용할 수 있다.


ColumnFamily TTL 과 Cell TTL 의 차이점은

- Cell TTL 은 초 가 아니라 밀리세컨드 단위로 동작한다.
- Cell TTL 은 ColumnFamily TTL 보다 길 수 없다

## Keeping Deleted Cells
기본적으로 delete mark 는 시작 시간까지 연장된다. delete mark 가 적용되기 전의 시간 범위를 나타내는 경우에도 Get,Scan 시에 delete 된 cell 은 볼 수 없다.

> 원문이 이건데 잘 해석이 되지 않는다. By default, delete markers extend back to the beginning of time. Therefore, Get or Scan operations will not see a deleted cell (row or column), even when the Get or Scan operation indicates a time range before the delete marker was placed.


ColumnFamiliy 는 옵션으로 삭제된 cell 을 보관할 수 있다. 이 경우 삭제된 cell 은 검색할 수 있다.


> [여기](https://hbase.apache.org/book.html#cf.keep.deleted) 는 해석하기가 좀 어렵다. 잘 이해가 안감...


Raw 스캔을 통해서 deleted cell 검색이 가능하다 뭐 그런 의미 같은데...
(일반 툼스톤 마커를 이용한 삭제가 아니라 KEEP_DELETED_CELLS 옵션을 통한 그런 것?)

```
create 'test', {NAME=>'e', VERSIONS=>2147483647}
put 'test', 'r1', 'e:c1', 'value', 10
put 'test', 'r1', 'e:c1', 'value', 12
put 'test', 'r1', 'e:c1', 'value', 14
delete 'test', 'r1', 'e:c1',  11

hbase(main):017:0> scan 'test', {RAW=>true, VERSIONS=>1000}
ROW                                              COLUMN+CELL
 r1                                              column=e:c1, timestamp=14, value=value
 r1                                              column=e:c1, timestamp=12, value=value
 r1                                              column=e:c1, timestamp=11, type=DeleteColumn
 r1                                              column=e:c1, timestamp=10, value=value
1 row(s) in 0.0120 seconds

hbase(main):018:0> flush 'test'
0 row(s) in 0.0350 seconds

hbase(main):019:0> scan 'test', {RAW=>true, VERSIONS=>1000}
ROW                                              COLUMN+CELL
 r1                                              column=e:c1, timestamp=14, value=value
 r1                                              column=e:c1, timestamp=12, value=value
 r1                                              column=e:c1, timestamp=11, type=DeleteColumn
1 row(s) in 0.0120 seconds

hbase(main):020:0> major_compact 'test'
0 row(s) in 0.0260 seconds

hbase(main):021:0> scan 'test', {RAW=>true, VERSIONS=>1000}
ROW                                              COLUMN+CELL
 r1                                              column=e:c1, timestamp=14, value=value
 r1                                              column=e:c1, timestamp=12, value=value
1 row(s) in 0.0120 seconds
```

위 예제를 보니까 원래는 delete 한 특졍 버젼 미만은 안보여지고 major_compact 를 돌리고 나면 delete 한 것도 사라지는 것 같다.

근데 `KEEP_DELETED_CELLS` 옵션을 적용하게 된다면 아래처럼 되는 것임

```
hbase(main):005:0> create 'test', {NAME=>'e', VERSIONS=>2147483647, KEEP_DELETED_CELLS => true}

put 'test', 'r1', 'e:c1', 'value', 10
put 'test', 'r1', 'e:c1', 'value', 12
put 'test', 'r1', 'e:c1', 'value', 14
delete 'test', 'r1', 'e:c1',  11

hbase(main):010:0> scan 'test', {RAW=>true, VERSIONS=>1000}
ROW                                                                                          COLUMN+CELL
 r1                                                                                          column=e:c1, timestamp=14, value=value
 r1                                                                                          column=e:c1, timestamp=12, value=value
 r1                                                                                          column=e:c1, timestamp=11, type=DeleteColumn
 r1                                                                                          column=e:c1, timestamp=10, value=value
1 row(s) in 0.0550 seconds

hbase(main):011:0> flush 'test'
0 row(s) in 0.2780 seconds

hbase(main):012:0> scan 'test', {RAW=>true, VERSIONS=>1000}
ROW                                                                                          COLUMN+CELL
 r1                                                                                          column=e:c1, timestamp=14, value=value
 r1                                                                                          column=e:c1, timestamp=12, value=value
 r1                                                                                          column=e:c1, timestamp=11, type=DeleteColumn
 r1                                                                                          column=e:c1, timestamp=10, value=value
1 row(s) in 0.0620 seconds

hbase(main):013:0> major_compact 'test'
0 row(s) in 0.0530 seconds

hbase(main):014:0> scan 'test', {RAW=>true, VERSIONS=>1000}
ROW                                                                                          COLUMN+CELL
 r1                                                                                          column=e:c1, timestamp=14, value=value
 r1                                                                                          column=e:c1, timestamp=12, value=value
 r1                                                                                          column=e:c1, timestamp=11, type=DeleteColumn
 r1                                                                                          column=e:c1, timestamp=10, value=value
1 row(s) in 0.0650 seconds
```

major_compact 를 돌려도 삭제 되지 않는다. 오직 삭제 되려면 major version 보다 커져야만? 삭제가 되는 것 같다. (또는 TTL 에 의해 만료 되거나)

## Secondary Indexes and Alternate Query Paths


## 참고
- [https://hbase.apache.org/book.html#regionserver_sizing_rules_of_thumb](https://hbase.apache.org/book.html#regionserver_sizing_rules_of_thumb)
- [http://hadoop-hbase.blogspot.com/2013/01/hbase-region-server-memory-sizing.html](http://hadoop-hbase.blogspot.com/2013/01/hbase-region-server-memory-sizing.html)
- [https://www.slideshare.net/deview/211-hbase](https://www.slideshare.net/deview/211-hbase)
- [https://phoenix.apache.org/salted.html](https://phoenix.apache.org/salted.html)
- [https://itpeernetwork.intel.com/discussion-on-designing-hbase-tables/#gs.pw5s14](https://itpeernetwork.intel.com/discussion-on-designing-hbase-tables/#gs.pw5s14)
- [https://issues.apache.org/jira/browse/HBASE-11682](https://issues.apache.org/jira/browse/HBASE-11682)
