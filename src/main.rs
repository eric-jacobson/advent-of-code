use std::fs;

fn main() {
    println!("********** Day 1 **********");
    day_1();

    println!("\n********** Day 2 **********");
    day_2();
}


fn day_2() {
    let input = fs::read_to_string("inputs/rock-paper-scissors").expect("failed to read the file");

    /*
      rock = A 1 paper = B 2 scissors = C 3
      x = lose y = draw z = win
    */
    let round_result = |result: Vec<&str>| match result[..] {
        ["A", "X"] => 3 + 0,    // ["A", "X"] => 1 + 3,
        ["A", "Y"] => 1 + 3,    // ["A", "Y"] => 2 + 6,
        ["A", "Z"] => 2 + 6,    // ["A", "Z"] => 3 + 0,
        ["B", "X"] => 1 + 0,    // ["B", "X"] => 1 + 0,
        ["B", "Y"] => 2 + 3,    // ["B", "Y"] => 2 + 3,
        ["B", "Z"] => 3 + 6,    // ["B", "Z"] => 3 + 6,
        ["C", "X"] => 2 + 0,    // ["C", "X"] => 1 + 6,
        ["C", "Y"] => 3 + 3,    // ["C", "Y"] => 2 + 0,
        ["C", "Z"] => 1 + 6,    // ["C", "Z"] => 3 + 3,
        _ => {
            panic!();
        }
    };

    let total_score = input
        .lines()
        .map(|line| line.split(" ").collect())
        .map(|line| round_result(line))
        .sum::<u32>();

    println!("Total score: {total_score}");
}

fn day_1() {
    let input = fs::read_to_string("inputs/calorie-counts").expect("failed to read the file");
    let mut totals = input
        .split("\n\n")
        .collect::<Vec<&str>>()
        .iter()
        .map(|lines| {
            lines
                .split("\n")
                .map(|line| line.parse::<u32>().unwrap())
                .sum::<u32>()
        })
        .collect::<Vec<u32>>();

    // part 1 - max individual calories
    let total = totals.iter().max().unwrap();
    println!("Max calories: {total}");

    // part 2 - total of top 3 calorie counts
    totals.sort();
    totals.reverse();
    let top_three: u32 = totals[0..3].iter().sum();
    println!("Total of top three = {top_three}");
}
