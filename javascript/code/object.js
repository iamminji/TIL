// 객체 생성 방법

// #1 new 사용
var book = new Object();
book.subject = "javascript";
book.author = "Unknown";

console.log(book.subject + " " + book.author);

// #2 객체 리터럴
var book2 = {
    "subject": "javascript",
    "author": "Unknown"
};

console.log(book2.subject + " " + book2.author);

var book2_1 = {
    subject: "javascript",
    author: "Unknown"
};

console.log(book2_1.subject + " " + book2_1.author);

// #3 사용자 정의 객체
function Book3(subject, author) {
    this.subject = subject;
    this.author = author;
}

var book3 = new Book3();
book3.subject = "javascript";
book3.author = "Unknown";
console.log(book3.subject + " " + book3.author);


// 생성자에 멤버 추가
function obj() {
}

obj.wow = '히히';
console.log(obj.wow);


function Person() {

}

var p1 = new Person();
p1.name = "김땡땡";
p1.age = 1;
console.log(p1.name + " " + p1.age);
delete p1.name;
console.log(p1.name + " " + p1.age);

var p2 = new Person();
p2.name = "홍길동";
console.log(p2.name + " " + p2.age);

// prototype example
Person.prototype.getName = function () {
    return this.name;
};

Person.prototype.getInfo = function () {
    return "Info:" + this.name + ' ' + this.age
};

console.log(p2.getName(), p2.getInfo());

// 덮어쓰기
p1.getInfo = function () {
    return "test"
};
console.log(p1.getName(), p1.getInfo());

// 공개, 비공개 멤버 예제
// #1
function A() {
    // 내부 지역 변수
    var _localX = 7;
    this.getX = function () {
        return _localX;
    };
    this.setX = function (x) {
        if (x < 10) {
            _localX = x;
            return _localX;
        }
    }
}

var a = new A();
console.log(a.getX());
a.setX(1);
console.log(a.getX());

// #2 클로저
function outer() {
    var _x = 0;
    
    function _private01() {
        return ++_x;
    }
    
    function _private02() {
        return (_x += 2);
    }
    
    return {public01: _private01, public02: _private02}
}

var o1 = outer();
console.log(o1.public01());
console.log(o1.public02());
var o2 = outer();
console.log(o2.public01());
console.log(o2.public02());

//
function dog(name) {
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


var testName = "before";
function setNewName2(newName){
    this.testName = newName;
    console.log("inner", this.testName);
}

setNewName2("after");
console.log("outer", testName);

var d = new dog("댕댕");
d.setNewName("멍멍");

var c = new cat("야옹");
c.setNewName("나비");

console.log(d.name);
console.log(c.name);

//
function animal(name) {
    this.name = name;
}

animal.prototype.setNewName = function (newName) {
    this.name = newName;
};

animal.prototype.setNewName("강아지");

var a = new animal("멍멍이");
console.log(a.name);
a.setNewName("멍멍이2");
console.log(a.name);