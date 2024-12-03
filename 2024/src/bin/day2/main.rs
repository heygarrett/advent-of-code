use std::fs;

fn main() -> Result<(), Box<dyn std::error::Error>> {
	let input = fs::read_to_string("src/bin/day2/input.txt")?;

	let answer = part1(input);
	println!("part 1: {}", answer);

	Ok(())
}

fn part1(input: String) -> usize {
	input
		.lines()
		.filter(|&line| {
			let numbers: Vec<isize> = line
				.split_whitespace()
				.map(|num_str| num_str.parse().expect("should be a number"))
				.collect();

			let ascending = numbers.windows(2).all(|w| w[0] < w[1]);
			let descending = numbers.windows(2).all(|w| w[0] > w[1]);
			let valid_level_diff = numbers
				.windows(2)
				.all(|w| matches!((w[0] - w[1]).abs_diff(0), 1..=3));

			(ascending || descending) && valid_level_diff
		})
		.count()
}

#[cfg(test)]
mod tests {
	use super::*;

	fn get_example_input() -> String {
		fs::read_to_string("src/bin/day2/example.txt").expect("should be example input")
	}

	#[test]
	fn test_part1() {
		assert_eq!(part1(get_example_input()), 2);
	}
}
