use std::{fs, iter::zip};

fn main() -> Result<(), Box<dyn std::error::Error>> {
	let contents = fs::read_to_string("src/bin/day1/input.txt")?;
	let answer = part1(contents);
	println!("{}", answer);
	Ok(())
}

fn part1(input: String) -> u32 {
	let mut list1: Vec<u32> = vec![];
	let mut list2: Vec<u32> = vec![];

	input.lines().for_each(|line| {
		let numbers: Vec<u32> = line
			.split_whitespace()
			.map(|num_str| num_str.parse().expect("should be a number"))
			.collect();

		list1.push(numbers[0]);
		list2.push(numbers[1]);
	});

	list1.sort();
	list2.sort();

	zip(list1, list2)
		.map(|pair| u32::abs_diff(pair.0, pair.1))
		.sum()
}

#[cfg(test)]
mod tests {
	use super::*;

	#[test]
	fn test_part1() {
		let contents =
			fs::read_to_string("src/bin/day1/example.txt").expect("expecting example.txt");
		assert_eq!(part1(contents), 11);
	}
}
