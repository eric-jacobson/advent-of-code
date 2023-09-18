use itertools::Itertools;
use std::fs;

fn main() {
    let input = fs::read_to_string("inputs/rucksack-reorganization").expect("error reading file");

    let part_one: u32 = input
        .lines()
        .map(|line| line.split_at(line.len() / 2))
        .map(|(front, back)| {
            [front, back].map(|s| {
                s.chars()
                    .map(|c| get_numeric_val_from_char(c))
                    .fold(0, |acc, num| acc | (1 << num))
            })
        })
        .map(|[front, back]| u64::trailing_zeros(front & back))
        .sum();

    println!("Part one: {:?}", part_one);

    let part_two: u32 = input
        .lines()
        .tuples()
        .filter_map(|(line_one, line_two, line_three)| {
            line_one
                .chars()
                .find(|c| line_two.contains(*c) && line_three.contains(*c))
        })
        .map(|c| get_numeric_val_from_char(c))
        .sum();

    println!("Part two: {:?}", part_two);
}

fn get_numeric_val_from_char(c: char) -> u32 {
    match c {
        'a'..='z' => 1 + c as u32 - 'a' as u32,
        'A'..='Z' => 27 + c as u32 - 'A' as u32,
        _ => unreachable!(),
    }
}
