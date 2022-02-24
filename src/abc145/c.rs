use proconio::input;

fn main() {
    input! {
      n: usize,
      x: [(i32, i32); n],
    }

    let mut d = 0.0;

    for i in 0..n - 1 {
        for j in i + 1..n {
            d += dist(x[i], x[j]);
        }
    }

    println!("{}", d * 2.0 / n as f64);
}

fn double(x: i32) -> i32 {
    x * x
}

fn dist(i: (i32, i32), j: (i32, i32)) -> f64 {
    (((double(i.0 - j.0)) + double(i.1 - j.1)) as f64).sqrt()
}
