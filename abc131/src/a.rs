use proconio;

pub fn run() {
    proconio::input! {
        s: proconio::marker::Chars,
    }
    for i in 0..3 {
        if s[i] == s[i + 1] {
            println!("Bad");
            return;
        }
    }
    println!("Good");
}
