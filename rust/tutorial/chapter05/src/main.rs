fn main() {
    _ex5()
}

fn _ex1() {
    let mut user1 = User {
        email: String::from("sample@example.com"),
        username: String::from("anonymous"),
        active: true,
        sign_in_count: 1,
    };
    user1.username = String::from("unknown");

    let user2 = _build_user(
        String::from("helloworld@example.com"),
        String::from("hellowold"),
    );

    let user3 = User {
        email: String::from("user3@example.com"),
        username: String::from("user3"),
        // 나머지 필드에는 user1의 값을 사용하라는 문법
        ..user1
    };
}

fn _build_user(email: String, username: String) -> User {
    User {
        email,
        username,
        active: true,
        sign_in_count: 1,
    }
}

struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}

struct Color(i32, i32, i32);
struct Point(i32, i32, i32);

fn _ex2() {
    let black = Color(0, 0, 0);
    let origin = Point(0, 0, 0);
}

fn _ex3() {
    let rect1 = (30, 50);
    println!("사각형의 면적: {} 제곱 픽셀", area(rect1));
}

fn area(dimensions: (u32, u32)) -> u32 {
    dimensions.0 * dimensions.1
}

// 러스트에서는 derive 어노테이션을 이용해 사용자 정의 타입에 유용한 동작을 적용할 수 있다.
#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

fn _ex4() {
    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };
    println!("사각형의 면적: {} 제곱 픽셀", area_struct(&rect1));
    println!("rect1: {:?}", rect1);
}

fn area_struct(rectangle: &Rectangle) -> u32 {
    // 이 매개변수 rectangle은 Rectangle 구조체의 불변 인스턴스에 대한 borrow 다
    // 이 때는 구조체의 소유권을 가져오는 것보다는 잠깐 빌려 쓰는 편이 낫다.
    // 이렇게 하면 main 함수는 여전히 소유권을 가지고 rect1 변수를 계속 사용할 수 있다.
    rectangle.width * rectangle.height
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }
    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }
    fn square(size: u32) -> Rectangle {
        Rectangle {
            width: size,
            height: size,
        }
    }
}

fn _ex5() {
    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };
    println!("사각형의 면적: {} 제곱 픽셀", rect1.area());
}
