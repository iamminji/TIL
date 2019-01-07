## Javascript Tutorial

자바스크립트로 작성된 코드는 **<script>...</script>** 구문을 HTML 안에 놓으면 된다. **<script>** 구문이 어디있든 실행되지만 보통은 **<head>** 태그에 있는 것을 추천한다.

<pre><code><script ...>
    Javascript Code
</script>
</code></pre>

#### 특징
- 자바스크립트는 공백, 탭, 그리고 뉴 라인(캐리지 리턴) 을 무시한다.
- 자바스크립트의 세미콜론은 옵션이다. 있어도 되고, 없어도 된다. 하지만 별개의 라인을 한 줄로 작성하고 싶을 땐 붙여야 한다. (Go나 Scala와 비슷)
- 자바스크립트는 Case Sensitivity 한 언어이기 때문에 Time 과 TIME 은 별개의 것으로 간주한다.
- 자바스크립트의 주석은 C-style 이다.

#### 자바스크립트의 위치
앞서 이야기 했듯이 자바스크립트는 HTML 구문 어디서든 실행 되지만 아래의 위치를 선호한다.
1. <head> ... </head>
2. <body> ... </body>
3. <body> ... </body> 와 <head> ... </head>
4. 외부 파일을 <head> ... </head> 에 첨부

#### 자바스크립트 Syntax

##### 변수
자바스크립트에서 변수를 선언할 때는 `var` 를 사용한다. 자바스크립트의 `var` 에는 기본적인 타입 외에도 사용자 정의 객체도 할당할 수 있다. 이것이 가능한 이유는 다른 언어와는 다른 타입 체계를 채용했기 때문이다.

자바스크립트를 흔히 __약한 타입의 언어 Weakly typed language__ 라고 한다.

##### 연산자

* 비교 연산자

비교 연산자에는 `==` 와 `===` 가 있다. `==` 연산자는 두 피연산자의 타입이 다른 경우는 타입을 일치 시키고 비교한다. 다음의 예제를 살펴 보자.
<pre><code>var b = (1 == "1");
console.log(b);
</code></pre>

해당 코드에서 `b`는 __true__ 를 갖게 된다. (자바스크립트 규칙에 따라) 숫자 1을 _문자열_ 로 변환한 다음 비교했기 때문이다. 그러므로 타입 까지 비교하고
 싶다면 `===` 을 쓰면 된다.

객체는 또 다르다. 

<pre><code>var s = "hello";
var o = new Object("hello");
console.log(s == o);
</code></pre>

위와 같은 결과에서 답은 `true` 가 된다. 문자열, 숫자, 불린값과 객체를 비교하면 객체의 값(valueOf()) 을 비교하기 때문이다. 
두 연산자 모두 객체라면 참조값을 비교하게 된다.

<pre><code>var o1 = new String("hello");
var o2 = new Object("hello");
console.log(o1 == o2);
</code></pre>

이러면 결과는 `false` 다.

즉, 피 연산자로 객체가 오게 될 경우 참조 값이 일치하면, 다시 말하자면 두 참조값이 가리키는 메모리 상의 객체의 위치가 같다면 `true` 가 반환된다.

아래는 여러 가지 헷갈리는 동등 연산자 예제들이다.
<pre><code>var a = ["1", "2", "3"];
var b = ["1", "2", "3"];
// false
console.log(a == b);

var b2 = [a[0], a[1], a[2]];

// false
console.log(a == b2);

// true
console.log(a[0] == b[0]);
console.log(a[0] === b[0]);
console.log(a[0] === b2[0]);
console.log(b[0] === b2[0]);

// false
console.log(["1", "2", "3"] == ["1", "2", "3"]);

var p = new Array(1, 2, 3);
var q = new Array(1, 2, 3);
// false
console.log(p == q);

// false
console.log(NaN == NaN);

// true
console.log(undefined == null);
console.log(null == null);
console.log(undefined == undefined);
</code></pre>

* 논리 연산자

자바스크립트에서는 `0`, `""`, `null`, `undefined`, `NaN` 과 같은 값을 제외하고는 모두 참을 반환한다. 이 논리는 자바스크립트에서는 
논리 연산의 양쪽에 객체가 피연산자로 나와서 평가될 수 있음을 의미한다. 즉, 객체가 피 연산자로 올 경우엔 반환 되는 값은 true, false 가 아니라 
그 피연산자의 최종값이 된다.

<pre><code>var a = (1+2) || {};</code></pre>

이와 같은 코드에서 a는 `true` 가 아니라 좌측 피연산자의 최종 값인 3이 된다. OR 연산자는 좌측 피 연산자가 거짓이면 우측 피연산자의 최종 값을 반환하고, 
좌측 피연산자가 참이면 그 값이 반환된다.

AND 연산자의 경우는 좌측 피연산자의 평가가 거짓이면 바로 좌측을 반환한다. 좌측 값이 참이면 우측 피연산자 값을 그대로 반환한다.

부정 연산자는 피연산자의 최종 값과는 관계가 없고, 흔히 아는 true/false boolean 값 만을 갖게 된다.

##### if ... else

##### switch

##### while loop

##### for loop

##### for in

##### function
자바스크립트에서는 함수를 사용할 때 **function** 키워드를 사용한다.

<pre><code><script type="text/javascript">
    function functionname(parameter-list) {
        statements
    }
</script>
</code></pre>

자바스크립트에서는 **return** 구문은 옵션이다.

##### void
**void** 는 말 그래도 아무것도 리턴하지 않는 값이다.

<pre><code>function getValue() {
    var a, b, c;
    a = void (b = 5, c = 7);
}
</code></pre>

여기서 a 는 undefined 가 된다.

#### 자바스크립트의 Objects
자바스크립트는 객체지향 프로그래밍 언어이다.

다음은 자바스크립트의 built-in 객체들이다.

- Number
- Boolean
- String
- Array
- Date
- Math
- Regexp
