
RDD
==


__RDD__ 는 단순하게 분산되어 존재하는 데이터 요소들의 모임이다. 스파크에서의 모든 작업은 새로운 RDD를 만들거나, 존재하는 RDD를 변형하거나,
결과 계산을 위해 RDD에서 연산을 호출하는 것 중의 하나로 표현된다.
그리고 내부적으로 스파크는 자동으로 RDD에 있는 데이터들을 클러스터에 분배하며 클러스터 위에서 수행하는 연산들을 병렬화한다.

#### RDD 기초
스파크의 RDD는 단순하게 말하면 분산되어 있는 변경 불가능한 객체 모음이다. 각 RDD는 클러스터의 서로 다른 노드들에게서 연산 가능하도록
여러 개의 파티션(partition) 으로 나뉜다. RDD는 사용자 정의 클래스를 포함해서 파이썬, 자바, 스칼라의 어떤 타입의 객체든 가질 수 있다.

RDD는 외부 데이터세트를 로드하거나 드라이버 프로그램에서 객체 컬렉션을 분산시키는 두 가지 방법 중의 하나로 만들 수 있다.

한 번 만들어진 RDD는 아래의 두 가지 타입의 연산을 지원한다.
+ 트랜스포메이션
    - 트랜스포메이션은 존재하는 RDD에서 새로운 RDD를 만들어 낸다.
+ 액션
    - 액션은 RDD를 기초로 결과 값을 계산하며 그 값을 드라이버 프로그램에 되돌려 주거나 외부 스토리지에 저장하기도 한다.


비록 아무 때나 새로운 RDD를 정의할 수는 있지만, 스파크는 그것들을 늘 여유로운 방식 _lazy evaluation_ 으로 처음 액션을 사용하는 시점에 처맇나다.

마지막으로, 스파크의 RDD들은 기본적으로 액션이 실행될 때 마다 매번 새로 연산을 한다. 만약 여러 액션에서 RDD 하나를 재 사용하고 싶으면 스파크에게
<code>RDD.persist()</code> 를 사용하여 계속 결과를 유지하도록 요청할 수 있다.

#### RDD 생성하기
스파크는 RDD 를 만드는 두 가지 방법, 즉 외부 데이터세트의 로드와 직접 만든 드라이버 프로그램에서 데이터 집합을 병렬화 하는 것을 제공한다.

+ RDD를 만드는 가장 간단한 방법은 당신이 만든 프로그램에 있는 데이터세트를 가져다가 <code>SparkContext</code>의 <code>paralleize()</code>
메소드에 넘겨주는 것이다.

<pre><code>val lines = sc.paralleize(List("pandas", "i like pandas"))
</code></pre>

+ RDD를 만드는 더욱 일반적인 방법은 외부 스토리지에서 데이터를 불러오는 것이다.

<pre><code>val lines = sc.textFile("/path/to/README.md")
</code></pre>

#### RDD의 연산
RDD는 두 가지 타입의 연산 작업, 즉 트랜스포메이션과 액션을 지원한다.

+ 트랜스포메이션
    - <code>map()</code> 이나 <code>filter()</code> 처럼 새로운 RDD를 만들어 내는 연산들이다.
+ 액션
    - <code>count()</code> 나 <code>first()</code> 처럼 드라이버 프로그램에 결과를 되돌려 주거나 스토리지에 결과를 써 넣는 연산들이며, 실제 계산을 수행 한다.

##### 트랜스포메이션
트랜스포메이션은 새로운 RDD를 만들어 돌려주는 RDD의 연산 방식이다. 트랜스포메이션된 RDD는 실제로 액션을 사용하는 다소 늦은 시점에 계산된다.
많은 트랜스포메이션들을 __(데이터) 요소 위주__, 다시 말하면 한 번에 하나의 요소에서만 작업이 이루어진다. 하지만 모든 트랜스포메이션이 그렇지는 않다.

##### 액션
액션은 드라이버 프로그램에 최종 결과 값을 되돌려 주거나 외부 저장소에 값을 기록하는 연산 작업이다. 액션은 실제로 결과 값을 내어야 하므로 트랜스포메이션이
계산을 수행하도록 강제한다.

##### 여유로운 수행 방식 (lazy evaluation)
스파크는 액션을 만나기 전까지는 실제로 트랜스포메이션을 처리 하지 않는다.

즉, RDD에 대한 트랜스포메이션을 호출할 때(예: map()) 그 연산이 즉시 수행되는 것은 아니다. 대신 내부적으로 스파크는 메타데이터에 연산이 요청되었다는
사실만을 기록하다. 그러므로 RDD가 실제로는 어떤 특정 데이터를 갖고 있는 게 아니라 트랜스포메이션들이 생성한 데이터를 어떻게 계산할지에 대한 명령어들을
갖고 있다고 생각하는 것이 구조를 이해하기에 더 쉽다.

