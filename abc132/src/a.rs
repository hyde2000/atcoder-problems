use proconio;

pub fn run() {
    proconio::input! {
      mut s: proconio::marker::Chars,
    }
    s.sort();
    if s[0] == s[1] && s[2] == s[3] && s[0] != s[2] {
        println!("Yes");
        return;
    }
    println!("No");
}
