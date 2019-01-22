
Chapter04. 키/값 페어로 작업하기
==


#### 배경

스파크는 키/값 쌍을 가지고 있는 RDD에 대해 특수한 연산들을 제공한다. 이 RDD들은 페어 RDD라 불린다.

#### 페어 RDD 생성
스파크에서 페어 RDD를 만드는 방법은 다양하다. 키/값 데이터를 곧바로 페어 RDD로 만들어 리턴할 수도 있지만 보통은 일반 RDD를 이용한다.
이를 위해서는 키/값 쌍을 되돌려 주는 <code>map()</code> 함수를 써서 변환할 수 있다.

스칼라에서는 키를 가지는 데이터를 위한 함수들을 위해 튜플을 리턴해야 한다. 튜플의 RDD에는 추가적인 키/값 함수들을 제공하기 위해 묵시적인 변환이 있다.

<pre><code>val pairs = lines.map(x => (x.split(" ")(0), x))
</code></pre>

위의 코드는 여러 라인의 텍스트를 가진 RDD로 시작해 각 라인의 첫 단어를 키로 해서 데이터를 만든다.

메모리에 있는 데이터로부터 페어 RDD를 만들어내려면 페어 데이터세트에서 <code>SparkContext.parallelize()</code>만 호출하면 된다.

#### 페어 RDD의 트랜스포메이션
페어 RDD는 기본 RDD에서 가능한 모든 트랜스포메이션을 사용할 수 있다. 단, 페어 RDD는 튜플을 가지므로 개별 데이터를 다루는 함수 대신 튜플을
처리하는 함수를 전달해야 한다.

##### 페어 RDD의 트랜스포메이션 (예: {(1, 2), (3, 4), (3, 6)})

|    함수 이름    |       목적                                  | 예 | 결과 |
| ------------- |:-----------------------------------------------------:| ------ | ------|
|     reduceByKey(func)    | 동일 키에 대한 값들을 합친다. | rdd.reduceByKey((x, y) => x+y)  |  {(1, 2), (3, 10)} |
|     groupByKey()    | 동일 키에 대한 값들을 그룹화한다. | rdd.groupByKey()  |  {(1, [2]), (3, [4, 6])} |
|     combineByKey(createCombiner, mergeValue, mergeCombiners, partitioner)    | 다른 결과 타입을 써서 동일 키의 값들을 합친다. | - | - |
|     mapValues(func)    | 키의 변경 없이 페어 RDD의 각 값에 함수를 적용한다. | rdd.mapValues(x => x+1) | {(1, 3), (3, 5), (3, 7)} |
|     flatMapValues(func)    | 페어 RDD의 각 값에 대해 반복자를 리턴하는 함수를 적용하고, 리턴받는 값들에 대해 기존 키를 써서 키/값 쌍을 만든다. 종종 토큰 분리에 쓰인다. | rdd.flatMapValues(x => (x to 5)) | {(1, 2), (1, 3), (1, 4), (1, 5), (3, 4), (3, 5)} |
|     keys()    | RDD가 가진 키들만을 되돌려 준다. | rdd.keys() | {1, 3, 3} |
|     values()    | RDD가 가진 값들만을 되돌려 준다. | rdd.values() | {2, 4, 6} |
|     sortByKey()    | 키로 정렬된 RDD를 되돌려 준다. | rdd.sortByKey() | {(1, 2), (3, 4), (3, 6)} |

##### 두 페어 RDD에 대한 트랜스포메이션(rdd = {(1, 2), (3, 4), (3, 6)}, other = {(3, 9)})

|    함수 이름    |       목적                                  | 예 | 결과 |
| ------------- |:-----------------------------------------------------:| ------ | ------|
|     subtractByKey    | 다른 쪽 RDD에 있는 키를 써서 RDD의 데이터를 삭제한다. | rdd.subtractByKey(other) | {(1, 2)} |
|     join    | 두 RDD에 대해 이너 조인(inner join)을 수행한다. | rdd.join(other) | {(3, (4, 9)), (3, (6, 9))} |
|     rightOuterJoin    | 첫 번째 RDD에 있는 키들을 대상으로 두 RDD 간에 조인을 수행한다. | rdd.rightOuterJoin(other) | {(3, (Some(4), 9)), (3, (Some(6), 9)))} |
|     leftOuterJoin    | 다른 쪽 RDD에 있는 키들을 대상으로 두 RDD 간에 조인을 수행한다. | rdd.leftOuterJoin(other) | {(3, (Some(4), 9)), (3, (Some(6), 9)))} |
|     cogroup    | 동일 키에 대해 양쪽 RDD를 그룹화한다. | rdd.cogroup(other) | {(1, ([2], [])), 3, ([4, 6], [9])))} |

