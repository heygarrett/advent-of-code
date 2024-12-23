use std::fs;

fn main() {
	let input = fs::read_to_string("src/bin/day06/input.txt")
		.ok()
		.unwrap_or_default();

	let answer = part1(&input);
	println!("part 1: {}", answer);
}

fn part1(input: &str) -> u32 {
	41
}

#[cfg(test)]
mod tests {
	use super::*;

	fn get_examlpe_input() -> String {
		fs::read_to_string("src/bin/day06/example.txt")
			.ok()
			.unwrap_or_default()
	}

	#[test]
	fn test_part1() {
		assert_eq!(part1(&get_examlpe_input()), 41);
	}
}
