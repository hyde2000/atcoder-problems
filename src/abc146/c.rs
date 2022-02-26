use proconio::input;

fn main() {
    input! {
      a: u128,
      b: u128,
      x: u128,
    }

    let mut left = 0;
    let mut right = 1e9 as u128 + 1;

    while left + 1 != right {
        let median = (left + right) / 2;
        let price = a * median + b * digit(median);
        if price <= x {
            left = median;
        } else {
            right = median;
        }
    }
    println!("{}", left);
}

fn digit(x: u128) -> u128 {
    if x / 10 == 0 {
        return 1;
    } else {
        return digit(x / 10) + 1;
    }
}
