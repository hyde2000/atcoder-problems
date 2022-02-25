use proconio::input;

fn main() {
    input! {
      n: u8,
      s: String,
    }

    for c in s.chars() {
        print!(
            "{}",
            std::char::from_u32(((c as u8 - b'A' + n) % 26 + b'A') as u32).unwrap()
        );
    }
}
