use proconio::input;

fn main() {
    input! {
      s: String
    }

    let mut end = s.len() - 1;
    let mut count = 0;
    for i in 0..(s.len() - 1) {
        if i >= end {
            break;
        }
        if s.chars().nth(i).unwrap() != s.chars().nth(end).unwrap() {
            count += 1;
        }
        end -= 1;
    }
    println!("{}", count);
}
