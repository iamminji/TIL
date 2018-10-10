
Chapter03. RDD로 프로그래밍하기
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
  
##### 기본 RDD
  
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
+ intersection
    - 양쪽 RDD에 동시에 존재하는 요소만 되돌려 준다.
    - 동작하면서 모든 중복을 제거한다.
    - 중복을 찾기 위해 셔플링이 수반되므로 union() 보다 성능이 훨씬 떨어진다.
+ subtract()
    - 다른 RDD를 인자로 받아 첫 번째 RDD의 항복 중 두 번째 RDD에 있는 항목을 제외한 항목들을 가진 RDD를 되돌려 준다.
    - 이것도 intersect()와 마찬가지로 셔플링을 수행한다.
  
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
|     aggregate(zeroValue)<seqOp, combOp)   | reduce()와 유사하나 다른 타입을 리턴한다. | rdd.aggregate(((0, 0)))((x, y) => (x._1 + y, x._2 + 1), (x, y) => (x._1 + y._1, x._2 + y._2))  | (9, 4) |
  
#### RDD 타입 간 변환하기
<code>mean()</code> 이나 <code>variance()</code>, 키/값 페어 <code>join()</code> 같은 것들은 특정한 타입의 RDD에서만 쓸 수 있다.

#### 영속화(캐싱)
스파크 RDD는 여유로운 방식으로 수행되지만, 때때로는 동일한 RDD를 여러 번 사용하고 싶을 때도 있을 것이다. 대신 여러 번 반복 연산을 피하고 싶다면 
스파크에 데이터 영속화(persist/persistence) 요청을 할 수 있다. RDD 영속화에 대한 요청을 하면 RDD를 계산한 노드들은 그 파티션들을 저장하고 있게 된다.
  
만약 메모리에 많은 데이터를 올리려고 시도하면 스파크는 LRU 캐시 정책에 따라 오래된 파티션들을 자동으로 버린다.

