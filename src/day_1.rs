use std::fs;

pub fn counting_calories() {
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
