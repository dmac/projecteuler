use std::num::sqrt;
use std::f64::floor;
use std::iter::range_inclusive;
use std::hashmap::HashSet;
use std::iter::AdditiveIterator;

// Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.

fn divisors(n: int) -> HashSet<int> {
    let mut ds = HashSet::new();
    for i in range_inclusive(1, floor(sqrt(n as f64)) as int) {
        if n % i == 0 {
            ds.insert(i);
            if n / i < n {
                ds.insert(n / i);
            }
        }
    }
    return ds;
}

fn is_abundant(n: int) -> bool {
    n < divisors(n).iter().map(|&i| i).sum()
}

fn get_abundant_numbers() -> HashSet<int> {
    let mut h = HashSet::new();
    for i in range_inclusive(0, 28123) {
        if is_abundant(i) {
            h.insert(i);
        }
    }
    h
}

fn has_abundant_pair(n: int, h: &HashSet<int>) -> bool {
    for &i in h.iter() {
        if h.contains(&(n-i)) {
            return true;
        }
    }
    return false;
}

fn main() {
    let h = get_abundant_numbers();
    let mut sum = 0;
    for i in range_inclusive(0, 28123) {
        if !has_abundant_pair(i, &h) {
            sum += i;
        }
    }
    println!("{:d}", sum);
}

#[test]
fn test_divisors() {
    let h4: HashSet<int> = [1, 2].iter().map(|&i| i).collect();
    assert_eq!(h4, divisors(4));
    let h12: HashSet<int> = [1, 2, 3, 4, 6].iter().map(|&i| i).collect();
    assert_eq!(h12, divisors(12));
    let h28: HashSet<int> = [1, 2, 4, 7, 14].iter().map(|&i| i).collect();
    assert_eq!(h28, divisors(28));
}

#[test]
fn test_is_abundant() {
    assert!(!is_abundant(2));
    assert!(is_abundant(12));
    assert!(!is_abundant(28));
}

#[test]
fn test_has_abundant_pair() {
    let h = get_abundant_numbers();
    assert!(has_abundant_pair(24, &h));
    assert!(!has_abundant_pair(25, &h));
}
