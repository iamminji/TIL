/* 비교 연산자 */
var s = "hello";
var o = new Object("hello");
// true
console.log(s == o);

var o1 = new String("hello");
var o2 = new Object("hello");
//false
console.log(o1 == o2);

// false
console.log(o == o2);

var o3 = o2;
// true
console.log(o3 == o3);

var a = ["1", "2", "3"];
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

///////////////////////////////////////////////////////////////////
console.log("######################### 논리 연산자 ##########################");
/* 논리 연산자 */

console.log("OR 연산자");
// 왼쪽이 참이면 바로 반환, 거짓이면 오른쪽 반환
var res1 = {} || (1+2);
// {}
console.log(res1);
var res2 = (1+2) || (2+3);
// 3
console.log(res2);
var res3 = (1+2) || {};
// 3
console.log(res3);
var res4 = false || (1+2);
// 3
console.log(res4);
var res5 = (1+2) || false;
// 3
console.log(res5);
var res6 = true || (1+2);
// true
console.log(res6);
var res7 = (1+2) || true;
// 3
console.log(res7);

console.log("AND 연산자");
// 왼쪽이 거짓이면 왼쪽 반환, 참이면 오른쪽 반환
var res8 = {} && (1+2);
// 3
console.log(res8);
var res9 = (1+2) && (2+3);
// 5
console.log(res9);
var res10 = (1+2) && {};
// {}
console.log(res10);
var res11 = false && (1+2);
// false
console.log(res11);
var res12 = (1+2) && false;
// false
console.log(res12);
var res13 = true && (1+2);
// 3
console.log(res13);
var res14 = (1+2) && true;
// true
console.log(res14);