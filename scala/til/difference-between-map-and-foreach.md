## map 과 foreach 의 차이점

[https://stackoverflow.com/questions/354909/is-there-a-difference-between-foreach-and-map](https://stackoverflow.com/questions/354909/is-there-a-difference-between-foreach-and-map)


기본적으로 map은 다른 리스트를 리턴해주고 foreach는 안해줌. foreach는 값 가지고 딴데 쓰고 싶을때(디비에 쓴다던가, 뭔가를 하고 싶을때) 사용하고 map은 새로운 컬렉션을 만들고 싶을때 쓴다.


간단한 예제
```
scala> val k = t.foreach(a => a+"!")
k: Unit = ()

scala> val k = t.map(a => a+"!")
k: scala.collection.mutable.SortedSet[String] = TreeSet(a!, b!, c!)
```