#### 집합 연산
데이터 세트가 키/값 쌍으로 표현될 때 동일 키에 대해 집계된 통계를 산출하는 작업은 매우 흔한 일이다. 그간 <code>fold()</code>, <code>aggregate()</code>,
<code>reduce()</code> 같은 함수들을 봤는데 페어 RDD에도 이것들과 유사한 각 키별 트랜스포메이션이 존재한다.

<code>reduceByKey()</code>는 <code>reduce()</code>와 매우 유사하다. 둘 다 함수를 받아 값들을 합치는 데에 사용한다.
<code>reduceByKey()</code>는 여러 개의 병합 연산을 실행하는데 하나의 연산은 하나의 키에 대한 작업이 되고 각 작업은 동일 키에 대한 값을 하나로 합치게 된다.
데이터 세트의 키가 매우 많을 수도 있으므로 <code>reduceByKey()</code>는 값 하나를 사용자 프로그램에 되돌려 주는 액션이 아니라, 각 키와 키에 대해 합쳐진
값으로 구성된 새로운 RDD를 리턴한다.

##### 스칼라에서 단어 세기

<pre><code>val input = sc.textFile("/파일/경로/")
val words = input.flatMap(x => x.split(" "))
val result = words.map(x => (x, 1)).reduceByKey((x, y) = x + y)
</code></pre>

<code>combineByKey()</code>는 키 별 집합 연산 함수 중 가장 일반적으로 쓰인다. <code>aggregate()</code>와 마찬가지로 입력 데이터와
동일한 타입의 값을 되돌려 줄 필요는 없다.

##### 스칼라에서 combineByKey()를 사용한 키별 평균
<pre><code>val result = input.combineByKey(
    (v) => (v, 1),
    (acc: (Int, Int), v) => (acc._1 + v, acc._2 + 1),
    (acc1: (Int, Int), acc2: (Int, Int)) => (acc1._1 + acc2._1, acc1._2 + acc2._2)).map {
        case (key, value) => (key, value._1 / value._2.toFloat)
    }
result.collectAsMap().map(println(_))
</code></pre>

##### 병렬화 수준 최적화
모든 RDD는 고정된 개수의 파티션을 갖고 있으며 이것이 RDD에서 연산이 처리될 때 동시 작업의 수준을 결정하게 된다.
<pre><code>val data = Seq(("a", 3), ("b", 4), ("a", 1))
sc.parallelize(data).reduceByKey((x, y) => x + y) // 기본 병렬화 수준 사
sc.parallelize(data).reduceByKey((x, y) => x + y, 10) // 병렬화 수준 지정
</code></pre>

#### 데이터 그룹화
키를 가진 데이터를 일반적으로 활용하는 방법 중 하나는 키에 의해 데이터를 그룹짓는 것이다.
만약 데이터가 이미 원하는 값을 키로 쓰고 있을 경우, <code>groupByKey()</code>를 쓰면 RDD의 키를 사용해서 데이터를 그룹화한다.
K 타입의 키와 V 타입의 값을 가진 RDD라면 __[K, Iterable[V]]__ 타입의 RDD를 되돌려 준다.

<code>groupBy()</code>는 쌍을 이루지 않았거나 현재 키와 관계되지 않은 다른 조건을 써서 데이터를 그룹화하고자 하는 경우에 쓰인다.

