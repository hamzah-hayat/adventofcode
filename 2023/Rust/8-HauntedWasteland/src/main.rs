extern crate num;

use std::{fs, collections::HashMap, sync::{mpsc, Arc, Mutex}, thread};
use regex::Regex;
use num::integer::lcm;

fn main() {
    let file_path = "input.txt";

    let contents = fs::read_to_string(file_path)
        .expect("Should have been able to read the file");

    let result_1 = part1(contents.clone());
    print!("The answer to Part 1 is {}\n",result_1);

    let result_2 = part2(contents);
    print!("The answer to Part 2 is {}\n",result_2);
}

// For each Node
// 0 is name of this Node
// 1 is name of left Node
// 2 is name of right Node
struct Node(String,String,String);

fn part1(input:String) -> usize {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    // Get the instructions
    let directions = split_input[0];
    let direction_chars:Vec<char> = directions.chars().collect();
    

    let mut nodes= HashMap::new();
    // Read in our nodes
    for line in 1..split_input.len() {

        // Read races
        let re: Regex = Regex::new(r"(\w+) = \((\w+), (\w+)\)").unwrap();
        for (_, [node,node_left,node_right]) in re.captures(split_input[line]).map(|c| c.extract()) {
            nodes.insert(node.to_string(),Node(node.to_string(),node_left.to_string(),node_right.to_string()));
        }

    }

    // Run the instructions until we reach ZZZ
    let mut current_node = nodes.get("AAA").unwrap();
    while current_node.0 != "ZZZ" {
        let index = result % direction_chars.len();

        if direction_chars[index] == 'L' {
            current_node = nodes.get(&current_node.1).unwrap();
        } else {
            current_node = nodes.get(&current_node.2).unwrap();
        }

        result += 1;
    }



    return result.try_into().unwrap();
}

fn part2(input:String) -> usize {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    // Get the instructions
    let directions = split_input[0];
    let direction_chars:Vec<char> = directions.chars().collect();
    

    let mut nodes= HashMap::new();
    // Read in our nodes
    for line in 1..split_input.len() {

        // Read races
        let re: Regex = Regex::new(r"(\w+) = \((\w+), (\w+)\)").unwrap();
        for (_, [node,node_left,node_right]) in re.captures(split_input[line]).map(|c| c.extract()) {
            nodes.insert(node.to_string(),Node(node.to_string(),node_left.to_string(),node_right.to_string()));
        }
    }

    // Run the instructions until we reach a Z Node
    let mut finish_nums = vec![];
    for n in &nodes {

        // Start from this node
        if n.0.ends_with("A") {
            let mut finish = 0;
            let mut current_node = nodes.get(n.0).unwrap();
            while !current_node.0.ends_with("Z") {
                let index = finish % direction_chars.len();
        
                if direction_chars[index] == 'L' {
                    current_node = nodes.get(&current_node.1).unwrap();
                } else {
                    current_node = nodes.get(&current_node.2).unwrap();
                }
        
                finish += 1;
            }
            finish_nums.push(finish);
        }
    }

    // Find LCM
    let mut least_common_multiple = 1;
    for num in 0..finish_nums.len() {
        least_common_multiple = lcm(finish_nums[num],least_common_multiple);
    }

    return least_common_multiple;
}

fn part2_concurrency_brute_froce_doesnt_work(input:String) -> u32 {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    // Get the instructions
    let directions = split_input[0];
    let direction_chars:Vec<char> = directions.chars().collect();
    

    let mut nodes= HashMap::new();
    // Read in our nodes
    for line in 1..split_input.len() {
        // Read races
        let re: Regex = Regex::new(r"(\w+) = \((\w+), (\w+)\)").unwrap();
        for (_, [node,node_left,node_right]) in re.captures(split_input[line]).map(|c| c.extract()) {
            nodes.insert(node.to_string(),Node(node.to_string(),node_left.to_string(),node_right.to_string()));
        }
    }

    // Use concurrent threads to check each Start (anything ending with A) and if they reach an End (anything ending with Z)
    // We start a number of threads equal to the number of starting locations, then each loop we check if they all return true or false.
    // If any return false, we keep going, if all return true, we are done


    // Channel
    let (tx_main, rx_main) = mpsc::channel();
    let mut tx_workers = vec![];
    let mut number_of_threads = 0;

    let mut answers: Vec<bool> = vec![];
    for n in nodes {
        if n.0.ends_with("A") {
            number_of_threads += 1;

            let (tx_worker, rx_worker) = mpsc::channel();
            tx_workers.push(tx_worker);
            let tx = tx_main.clone();

            let new_input = input.clone();

            answers.push(false);

            thread::spawn(move || {

                // BRUH just remake the node list in here, FUK IT
                let split_input: Vec<&str> = new_input.split("\n").collect();
                let mut nodes= HashMap::new();
                // Read in our nodes
                for line in 1..split_input.len() {
                    // Read races
                    let re: Regex = Regex::new(r"(\w+) = \((\w+), (\w+)\)").unwrap();
                    for (_, [node,node_left,node_right]) in re.captures(split_input[line]).map(|c| c.extract()) {
                        nodes.insert(node.to_string(),Node(node.to_string(),node_left.to_string(),node_right.to_string()));
                    }
                }
                
                let mut current_node = nodes.get(&n.0).unwrap();
                loop {
                    let direction = rx_worker.recv().unwrap();
            
                    if direction == true {
                        current_node = nodes.get(&current_node.1).unwrap();
                    } else {
                        current_node = nodes.get(&current_node.2).unwrap();
                    }

                    if current_node.0.ends_with("Z") {
                        tx.send(true).unwrap();
                    } else {
                        tx.send(false).unwrap();
                    }
                }
            });
        }
    }

    let mut not_finished = true;
    while not_finished {
        not_finished = false;

        let index = result % direction_chars.len();
        for w in &tx_workers {
            if direction_chars[index] == 'L' {
                w.send(true);
            } else {
                w.send(false);
            }
        }

        for i in 0..number_of_threads {
            answers[i] = rx_main.recv().unwrap();
        }

        for i in &answers {
            if !*i {
                not_finished = true;
            }
        }

        result += 1;
    }

    return result.try_into().unwrap();
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_p1_example() {
        let test = String::from("RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)");
        let answer = part1(test);
        assert_eq!(answer, 2);
    }

    #[test]
    fn test_p1_example_2() {
        let test = String::from("LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)");
        let answer = part1(test);
        assert_eq!(answer, 6);
    }

    #[test]
    fn test_p2_example() {
        let test = String::from("LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)");
        let answer = part2(test);
        assert_eq!(answer, 6);
    }

}