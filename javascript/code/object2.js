

var name = "a";
function setNewName(newName){
    name = newName;
}


setNewName("b");
console.log(name);


// inheritance
function Animal(type) {
    this.type = type;
    this.show = function() {console.log("앗 깜짝이야")};
}

Animal.prototype.greeting = function() {
    console.log("으르렁");
};

function Dog(name, type) {
    Animal.call(this, name);
    this.type = type;
}

Dog.prototype = Object.create(Animal.prototype);
console.log(Dog.prototype);
Dog.prototype.constructor = Dog;
console.log(Dog.prototype);


var dog = new Dog();
dog.show();
dog.greeting();
