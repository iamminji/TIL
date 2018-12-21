// javascript map example

const numbers = [1, 2, 3, 4, 5];

// javascript
var processed = numbers.map(function (num) {
    return num * num;
});

console.log(processed);

// ES6
const result = numbers.map(num => num * num);
console.log(result);
