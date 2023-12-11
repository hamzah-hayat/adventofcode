use std::{fs, u32};
use regex::Regex;

fn main() {
    let file_path = "input.txt";

    let contents = fs::read_to_string(file_path)
        .expect("Should have been able to read the file");

    let result_1 = part1(contents.clone());
    print!("The answer to Part 1 is {}\n",result_1);

    let result_2 = part2(contents);
    print!("The answer to Part 2 is {}\n",result_2);
}

fn part1(input:String) -> u32 {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    for n in 0..split_input.len() {

        // Process each game
        let game_split: Vec<&str> = split_input[n].split(":").collect();
        let mut possible = true;

        // Get game number
        let re = Regex::new(r"Game ([0-9]+)").unwrap();
        let mut game_result = vec![];
        for (_, [number]) in re.captures_iter(game_split[0]).map(|c| c.extract()) {
            game_result.push(number.parse().unwrap());
        }
        let game_number:u32 = game_result[0];

        // Now regex and check each bag number
        let re2 = Regex::new(r"(\d+) (red|green|blue)").unwrap();
        for (_, [number,color]) in re2.captures_iter(game_split[1]).map(|c| c.extract()) {
            let num:u32 = number.parse().unwrap();
            if color=="red" && num>12 {
                possible = false;
            }
            if color=="green" && num>13 {
                possible = false;
            }
            if color=="blue" && num>14 {
                possible = false;
            }
        }

        if possible {
            result += game_number;
        }

    }

    return result;
}

fn part2(input:String) -> u32 {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    for n in 0..split_input.len() {

        // Process each game
        let game_split: Vec<&str> = split_input[n].split(":").collect();
        let mut highest_red = 0;
        let mut highest_green = 0;
        let mut highest_blue = 0;

        // Get game number
        let re = Regex::new(r"Game ([0-9]+)").unwrap();
        let mut game_result = vec![];
        for (_, [number]) in re.captures_iter(game_split[0]).map(|c| c.extract()) {
            game_result.push(number.parse().unwrap());
        }
        let game_number:u32 = game_result[0];

        // Now regex and check each bag number
        let re2 = Regex::new(r"(\d+) (red|green|blue)").unwrap();
        for (_, [number,color]) in re2.captures_iter(game_split[1]).map(|c| c.extract()) {
            let num:u32 = number.parse().unwrap();
            if color=="red" && num > highest_red {
                highest_red = num;
            }
            if color=="green" && num > highest_green {
                highest_green = num;
            }
            if color=="blue" && num > highest_blue {
                highest_blue = num;
            }
        }

        result += highest_red*highest_green*highest_blue;

    }

    return result;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_p1_basic() {
        let test = String::from("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green");
        let answer = part1(test);
        assert_eq!(answer, 8);
    }

    #[test]
    fn test_p2_basic() {
        let test = String::from("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green");
        let answer = part2(test);
        assert_eq!(answer, 2286);
    }
}