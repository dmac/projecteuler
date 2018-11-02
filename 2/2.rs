fn main() {
    let mut f1 = 0;
    let mut f2 = 1;
    let mut f3 = f1 + f2;
    let mut sum = 0;
    while f3 < 4000000 {
        if f3%2 == 0 {
            sum += f3;
        }
        f1 = f2;
        f2 = f3;
        f3 = f1 + f2;
    }
    println!("{}", sum);
}
