
정렬 (Sorting)
==

###### 최초 작성일: 2018-10-01


정렬
--

맵리듀스는 기본적으로 입력 데이터의 키를 기준으로 정렬되기 때문에 하나의 Reduce 태스크만 실행되게 한다면 정렬을 쉽게 해결할 수 있다.
하지만 여러 데이터노드가 구성된 상황에서 하나의 Reduce 태스크만 실행하는 것은 분산 환경의 장점을 살리지 않은 것이며,
대량의 데이터를 정렬하게 될 경우 네트워크 부하도 상당할 것이다.

  
하둡은 크게 다음과 같은 정렬이 있다.

- 보조 정렬 (Secondary Sort)
- 부분 정렬 (Partial Sort)
- 전체 정렬 (Total Sort)


### 보조정렬

보조 정렬은 키의 값들을 그룹핑하고, 그룹핑된 레코드에 순서를 부여하는 방식이다.
아래와 같은 순서대로 진행된다.

1. 기존 키의 값들을 조합한 복합키(Composite Key) <code>Nature Key +  Secondary Key</code> 를 정의한다. 이 때 키의 값 중에서 어떤 키를 그룹핑 키로 사용할지 결정한다.
2. 복합키의 레코드를 정렬하기 위한 비교기(Comparator) 를 정의한다.
3. 그룹핑 키를 파티셔닝할 파티셔너(Partitioner) 를 정의한다.
4. 그룹핑 키를 정렬하기 위한 비교기(Comparator) 를 정의한다.


#### 복합키 정의

복합키는 <code>WritableComparable</code> 인터페이스를 구현한다.

우선, 키로 사용할 멤버 변수를 지정한다.

<pre><code>public class MyKey implements WritableComparable&lt;MyKey&gt; {
    private String natureKey = "";
    private String secondaryKey = "";
}
</code></pre>
그리고 <code>WritableComparable</code> 엔 반드시 <code>readFields</code>, <code>write</code>, <code>compareTo</code> 메서드가 구현되어 있어야 한다.

<pre><code>@0verride
public void readFields(Datalnput in) throws IOException {
/* 입력 스트림에서 데이터 조회 */
}

@0verride
publi⊂ void write(DataOutput out) throws IOException {
/* 출력 스트림에 데이터 출력*/
}

@0verride
public int compareTo(DateKey key) {
/* 복합키와 복합키를 비교해 순서를 정함 */
}
</code></pre>


#### 복합키 비교기 구현 (Comparator)

<code>org.apache.hadoop.io.Writable</code> 인터페이스를 구현하려면 반드시 <code>org.apache.hadoop.io.WritableComparator</code>를 상속받아 구현해야 한다.

<code>WritableComparator</code>에서 _compare_ 는 이미 구현되어 있지만 객체를 스트림에서 조회한 값을 비교하게 되므로 정확한 정렬 순서를 부과할 수 없다.
따라서 메서드를 재정의해서 멤버 변수를 비교하는 로직을 구현해야 한다. 


#### 그룹키 파티셔너 구현 (Partitioner)

파티셔너는 맵 태스크의 출력 데이터를 리듀스 태스크의 입력 데이터로 보낼지 결정하고, 이렇게 파티셔닝된 데이터는 맵 태스크의 출력 데이터의 키의 값에 따라 정렬된다.
(_Map_ 과 _Reduce_ 사이에 수행된다.)

파티셔너는 반드시 <code>org.apache.hadoop.mapreduce.Partitioner</code> 를 상속 받아 구현해야 한다.

<pre><code>public class GroupKeyPartitioner extends Partitioner<MyKey, IntWritable> {
}
</code></pre>

파티셔너에 설정하는 두 개의 파라미터는 _Mapper_ 의 출력 데이터의 키와 값에 해당하는 파라미터이다.
파티셔너는 <code>getPartition</code> 메서드를 호출해 파티셔닝 번호를 조회한다.

<pre><code>@Override
public int getPartition(MyKey key, IntWritable val, int numPartitions) {
    int partition = blah blah blah;
    return partition;
}
</code></pre>

##### 질문1. 왜 파티셔너를 사용할까?
TODO. 좀 더 자세히 적어볼 것
같은 키를 가진 데이터가 같은 리듀서로 가는것을 확인하기 위해서.

##### 질문2. 파티셔너를 언제 사용하면 좋을까?
TODO. 추후에 직접 사용해보고 적어볼 것



#### 그룹키 비교기 구현 (Comparator)

리듀서는 그룹키 비교기를 사용해 같은 연도에 해당하는 모든 데이터를 하나의 Reducer 그룹에서 처리할 수 있다.


* 참고
- 시작하세요 하둡 프로그래밍 (ISBN : 9791158390389)
- https://www.quora.com/Why-is-Partitioner-used-in-Hadoop-MapReduce
- https://data-flair.training/blogs/hadoop-partitioner-tutorial/