# 유용한 스칼라 활용법

## 리스트 값들을 카운팅 하기

```
scala> val list = List(1,2,1,2,3,4,5,1,2)
list: List[Int] = List(1, 2, 1, 2, 3, 4, 5, 1, 2)

scala> list.groupBy(identity).mapValues(_.length)
res0: scala.collection.immutable.Map[Int,Int] = Map(5 -> 1, 1 -> 3, 2 -> 3, 3 -> 1, 4 -> 1)
```

#### ref
- [https://stackoverflow.com/questions/11448685/scala-how-can-i-count-the-number-of-occurrences-in-a-list](https://stackoverflow.com/questions/11448685/scala-how-can-i-count-the-number-of-occurrences-in-a-list)
- [identity function](https://en.wikipedia.org/wiki/Identity_function)
- [https://stackoverflow.com/questions/28407482/what-does-predef-identity-do-in-scala](https://stackoverflow.com/questions/28407482/what-does-predef-identity-do-in-scala)


## 자바 컬렉션 스칼라에서 사용하기
scala 2.11 에는 `TreeMap` 이 없어서 자바의 `java.util.TreeMap` 을 사용하려고 하였다. 그러나 `java.util.TreeMap` 에는 `map` 멤버가 없었다. 자바의 컬렉션도 스칼라 처럼 사용하고 싶었고 그래서 찾아보니 아래의 패키지를 사용하면 되었다.

```
import scala.collection.JavaConversions._
import scala.collection.JavaConverters._
```

`JavaConversions` 를 쓰면 자동으로 컬렉션에서 scala 의 컬렉션 처럼 사용할 수 있게 되고 (예를 들면 `map` 이라든지 `foreach` 같은 멤버 함수들) `JavaConverters`를 쓰면 기존의 컬렉션 뒤에 `asScala` 를 붙여서 사용하면 된다.


```
scala> val k = new java.util.TreeMap[String, Integer]()
k: java.util.TreeMap[String,Integer] = {}

scala> import scala.collection.JavaConverters._
import scala.collection.JavaConverters._

scala> k.put("a", 1)
res0: Integer = null

scala> k.put("2", 1)
res1: Integer = null

scala> k.put("3", 2)
res3: Integer = null

scala> k.asScala.map {
     | x =>
     | println(x)
     | }
(2,1)
(3,2)
(a,1)
res4: scala.collection.mutable.Iterable[Unit] = ArrayBuffer((), (), ())

scala> import scala.collection.JavaConversions._
import scala.collection.JavaConversions._

scala> k.map {x=>println(x)}
<console>:19: warning: object JavaConversions in package collection is deprecated (since 2.12.0): use JavaConverters
       k.map {x=>println(x)}
       ^
(2,1)
(3,2)
(a,1)
res5: scala.collection.mutable.Iterable[Unit] = ArrayBuffer((), (), ())
```

### JavaConversion VS JavaConverters
2.12.X 부터는 `JavaConversion` 은 중단되었으므로 `JavaConverters`의 사용을 권장하고 있다.


#### ref
- [https://stackoverflow.com/questions/8301947/what-is-the-difference-between-javaconverters-and-javaconversions-in-scala](https://stackoverflow.com/questions/8301947/what-is-the-difference-between-javaconverters-and-javaconversions-in-scala)

### isBlank vs isEmpty

StringUtils.isBlank()

```
StringUtils.isBlank(null)      = true
StringUtils.isBlank("")        = true  
StringUtils.isBlank(" ")       = true  
StringUtils.isBlank("bob")     = false  
StringUtils.isBlank("  bob  ") = false
```

StringUtils.isEmpty()

```
StringUtils.isEmpty(null)      = true
StringUtils.isEmpty("")        = true  
StringUtils.isEmpty(" ")       = false  
StringUtils.isEmpty("bob")     = false  
StringUtils.isEmpty("  bob  ") = false
```

### 스칼라에서 서로 다른 사이즈의 리스트를 동시에 순회하는 법
`zipAll` 이라는 리스트에 멤버가 있다.

```
scala> val list1 = List("a", "b", "c")
list1: List[String] = List(a, b, c)

scala> val list2 = List("x", "y")
list2: List[String] = List(x, y)

scala> list1.zipAll(list2, "", "")
res0: List[(String, String)] = List((a,x), (b,y), (c,""))
```

`zipAll` 은 세가지를 arguments 로 받는다.
- 같이 순회할 다른 리스트
- zipAll 멤버를 호출한 리스트의 길이가 짧은 경우 사용할 기본 값
- 다른 리스트의 길이가 짧은 경우 사용할 기본 값

#### ref
- [https://stackoverflow.com/questions/35405607/zip-two-lists-of-different-lengths-with-default-element-to-fill](https://stackoverflow.com/questions/35405607/zip-two-lists-of-different-lengths-with-default-element-to-fill)

### Map 에서 Value 만 가져오는 법

```
scala> val testmap = scala.collection.immutable.HashMap[String, List[String]]("1" -> List("hello", "world"), "2" -> List("test"))
testmap: scala.collection.immutable.HashMap[String,List[String]] = Map(1 -> List(hello, world), 2 -> List(test))

scala> testmap.values
res0: Iterable[List[String]] = MapLike.DefaultValuesIterable(List(hello, world), List(test))

scala> testmap.map(_._2)
res1: scala.collection.immutable.Iterable[List[String]] = List(List(hello, world), List(test))

scala> testmap.values.flatten
res2: Iterable[String] = List(hello, world, test)

scala> testmap.map(_._2).flatten
res3: scala.collection.immutable.Iterable[String] = List(hello, world, test)

scala> testmap.flatMap(_._2)
res4: scala.collection.immutable.Iterable[String] = List(hello, world, test)
```

### 문자열에서 가장 마지막 글자만 제거하는 법

```
scala> var t = "test123!"
t: String = test123!

scala> t.dropRight(1)
res5: String = test123

scala>
```

### 순환 참조

부모 클래스가 하위 클래스를 참조하는 경우, 순환 참조 문제가 생겨 컴파일이 불가능하다. 이 때 모든 클래스는 하나의 `.scala` 파일에 넣어서 컴파일해야 한다.

### custom 정렬하기

기본적으로 scala List 에는 `sortWith` 이라는 메서드가 있다. 해당 메서드의 인자 값으로 키 (정확히는 정렬 함수) 를 넣어주면 custom 정렬이 가능하다.

아래의 예는 path 를 가장 긴 길이 순으로 정렬한 것이다.

```
val rules = scala.collection.mutable.SortedSet.empty[String]
rules += "/hello/world"
rules += "/hello/world/1234"
rules += "/hello/world/1234/4"
rules += "/hello/world/456/4/0/2"
rules += "/hello/4"
rules += "/hi/4/1/2/3/4/5/6"
rules += "/hi/4/1/2/3/4/5"
rules += "/hi/4/1/2/3/4/5/8"
rules += "/hi/4/1/2"
rules += "/hello/4/1/2/3/4/5/6"

println(rules mkString " ")
println("------------------------------------------------")
rules.foreach(
    x =>
        println(x, x.count(_ == '/'))
)
println("------------------------------------------------")
val result = rules.toList.sortWith(_.count(_ == '/') > _.count(_ == '/'))
println(result mkString " ")
println("------------------------------------------------")
result.foreach(
    x =>
        println(x, x.count(_ == '/'))
)
println("------------------------------------------------")
}
```

결과 값은 아래처럼 출력된다. 같은 길이 ('/' 을 포함한 문자 개수) 를 가졌지만 hi 를 가지고 있는 path 보다 hello 가 먼저 등장했다.
그 이유는 스칼라의 `sorting` 은 기본적으로 `stable` 하기 때문이다. (코드 보니, 내부적으로 `sortWith` 가 `sortBy`를 사용하고 있었다.)

```
/hello/4 /hello/4/1/2/3/4/5/6 /hello/world /hello/world/1234 /hello/world/1234/4 /hello/world/456/4/0/{VAR_NUM} /hi/4/1/2 /hi/4/1/2/3/4/5 /hi/4/1/2/3/4/5/6 /hi/4/1/2/3/4/5/8
------------------------------------------------
(/hello/4,2)
(/hello/4/1/2/3/4/5/6,8)
(/hello/world,2)
(/hello/world/1234,3)
(/hello/world/1234/4,4)
(/hello/world/456/4/0/2,6)
(/hi/4/1/2,4)
(/hi/4/1/2/3/4/5,7)
(/hi/4/1/2/3/4/5/6,8)
(/hi/4/1/2/3/4/5/8,8)
------------------------------------------------
/hello/4/1/2/3/4/5/6 /hi/4/1/2/3/4/5/6 /hi/4/1/2/3/4/5/8 /hi/4/1/2/3/4/5 /hello/world/456/4/0/{VAR_NUM} /hello/world/1234/4 /hi/4/1/2 /hello/world/1234 /hello/4 /hello/world
------------------------------------------------
(/hello/4/1/2/3/4/5/6,8)
(/hi/4/1/2/3/4/5/6,8)
(/hi/4/1/2/3/4/5/8,8)
(/hi/4/1/2/3/4/5,7)
(/hello/world/456/4/0/2,6)
(/hello/world/1234/4,4)
(/hi/4/1/2,4)
(/hello/world/1234,3)
(/hello/4,2)
(/hello/world,2)
------------------------------------------------
```

ref
- [https://stackoverflow.com/questions/35369243/is-scala-sorting-stable/35369500](https://stackoverflow.com/questions/35369243/is-scala-sorting-stable/35369500)
