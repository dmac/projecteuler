fn main() {
    let mut max = 0;
    for i in range(100 as i64, 1000) {
        for j in range(100 as i64, 1000) {
            let p = i * j;
            if is_palindrome(p) && p > max {
                max = p;
            }
        }
    }
    println!("{}", max);
}

fn is_palindrome(n: i64) -> bool {
    let v = n.to_str().into_bytes();
    for i in range(0, v.len()) {
        if v[i] != v[v.len()-1-i] {
            return false;
        }
    }
    true
}
