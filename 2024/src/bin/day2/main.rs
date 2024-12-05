use std::fs;

fn main() -> Result<(), Box<dyn std::error::Error>> {
	let input = fs::read_to_string("src/bin/day2/input.txt")?;
	let parsed_input = parse_input(input);

	let answer = part1(&parsed_input);
	println!("part 1: {}", answer);

	let answer = part2(&parsed_input);
	println!("part 2: {}", answer);

	Ok(())
}

fn part1(lines: &[Vec<isize>]) -> usize {
	lines.iter().filter(|&line| get_safety(line)).count()
}

fn part2(lines: &[Vec<isize>]) -> usize {
	lines
		.iter()
		.filter(|&line| {
			if get_safety(line) {
				return true;
			}

			for index in 0..line.len() {
				let mut line_subset = line.to_vec();
				line_subset.remove(index);
				if get_safety(&line_subset) {
					return true;
				}
			}

			false
		})
		.count()
}

fn parse_input(input: String) -> Vec<Vec<isize>> {
	input
		.lines()
		.map(|line| {
			line.split_whitespace()
				.map(|num_str| num_str.parse().expect("should be a number"))
				.collect()
		})
		.collect()
}

fn get_safety(numbers: &[isize]) -> bool {
	let ascending = numbers.windows(2).all(|w| w[0] < w[1]);
	let descending = numbers.windows(2).all(|w| w[0] > w[1]);
	let valid_level_diff = numbers
		.windows(2)
		.all(|w| matches!((w[0] - w[1]).abs_diff(0), 1..=3));

	(ascending || descending) && valid_level_diff
}

#[cfg(test)]
mod tests {
	use super::*;

	fn get_example_input() -> String {
		fs::read_to_string("src/bin/day2/example.txt").expect("should be example input")
	}

	#[test]
	fn test_part1() {
		assert_eq!(part1(&parse_input(get_example_input())), 2);
	}

	#[test]
	fn test_part2() {
		assert_eq!(part2(&parse_input(get_example_input())), 4);
	}
}
