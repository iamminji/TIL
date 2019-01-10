function add(x, y) {
    //console.log(description); // runtime error!
    console.log(add.description);
    return x + y;
}

add.description = "i am function";
add(1, 2);

var a = 1;

function f() {
    if (true) {
        var c = 2;
    }
    return c;
}

console.log(f());

var g1 = "전역 변수#1";

function f2() {
    g2 = "전역 변수#2";
}

console.log(g1);
// console.log(g2); // not defined!
f2();
console.log(g2);

function outer() {
    var a = 1;
    console.log(a);
    
    function inner() {
        a = 2;
        console.log(a);
    }
    
    inner();
    console.log(a);
}

outer();


var x = "global";

function f3() {
    console.log("#1", x);
    var x = "local";
    console.log("#2", x);
}

f3();

///////////////////////////////////////////
// Functions
///////////////////////////////////////////

var adder = new Function('a', 'b', 'return a+b');
console.log(adder(1,2));
console.log(adder.constructor === Function);