use std::fs;

fn main() -> Result<(), Box<dyn std::error::Error>> {
	let input = fs::read_to_string("src/bin/day3/input.txt")?;

	let answer = part1(&input);
	println!("part 1: {}", answer);

	let answer = part2(&input);
	println!("part 2: {}", answer);

	Ok(())
}

fn part1(input: &str) -> usize {
	parse_and_calculate(input)
}

fn part2(input: &str) -> usize {
	let mut sum = 0;
	for split in input.split("do").filter(|&s| &s[..5] != "n't()") {
		sum += parse_and_calculate(split);
	}

	sum
}

fn parse_and_calculate(input: &str) -> usize {
	input
		.split("mul")
		.filter(|&s| &s[..1] == "(")
		.filter_map(|s| s[1..].splitn(2, ")").take(1).next())
		.filter_map(|s| s.split_once(","))
		.filter_map(|(x, y)| {
			x.parse::<usize>().ok().and_then(|parsed_x| {
				y.parse::<usize>().ok().map(|parsed_y| parsed_x * parsed_y)
			})
		})
		.sum()
}

#[cfg(test)]
mod tests {
	use super::*;

	#[test]
	fn test_part1() {
		let example_input = fs::read_to_string("src/bin/day3/example.txt")
			.expect("should be example input");
		assert_eq!(part1(&example_input), 161);
	}

	#[test]
	fn test_part2() {
		let example_input = fs::read_to_string("src/bin/day3/example2.txt")
			.expect("should be example input");
		assert_eq!(part2(&example_input), 48);
	}
}
