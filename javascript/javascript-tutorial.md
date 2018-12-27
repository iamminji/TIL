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
