use std::{fs, iter::zip};

fn main() -> Result<(), Box<dyn std::error::Error>> {
	let contents = fs::read_to_string("src/bin/day1/input.txt")?;
	let answer = calculate_distance(contents);
	println!("{}", answer);
	Ok(())
}

fn calculate_distance(input: String) -> u32 {
	let mut list1: Vec<u32> = vec![];
	let mut list2: Vec<u32> = vec![];
	for line in input.lines() {
		let mut nums = line.split_whitespace();
		list1.push(nums.next().expect("number").parse::<u32>().expect("number"));
		list2.push(nums.next().expect("number").parse::<u32>().expect("number"));
	}
	list1.sort();
	list2.sort();

	zip(list1, list2)
		.map(|tup| u32::abs_diff(tup.0, tup.1))
		.sum()
}

#[cfg(test)]
mod tests {
	use super::*;

	#[test]
	fn test_calculate_distance() {
		let contents =
			fs::read_to_string("src/bin/day1/example.txt").expect("expecting example.txt");
		assert_eq!(calculate_distance(contents), 11);
	}
}
