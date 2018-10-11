
unit23. 맵 사용하기
==
  
  
맵은 다른 언어와 마찬가지로 키-값 형태로 자료를 저장한다. 또한, 슬라이스와 마찬가지로 레퍼런스 타입이다.
  
+ var 맵명 map[키_자료형]값_자료형
  
<pre><code>var a map[string]int // 키는 string, 값은 int인 맵 선언
</code></pre>
  
맵은 __make__ 함수를 사용하여 공간을 할당해야 값을 넣을 수 있다. 여기서 맵 선언과 동시에 make 함수를 사용하면 map 키워드와 자료형을 생략할 수 있다.
  
+ make(map[키_자료형]값_자료형)
  
<pre><code>var a map[string]int = make(map[string]int)
var b = make(map[string]int)
c := make(map[string]int)
</code></pre>
  
맵을 생성하면서 키와 값을 초기화하려면 __{}__ 중괄호를 사용한다. 중괄호에서 키와 값은 한 줄로 나열해도 되고 여러 줄로 나열해도 된다.
  
23.1 맵에 데이터 저장하고 조회하기
--
<pre><code>solarSystem := make(map[string]float32)
solarSystem["Mercury"] = 87.909

fmt.Println(solarySystem["Mercury"])
</code></pre>
  
맵에 존재하지 않는 키를 조회했을 때는 빈 값이 출력된다. 문자열이면 "" 을, 숫자형이면 0
  
키가 존재하는지 여부를 보려면 리턴값을 활용하면 되는데, 리턴값에서 두 번째 변수에 키의 존재 여부가 저장된다.
  
<pre><code>value, ok := solarSystem["Pluto"]

if value, ok := solarSystem["Saturn"]; ok {
}
</code></pre>
  
23.2 맵 순회하기
--
+ for 키, 값 := range 맵 {}

<pre><code>for key, value := range solarSystem {
}

// 키 안 쓰고 싶으면
for _, value := range solarSystem {
}
</code></pre>
  
23.3 맵에서 데이터 삭제하기
--
맵에서 값을 삭제하려면 __delete__ 함수를 사용한다.
  
+ delete(맵, 삭제할_키)
  
<pre><code>a := map[string]int{"Hello": 10, "World": 20}
delete(a, "World")
</code></pre>
  
23.4 맵 안에 맵 만들기
--
맵의 값 안에는 일반 자료형뿐만 아니라 맵 자체도 들어갈 수 있다.
  
+ map[키_자료형]map[키_자료형]값_자료형
  