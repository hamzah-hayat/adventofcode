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

    // First grab seeds
    let re: Regex = Regex::new(r"(\d+)").unwrap();
    let mut seeds:Vec<usize> = vec![];
    for (_, [number]) in re.captures_iter(split_input[0]).map(|c| c.extract()) {
        seeds.push(number.parse().unwrap());
    }

    // Grab the lines for each map
    let mut map_line_nums:Vec<usize> = vec![];
    for n in 0..split_input.len() {
        if split_input[n].contains("map") {
            map_line_nums.push(n.try_into().unwrap());
        }
    }
    map_line_nums.push(split_input.len().try_into().unwrap());

    for map in 0..map_line_nums.len()-1 {
        println!("{}",split_input[map_line_nums[map]]);

        // Each seed can only be changed by one map per category
        let mut change_made:Vec<bool> = vec![];
        for s in 0..seeds.len() {
            change_made.push(false);
        }

        // Process each map
        for map_line in map_line_nums[map]+1..map_line_nums[map+1]-1 {
            let split_map: Vec<&str> = split_input[map_line].split(" ").collect();
            let destination: usize = split_map[0].parse().unwrap();
            let source: usize = split_map[1].parse().unwrap();
            let range: usize = split_map[2].parse().unwrap();
            println!("{:?}",split_map);

            for s in 0..seeds.len() {
                if !change_made[s] {
                    // If seed is between source - source + range, then we can do seed -> destination 
                    if seeds[s] >= source && seeds[s] <= source + range {
                        let offset = seeds[s] - source;
                        seeds[s] = destination + offset;
                        change_made[s] = true
                    }
                }
            }
        }
        println!("Seeds: {:?}",seeds)
    }

    result = *seeds.iter().min().unwrap();

    return result.try_into().unwrap();
}

struct Pair(usize, usize);

fn part2(input:String) -> u32 {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    // First grab seeds
    let re: Regex = Regex::new(r"(\d+)").unwrap();
    let mut raw_seed_ranges:Vec<usize> = vec![];
    for (_, [number]) in re.captures_iter(split_input[0]).map(|c| c.extract()) {
        raw_seed_ranges.push(number.parse().unwrap());
    }

    let mut pairs:Vec<Pair> = vec![];
    for s in (0..raw_seed_ranges.len()).step_by(2) {
        pairs.push(Pair(raw_seed_ranges[s],raw_seed_ranges[s]+raw_seed_ranges[s+1]));
    }

    // Grab the lines for each map
    let mut map_line_nums:Vec<usize> = vec![];
    for n in 0..split_input.len() {
        if split_input[n].contains("map") {
            map_line_nums.push(n.try_into().unwrap());
        }
    }
    map_line_nums.push(split_input.len().try_into().unwrap());

    // Start processing maps
    for map in 0..map_line_nums.len()-1 {
        println!("{}",split_input[map_line_nums[map]]);

        let mut new_pairs:Vec<Pair> = vec![];

        while pairs.len() != 0 {
            // With each Seed Pair, we check
            // If the entire range is covered by this map, if so, simply offset and continue
            // Otherwise, break relevent pair of Pair off and offset, while keeping the rest
            let mut change_made = false;
            let current_pair: Pair = pairs.pop().unwrap();
            // pairs[s].0 is start of range
            // pairs[s].1 is end of range

            // Process each map
            println!("Processing Pair:{},{}",current_pair.0,current_pair.1);
            for map_line in map_line_nums[map]+1..map_line_nums[map+1]-1 {
                let split_map: Vec<&str> = split_input[map_line].split(" ").collect();
                let destination: usize = split_map[0].parse().unwrap();
                let source: usize = split_map[1].parse().unwrap();
                let range: usize = split_map[2].parse().unwrap();
                println!("Map is Dest:{}, Source:{} ,Range:{}",destination,source,range);

                // Pair is entirely inside map
                if current_pair.0 >= source && current_pair.1 <= source + range {
                    let offset = current_pair.0 - source;
                    new_pairs.push(Pair(destination + offset,destination + offset + current_pair.1 - current_pair.0));
                    change_made = true;
                }

                // Start of Seed Pair is in map, but rest is outside
                if current_pair.0 >= source && current_pair.0 < source + range && current_pair.1 > source + range {
                    let offset = current_pair.0 - source;
                    new_pairs.push(Pair(destination + offset, destination + range));
                    pairs.push(Pair(source+range,current_pair.1));
                    change_made = true;
                }

                // Start of Seed Pair is not in map, but Seed Pair end is
                if current_pair.0 < source && current_pair.1 <= source + range && current_pair.1 > source {
                    let offset = current_pair.1 - source;
                    new_pairs.push(Pair(destination,destination+offset));
                    pairs.push(Pair(current_pair.0,source));
                    change_made = true;
                }
            }
            if !change_made {
                new_pairs.push(Pair(current_pair.0,current_pair.1));
            }
        }
        // for p in &new_pairs {
        //     println!("Pair:{},{}",p.0,p.1);
        // }
        // println!("______________");
        pairs = new_pairs;
    }

    let mut lowest = 100000000000;
    for pair in pairs {
        if pair.0 < lowest {
            lowest = pair.0;
        }
    }

    result = lowest;

    return result.try_into().unwrap();
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_p1_basic() {
        let test = String::from("seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
");
        let answer = part1(test);
        assert_eq!(answer, 35);
    }

    #[test]
    fn test_p2_basic() {
        let test = String::from("seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48
                
soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15
                
fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4
                
water-to-light map:
88 18 7
18 25 70
                
light-to-temperature map:
45 77 23
81 45 19
68 64 13
                
temperature-to-humidity map:
0 69 1
1 0 69
                
humidity-to-location map:
60 56 37
56 93 4
");
        let answer = part2(test);
        assert_eq!(answer, 46);
    }
}