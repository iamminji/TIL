# 데이터 엔지니어링

## 분산 시스템
### 고가용성
주로 SPOF 를 피하기 위해서 고가용성을 이야기 하곤 한다.

#### 쿼럼
고가용성에 대해 알아보면 `쿼럼` 이란 단어가 자주 등장한다. 분산 시스템을 다룰 때 주로 발생하는 문제는 실패나 장애가
발생한 상황에서도 서로 다른 머신에서 실행 중인 프로세스가 어떤 값에 동의할 수 있는 방법을 만드는 것이다.

이를 분산 합의(Consensus) 라고 부른다.

### CAP
`CAP` 이론은 빅데이터 플랫폼에 자주 등장하는 용어로 분산 시스템에서 일관성(Consistency), 가용성(Availability), 분할 용인(Partition tolerance)이라는 세 가지 조건을 모두 만족하는 것은 없다는 이론이다.

- Consistency
모든 `읽기` 작업은 가장 최신의 데이터(또는 에러)를 가져올 수 있어야 한다.

- Availability
모든 `요청` 작업은 성공 해 응답을 받을 수 있어야 한다. 이 때 받는 응답이 가장 최신이지는 않아도 된다.

- Partition tolerance
시스템이 partition 상태에서도 정상 동작해야 한다. 즉, 어떤 이유에 의해 paritition 으로 분리 되어도 시스템은 계속 동작해야 한다. 여기서 발생하는 원인으로는 네트워크 오류 뿐만 아니라 일부 메시지가 전달 되지 않는 상황도 포함된다.


파티션 오류가 발생하면 아래의 두 가지 중에 선택해야 한다.

1. 작업을 취소하여 availability (가용성) 을 낮추고 consistency (일관성) 을 보장한다.
2. 작업을 진행하여 availaility (가용성) 을 제공하고 consistency (일관성) 을 보장하지 않는다.


#### CAP 를 사용하는 예
주키퍼는 일관성과 분할 용인 (CP) 을 HBase 도 역시 CP를 택하고 있다.

CAP 는 위에서 언급했듯이 3가지를 모두 만족할 수 없다. 주로 C 또는 A 를 택하고 P 를 같이 가져가는 형식이다.

#### 참고
- [https://blog.seulgi.kim/2018/03/cap-theorem.html](https://blog.seulgi.kim/2018/03/cap-theorem.html)


### Split Brain 이란?

시스템의 두 부분 이상이 독립적으로 진행되어 시스템이 일관되지 않게 동작하는 것을 말한다.

분산 시스템에서 마스터-슬레이브 상태에서 네트워크 이상으로 인해 슬레이브는 마스터가 이상이 있다고 판단한다.
때로는 맞을 수도 있고 때로는 오탐일 수도 있다. 만약 잘못된 판단임에도 슬레이브 중 하나가 마스터로 선출이 되면 두 개의 마스터가 존재하게 된다. 이런 경우를 Split Brain 스플릿 브레인이라고 부른다.

> split brain 을 해결하는 방법 중 하나는 fencing 이 있다. stand by master 가 마스터 (또는 리더) 로 선출 되기 전에 active master 를 한번 더 (?) 죽이는 것이다. active master 가 정말로 죽었는지 모르기 때문에 하는 방법이다. (보통은 quroum 을 이용하긴 함)


### 쿼럼 (정족수)
- 합의를 위한 최소한의 노드 수
- 과반수 서버 노드를 채택하여 쿼럼을 구성한다.
- 과반수를 채택하는 이유는 분산 코디네이션 환경에서 예상치 못한 장애가 발생해도 분산 시스템의 일관성을 유지하기 위해서이다.

### 샤딩과 레플리케이션의 차이
- 샤딩은 데이터를 분산해서 저장
- 레플리케이션은 같은 데이터를 저장

### 구글 논문
- [The Google file system](https://static.googleusercontent.com/media/research.google.com/ko//archive/gfs-sosp2003.pdf)

확장 가능한 분산 파일 시스템인 GFS 는 범용 하드웨어로 대규모 데이터를 저장하는 클러스터를 구축한다. 이 파일 시스템은 노드 간에 데이터 사본을 저장하여 저장 서버 하나를 잃더라도 데이터 가용성에는 영향이 없었다.

- [MapReduce: Simplified Data Processing on Large Clusters](https://static.googleusercontent.com/media/research.google.com/ko//archive/mapreduce-osdi04.pdf)

맵리듀스 논문으로 인해 대규모 클러스터 데이터 처리가 간소화 되었다.

- [Bigtable: A Distributed Storage System for Structured Data](https://static.googleusercontent.com/media/research.google.com/ko//archive/bigtable-osdi06.pdf)

HBase 의 뼈대가 되는 논문으로, 거대하게 화장할 수 있도록 구조화된 데이터를 관리하는 분산 저장 시스템이 바로 빅 테이블이다.