#### 스파크에 함수 전달하기

<pre><code>class SearchFunctions(val query: String) {

   def isMatch(s: String): Boolean = {
       s.contains(query)
   }

   def getMatchesFunctionReference(rdd: RDD[String]): RDD[Boolean] = {
       rdd.map(isMatch)
   }

   def getMatchesFieldReference(rdd: RDD[String]): RDD[Array[String]] = {
       rdd.map(x => x.split(query))
   }

   def getMatchesNoReference(rdd: RDD[String]): RDD[Array[String]] = {
       val query_ = this.query
       rdd.map(x => x.split(query_))
   }
}
</code></pre>

#### 많이 쓰이는 트랜스포메이션과 액션

###### 데이터 요소 위주 트랜스포메이션

+ map()
    - RDD의 각 제이터에 적용하고 결과 RDD에 각 데이터의 새 결과 값을 담는다.
    - 반환 타입이 입력 타입과 같지 않아도 된다.
    - flatMap()
        = 간혹 각 입력 데이터에 대해 여러 개의 아웃풋 데이터를 생성해야 할 때도 있다. 이를 위한 연산은 __flatMap()__ 이라 불린다.
        = flatMap() 은 반환받은 반복자들을 _펼쳐 놓는다_ 라고 생각하면 쉽다.
+ filter()
    - 함수를 받아 filter() 함수를 통과한 데이터만 RDD에 담아 리턴한다.

###### 가상 집합 연산
RDD는 합집합(union), 교집합(intersect) 같은 다양한 수학적 집합 연산을 지원한다. 심지어 RDD가 집합의 형태가 아닌 경우에도 지원한다.

+ distinct()
    - 중복이 없는 데이터 세트를 원할 경우에 사용한다.
+ union()
    - 양쪽의 데이터를 합해서 되돌려준다.
    - 스파크에서는 중복을 유지한다.
+ intersection()
    - 양쪽 RDD에 동시에 존재하는 요소만 되돌려 준다.
    - 동작하면서 모든 중복을 제거한다.
    - 중복을 찾기 위해 셔플링이 수반되므로 union() 보다 성능이 훨씬 떨어진다.
+ subtract()
    - 다른 RDD를 인자로 받아 첫 번째 RDD의 항복 중 두 번째 RDD에 있는 항목을 제외한 항목들을 가진 RDD를 되돌려 준다.
    - 이것도 intersection()과 마찬가지로 셔플링을 수행한다.

그 밖에도 두 RDD에 대한 카테시안 곱(Cartesian Product) 을 계산할 수도 있다. __cartesian__ 트랜스포메이션은 첫 번째 RDD에 있는 데이터 a와
두 번째 RDD에 있는 데이터 b에 대해 모든 가능한 쌍 (a, b)를 되돌려준다.
카테시안 곱은 모든 사용자들에 대해 흥미 있어 할 만한 제안을 파악하는 등, 가능한 쌍들에 대한 유사성을 파악하고 싶은 경우에 쓸 수 있다. 대신 비용이 매우 큰 작업이다.

_최종 정리_

rdd 에 List(1, 2, 3, 3) 가 있다고 가정한다.
<pre><code> val rdd = sc.parallelize(List(1, 2, 3, 3))
</code></pre>

|    함수 이름    |       용도                                      | 예 | 결과 |
| ------------- |:-----------------------------------------------------:| ------ | ------|
|     map()     |  RDD의 각 요소에 함수를 적용하고 결과 RDD를 되돌려 준다. | rdd.map(x => x+1)  |  {2, 3, 4, 4} |
|     flatMap()     |  RDD의 각 요소에 함수를 적용하고 반환된 반복자의 내용들로 이루어진 RDD를 되돌려 준다. 종종 단어 분해를 위해 쓰인다. | rdd.flatMap(x => x.to(3))  |  {1, 2, 3, 2, 3, 3, 3} |
|     filter()     |  filter()로 전달된 함수의 조건을 통과한 값으로만 이루어진 RDD를 되돌려 준다. | rdd.filter(x => x != 1)  |  {2, 3, 3} |
|     distinct()     |  중복 제거 | rdd.distinct()  |  {1, 2, 3} |
|     sample(withReplacement, fraction, [seed])   | 복원 추출(withReplacement=true)이나 비복원 추출로 RDD에서 표본을 뽑아낸다.   | rdd.sample(false, 0.5) | 생략 |

예제를 실행하고 결과를 보고 싶으면 <code>.collect().foreach(println)</code> 을 결과 rdd 값에 실행하면 된다.

아래와 같은 RDD를 가지고 있을 경우의 정리이다.
<pre><code> val rdd = sc.parallelize(List(1, 2, 3))
val other = sc.parallelize(List(3, 4, 5))
</code></pre>

