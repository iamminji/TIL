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

