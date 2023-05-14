fn main() {
    ex4()
}

fn ex1() {
    let number = 3;

    if number < 5 {
        println!("조건이 일치합니다!");
    } else {
        println!("조건이 일치하지 않습니다.");
    }
}

fn ex2() {
    let condition = true;
    let number = if condition { 5 } else { 6 };

    println!("number의 값:{}", number);
}

fn ex3() {
    loop {
        println!("다시 실행!");
        break;
    }
}

fn ex4() {
    let mut counter = 0;
    let result = loop {
        counter += 1;
        if counter == 10 {
            // break 자체가 return 구문 역할? 같은걸 해서
            // 세미콜론이 있어도 되고 없어도 되고 상관 없음
            break counter * 2;
        }
    };

    println!("The result is {}", result);
}

fn ex5() {
    let mut number = 3;

    while number != 0 {
        println!("{}!", number);
        number = number - 1;
    }

    println!("발사!");
}

fn ex6() {
    let a = [10, 20, 30, 40, 50];
    let mut index = 0;

    while index < 5 {
        println!("요소의 값: {}", a[index]);
        index = index + 1;
    }
}

fn ex7() {
    let a = [10, 20, 30, 40, 50];

    for element in a.iter() {
        println!("요소의 값: {}", element);
    }
}

fn ex8() {
    for number in (1..4).rev() {
        println!("{}!", number);
    }
    println!("발사!");
}
