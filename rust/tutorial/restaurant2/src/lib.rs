#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }
}

mod front_of_house {
    pub mod hosting {
        pub fn add_to_waitlist() {}
        fn seat_at_table() {}
    }
}

pub fn eat_at_restaurant() {
    // 절대 경로
    crate::front_of_house::hosting::add_to_waitlist();

    // 상대 경로
    front_of_house::hosting::add_to_waitlist();

    let mut meal = back_of_house::Breakfast::summer("호밀빵");
    // fruit 은 비공개라 값 변경 불가능함
    meal.toast = String::from("밀빵");
    println!("{} 토스트로 주세요", meal.toast);

    // enum 을 pub 으로 공개하면 필드 접근 전부 가능하다.
    let order1 = back_of_house::Appetizer::Soup;
    let order2 = back_of_house::Appetizer::Salad;
}

fn serve_order() {}

mod back_of_house {

    pub struct Breakfast {
        pub toast: String,
        seasonal_fruit: String,
    }

    impl Breakfast {
        pub fn summer(toast: &str) -> Breakfast {
            Breakfast {
                toast: String::from(toast),
                seasonal_fruit: String::from("복숭아"),
            }
        }
    }

    pub enum Appetizer {
        Soup,
        Salad,
    }

    fn fix_incorrect_order() {
        cook_order();
        super::serve_order();
    }

    fn cook_order() {}
}

// use 키워드 사용
// 파일 시스템에서 심볼릭 링크를 생성하는 것과 유사하다.
use crate::front_of_house::hosting;
// 상대 경로로 지정하는 법
// use self::front_of_house::hosting;

// 다시 내보내기
// pub use crate::front_of_house::hosting;

pub fn eat_at_restaurant2() {
    hosting::add_to_waitlist();
}

use std::fmt::Result;
use std::io::Result as IoResult;
//
// fn function1() -> Result {
// }
//
// fn function2() -> IoResult<()> {
// }
