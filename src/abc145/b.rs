use std::process::exit;

use proconio::input;

fn main() {
    input! {
      a: u16,
      b: String
    }
    if a % 2 != 0 {
        println!("No");
        exit(0);
    }
    let mut left = String::new();
    let mut right = String::new();
    for i in b.chars().enumerate() {
        if i.0 <= (a / 2 - 1).into() {
            left.push(i.1);
        } else {
            right.push(i.1);
        }
    }
    if left == right {
        println!("Yes");
    } else {
        println!("No");
    }
}
