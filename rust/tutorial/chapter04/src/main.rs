fn main() {
    _ex9()
}

fn _ex1() {
    let mut s = String::from("hello");
    s.push_str(", world!");
    println!("{}", s);
}

fn _ex2() {
    // rust 에서는 아래 코드가 에러가 난다.
    let s1 = String::from("hello");
    let s2 = s1;

    // println!("{}, world", s1);
    println!("{}, world", s2);

    // 러스트에서는 변수가 범위를 벗어나면 자동으로 drop 함수를 호출하여 메모리를 해제 함
    // 그런데 두 변수가 같은 메모리를 참조하고 있기 때문에 실행 시 에러가 난다
    // 이를 dobule free error 라고 한다.
}

fn _ex3() {
    let s1 = String::from("hello");
    let s2 = s1.clone();

    println!("s1 = {}, s2 = {}", s1, s2)
}

fn _ex4() {
    let s = String::from("hello");
    _takes_ownership(s);
    // 이후 s 를 사용하려고 하면 에러가 난다.
    // println!("wow, {}", s);
    let x = 5;
    _makes_copy(x);
}

fn _takes_ownership(some_string: String) {
    println!("{}", some_string);
}

fn _makes_copy(some_integer: i32) {
    println!("{}", some_integer);
}

fn _ex5() {
    // gives_ownership 함수의 리턴 값이 변수 s1으로 옮겨진다.
    let s1 = _gives_ownership();
    println!("{}", s1);
    // 변수 s2가 범위 내에 생성된다.
    let s2 = String::from("hello");
    println!("{}", s2);
    // 변수 s2가 takes_and_gives_back 함수로 옮겨간 후 리턴 값은 변수 s3로 옮겨진다.
    let s3 = _takes_and_gives_back(s2);
    println!("{}", s3);
}

fn _gives_ownership() -> String {
    let some_string = String::from("hello");
    some_string
}

fn _takes_and_gives_back(a_string: String) -> String {
    a_string
}

fn _ex6() {
    let s1 = String::from("hello");
    // & 기호를 사용하여 소유권을 가져오지 않고도 값을 참조할 수 있다.
    // 참조는 소유권을 갖지 않기 때문에 참조가 가리키는 값은 참조가 범위를 벗어나더라도 drop 함수가 호출되지 않는다.
    let len = _calculate_length(&s1);
    println!("'{}'의 길이는 {}입니다.", s1, len);

    // 참조에 mut 키워드를 추가하면 수정이 가능하다.
    let mut s2 = String::from("hello");
    _change(&mut s2);
    println!("{}", s2)
}

fn _calculate_length(s: &String) -> usize {
    s.len()
    // 이 값은 참조/대여했기 때문에 s를 수정할 수는 없다.
    // 기본적으로 불변이다.
}

fn _change(some_string: &mut String) {
    some_string.push_str(", world");
}

fn _ex7() {
    // let reference_to_nothing = _dangle();
}

// fn _dangle() -> &String {
// 이 참조는 리턴 후 해제되므로 에러가 발생한다.
// let s = String::from("hello");
// &s
// }

fn _ex8() {
    let mut s = String::from("hello world");
    let word = _first_word(&s);
    println!("{}", word);
    s.clear();
}

fn _first_word(s: &String) -> usize {
    let bytes = s.as_bytes();
    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return i;
        }
    }
    s.len()
}

fn _ex9() {
    let s = String::from("hello world");
    let hello = &s[0..5];
    let world = &s[6..11];
    println!("{} {}", hello, world)
}
