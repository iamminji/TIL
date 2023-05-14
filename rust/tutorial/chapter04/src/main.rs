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

    println!("{}, world", s2);

    // 러스트에서는 변수가 범위를 벗어나면 자동으로 drop 함수를 호출하여 메모리를 해제 함
    // 그런데 두 변수가 같은 메모리를 참조하고 있기 때문에 실행 시 에러가 난다
    // 이를 dobule free error 라고 한다.
    // 변수 복사는, 얕은 복사처럼 보이지만 러스트는 첫 번째 변수 s1을 무효화 시키므로 복사가 아니라 move라고 한다.)

    // s1 을 사용하려고 하면 value borrowed here after move 라는 에러가 뜬다.
    // println!("{}, world", s1);

    // s1은 유효하지 않기 때문에
    // scope를 벗어나면 s2만 drop 된다.
}

fn _ex2_2() {
    // 정수형 타입은 스택에 저장되어서 move 할 필요가 없음
    // 이를 Copy 라고 한다. (스택 데이터 복사)
    let x = 5;
    let y = x;

    println!("{}", x);
    println!("{}", y);

    // Copy 트레이트를 사용하는 일반적인 경우는
    // integer, boolean, char 등이 있다.
    // Drop 트레이트가 적용되어 있으면 Copy 트레이트는 못한다고 보면 된다.
}

fn _ex3() {
    // 스택이 아니라 힙 데이터를 복사하려면 Clone을 쓴다.
    let s1 = String::from("hello");
    let s2 = s1.clone();

    // 힙 데이터가 복사되었으므로 s1, s2 둘 다 drop 가능하다.
    println!("s1 = {}, s2 = {}", s1, s2)
}

fn _ex4() {
    let s = String::from("hello");
    _takes_ownership(s);
    // s가 함수 내로 이동되었기 때문에 더 이상 유효하지 않다.
    // println!("Not worked, {}", s);
    let x = 5;
    _makes_copy(x);
    // x는 i32 (스택 복사 Copy 트레이트가 구현되어 있음) 이므로
    // 함수 내로 이동되었어도 유효하다.
    println!("It worked!, {}", x);
}

fn _takes_ownership(some_string: String) {
    println!("{}", some_string);
    // some_string 은 여기서 메서드 벗어날 때 drop 된다.
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
    // 여기서 s2을 사용하면 move 되었다고 해서 컴파일 에러가 난다.
    // 함수에 s2를 넘기고 그걸 다시 사용하고 싶은데, 이럴 때 마다 _takes_and_gives_back 같은걸 사용하는 건 너무 귀찮은 일이다.
    // 그래서 참조 reference 개념이 등장했다.
    println!("{}", s3);

    // 함수를 벗어나면 s1, s3(소유권을 받음) 는 drop되지만 s2는 _takes_and_gives_back로 이동되었으므로
    // s2는 drop 호출되지 않는다.
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
    // 가변 참조를 만드는 것이다.
    let mut s2 = String::from("hello");
    let r1 = &mut s2;
    // 가변 참조도 제약이있는데 같은 데이터 (여기서는 s2) 에 대한 가변 참조는 하나만 만들 수 있다.
    // let r2 = &mut s2;
    _change(r1);
    println!("{}", s2);
}

fn _calculate_length(s: &String) -> usize {
    // s를 borrow 한다 라고 표현한다.
    s.len()
    // 이 값은 참조/대여했기(소유권이 없음) 때문에 s를 수정할 수는 없다.
    // 그래서 s를 다시 주기 위해서 리턴할 필요가 없다. (그 전에는 튜플같은 걸 써서 원래 값과 필요한 값 / 여기서는 길이 등을 같이 리턴해야했음)
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
