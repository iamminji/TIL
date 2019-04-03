
function outer() {
    var x = 0;
    return function() {
        return ++x;
    }
}

var x = -1;
var f = outer();
console.log(f());
console.log(f());
console.log(f());
console.log(f());

