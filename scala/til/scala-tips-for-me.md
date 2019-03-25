# 유용한 스칼라 활용법

## 리스트 값들을 카운팅 하기

```
scala> val list = List(1,2,1,2,3,4,5,1,2)
list: List[Int] = List(1, 2, 1, 2, 3, 4, 5, 1, 2)

scala> list.groupBy(identity).mapValues(_.length)
res0: scala.collection.immutable.Map[Int,Int] = Map(5 -> 1, 1 -> 3, 2 -> 3, 3 -> 1, 4 -> 1)
```

### ref
- [https://stackoverflow.com/questions/11448685/scala-how-can-i-count-the-number-of-occurrences-in-a-list](https://stackoverflow.com/questions/11448685/scala-how-can-i-count-the-number-of-occurrences-in-a-list)
- [identity function](https://en.wikipedia.org/wiki/Identity_function)
- [https://stackoverflow.com/questions/28407482/what-does-predef-identity-do-in-scala](https://stackoverflow.com/questions/28407482/what-does-predef-identity-do-in-scala)

