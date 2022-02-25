use proconio::input;

const MOD: u64 = 1_000_000_007;

fn main() {
    input! {
      x: u64,
      y: u64,
    }

    let res = if (x + y) % 3 != 0 {
        0
    } else if !(y >= x / 2 && x >= y / 2) {
        0
    } else {
        let mut n1 = 0;
        let mut n2 = 0;
        for i in 0..=x {
            n1 = i;
            n2 = (x - n1) / 2;
            let cx = n1 * 1 + n2 * 2;
            let cy = n2 * 1 + n1 * 2;
            if x == cx && y == cy {
                break;
            }
        }
        comb(n1 + n2, n1.min(n2)) % MOD
    };
    println!("{}", res);
}

fn comb(n: u64, r: u64) -> u64 {
    let mut nu = 1;
    let mut de = 1;
    for i in 0..r {
        nu = nu * (n - i) % MOD;
        de = de * (i + 1) % MOD;
    }
    (nu * pow(de, MOD - 2)) % MOD
}

fn pow(mut x: u64, mut n: u64) -> u64 {
    let mut ret: u64 = 1;
    while n > 0 {
        if (n & 1) == 1 {
            ret = ret * x % MOD;
        }
        x = x * x % MOD;
        n >>= 1;
    }
    ret
}
