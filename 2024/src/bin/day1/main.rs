use std::{fs, iter::zip};

fn main() -> Result<(), Box<dyn std::error::Error>> {
	let contents = fs::read_to_string("src/bin/day1/input.txt")?;

	let answer = part1(&contents);
	println!("part 1: {}", answer);

	let answer = part2(&contents);
	println!("part 2: {}", answer);

	Ok(())
}

fn part1(input: &str) -> u32 {
	let (mut list1, mut list2) = generate_lists(input);

	list1.sort();
	list2.sort();

	zip(list1, list2)
		.map(|pair| u32::abs_diff(pair.0, pair.1))
		.sum()
}

fn part2(input: &str) -> u32 {
	let (list1, list2) = generate_lists(input);

	list1
		.iter()
		.map(|num1| num1 * list2.iter().filter(|&num2| num1 == num2).count() as u32)
		.sum()
}

fn generate_lists(input: &str) -> (Vec<u32>, Vec<u32>) {
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

	(list1, list2)
}

#[cfg(test)]
mod tests {
	use super::*;

	#[test]
	fn test_part1() {
		let contents =
			fs::read_to_string("src/bin/day1/example.txt").expect("expecting example.txt");
		assert_eq!(part1(&contents), 11);
	}

	#[test]
	fn test_part2() {
		let contents =
			fs::read_to_string("src/bin/day1/example.txt").expect("expecting example.txt");
		assert_eq!(part2(&contents), 31);
	}
}
