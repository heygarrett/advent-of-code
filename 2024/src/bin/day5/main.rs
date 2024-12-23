use std::{cmp::Ordering, fs};

fn main() -> Result<(), Box<dyn std::error::Error>> {
	let input = fs::read_to_string("src/bin/day5/input.txt")?;

	let answer = part1(&input);
	println!("part 1: {}", answer);

	let answer = part2(&input);
	println!("part 2: {}", answer);

	Ok(())
}

fn part1(input: &str) -> u32 {
	let (rules, updates) = parse_input(input);
	updates
		.iter()
		.filter(|&update| rules.iter().all(|rule| validate(rule, update)))
		.map(|update| {
			let length = update.len();
			update[length / 2]
		})
		.sum()
}

fn part2(input: &str) -> u32 {
	let (rules, mut updates) = parse_input(input);
	updates
		.iter_mut()
		.filter(|update| !rules.iter().all(|rule| validate(rule, update)))
		.map(|update| {
			sort(&rules, update);
			let length = update.len();
			update[length / 2]
		})
		.sum()
}

fn parse_input(input: &str) -> (Vec<(u32, u32)>, Vec<Vec<u32>>) {
	let (raw_rules, raw_updates) = input.split_once("\n\n").unwrap_or_default();
	let rules = parse_ordering_rules(raw_rules).unwrap_or_default();
	let updates = parse_page_updates(raw_updates).unwrap_or_default();

	(rules, updates)
}

fn parse_ordering_rules(raw_rules: &str) -> Option<Vec<(u32, u32)>> {
	raw_rules
		.lines()
		.map(|line| {
			line.split_once("|").and_then(|(x, y)| {
				x.parse().ok().and_then(|parsed_x| {
					y.parse().ok().map(|parsed_y| (parsed_x, parsed_y))
				})
			})
		})
		.collect()
}

fn parse_page_updates(raw_updates: &str) -> Option<Vec<Vec<u32>>> {
	raw_updates
		.lines()
		.map(|line| {
			line.split(",")
				.map(|raw_number| raw_number.parse().ok())
				.collect()
		})
		.collect()
}

fn validate(rule: &(u32, u32), update: &[u32]) -> bool {
	update
		.iter()
		.position(|&e| e == rule.0)
		.and_then(|first_index| {
			update
				.iter()
				.position(|&e| e == rule.1)
				.map(|second_index| first_index < second_index)
		})
		.unwrap_or(true)
}

fn sort(rules: &Vec<(u32, u32)>, update: &mut [u32]) {
	update.sort_by(|&a, &b| {
		for &rule in rules {
			if rule == (a, b) {
				return Ordering::Less;
			} else if rule == (b, a) {
				return Ordering::Greater;
			}
		}

		Ordering::Equal
	})
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

	#[test]
	fn test_part2() {
		assert_eq!(part2(&get_example_input()), 123);
	}
}
