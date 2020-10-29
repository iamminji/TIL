# HBase

## Catalog Table
hbase shell 에서 `list` 명령어를 치면 안나와서 헷갈릴 수 있지만 catalog (이하 __hbase:meta__) 도 테이블 중에 하나이다.

### hbase:meta
이전에는 `.META.` 라고 불렸던 `hbase:meta` 테이블은 모든 리젼에 대한 정보를 가지고 있다.
그리고 이 `hbase:meta` 는 주키퍼에 저장되어 있다.

`hbase:meta` 테이블 구조는 다음과 같다.

#### key
`[table],[region start key],[region id]` 형식으로 되어 있다.

#### values

- `info:regioninfo`
region 인스턴스 (HRegionInfo) 가 직렬화되어 저장되어 있다.
- `info:server`
region 정보가 있는 regionServer 의 server:port 정보가 있다.
- `info:serverstartcode`
region 정보가 있는 regionServer 의 시작 시간이 있다.

테이블이 분리될 때`info:splitA` 와 `info:splitB` 라는 두 개의 컬럼이 생성된다. 이 컬럼은 두개의 자식 리젼을 의미하고 Region 이 분할 되면 이row 들도 삭제 된다.

```
# 이런식으로 정보 확인 가능
hbase(main):004:0> scan 'hbase:meta'
```

HBase 부트 시 주키퍼에서 `hbase:meta` 위치를 찾고 `hbase:meta` 는 server 와 startcode 값을 업데이트 한다. 
(주키퍼에서 `/hbase/meta-region-server` 라는 곳에 위치해 있다.)

##### 참고
- https://hbase.apache.org/2.0/devapidocs/org/apache/hadoop/hbase/MetaTableAccessor.html


## Region-RegionServer 할당 방식
Region 이 어떻게 RegionServer 에 할당 되는걸까?

## StartUp (시작 시)
1. Master 는 `AssignmentManager` 를 실행 시킨다.
2. `AssignmentManager` 는 `hbase:meta` 에서 (이미 있는) `region` 정보를 본다.
3. 만약 `region` 이 여전히 유효하다면 (예를 들면 `RegionServer` 가 살아있는 상태다.) 할당을 한다.
4. 만약 (할당이) 유효하지 않다면 `LoadBalancerFactory` 가 region 을 할당 시킨다. __로드 밸런서가 Region 을 RegionServer 에 할당 시키는 것이다.__
5. `hbase:meta` 에 RegionServer 의 정보를 업데이트 한다. (필요하다면 RegionServer의 start code 도 저장한다.) 

## FailOver (실패 시)
1. RegionServer 가 다운되었기 때문에 해당 RegionServer 에 있는 region 은 즉시 이용 못하게(unavailable) 된다.
2. Master 는 RegionServer 의 fail 상태를 감지할 것이다.
3. region 할당이 유효하지 않아, 재 할당 되는데 시작 순서(startup)와 똑같이 진행된다.
4. 실행중인 쿼리는 재 시작되고 손실 되진 않는다.
5. 다음 시간 내에 새로운 RegionServer 에서 운영 가능해 진다.

```
ZooKeeper session timeout + split time + assignment/replay time
```

## Region Load Balancing
리젼은 주기적으로 `LoadBalancer(다운된 RegionServer 의 담당하는) ` 에 의해 옮겨진다.

### Region State Transition
HBase 는 각 region 의 상태를 `hbase:meta` 에 유지한다.
`hbase:meta` 에 있는 region 상태는 주키퍼에서 유지 한다.(??)

#### OFFLINE
region 이 오프라인 상태이다. (open 되지 않음)
#### OPENING
region 이 open 되는 중이다.
#### OPEN
region 이 open 되고 RegionServer 가 master 에게 알린다.
#### FAILED_OPEN
RegionServer 가 region 을 open 하는데 실패했다.
#### CLOSING
region 이 close 되는 중이다.
#### CLOSED
RegionServer 가 region 을 close 시키고 master 에게 알린다.
#### FAILED_CLOSE
RegionServer 가 region 을 close 시키는데 실패했다.
#### SPLITTING
RegionServer가 master 에게 region 이 분리(split) 되는 중인걸 알린다.
#### SPLIT
RegionServer 가 master 에게 region split 이 완료되었다고 알린다.
#### SPLITTING_NEW
region 이 split 되어 생성되는 중이다.
#### MERGING
RegionServer 가 master 에게 region 이 다른 region 와 merge 되었다고 알린다.
#### MERGED
RegionServer가 master 에게 region 이 merge 되었음을 알린다.
#### MERGING_NEW
region 이 두개의 region 에서  합쳐지는 중이다.

## Region-RegionServer Locality


##### 참고
- https://hbase.apache.org/book.html#regions.arch.assignment


## HBase 의 Locality
테이블당 리젼 여러개를 갖고 있고 각 리젼은 패밀리 별로 HDFS 3개 (Replication) 으로 이루어져있다. 
만약 각 패밀리에서 가지고 있는 (하나의) HDFS 가 존재하는 서버와 리젼을 가지고 있는 서버랑
동일하면 로컬리티는 1이고 그렇지 않다면 전체 패밀리 블럭 사이즈에서 없는 패밀리 블럭 사이즈만큼 1 에서 빠진다.

## HBase Balancer

