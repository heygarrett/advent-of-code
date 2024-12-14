use std::fs;

fn main() -> Result<(), Box<dyn std::error::Error>> {
	let input = fs::read_to_string("src/bin/day4/input.txt")?;

	let answer = part1(&input);
	println!("part 1: {}", answer);

	let answer = part2(&input);
	println!("part 2: {}", answer);

	Ok(())
}

fn part1(input: &str) -> u32 {
	input
		.lines()
		.enumerate()
		.map(|(line_index, line)| {
			line.chars()
				.enumerate()
				.map(|(ch_index, _)| {
					let directions: [(i32, i32); 8] = [
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
							"XMAS".chars().enumerate().all(|(depth, letter)| {
								let coordinates = (
									ch_index as i32 + depth as i32 * x,
									line_index as i32 + depth as i32 * y,
								);

								if !((0..line.len()).contains(&(coordinates.0 as usize))
									&& (0..input.lines().count())
										.contains(&(coordinates.1 as usize)))
								{
									return false;
								}

								Some(letter)
									== input.lines().nth(coordinates.1 as usize).and_then(
										|nth_line| {
											nth_line.chars().nth(coordinates.0 as usize)
										},
									)
							})
						})
						.count() as u32
				})
				.sum::<u32>()
		})
		.sum()
}

fn part2(input: &str) -> u32 {
	input
		.lines()
		.enumerate()
		.map(|(line_index, line)| {
			line.chars()
				.enumerate()
				.filter(|(ch_index, ch)| {
					if *ch != 'A' {
						return false;
					}
					let directions: [(i32, i32); 4] =
						[(1, -1), (1, 1), (-1, 1), (-1, -1)];

					["MMSS", "SMMS", "SSMM", "MSSM"]
						.iter()
						.any(|&valid_orientation| {
							Some(valid_orientation.to_string())
								== directions
									.iter()
									.map(|(x, y)| {
										let coordinates =
											(*ch_index as i32 + x, line_index as i32 + y);
										input
											.lines()
											.nth(coordinates.1 as usize)
											.and_then(|line| {
												line.chars().nth(coordinates.0 as usize)
											})
									})
									.collect::<Option<String>>()
						})
				})
				.count() as u32
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
		assert_eq!(part1(&example_input), 18);
	}

	#[test]
	fn test_part2() {
		let example_input = get_example_input();
		assert_eq!(part2(&example_input), 9);
	}
}
