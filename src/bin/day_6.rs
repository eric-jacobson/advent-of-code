use std::fs;

use itertools::Itertools;

fn main() {
    let input = fs::read_to_string("inputs/tuning-trouble").expect("failed to read the file");

    println!("Part one: {:?}", 4 + find_marker(4, &input));
    println!("Part two: {:?}", 14 + find_marker(14, &input));
}

fn find_marker(window_size: usize, input: &str) -> usize {
    input
        .as_bytes()
        .windows(window_size)
        .position(|window| window.iter().tuple_combinations().all(|(a, b)| a != b))
        .unwrap()
}
