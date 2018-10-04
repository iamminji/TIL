
Chapter05. 기본 타입과 연산
==


5.1 기본 타입
--
  
|    기본 타입    |       범위                                             |
| ------------- |:-----------------------------------------------------:|
|     Byte      | 8비트 2의 보수 표현을 사용하는 부호 있는 정수(-2^7 ~ 2^7-1)    |
|     Short     | 16비트 2의 보수 표현을 사용하는 부호 있는 정수(-2^15 ~ 2^15-1) |
|     Int       | 32비트 2의 보수 표현을 사용하는 부호 있는 정수(-2^31 ~ 2^31-1) |
|     Long      | 16비트 2의 보수 표현을 사용하는 부호 있는 정수(-2^63 ~ 2^63-1) |
|     Char      | 16비트 부호 없는 유니코드(Unicode) 문자 (0 ~ 2^16-1)        |
|     String    | Char의 시퀀스                                           |
|     Float     | 32비트 IEEE 754 단일 정밀도 부동소수점 수                   |
|     Double    | 64비트 IEEE 754 2배 정밀도 부동소수점 수                    |
|     Boolean   | true 혹은 false                                        |

5.2 리터럴
--
리터럴(Literal) 은 상수 값을 코드에 직접 적는 방법을 의미한다.
  
#### 정수 리터럴
  
#### 부동 소수점 리터럴
  
#### 문자 리터럴
  
#### 문자열 리터럴
  
문자열 리터럴은 큰 따옴표 <code>"</code>로 둘러싼 문자들로 이뤄진다. 스칼라에는 raw 문자열을 위한 특별한 문법을 추가했다.
raw 문자열은 큰 따옴표를 3개 연속으로 사용해 시작한다. _파이썬도 이와 같은 문법이 있지만 작은 따옴표와 큰 따옴표 구분이 없고 여러 줄을 쓰기 위해서만 사용한다._ 
  
<pre><code>println("""Welcome To Ultamix 3000.
                      Type "HELP" for help.""")
</code></pre>
  
raw 문자열 내부에는 큰 따옴표를 3개 연속으로 사용하는 것을 제외하면 개행 문자, 따옴표, 큰 따옴표, 특수 문자등 모든 문자를 넣을 수 있다.
  
그러나 위 코드를 실행하면 큰 따옴표 안에 있는 띄어쓰기가 먹혀, 저 문자열 그대로 출력된다. 즉 아래처럼 출력된다.
  
<pre><code>Welcome To Ultamix 3000.
           Type "HELP" for help.
</code></pre>           
  
앞에 개행을 없애고 싶다면 파이프 문자를 각 줄의 시작 부분에 넣는다.
  
<pre><code>println("""|Welcome To Ultamix 3000.
           |Type "HELP" for help.""")
</code></pre>

결과
<pre><code>Welcome To Ultamix 3000.
Type "HELP" for help.
</code></pre>
  
#### 심볼 리터럴
  
심볼 리터럴은 <code>'ident</code> 처럼 쓴다. 작은 따옴표 뒤에 오는 식별자 부분은 알파벳과 숫자를 혼합한 올바른 식별자라면 아무것이나 가능하다. 
스칼라는 이런 리터럴은 <code>scala.Symbol</code> 이라는 클래스의 인스턴스로 매핑한다. 즉 저 심볼은 <code>Symbol("ident")</code>와 같다.
  
만약 같은 심볼 리터럴을 두 번 사용하면, 두 표현식 모두 완전히 동일한 Symbol 객체를 참조할 것이다.
  
#### 불리언 리터럴
  
5.3 문자열 인터폴레이션
--
  
<pre><code>val name="reader"
println(s"Hello, $name!")
</code></pre>
  
여기서 <code>s"Hello, $name!"</code> 은 문자열 리터럴로 처리된다. 이를 s 인터폴레이터라 부른다.
  
스칼라는 디폴트로 s 인터폴리터 외에 raw 와 f 라는 두 가지 인터폴레이터를 더 제공한다. raw 문자열 인터폴레이터는 s 처럼 작동하지만 문자열 이스케이프를 인식하지 못한다.
  
f 문자열 인터폴레이터를 사용하면 내장된 표현식에 대해 printf 스타일의 형식 지정을 사용할 수 있다.
  
5.4 연산자는 메소드다
--
  
<code>1 + 2</code> 는 <code>1.+(2)</code> 와 같다.
  
어떤 메소드라도 연산자 표기법을 사용할 수 있다. 즉 String 클래스의 인스턴스 s의 indexOf 메서드는 <code>s indexOf 'o'</code> 와 같이 쓸 수 있는 것이다. 
물론 <code>s.indexOf('o')</code> 도 가능하다.
  
5.5 산술 연산
--
  
5.6 관계 연산과 논리 연산
--
  
5.7 비트 연산
--
  
5.8 객체 동일성
--
  
두 객체가 같은지 비교하고 싶다면 <code>==</code> 를 사용할 수 있고, 같지 않은지를 비교하려면 <code>!=</code>를 사용한다.
  
<pre><code>scala> 1 == 2
res31: Boolean = false
</code></pre>
  
이런 연산은 실제로는 기본 타입의 객체 뿐ㄴ 아니라 모든 객체에 적용할 수 있다. 예를 들어, 리스트를 비교하는 데 <code>==</code>를 쓸 수 있다.
  
<pre><code>scala> List(1, 2, 3) == List(1, 2, 3)
res31: Boolean = true
</code></pre>
  
더 나아가서, 타입이 각기 다른 두 객체도 비교 가능하다.
  
<pre><code>scala> List(1, 2, 3) == "hello"
res31: Boolean = false
</code></pre>
  
객체를 <code>null</code>과 비교하거나, 역으로 <code>null</code>을 객체와 비교할 수 있다. 예외는 발생하지 않는다.
  
<pre><code>List(1, 2, 3) == null
res38: Boolean = false
</code></pre>
  
스칼라의 <code>==</code>는 먼저 좌항이 <code>null</code>인지 검사하고, 좌항이 <code>null</code>이 아니라면 해당 객체의 <code>equals</code> 메서드를 호출한다.
<code>equals</code> 메서드이기 때문에 실제 어떤 비교를 수행할지는 왼쪽 인자의 타입에 따라 달라진다. 자동으로 <code>null</code>을 체크하기 떄문에 직접 <code>null</code>을 검사할 필요가 없다.
  
이런 식이면 객체가 달라도 그 내용이 동일하다면 <code>true</code>를 리턴할 것이다. 바로 아래 처럼.
<pre><code>scala> ("he" + "llo") == "hello"
res40: Boolean = true
</code></pre>
  
5.9 연산자 우선순위와 결합 법칙
--
  
5.10 풍부한 래퍼
--
  
|          코드                |      결과              |
| --------------------------- |:---------------------:|
|     0 max 5                 |     5                 |
|     0 min 5                 |     0                 |
|     -2.7 abs                |     2.7               |
|     -2.7 round              |     -3L               |
|     1.5 isInfinity          |     false             |
|     (1.0 / 0) isInfinity    |     true              |
|     4 to 6                  |     Range(4, 5, 6)    |
|     "bob" capitalize        |     "Bob"             |
|     "robert" drop 2         |     "bert"            |
