fn main() {
    // value_in_cents(Coin::Quarter(UsState::Alaska));
    _ex4();
}

enum IpAddrKind {
    V4,
    V6,
}

struct IpAddr {
    kind: IpAddrKind,
    address: String,
}

fn _ex1() {
    let home = IpAddr {
        kind: IpAddrKind::V4,
        address: String::from("127.0.0.1"),
    };

    let loopback = IpAddr {
        kind: IpAddrKind::V6,
        address: String::from("::1"),
    };
}

enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

struct QuitMessage;
struct MoveMessage {
    x: i32,
    y: i32,
}
struct WriteMessage(String);
struct ChangeColorMessage(i32, i32, i32);

impl Message {
    fn call(&self) {
        // 메서드 본분 작성
    }
}

fn _ex2() {
    let m = Message::Write(String::from("Hello"));
    m.call()
}

enum Coin {
    Penny,
    Nickle,
    Dime,
    Quarter(UsState),
}

fn value_in_cents(coin: Coin) -> u32 {
    match coin {
        Coin::Penny => {
            println!("페니!");
            1
        }
        Coin::Nickle => 5,
        Coin::Dime => 10,
        Coin::Quarter(state) => {
            println!("State quarter from {:?}!", state);
            25
        }
    }
}

#[derive(Debug)]
enum UsState {
    Alabama,
    Alaska,
}

fn plus_one(x: Option<i32>) -> Option<i32> {
    match x {
        None => None,
        Some(i) => Some(i + 1),
    }
}

fn _ex3() {
    let five = Some(5);
    let six = plus_one(five);
    let none = plus_one(None);
}

fn _ex4() {
    // match 구문 한개일 때는 if-let 을 사용할 수 있다.
    let some_u8_value = Some(0u8);
    match some_u8_value {
        Some(3) => println!("three!"),
        _ => (),
    }

    // 이렇게 사용
    // if let 문법은 주어진 값에 대해 하나의 패턴만 검사하고 나머지 값은 무시하는 match 표현식을 더 쉽게 사용하기 위한 syntax sygar 이다.
    if let Some(3) = some_u8_value {
        println!("three!");
    }

    let mut count = 0;
    match coin {
        Coin::Quarter(state) => println!("{:?}주의 25센트 동전!", state),
        _ => count += 1,
    }

    if let Coin::Quarter(state) = coin {
        println!("{:?}주의 25센트 동전!", state);
    } else {
        count += 1;
    }
}
