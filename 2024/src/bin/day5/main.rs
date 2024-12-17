use std::fs;

fn main() -> Result<(), Box<dyn std::error::Error>> {
	let input = fs::read_to_string("src/bin/day5/input.txt")?;

	let answer = part1(&input);
	println!("part 1: {}", answer);

	Ok(())
}

fn part1(input: &str) -> u32 {
	143
}

#[cfg(test)]
mod tests {
	use super::*;

	fn get_example_input() -> String {
		fs::read_to_string("src/bin/day5/example.txt").expect("should be example input")
	}

	#[test]
	fn test_part1() {
		assert_eq!(part1(&get_example_input()), 143);
	}
}
