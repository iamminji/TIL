# 상속

## 객체 상속

먼저 자바스크립트에서 정의하고 있는 프로토타입 객체의 실체는 `Object` 객체라는 사실을 알고있어야 한다. 이는 자바스크립트의 상속 관계에서 최상위 부모가 
`Object` 임을 의미한다.

즉, 만약에 Animal 이라는 생성자를 정의한다고 보면

<pre><code>Animal.prototype = new Object(); // Animal의 프로토타입 객체는 Obejct 인스턴스
Animal.prototype.constructor = Animal; // 생성된 인스턴스의 생성자는 Animal
</code></pre>

이라는 과정을 거친다는 것이다.

만약에 

<pre><code>function Animal(type) {
   this.type = type;
}

Animal.prototype.greeting = function() {
    console.log("으르렁");
};

function Dog(name, type) {
   Animal.call(this, name);
   this.type = type;
}
</code></pre>

이런식으로 하면 Dog에서 Animal을 호출하는 `this` 는 Dog의 인스턴스이기 때문에 이 `this`를 Animal이 받아서 사용하면 
Animal의 멤버가 Dog의 멤버로 된다.

상속을 할 때는 위처럼 생성자를 만들고 그 생성자를 하위 상속받을 Function 에서 `call` 을 이용해 호출하면 된다. 전달할 매개변수가 없다면 그냥 `this` 만 하면 된다.

그러나 문제가 하나 있다. Dog 가 greeting 을 상속받지 못하는 것이다. (물론 생성자 내에서 정의하는 것이면 상속이 가능하다.) 이럴 경우엔 Dog 의 
prototype 객체가 Animal 의 prototype 객체가 아니여서 그런 것이다. 그러므로 새 객체를 생성해서 바꿔치기 해준다.

<pre><code>Dog.prototype = Object.create(Animal.prototype);
console.log(Dog.prototype); // Animal
Dog.prototype.constructor = Dog;
console.log(Dog.prototype); // Dog! 자기 자신을 가르키게 해주어야 한다!
</code></pre>

### this

생성자 내부에서 사용되는 `this` 는 최초 생성자를 사용해서 만들어지는 인스턴스를 가리킨다. 자바스크립트에서는 메서드 코드가 생성자 외부에 존재할 수 있다는 사실 때문에 
`this` 를 이해하기 힘들 때가 있다.

<pre><code>function dog(name) {
    this.name = name;
    this.setNewName = setNewName;
}

function cat(name) {
    this.name = name;
    this.setNewName = setNewName;
}

function setNewName(newName) {
    this.name = newName;
}

var d = new dog("댕댕");
d.setNewName("멍멍");

var c = new cat("야옹");
c.setNewName("나비");

console.log(d.name);
console.log(c.name);
</code></pre>

이처럼 메서드를 생성자 밖에 지정하고 `this` 로 지정하면 `this` 는 호출하는 객체가 된다.

만약 `setNewName` 메서드를 일반 함수처럼 사용한다면 `this`는 루트 객체가 된다. 

<pre><code>var name = "a";
function setNewName(newName) {
   this.name = newName;
}

setNewName("b");
console.log(name);
</code></pre>

위와 같이 일반 함수처럼 작성하면 name 은 루트의 객체의 멤버라 "b" 로 엎어쳐진다. 
_참고로 nodejs 에선 안됨 브라우저만 됨. nodejs의 scope 가 브라우저랑 달라서 그럼_

### __proto__

자바스크립는 인스턴스 내부에 그 생성자의 프로토 타입 객체를 참조하기 위해 `__proto__` 라는 이름의 속성을 추가해 인스턴스가 프로토타입 객체의 멤버를 상속한다.
이 `__proto__` 라는 속성은 비공개 속성이다.

앞서 이야기했듯이 프로토타입 멤버는 해당 생성자로 생성된 모든 인스턴스가 공유할 수 있는데 이것이 가능한 이유는 모든 인스턴스가 `__proto__` 속성을 이용해 
자신의 프로토타입 객체에 접근해서 멤버를 검색할 수 있기 때문이다.

아래와 같은 생성자를 작성했다고 보자.

<pre><code>function Person(name){
    this.name = name;
}
var person = new Person("홍길동");
person.valueOf() // 어떻게 될까?
</code></pre>

생성자에 없는 메서드 <code>valueOf()</code> 를 호출한다고 볼 때 에러가 날 것 같지만, 실제로는 그렇지 않다. (참고로 <code>valueOf()</code> 는 
<code>Object</code> 에 있는 메서드다.)

- 브라우저는 우선 `person` 객체가 `valueOf()` 메서드를 가지고 있는지 체크한다.
- 없으므로 `person` 의 프로토타입 객체(`Person()` 생성자의 프로토타입) 에 `valueOf()` 메서드가 있는지 체크한다.
- 여전히 없으므로 `Person()` 생성자의 프로토타입 객체의 프로토타입 객체(`Object()` 생성자의 프로토타입) 가 `valueOf()` 메서드를 가지고 있는지
 체크한다. 여기에 있으니까 호출이 끝난다!

이를 __프로토타입 체인__ 이라 한다.

>모든 객체의 최상위 부모는 Object다. 자바스크립트의 상속은 프로토타입 기반의 상속이다.

그러나 `Object()` 에 있는 멤버라고 모두 상속 받을 수 있는 것은 아니다. 정확히는 `Object.prototype.` 으로 시작하는, `prototype` 속성에 
정의되어 있는 멤버들만 가능하다. `prototype` 속성도 하나의 객체이며 __프로토타입 체인__ 을 통해 상속하고자 하는 속성과 메서드를 담아주는 버킷으로 
주로 사용하는 객체다.

## Function 상속

`Object` 처럼 `Function` 도 상속할 수 있다. 모든 `Object` 객체가 `Object` 의 프로토타입 멤버를 상속하듯이 모든 함수도 `Function` 의 프로토타입 멤버를 상속한다.

> 모든 함수는 Function의 프로토타입 멤버를 상속받는다.
