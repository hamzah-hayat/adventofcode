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

fn part1(input:String) -> u32 {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    for n in 0..split_input.len() {
        // Current score
        let mut current_score = 0;

        // Process each game
        let game_split: Vec<&str> = split_input[n].split(":").collect();
        let numbers_split: Vec<&str> = game_split[1].split("|").collect();

        // Get Winning and Card bumbers
        let re: Regex = Regex::new(r"(\d+)").unwrap();
        let mut winning_numbers:Vec<u32> = vec![];
        for (_, [number]) in re.captures_iter(numbers_split[0]).map(|c| c.extract()) {
            winning_numbers.push(number.parse().unwrap());
        }

        let mut card_numbers:Vec<u32> = vec![];
        for (_, [number]) in re.captures_iter(numbers_split[1]).map(|c| c.extract()) {
            card_numbers.push(number.parse().unwrap());
        }


        // Check if each card number is in winning numbers
        for card_num in card_numbers {
            if winning_numbers.contains(&card_num) {
                if current_score == 0 {
                    current_score = 1;
                } else {
                    current_score = current_score * 2;
                }
            }
        }

        result += current_score;
    }

    return result;
}

fn part2(input:String) -> u32 {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    let mut games: HashMap<u32, u32> = HashMap::new();

    // We start with one "copy" of every game
    for g in 1..split_input.len()+1 {
        games.insert(g.try_into().unwrap(), 1);
    }

    for n in 0..split_input.len() {
        // Current score
        let mut current_score = 0;

        // Process each game
        let game_split: Vec<&str> = split_input[n].split(":").collect();
        let numbers_split: Vec<&str> = game_split[1].split("|").collect();

        // Get Game number
        let game_re: Regex = Regex::new(r"(\d+)").unwrap();
        let mut game_match:Vec<u32> = vec![];
        for (_, [number]) in game_re.captures_iter(game_split[0]).map(|c| c.extract()) {
            game_match.push(number.parse().unwrap());
        }
        let current_game = game_match[0];

        // Get Winning and Card bumbers
        let re: Regex = Regex::new(r"(\d+)").unwrap();
        let mut winning_numbers:Vec<u32> = vec![];
        for (_, [number]) in re.captures_iter(numbers_split[0]).map(|c| c.extract()) {
            winning_numbers.push(number.parse().unwrap());
        }

        let mut card_numbers:Vec<u32> = vec![];
        for (_, [number]) in re.captures_iter(numbers_split[1]).map(|c| c.extract()) {
            card_numbers.push(number.parse().unwrap());
        }


        // Check if each card number is in winning numbers
        for card_num in card_numbers {
            if winning_numbers.contains(&card_num) {
                current_score += 1
            }
        }

        // Now add copies to subsequent games based on our score (and number of copies)
        let copies: u32 = *games.get(&current_game).unwrap();
        for game in current_game+1..current_game+current_score+1 {
            let current_copies = games.get(&game).unwrap();
            games.insert(game, copies+current_copies);
        }
    }

    for g in games {
        result += g.1;
    }

    return result;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_p1_basic() {
        let test = String::from("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11");
        let answer = part1(test);
        assert_eq!(answer, 13);
    }

    #[test]
    fn test_p2_basic() {
        let test = String::from("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11");
        let answer = part2(test);
        assert_eq!(answer, 30);
    }
}