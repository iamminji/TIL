# 자바 이모저모

### 자바 HashMap, HashTable, ConcurrentHashMap, TreeMap
자바 Map 4총사

#### HashMap
- 자바의 HashMap 에서는 해시 충돌을 방지하기 위하여 Seperate Chaning 을 사용한다
- 자바8 부터는 해시 충돌 시 링크드리스트 대신 트리를 사용한다.
- non thread-safe 하다.
- 키/값으로 null 을 허용한다.
- 정렬되어 있지 않다.
- get/put/containseKey 연산이 O(1) 의 시간 복잡도를 갖는다.
#### TreeMap
- 키 값을 기준으로 정렬되어 있다.
- 내부적으로 red-black tree 를 사용해서 구현을 했다.
- get/put/containseKey 연산이 O(logn) 시간 복잡도를 갖는다.
#### HashTable
- thread-safe 하다.
- 키/값 으로 null 을 허용하지 않는다. (null 은 객체가 아니므로 hashCode 연산이 불가능하다.)
- HashMap 이 HashTable 에서 더 나중에 나온 발전 된 자료구조이다
#### ConcurrentHashMap
- HashTable 의 대안으로 나왔다.
- HashTable 도 thread-safe 하지만 ConcurrentHashMap 은 더 나은 성능을 갖고 있다. 그 이유는 ConcurrentHashMap 은 map 의 일부에만 lock 을 걸고 Hashtable은 map 전체에만 lock 을 걸기 때문이다.
- 쓰기보다 읽기가 많을 때 가장 적합하다.

### 자바에서 parseInt와 valueOf 의 차이점

parseInt는 primitive int를 반환하고, valueOf(String) 은 new Integer()를 반환한다.

그래서 Integer.valueOf(s) 는 new Integer(Integer.parseInt(s)) 이것과 동일하다고 한다.

보면서, 자바는 그러면 왜 integer를 primitive와 object로 둘다 두어서 사람을 헷갈리게 하는건지? 언제 object를 써야하고 primitive를 써야하는건지? 잘 모르겠다.

자바를 안 써서 피부로 와닿지 않는 것 같기도 하다.

* https://softwareengineering.stackexchange.com/questions/203970/when-to-use-primitive-vs-class-in-java