<code>cogroup()</code>는 여러 RDD에 동일 키를 공유해 데이터를 그룹화할 수도 있다. 동일한 키 타입 K와 데이터 타입 V와 W를 각각 쓰는
두 RDD에 <code>cogroup</code>을 적용하면 __[(K, (Iterable[V], Iterable[W])]__ 타입의 RDD를 리턴한다. 만약 한쪽 RDD에만
해당 키에 해당하는 값이 존재하고 다른 쪽 RDD에는 없다면, 결과 RDD에서의 해당 반복자는 그냥 비어있다. <code>cogroup()</code>은 다음 절에서
다룰 조인(join)을 위한 블록 구성을 위해 사용된다.

#### 조인
키 값을 가진 데이터에 대해 가장 유용한 연산 중의 하나는 이를 다른 키 값을 가진 데이터와 함께 사용하는 것이다.
스파크는 좌/우 외부 조인(left/right outer join), 교차 조인(cross join), 내부 조인(inner join) 등의 모든 범위의 조인을 지원한다.

##### 스칼라 셸 내부 조인 예제
<pre><code>storeAddress = {
    (Store("Ritual"), "1026 Valencia St"), (Store("Philz"), "748 Van Ness Ave"),
    (Store("Philz"), "3101 24th St"), (Store("Starbucks"), "Seattle")
}

storeRating = {
    (Store("Ritual"), 4.9), (Store("Philz"), 4.8)
}

storeAddress.join(storeRating) == {
    (Store("Ritual"), ("1026 Valencia St", 4.9)),
    (Store("Philz"), ("748 Van Ness Ave", 4.8)),
    (Store("Philz"), ("3101 24th St", 4.8)),
}
</code></pre>

##### leftOuterJoin()과 rightOuterJoin()
<pre><code>storeAddress.leftOuterJoin(storeRating) == {
    (Store("Ritual"), ("1026 Valencia St", Some(4.9))),
    (Store("Philz"), ("748 Van Ness Ave", Some(4.8))),
    (Store("Philz"), ("3101 24th St", Some(4.8))),
    (Store("Starbucks"), ("Seattle", None))
}

storeAddress.rightOuterJoin(storeRating) == {
    (Store("Ritual"), (Some("1026 Valencia St"), 4.9)),
    (Store("Philz"), (Some("748 Van Ness Ave"), 4.8)),
    (Store("Philz"), (Some("3101 24th St"), 4.8)),
}
</code></pre>

#### 데이터 정렬
데이터 정렬은 많은 경우, 특히 다운스트 데이터를 생성하는 경우에 유용하다. 한 번 데이터를 정렬하고 나면 이후에 데이터에 대한
<code>collect()</code>나 <code>save()</code> 같은 함수 호출들은 결과가 정렬되어 나오게 된다.

##### 스칼라에서의 사용자 지정 정렬, 정수를 문자열인 것 처럼 정렬하기
<pre><code>
val input: RDD[(Int, Venue)] = ...
implicit val sortIntegerByString = new Ordering[Int] {
    override def compare(a: Int, b: Int) = a.toString.compare(b.toString)
}
</code></pre>

#### 페어 RDD에서 쓸 수 있는 액션
트랜포메이션들과 마찬가지로 기본 RDD에서 가능한 모든 전통적인 액션 또한 페어RDD에서 사용할 수 있다.

##### 페어 RDD의 액션(예 ({(1, 2), (3, 4), (3,6)}))
|    함수 이름    |       목적                                  | 예 | 결과 |
| ------------- |:-----------------------------------------------------:| ------ | ------|
|     countByKey()    | 각 키에 대한 값의 개수를 센다. | rdd.countByKey() | {(1, 1), (3, 2)} |
|     collectAsMap()    | 쉬운 검색을 위해 결과를 맵 형태로 모은다. | rdd.collectAsMap() | Map{(1, 2), (3, 4), (3, 6)} |
|     lookup(key)    | 들어온 키에 대한 모든 값을 되돌려 준다. | rdd.lookup(3) | [4, 6] |

#### 데이터 파티셔닝
분산 프로그램에서 통신은 비용이 매우 크므로 네트워크 부하를 최소화할 수 있는 데이터 배치는 프로그램 성능을 비약적으로 향상시킬 수 있다.
파티셔닝은 조인 같은 키 중심의 연산에서 데이터 세트가 __여러 번__ 재활용 될 때만 의미가 있다.
스파크의 파티셔닝은 모든 RDD의 키/값 쌍에 대해 가능하며, 시스템이 각 키에 제공된 함수에 따라 값들을 그룹화하도록 한다.
비롯 스파크는 각 키가 어떤 노드로 전달되는지 같은 명시적인 제어를 제공하지는 않지만 _특정 노드가 장애가 나더라도 전체적으로는 동작할 수 있도록 설계_,
어떤 키의 모음들이 __임의의 노드__ 에 함께 모여있게 되는 것은 보장해 준다.

<pre><code>/*  코드 초기화, 사용자 정보를 HDFS의 하둡 시퀀스 파일에서 읽어 온다.
* 이는 찾아낸 HDFS 블록에 맞춰 userData의 요소들을 적절히 분산시키지만,
* 특별히 스파크에 어떤 ID가 어떤 파티션에 존재하는지에 대해서 정보를 제공하지는 않는다.
*/
val sc = new SparkContext(...)
val userData = sc.sequenceFile[UserID, UserInfo]("hdfs://...").persist()

/* 지난 5분간의 이벤트 로그 파일을 처리하기 위해 주기적으로 불리는 함수.
* 여기서 처리하는 시퀀스 파일이 (UserID, LinkInfo) 쌍을 갖고 있다고 가정한다.
*/
def processNewLogs(logFileName: String) {
    val events = sc.sequenceFile[UserID, LinkInfo](logFileName)
    /* (UserID, (UserInfo, LinkInfo)) 를 페어로 가지는 RDD */
    val joined = userData.join(events)

    val offTopicVisits = joined.filter {
        /* 각 아이템을 그 컴포넌트들로 확장한다. */
        case (userId, (userInfo, linkInfo)) =>
            !userInfo.topics.contains(linkInfo.topic)
    }.count()
    println("Number of visits to non-subscribe topics: " + offTopicVisits)
}
</code></pre>

위의 코드는 효율적이지는 않다. 이는 <code>processNewLogs()</code>가 매번 호출될 때마다 불리는 <code>join()</code>이
데이터 세트에서 키가 어떻게 파티션 되어 있는지에 대해 모르기 때문에 이 연산은 __양쪽__ 데이터 세트를 모두 해싱하고 동일 해시키의 데이터끼리
네트워크로 보내 동일 머신에 모이도록 한 후 해당 머신에서 동일한 키의 데이터끼리 조인을 수행한다. 이러면 userData에 아무 변경이 없어도 함수가
호출될 때마다 userData는 해싱 작업을 거쳐 셔플링 된다.

1. userData에서 해싱해서 동일 머신에 보냄
2. 머신에서 조인함

이를 개선하려면 프로그램 시작 때 한 번만 해시 파티션하도록 하면 된다.

<pre><code>val sc = new SparkContext(...)
val userData = sc.sequenceFile[UserId, UserInfo]("hdfs://...")
                 .partitionBy(new HashPartitioner(100)) // 파티션 100개 생성
                 .persist()
</code></pre>

usrData를 만들면서 <code>partitionBy()</code>를 호출했으므로 스파크는 userData가 해시 파티션되어 있음을 알고 이 정보를 최대한 활용할 것이다.

<code>userData.join(events)</code>를 호출할 때 스파크는 오직 events RDD만 셔플해서 이벤트 데이터를 각각의 UserID와 맞는 userData 해시 파티션이 있는 머신으로 전송할 것이다.

> persist() 영속화 없는 파티션된 RDD의 사용은 의미가 없다.
persist()가 없으면 이후 다른 RDD 액션들이 결국엔 재연산을 하게 되기 때문이다.

#### RDD의 파티셔너 정하기

<pre><code>scala> val pairs = sc.parallelize(List((1, 1), (2, 2), (3, 3)))
pairs: org.apache.spark.rdd.RDD[(Int, Int)] = ParallelCollectionRDD[0] at parallelize at <console>:24

scala> pairs.partitioner
res6: Option[org.apache.spark.Partitioner] = None

scala> import org.apache.spark.HashPartitioner
import org.apache.spark.HashPartitioner

scala> val partitioned = pairs.partitionBy(new HashPartitioner(2))
partitioned: org.apache.spark.rdd.RDD[(Int, Int)] = ShuffledRDD[1] at partitionBy at <console>:26

scala> partitioned.partitioner
res7: Option[org.apache.spark.Partitioner] = Some(org.apache.spark.HashPartitioner@2)
</code></pre>

우선 pairs에 (Int, Int) 쌍의 RDD를 만들었으나 내부적으로 파티션 정보는 존재하지 않는다.(Option 객체가 None을 가짐)
그리고 첫 번째 RDD _pairs_ 에 해시 파티셔닝을 적용해 두 번째 RDD _partitioned_ 를 만든다.
만약 partitioned 를 나중의 연산에도 쓰고 싶다면 생성할 때 뒤에 <code>.persist()</code> 를 붙인다.
그 이유는 위에서 언급했던 것과 같다.

#### 파티셔닝이 도움이 되는 연산들
스파크의 많은 연산들이 네트워크를 통해 키별로 데이터를 셔플링하는 과정을 포함한다. 이런 종류의 연산들은 모두 파티셔닝에 의해 이득을 볼 수 있다.
파티셔닝에 의한 효과를 볼 수 있는 연산들을 <code>cogroup()</code>, <code>groupWith</code>, <code>join()</code>,
<code>leftOuterJoin</code>, <code>rightOuterJoin</code>, <code>groupByKey()</code>, <code>reduceByKey</code>,
<code>combineByKey</code>, <code>lookup()</code> 이다.

이미 파티션되어 있는 단일 RDD 위에서 동작하는 <code>reduceByKey()</code> 같은 연산들은 각 키에 관련된 모든 값을 단일 머신에서 처리하게 되며,
각 작업 노드에서 병합된 최종 결과 값은 마스터 노드로 보내진다.

두 개의 RDD를 필요로 하는 <code>cogroup()</code> 이나 <code>join()</code> 같은 연산들에서 미리 파티셔닝하는 것은 최소 하나 이상의 RDD가 셔플될 필요가 없게 해준다.

만약 두 RDD가 동일한 파티셔너를 가지고 있고, 동일한 머신들에 캐시되어 있거나 둘 중 하나가 연산이 아직 완료되지 않았다면, 네트워크를 통한 셔플은 발생하지 않는다.

#### 파티셔닝에 영향을 주는 연산들
지정된 파티셔닝을 쓰는 것이 보장되지 못하는 트랜스포메이션들에 대해서는 결과 RDD에 파티셔너가 세팅되지 않는다.

> 예를 들어 키/값 쌍을 가지는 해시 파티션된 RDD에 map() 을 호출한다면, map()에 전달되는 함수는 이론적으로 각 데이터 요소의 키를
변경할 가능성이 있으므로 결과 RDD는 파티셔너를 가지지 않는다.

스파크는 사용자가 작성한 함수가 키를 유지하는지에 대해 따로 검사하지 않는다. 대신 각 튜플의 키가 동일하게 유지되는 것을 보장하는
<code>mapValues()</code> 나 <code>flatMapValues()</code> 같은 연산을 제공한다.

다음은 결과 RDD에 파티셔너가 지정되는 모든 연산들이다.

<code>cogroup()</code>, <code>groupWith</code>, <code>join()</code>,
<code>leftOuterJoin</code>, <code>rightOuterJoin</code>, <code>groupByKey()</code>, <code>reduceByKey</code>,
<code>combineByKey</code>, <code>partitionBy()</code>, <code>sort()</code>, <code>mapValues()</code>,
<code>flatMapValues</code>, <code>filter()</code> 이다.

#### 사용자 지정 파티셔너
스파크는 사용자에게 자체적으로 __Partitioner__ 객체를 만들어서 어떤 식으로 파티셔닝을 수행할지 튜닝이 가능하도록 해 준다.
사용자 지정 파티셔너를 구현하기 위해서는  <code>org.apache.spark.Partitioner</code> 클래스를 상속하고 아래 세 개의 메소드를 구현해야 한다.

+ numParitions: Int
    = 생성할 파티션의 개수를 되돌려 준다.
+ getPartition(key: Any): Int
    = 주어진 키에 대한 파티션 ID를 되돌려 준다. 음수를 리턴하지 않도록 주의를 기울일 것!
+ equals()
    = 자바의 기본적인 동일값 객체 검사 메소드
