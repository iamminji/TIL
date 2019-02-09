# Scala Tips

내가 나중에 보고 싶어서 정리하는 스칼라 팁

## 리스트 List

### 리스트에 원소 추가하기

#### 원소 앞에 추가하기

```
scala> val l = List(1,2,3)
l: List[Int] = List(1, 2, 3)
scala> 5 :: l
res0: List[Int] = List(5, 1, 2, 3)
scala> 5 +: l
res1: List[Int] = List(5, 1, 2, 3)
```

#### 원소 뒤에 추가하기

리스트 뒤에 `:+` 을 해준다.

```
scala> l :+ 6
res0: List[Int] = List(1, 2, 3, 6)
```

### 리스트에서 짝수만 출력하기

우선 짝수일 때 true를 리턴하는 함수를 하나 만든다.
```
def isEven(n: Int) = (n % 2) == 0
```

그리고 쓸 때는 아래 처럼 쓴다.
```
List(1, 2, 3, 4) filter isEven foreach println
```

위의 코드와 아래의 코드들의 의미가 같다.
```
List(1, 2, 3, 4).filter((i: Int) => isEven(i)).foreach((i: Int) => println(i))
List(1, 2, 3, 4).filter(i => isEven(i)).foreach(i => println(i))
List(1, 2, 3, 4).filter(isEven).foreach(println)
```

## for 루프

### yield 로 값 만들기
for 루프를 통해 결과를 모아서 새로운 컬렉션을 만들 수 있다.

```
val languages = List("python", "java", "scala", "javascript", "go")
val filteredLanguages = for {
  lang <- languages
  if lang.contains("java")
  if !lang.contains("script")
} yield lang
```

가드를 통해 필터링 할 수 있고, 결과적으로 filteredLanguages 엔 java 만 들어간다.
