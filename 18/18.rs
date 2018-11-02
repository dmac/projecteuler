fn sum_chain(input: &[~[int]], chain: &[int]) -> int {
    let mut sum = 0;
    for i in range(0, chain.len()) {
        sum += input[i][chain[i]];
    }
    sum
}

fn step_chain(input: &[~[int]], chain: &[int]) -> Option<~[~[int]]> {
    if chain.len() == input.len() {
        return None;
    }
    let last_index = *chain.last().unwrap();
    let mut c1 = chain.clone().into_owned();
    let mut c2 = chain.clone().into_owned();
    c1.push(last_index);
    c2.push(last_index+1);
    Some(~[c1, c2])
}

fn main() {
    let input =
        ~[~[75],
          ~[95, 64],
          ~[17, 47, 82],
          ~[18, 35, 87, 10],
          ~[20, 04, 82, 47, 65],
          ~[19, 01, 23, 75, 03, 34],
          ~[88, 02, 77, 73, 07, 63, 67],
          ~[99, 65, 04, 28, 06, 16, 70, 92],
          ~[41, 41, 26, 56, 83, 40, 80, 70, 33],
          ~[41, 48, 72, 33, 47, 32, 37, 16, 94, 29],
          ~[53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14],
          ~[70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57],
          ~[91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48],
          ~[63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31],
          ~[04, 62, 98, 27, 23, 09, 70, 98, 73, 93, 38, 53, 60, 04, 23]];

    let mut finished_chains = ~[];
    let mut chains = ~[~[0]];
    loop {
        match chains.pop() {
            None => break,
            Some(chain) => {
                match step_chain(input, chain) {
                    None => finished_chains.push(chain),
                    Some(next_chains) => {
                        for next_chain in next_chains.move_iter() {
                            chains.unshift(next_chain);
                        }
                    }
                }
            }
        }
    }

    let mut max_sum = 0;
    for chain in finished_chains.move_iter() {
        let sum = sum_chain(input, chain);
        if sum > max_sum {
            max_sum = sum;
        }
    }

    println!("{:d}", max_sum);
}

#[test]
fn test_sum_chain() {
    let input = ~[~[1], ~[2, 3]];
    assert_eq!(4, sum_chain(input, [0, 1]));
}

#[test]
fn test_step_chain() {
    let input = ~[~[1], ~[2, 3], ~[4, 5, 6]];
    assert_eq!(Some(~[~[0, 1, 1], ~[0, 1, 2]]), step_chain(input, [0, 1]));
    assert_eq!(None, step_chain(input, [0, 1, 2]));
}
