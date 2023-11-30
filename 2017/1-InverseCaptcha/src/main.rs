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
    let input_chars: Vec<char> = input.chars().collect();
    let mut result = 0;

    for n in 0..input.len()-1 {
        if input_chars[n] == input_chars[n+1] {
            result += input_chars[n].to_digit(10).unwrap();
        }
    }

    // Check final pair
    if input_chars[input.len()-1]==input_chars[0] {
        result += input_chars[0].to_digit(10).unwrap();
    }

    return result;
}

fn part2(input:String) -> u32 {
    let input_chars: Vec<char> = input.chars().collect();
    let mut result = 0;

    for n in 0..input.len()-1 {
        let second = (n + (input.len()/2)) % input.len();

        if input_chars[n] == input_chars[second] {
            result += input_chars[n].to_digit(10).unwrap();
        }
    }

    // Check final pair
    let check1 = input.len()-1;
    let check2 = (input.len()/2)-1;


    if input_chars[input.len()-1]==input_chars[(input.len()/2)-1] {
        result += input_chars[input.len()-1].to_digit(10).unwrap();
    }

    return result;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_p1_basic() {
        let test = String::from("1122");
        let answer = part1(test);
        assert_eq!(answer, 3);
    }

    #[test]
    fn test_p1_only_ones() {
        let test = String::from("1111");
        let answer = part1(test);
        assert_eq!(answer, 4);
    }

    #[test]
    fn test_p1_no_matches() {
        let test = String::from("1234");
        let answer = part1(test);
        assert_eq!(answer, 0);
    }

    #[test]
    fn test_p1_edges() {
        let test = String::from("91212129");
        let answer = part1(test);
        assert_eq!(answer, 9);
    }

    // Part 2
    #[test]
    fn test_p2_all_match() {
        let test = String::from("1212");
        let answer = part2(test);
        assert_eq!(answer, 6);
    }

    #[test]
    fn test_p2_no_match() {
        let test = String::from("1221");
        let answer = part2(test);
        assert_eq!(answer, 0);
    }

    #[test]
    fn test_p2_two_matchs() {
        let test = String::from("123425");
        let answer = part2(test);
        assert_eq!(answer, 4);
    }

    #[test]
    fn test_p2_lots_match() {
        let test = String::from("123123");
        let answer = part2(test);
        assert_eq!(answer, 12);
    }
    
    #[test]
    fn test_p2_big_example() {
        let test = String::from("12131415");
        let answer = part2(test);
        assert_eq!(answer, 4);
    }


}