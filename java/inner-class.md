Nested Class (중첩 클래스)
==

자바에선 메서드나 변수 처럼 클래스 역시 다른 클래스의 멤버가 될 수 있다. 클래스 안에 다른 클래스를 작성하기만 하면 된다.
그리고 이를 __nested class__ 라고 부르고 해당 클래스를 감싸는 클래스를 __outer class__ 라 부른다.


표기는 아래와 같다.
<pre><code>class Outer_Demo {
    class Inner_Demo {
    }
}
</code></pre>

Nested Class 는 두 가지 타입으로 나누어진다.

1. non-static nested classes / Inner Classes
2. static nested classes


Non-static Nested Classes
--

Inner Class 는 자바에서 보안 매커니즘이다.
Inner Class 는 정의한 곳의 위치와 방법에 대해 아래 3가지로 부터 의존한다.

- Inner Class
- Method-local Inner Class
- Anonymous Inner Class

  
- Inner Class

Inner Class를 작성하는 방법은 단순하다. 그저 클래스 안에 클래스를 작성하기만 하면 되는 것이다.
다른 클래스와 다르게, Inner Class는 private으로 선언할 수 있고, private으로 선언하면 바깥의 다른 클래스에선 접근할 수 없다.

[소스 보기](https://github.com/iamminji/TIL/blob/master/java/code/InnerClass01.java)

해당 코드를 실행하면 아래와 같은 결과를 얻을 수 있다.

<code>This is an inner class</code>

Inner Class는 또한 클래스의 private member에 접근할 수 있다.

[소스 보기](https://github.com/iamminji/TIL/blob/master/java/code/InnerClass02.java)

해당 코드를 실행하면 아래와 같은 결과를 얻을 수 있다.

<code>This is the getnum method of the inner class
      175</code>

- Method-local Inner Class

자바에선 메서드 안에 클래스를 작성할 수 있다. 이는 지역변수 처럼 해당 메서드 안에서만 사용가능하다.

[소스 보기](https://github.com/iamminji/TIL/blob/master/java/code/InnerClass03.java)

해당 코드를 실행하면 아래와 같은 결과를 얻을 수 있다.

<code>This is method inner class: num => 23
</code>

- Anonymous Inner Class

Inner Class는 클래스 이름 없이 선언될 수 있는데 이를 Anonymous Inner Class라 한다.

[소스 보기](https://github.com/iamminji/TIL/blob/master/java/code/InnerClass04.java)

해당 코드를 실행하면 아래와 같은 결과를 얻을 수 있다.

<code>This is an example of anonymous class
</code>

- Anonymous Inner Class as Argument

[소스 보기](https://github.com/iamminji/TIL/blob/master/java/code/InnerClass05.java)

해당 코드를 실행하면 아래와 같은 결과를 얻을 수 있다.

<code>Hello, This is an example of anonymous inner class as an argument
      java8 hello, This is an example of anonymous inner class as an argument
</code>


    
Static Nested Classes
--

static inner class는 외부 클래스의 static member의 nested class 이다.
초기화 없이 접근할 수 있으며 이는 다른 static member와 같다.

[소스 보기](https://github.com/iamminji/TIL/blob/master/java/code/InnerClass06.java)

해당 코드를 실행하면 아래와 같은 결과를 얻을 수 있다.

<code>This is my nested class
</code>


#### Inner Class 의 장점

- 중첩 클래스는 private 을 포함한 outer class의 모든 멤버에 접근이 가능한 특별한 관계를 표현할 수 있다.

- 중첩 클래스는 가독성과 유지보수성에서 뛰어나다. 논리적으로 클래스 그리고 인터페이스가 한 곳에 있기 때문이다.

- 코드 작성량이 적다. (별개의 클래스 파일을 관리할 필요가 없다.)

참고 https://www.tutorialspoint.com/java/java_innerclasses.htm
