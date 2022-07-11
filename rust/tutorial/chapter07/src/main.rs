use std::{cmp::Ordering, io};
// use std::io 와 std::io::Write; 를 사용할 때는 self 키워드 사용
// use std::io::{self, Write};

// 다 가져올 때는 글롭 연산자 사용
use chapter07::eat_at_restaurant;
use std::collections::*;

fn main() {
    eat_at_restaurant();
}
