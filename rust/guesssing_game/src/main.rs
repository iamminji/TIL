use rand::Rng;
use std::cmp::Ordering;
use std::io;

fn main() {
    let secret_number = rand::thread_rng().gen_range(1, 101);

    println!("사용자가 맞혀야 할 숫자: {}", secret_number);

    loop {
        println!("정답이라고 생각하는 숫자를 입력하세요.");

        // 러스트에서 변수는 기본적으로 값을 변경할 수 없다.
        // mut 을 붙이면 mutable 변수이다.
        let mut guess = String::new();

        // 참조도 기본적으로 immutable 하기 때문에 &mut 키워드를 붙여주는 것이다.
        io::stdin()
            .read_line(&mut guess)
            .expect("입력한 값을 읽지 못했습니다.");

        // 32 비트 부호없는 정수라는 뜻
        // mutable 한 guess 를 위에서 선언했지만 러스트는 guess 변수가 보관하던 값을 새로운 변수의
        // 값으로 가려버린다.
        // 이 기능으로 변수의 타입 변환이 유연하다.
        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        match guess.cmp(&secret_number) {
            Ordering::Less => println!("입력한 숫자가 작습니다!"),
            Ordering::Greater => println!("입력한 숫자가 큽니다!"),
            Ordering::Equal => {
                println!("정답!");
                break;
            }
        }
        println!("입력한 값: {}", guess);
    }
}
