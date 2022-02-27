use proconio::input;

fn main() {
    input! {
        n: usize,
        edges: [(usize, usize); n-1],
    }
    let mut g = vec![vec![]; n];
    for (i, (a, b)) in edges.iter().enumerate() {
        g[a - 1].push((i, b - 1));
        g[b - 1].push((i, a - 1));
    }

    let mut c = vec![0; n - 1];
    dfs(&g, &mut c, 0, 0, 0);
    println!("{}", c.iter().max().unwrap());
    for cc in c {
        println!("{}", cc);
    }
}

fn dfs(g: &Vec<Vec<(usize, usize)>>, c: &mut Vec<usize>, pc: usize, v: usize, vp: usize) {
    let mut nextc = 1;
    for next in &g[v] {
        if nextc == pc {
            nextc += 1;
        }
        if next.1 == vp {
            continue;
        }

        c[next.0] = nextc;
        dfs(g, c, nextc, next.1, v);
        nextc += 1;
    }
}
