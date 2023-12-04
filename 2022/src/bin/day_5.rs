use std::{collections::VecDeque, fs};

use itertools::Itertools;

#[derive(Debug)]
struct Move {
    count: usize,
    from: usize,
    to: usize,
}

fn main() {
    let input = fs::read_to_string("inputs/supply-stacks").expect("failed to read the file");
    let data: (&str, &str) = input.split("\n\n").collect_tuple().unwrap();
    let moves: Vec<Move> = data
        .1
        .lines()
        .map(|line| line.split(" ").collect_tuple().unwrap())
        .map(|(_, count, _, from, _, to)| Move {
            count: count.parse::<usize>().unwrap(),
            from: from.parse::<usize>().unwrap(),
            to: to.parse::<usize>().unwrap(),
        })
        .collect();

    let mut stacks = build_stacks(data.0.lines().collect_vec());
    for item in &moves {
        for _ in 0..item.count {
            let item_to_move = stacks[item.from - 1].pop_front().unwrap();
            stacks[item.to - 1].push_front(item_to_move);
        }
    }
    let part_one = get_top_of_stacks(stacks);
    println!("Part 1: {:?}", part_one);

    let mut stacks = build_stacks(data.0.lines().collect_vec());
    for item in &moves {
        let mut items_to_move = VecDeque::new();
        for _ in 0..item.count {
            items_to_move.push_front(stacks[item.from - 1].pop_front().unwrap());
        }
        for c in items_to_move {
            stacks[item.to - 1].push_front(c);
        }
    }
    let part_two = get_top_of_stacks(stacks);
    println!("Part 2: {:?}", part_two);
}

fn get_top_of_stacks(stacks: Vec<VecDeque<char>>) -> String {
    stacks
        .iter()
        .map(|stack| stack.front().unwrap())
        .collect::<String>()
}

fn build_stacks(mut lines: Vec<&str>) -> Vec<VecDeque<char>> {
    let last = lines.pop().unwrap().replace(" ", "");

    let mut stacks: Vec<VecDeque<char>> = Vec::new();

    for _ in last.chars() {
        stacks.push(VecDeque::from(vec![]));
    }

    for line in &lines {
        let range = 1..line.len();
        let mut stack_index = 0;
        for c in range.step_by(4) {
            if line.chars().nth(c).unwrap().is_alphabetic() {
                stacks[stack_index].push_back(line.chars().nth(c).unwrap());
            }
            stack_index += 1;
        }
    }

    stacks
}
