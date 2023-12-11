use std::fs;

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
        let input_chars: Vec<char> = split_input[n].chars().collect();
        let mut first_num='z';
        let mut second_num='z';

        for char in 0..input_chars.len()  {
            if input_chars[char].is_digit(10) && first_num == 'z' {
                first_num = input_chars[char];
            }

            // The first num could also be the last!
            if input_chars[char].is_digit(10) {
                second_num = input_chars[char];
            }
        }
        let mut digit_num_str = String::from("");
        digit_num_str.push(first_num);
        digit_num_str.push(second_num);
        let digit_num = digit_num_str.parse::<u32>().unwrap();
        result += digit_num;
    }

    return result;
}

fn part2(input:String) -> u32 {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    for n in 0..split_input.len() {

        let mut current_line = split_input[n];
        let current_line = current_line.replace("oneight", "18");
        let current_line = current_line.replace("twone", "21");
        let current_line = current_line.replace("threeight", "38");
        let current_line = current_line.replace("fiveight", "58");
        let current_line = current_line.replace("sevenine", "79");
        let current_line = current_line.replace("eightwo", "82");
        let current_line = current_line.replace("eighthree", "83");
        let current_line = current_line.replace("nineight", "98");

        let current_line = current_line.replace("one", "1");
        let current_line = current_line.replace("two", "2");
        let current_line = current_line.replace("three", "3");
        let current_line = current_line.replace("four", "4");
        let current_line = current_line.replace("five", "5");
        let current_line = current_line.replace("six", "6");
        let current_line = current_line.replace("seven", "7");
        let current_line = current_line.replace("eight", "8");
        let current_line = current_line.replace("nine", "9");

        let input_chars: Vec<char> = current_line.as_str().chars().collect();
        let mut first_num='z';
        let mut second_num='z';

        for char in 0..input_chars.len()  {
            if input_chars[char].is_digit(10) && first_num == 'z' {
                first_num = input_chars[char];
            }

            // The first num could also be the last!
            if input_chars[char].is_digit(10) {
                second_num = input_chars[char];
            }
        }
        let mut digit_num_str = String::from("");
        digit_num_str.push(first_num);
        digit_num_str.push(second_num);
        let digit_num = digit_num_str.parse::<u32>().unwrap();
        result += digit_num;
    }

    return result;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_p1_basic() {
        let test = String::from("1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet");
        let answer = part1(test);
        assert_eq!(answer, 142);
    }

    #[test]
    fn test_p2_basic() {
        let test = String::from("two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen");
        let answer = part2(test);
        assert_eq!(answer, 281);
    }
}