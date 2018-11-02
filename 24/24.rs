fn permutations(ns: ~[int]) -> ~[~[int]] {
    if ns.len() <= 1 {
        return ~[ns];
    }
    let mut ps = ~[];
    for i in range(0, ns.len()) {
        let mut ns = ns.clone();
        let e = ns.remove(i).unwrap();
        for mut sub_p in permutations(ns).move_iter() {
            sub_p.unshift(e);
            ps.push(sub_p);
        }
    }
    ps
}

fn main() {
    let ps = permutations(~[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]);
    println!("{:?}", ps[1e6 as int - 1]);
}

#[test]
fn test_permutations() {
    assert_eq!(~[~[0, 1, 2], ~[0, 2, 1], ~[1, 0, 2], ~[1, 2, 0], ~[2, 0, 1], ~[2, 1, 0]],
               permutations(~[0, 1, 2]));
}
