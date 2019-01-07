// Scope 다양한 예제 정리

// 파싱 단계에서 변수 x가 메모리에 undefined로 정의됨
// 런타임에 0이 할당됨
var x =0;

// 파싱 단계에서 변수 add가 정의되고 함수의 코드 블록에 대한 참조가 할당됨
// 파싱 단계에서는 함수 구현 코드가 실행되지는 않음
function add(a, b) {
    // 런타임에 지역 변수 c에 a+b의 값이 할당됨
    var c = a+b;
    return c
}

console.log(add(1,2));
// undefined 에러가 난다.
// console.log(c);

// var 가 없으면 함수 내 변수는 루트 객체의 변수 스코프에 들어간다.
function add2(a, b){
    d = a+b;
    return d
}

console.log(add2(1,2));
console.log(d);

// 실행 시 파싱과 런타임은 별개가 된다.
// 함수 square가 실행됨
console.log(square(4));
// 파싱할 때 정의된 변수를 런타임에 덮어쓴다. (??)
var sqauare = 0;
function square(x) {
    return x*x;
}
// 변수 square가 실행됨
console.log(sqauare);