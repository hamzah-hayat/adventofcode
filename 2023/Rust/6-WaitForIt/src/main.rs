use std::{fs, u32};
use regex::Regex;
use std::collections::HashMap;

fn main() {
    let file_path = "input.txt";

    let contents = fs::read_to_string(file_path)
        .expect("Should have been able to read the file");

    let result_1 = part1(contents.clone());
    print!("The answer to Part 1 is {}\n",result_1);

    let result_2 = part2(contents);
    print!("The answer to Part 2 is {}\n",result_2);
}

// For each Race, 0 is the race time while 1 is the race record
struct Race(usize, usize);

fn part1(input:String) -> u32 {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 1;

    // Read races
    let re: Regex = Regex::new(r"(\d+)").unwrap();
    let mut times:Vec<u32> = vec![];
    let mut distances:Vec<u32> = vec![];
    for (_, [number]) in re.captures_iter(split_input[0]).map(|c| c.extract()) {
        times.push(number.parse().unwrap());
    }
    for (_, [number]) in re.captures_iter(split_input[1]).map(|c| c.extract()) {
        distances.push(number.parse().unwrap());
    }

    let mut wins: HashMap<u32, u32> = HashMap::new();
    for i in 0..times.len() {
        // For each race we want to figure out the number of ways we can win
        // we can choose a speed value from 0 - (total time of race - 1)
        // Then we have (total time of race - speed) amount of time, so we do speed * time and see if its greater then distance
        // If it is, we won, so record that as a win
        for speed in 0..times[i] {
            let remaining_time = times[i] - speed;
            if speed*remaining_time>distances[i] {
                wins.entry(i.try_into().unwrap()).and_modify(|counter| *counter += 1).or_insert(1);
            }
        }
    }

    for win in wins  {
        result = win.1 * result;
    }

    return result;
}

fn part2(input:String) -> u32 {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    // Read races
    let re: Regex = Regex::new(r"(\d+)").unwrap();
    let mut times:Vec<&str> = vec![];
    let mut distances:Vec<&str> = vec![];
    for (_, [number]) in re.captures_iter(split_input[0]).map(|c| c.extract()) {
        times.push(number);
    }
    for (_, [number]) in re.captures_iter(split_input[1]).map(|c| c.extract()) {
        distances.push(number);
    }

    let mut race_time = "".to_string();
    for t in times {
        race_time.push_str(t);
    }
    let mut race_distance = "".to_string();
    for d in distances {
        race_distance.push_str(d);
    }

    let race_time_num = race_time.parse().unwrap();
    let race_distance_num = race_distance.parse().unwrap();

    // For each race we want to figure out the number of ways we can win
    // we can choose a speed value from 0 - (total time of race - 1)
    // Then we have (total time of race - speed) amount of time, so we do speed * time and see if its greater then distance
    // If it is, we won, so record that as a win
    for speed in 0..race_time_num {
        let remaining_time: usize = race_time_num - speed;
        if speed*remaining_time>race_distance_num {
            result += 1;
        }
    }

    return result;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_p1_basic() {
        let test = String::from("Time:      7  15   30
        Distance:  9  40  200");
        let answer = part1(test);
        assert_eq!(answer, 288);
    }

    #[test]
    fn test_p2_basic() {
        let test = String::from("Time:      7  15   30
        Distance:  9  40  200");
        let answer = part2(test);
        assert_eq!(answer, 71503);
    }
}