|    함수 이름    |       용도                                      | 예 | 결과 |
| ------------- |:-----------------------------------------------------:| ------ | ------|
|     union()     |  두 RDD에 있는 데이터들을 합한 RDD를 생성한다. | rdd.union(other)  |  {1, 2, 3, 3, 4, 5} |
|     intersection()     | 양쪽 RDD에 모두 있는 데이터들만을 가진 RDD를 반환한다. | rdd.intersection(other)  |  {3} |
|     subtract()  | 한 RDD가 가진 데이터를 다른쪽에서 삭제한다. | rdd.subtract(other)  | {1, 2} |
|     cartesian()  | 두 RDD의 카테시안 곱 | rdd.cartesian(other)  | {(1, 3), (1, 4), ..., (3, 5)} |

###### 액션
아마도 가장 자주 쓸 만한 일반적 액션은 <code>reduce()</code> 함수인데 이 함수는 인자로 두 개의 데이터를 합쳐 같은 타입 데이터 하나를 반환하는
함수를 받는다. 그런 함수의 간단한 예로는 <code>+</code>를 사용할 수 있으며 이를 이용하면 RDD의 총합을 구할 수 있다.

<code>reduce()</code>처럼 <code>fold()</code>는 전달되는 것과 동일한 형태의 함수를 인자로 받으며, 거기에 추가로 각 파티션의 초기 호출에 쓰이는
__제로 밸류(zero value)__ 를 인자로 받는다. 이 값은 수행하는 연산에 대한 기본값 이어야 한다. (_멱등원_ 이라고도 부르는)

> fold() 에서는 대상이 되는 두 인자 중 첫 번째 인자의 값을 수정해서 되돌려 줌으로써 객체 생성을 최소화 할 수 있다. 하지만 두 번째 인자는 수정하지 않아야 한다.

다음은 RDD에 대한 기본 액션을 정리한 표이다. rdd 값은 List(1, 2, 3, 3) 라고 가정한다.

|    함수 이름    |       용도                                      | 예 | 결과 |
| ------------- |:-----------------------------------------------------:| ------ | ------|
|     collect()     |  RDD의 모든 데이터 요소 리턴 | rdd.collect()  |  {1, 2, 3, 3} |
|     count()     |  RDD의 요소 개수 리턴 | rdd.count()  |  4 |
|     countByValue()     |  RDD에 있는 각 값의 개수 리턴 | rdd.countByValue()  | {(1, 1), (2, 1), (3, 2)} |
|     take(num)     | RDD의 값들 중 num개 리턴  | rdd.take(2)  | {(1, 2)} |
|     top(num)     | RDD의 값들 중 상위 num개 리턴  | rdd.top(2)  | {(3, 3)} |
|     takeOrdered(num)(ordering)     | 제공된 ordering 기준으로 num개 값 리턴  | rdd.takeSamples(false, 1)  | 생략 |
|     takeSample(withReplacement, num, [seed])    | 무작위 값들 리턴  | rdd.takeSamples(false, 1)  | 생략 |
|     reduce(func)    | RDD의 값들을 병렬로 병합 연산한다.  | rdd.reduce((x, y) => x+y) | 9 |
|     fold(zero)(func)    | reduce()와 동일하나 제로 밸류를 넣어준다. | rdd.fold(0)((x, y) => x+y) | 9 |
|     aggregate(zeroValue)<seqOp, combOp)   | reduce()와 유사하나 다른 타입을 리턴한다. 처음 인자로는 리턴 받는 타입에 맞는 제로 벨류를 넣어주고, 두번째 인자엔 RDD의 값들을 누적 값에 연계해 주는 함수, 마지막은 두 개의 누적 값을 합쳐 주는 함수가 필요하다. | rdd.aggregate(((0, 0)))((x, y) => (x._1 + y, x._2 + 1), (x, y) => (x._1 + y._1, x._2 + y._2))  | (9, 4) |

#### RDD 타입 간 변환하기
<code>mean()</code> 이나 <code>variance()</code>, 키/값 페어 <code>join()</code> 같은 것들은 특정한 타입의 RDD에서만 쓸 수 있다.

#### 영속화(캐싱)
스파크 RDD는 여유로운 방식으로 수행되지만, 때때로는 동일한 RDD를 여러 번 사용하고 싶을 때도 있을 것이다. 대신 여러 번 반복 연산을 피하고 싶다면
스파크에 데이터 영속화(persist/persistence) 요청을 할 수 있다. RDD 영속화에 대한 요청을 하면 RDD를 계산한 노드들은 그 파티션들을 저장하고 있게 된다.

만약 메모리에 많은 데이터를 올리려고 시도하면 스파크는 LRU 캐시 정책에 따라 오래된 파티션들을 자동으로 버린다.

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
