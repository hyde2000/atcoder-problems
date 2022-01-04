use proconio;

pub fn run() {
    proconio::input! {
      n: i32,
      l: i32,
    }
    let mut min_abs = l + n - 1;
    let mut v = Vec::new();
    for idx in 1..n + 1 {
        let score = l + idx - 1;
        v.push(score);
        if score.abs() < min_abs.abs() {
            min_abs = score;
        }
    }
    let remove_idx = v.iter().position(|&r| r == min_abs).unwrap();
    v.remove(remove_idx);
    let sum: i32 = v.iter().sum();
    println!("{}", sum);
}
