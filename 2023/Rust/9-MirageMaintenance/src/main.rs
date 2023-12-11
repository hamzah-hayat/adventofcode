use std::fs;
use regex::Regex;

fn main() {
    let file_path = "input.txt";

    let contents = fs::read_to_string(file_path)
        .expect("Should have been able to read the file");

    // To high 1823861280
    let result_1 = part1(contents.clone());
    print!("The answer to Part 1 is {}\n",result_1);

    let result_2 = part2(contents);
    print!("The answer to Part 2 is {}\n",result_2);
}

fn part1(input:String) -> usize {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    // Solve each line
    for line in split_input {
        // Get the numbers
        let mut sequence: Vec<isize> = vec![];

        // Read sequences
        let re: Regex = Regex::new(r"(-*\d+)").unwrap();
        for (_, [number]) in re.captures_iter(line).map(|c| c.extract()) {
            sequence.push(number.parse().unwrap());
        }

        // Solve using recursion
        let last = sequence[sequence.len()-1];
        let last_change = solve_sequence(sequence);
        
        result += last_change + last
    }

    return result.try_into().unwrap();
}

fn solve_sequence(sequence: Vec<isize>) -> isize {

    // Check if all of new sequence is zero
    // If so, return the last element back
    let mut all_zero = true;
    for num in 0..sequence.len() {
        if sequence[num]!=0 {
            all_zero = false;
        }
    }

    // Create our new sequence
    let mut new_sequence: Vec<isize> = vec![];
    for i in 0..sequence.len()-1 {
        new_sequence.push(sequence[i+1]-sequence[i]);
    }

    let last = new_sequence[new_sequence.len()-1];

    if all_zero {
        return new_sequence[new_sequence.len()-1]
    } else {
        let below = solve_sequence(new_sequence);
        return last + below
    }
}

fn part2(input:String) -> usize {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    // Solve each line
    for line in split_input {
        // Get the numbers
        let mut sequence: Vec<isize> = vec![];

        // Read sequences
        let re: Regex = Regex::new(r"(-*\d+)").unwrap();
        for (_, [number]) in re.captures_iter(line).map(|c| c.extract()) {
            sequence.push(number.parse().unwrap());
        }

        // Solve using recursion
        let first = sequence[0];

        // reverse sequence
        let mut reverse_sequence: Vec<isize> = vec![];
        let mut counter = 0;
        for num in 0..sequence.len() {
            reverse_sequence.push(sequence[sequence.len()-num-1]);
            counter += 1;
        }


        let last_change = solve_sequence(reverse_sequence);
        
        result += last_change + first
    }

    return result.try_into().unwrap();
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_p1_example() {
        let test = String::from("0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45");
        let answer = part1(test);
        assert_eq!(answer, 114);
    }

    #[test]
    fn test_p2_example() {
        let test = String::from("0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45");
        let answer = part2(test);
        assert_eq!(answer, 2);
    }

}