use std::f64::floor;
use std::hashmap::HashSet;
use std::iter::range_inclusive;
use std::num::sqrt;

fn main() {
    let factors = prime_factors(600851475143 as f64);
    println!("{:f}", *factors.iter().max().unwrap());
}

fn prime_factors(n: f64) -> ~[f64] {
    let mut factors: HashSet<f64> = HashSet::new();
    for i in range_inclusive(2 as f64, sqrt(n)) {
        if n%floor(i) == 0 as f64 {
            let sub_factors = prime_factors(i);
            for &subf in sub_factors.iter() { factors.insert(subf); }
        }
    }
    if factors.is_empty() {
        return ~[n];
    }
    factors.iter().map(|&f| f).collect()
}
