fn main() {
    let mut n: u64 = 1;
    loop {
        if divisible_by_all(n, 20) {
            break;
        }
        n+=1;
    }
    println!("{:u}", n);
}

fn divisible_by_all(n: u64, upto: u64) -> bool {
    if n < upto {
        return false;
    }
    for i in range(1, upto+1) {
        if n%i != 0 {
            return false;
        }
    }
    return true;
}
