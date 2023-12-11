use std::{fs, u32};
use regex::Regex;
use array2d::Array2D;
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

    let mut rows = vec![];

    for n in 0..split_input.len() {
        let input_chars: Vec<char> = split_input[n].chars().collect();

        let mut add_line = vec![];
        for char in 0..input_chars.len()  {
            add_line.push(input_chars[char]);
        }
        rows.push(add_line)
    }

    // Create grid first
    let grid = Array2D::from_rows(&rows).unwrap();

    for x in 0..grid.row_len() {
        for y in 0..grid.column_len() {
            print!("{}",grid.get(x, y).unwrap());
        }
        println!();
    }


    // Check each set of numbers using regex + grid
    let re = Regex::new(r"(\d+)").unwrap();

    for n in 0..split_input.len() {

        let matchs = re.find_iter(split_input[n]);

        // For each number, check the grid around it
        for m in matchs {
            let mut valid = false;

            for x in m.start()..m.end() {
                let y = n;
                if check_around_square(&grid, y, x) {
                    valid = true;
                }
            }
            if valid {
                println!("{},{}",m.start(), m.as_str());
                result += m.as_str().parse::<u32>().unwrap();
            }
        }

    }

    return result;
}

fn part2(input:String) -> u32 {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    let mut rows = vec![];

    for n in 0..split_input.len() {
        let input_chars: Vec<char> = split_input[n].chars().collect();

        let mut add_line = vec![];
        for char in 0..input_chars.len()  {
            add_line.push(input_chars[char]);
        }
        rows.push(add_line)
    }

    // Create grid first
    let grid = Array2D::from_rows(&rows).unwrap();

    // Check each set of numbers using regex + grid
    let re = Regex::new(r"(\d+)").unwrap();

    let mut possible_gears= HashMap::new();

    for n in 0..split_input.len() {

        let matchs = re.find_iter(split_input[n]);

        // For each number, check the grid around it
        for m in matchs {

            for x in m.start()..m.end() {
                let y = n;
                let signed_x:isize = x.try_into().unwrap();
                let signed_y:isize = y.try_into().unwrap();
            
                // Check each square around grid[(x,y)]
                if is_gear_char(grid.get(y+1, x)) {
                    print!("{}",grid.get(y+1, x).unwrap());
                    let str = format!("{:?}", (y+1, x));
                    let new_str = m.as_str().to_owned() + " ";
                    possible_gears.entry(str).or_insert(String::new()).push_str(&new_str);
                    break;
                }
                if is_gear_char(grid.get(y+1, x+1)) {
                    print!("{}",grid.get(y+1, x+1).unwrap());
                    let str = format!("{:?}", (y+1, x+1));
                    let new_str = m.as_str().to_owned() + " ";
                    possible_gears.entry(str).or_insert(String::new()).push_str(&new_str);
                    break;
                }
                if is_gear_char(grid.get(y, x+1)) {
                    print!("{}",grid.get(y, x+1).unwrap());
                    let str = format!("{:?}", (y, x+1));
                    let new_str = m.as_str().to_owned() + " ";
                    possible_gears.entry(str).or_insert(String::new()).push_str(&new_str);
                    break;
                }
                if signed_y-1>=0 {
                    if is_gear_char(grid.get(y-1, x+1)) {
                        print!("{}",grid.get(y-1, x+1).unwrap());
                        let str = format!("{:?}", (y-1, x+1));
                        let new_str = m.as_str().to_owned() + " ";
                        possible_gears.entry(str).or_insert(String::new()).push_str(&new_str);
                        break;
                    }
                    if is_gear_char(grid.get(y-1, x)) {
                        print!("{}",grid.get(y-1, x).unwrap());
                        let str = format!("{:?}", (y-1, x));
                        let new_str = m.as_str().to_owned() + " ";
                        possible_gears.entry(str).or_insert(String::new()).push_str(&new_str);
                        break;
                    }
                }
                if signed_x-1>=0 {
                    if is_gear_char(grid.get(y, x-1)) {
                        print!("{}",grid.get(y, x-1).unwrap());
                        let str = format!("{:?}", (y, x-1));
                        let new_str = m.as_str().to_owned() + " ";
                        possible_gears.entry(str).or_insert(String::new()).push_str(&new_str);
                        break;
                    }
                    if is_gear_char(grid.get(y+1, x-1)) {
                        print!("{}",grid.get(y+1, x-1).unwrap());
                        let str = format!("{:?}", (y+1, x-1));
                        let new_str = m.as_str().to_owned() + " ";
                        possible_gears.entry(str).or_insert(String::new()).push_str(&new_str);
                        break;
                    }
                }
                if signed_y-1>=0 && signed_x-1>=0 {
                    if is_gear_char(grid.get(y-1, x-1)) {
                        print!("{}",grid.get(y-1, x-1).unwrap());
                        let str = format!("{:?}", (y-1, x-1));
                        let new_str = m.as_str().to_owned() + " ";
                        possible_gears.entry(str).or_insert(String::new()).push_str(&new_str);
                        break;
                    }
                }
            }
        }

    }

    for (_, gears) in possible_gears.iter() {
        let gears_split:Vec<&str> = gears.split(" ").collect();
        if gears_split.len() == 3 {
            result += gears_split[0].parse::<u32>().unwrap()*gears_split[1].parse::<u32>().unwrap();
        }
    }

    return result;
}

fn check_around_square(grid:&Array2D<char>,x:usize,y:usize) -> bool {

    let signed_x:isize = x.try_into().unwrap();
    let signed_y:isize = y.try_into().unwrap();

    // Check each square around grid[(x,y)]
    if is_special_char(grid.get(x+1, y)) {
        print!("{}",grid.get(x+1, y).unwrap());
        return true;
    }
    if is_special_char(grid.get(x+1, y+1)) {
        print!("{}",grid.get(x+1, y+1).unwrap());
        return true;
    }
    if is_special_char(grid.get(x, y+1)) {
        print!("{}",grid.get(x, y+1).unwrap());
        return true;
    }

    if signed_x-1>=0 {
        if is_special_char(grid.get(x-1, y+1)) {
            print!("{}",grid.get(x-1, y+1).unwrap());
            return true;
        }
        if is_special_char(grid.get(x-1, y)) {
            print!("{}",grid.get(x-1, y).unwrap());
            return true;
        }
    }

    if signed_y-1>=0 {
        if is_special_char(grid.get(x, y-1)) {
            print!("{}",grid.get(x, y-1).unwrap());
            return true;
        }
        if is_special_char(grid.get(x+1, y-1)) {
            print!("{}",grid.get(x+1, y-1).unwrap());
            return true;
        }
    }

    if signed_x-1>=0 && signed_y-1>=0 {
        if is_special_char(grid.get(x-1, y-1)) {
            print!("{}",grid.get(x-1, y-1).unwrap());
            return true;
        }
    }

    return false;
}

fn is_special_char(c:Option<&char>) -> bool {
    let u = *c.unwrap_or(&'.');

    if u=='*'||u=='='||u=='+'||u=='$'||u=='/'||u=='&'||u=='#'||u=='%'||u=='-'||u=='@' {
        return true;
    }
    return false;
}

fn is_gear_char(c:Option<&char>) -> bool {
    let u = *c.unwrap_or(&'.');

    if u=='*' {
        return true;
    }
    return false;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_p1_basic() {
        let test = String::from("467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..");
        let answer = part1(test);
        assert_eq!(answer, 4361);
    }

    #[test]
    fn test_p2_basic() {
        let test = String::from("467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..");
        let answer = part2(test);
        assert_eq!(answer, 467835);
    }
}