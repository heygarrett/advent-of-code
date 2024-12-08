use std::fs;

fn main() -> Result<(), Box<dyn std::error::Error>> {
	let input = fs::read_to_string("src/bin/day4/input.txt")?;

	let answer = part1(input);
	println!("part 1: {}", answer);

	Ok(())
}

fn part1(input: String) -> usize {
	input
		.lines()
		.enumerate()
		.map(|(line_index, line)| {
			line.chars()
				.enumerate()
				.map(|(ch_index, _)| {
					let directions: [(isize, isize); 8] = [
						(0, -1),
						(1, -1),
						(1, 0),
						(1, 1),
						(0, 1),
						(-1, 1),
						(-1, 0),
						(-1, -1),
					];
					directions
						.iter()
						.filter(|(x, y)| {
							"XMAS"
								.chars()
								.enumerate()
								.map(|(depth, letter)| {
									let coordinates = (
										depth as isize * x + ch_index as isize,
										depth as isize * y + line_index as isize,
									);
									if !(0..line.len())
										.contains(&(coordinates.0 as usize))
										|| !(0..input.lines().count())
											.contains(&(coordinates.1 as usize))
									{
										return false;
									}
									Some(letter)
										== input
											.lines()
											.nth(coordinates.1 as usize)
											.and_then(|nth_line| {
												nth_line
													.chars()
													.nth(coordinates.0 as usize)
											})
								})
								.all(|b| b)
						})
						.count()
				})
				.sum::<usize>()
		})
		.sum()
}

#[cfg(test)]
mod tests {
	use super::*;

	fn get_example_input() -> String {
		fs::read_to_string("src/bin/day4/example.txt").expect("should be example input")
	}

	#[test]
	fn test_part1() {
		let example_input = get_example_input();
		assert_eq!(part1(example_input), 18);
	}
}
