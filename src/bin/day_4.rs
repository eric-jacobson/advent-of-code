use itertools::Itertools;
use std::fs;

fn main() {
    let input = fs::read_to_string("inputs/camp-cleanup").expect("failed to read the file");

    let pairs: Vec<((&str, &str), (&str, &str))> = input
        .lines()
        .map(|line| line.split(",").collect_tuple().unwrap())
        .map(|(left, right)| {
            (
                left.split_once("-").unwrap(),
                right.split_once("-").unwrap(),
            )
        })
        .collect();

    // one range fully contains the other
    let part_one: usize = pairs
        .iter()
        .filter(|(left, right)| {
            (left.0.parse::<u16>().unwrap() <= right.0.parse::<u16>().unwrap()
                && left.1.parse::<u16>().unwrap() >= right.1.parse::<u16>().unwrap())
                | (left.0.parse::<u16>().unwrap() >= right.0.parse::<u16>().unwrap()
                    && left.1.parse::<u16>().unwrap() <= right.1.parse::<u16>().unwrap())
        })
        .count();

    // one range overlaps the other
    let part_two: usize = pairs
        .iter()
        .filter(|(left, right)| {
            left.0.parse::<u16>().unwrap() <= right.1.parse::<u16>().unwrap()
                && left.1.parse::<u16>().unwrap() >= right.0.parse::<u16>().unwrap()
        })
        .count();

    println!("Part one: {:?}", part_one);
    println!("Part two: {:?}", part_two);
}
