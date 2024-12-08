use std::fs;

fn main() -> Result<(), Box<dyn std::error::Error>> {
	let input = fs::read_to_string("src/bin/day3/input.txt")?;

	let answer = part1(input);
	println!("part 1: {}", answer);

	Ok(())
}

fn part1(input: String) -> usize {
	input
		.split("mul")
		.filter(|&s| s.chars().next().is_some_and(|c| c == '('))
		.filter_map(|s| s[1..].splitn(2, ")").take(1).next().take())
		.filter_map(|s| s.split_once(",").take())
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

	fn get_example_input() -> String {
		fs::read_to_string("src/bin/day3/example.txt").expect("should be example input")
	}

	#[test]
	fn test_part1() {
		assert_eq!(part1(get_example_input()), 161);
	}
}
