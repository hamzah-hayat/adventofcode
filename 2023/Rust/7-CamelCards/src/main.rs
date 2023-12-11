use core::num;
use std::{fs, u32, cmp::Ordering};
use fancy_regex::Regex;

fn main() {
    let file_path = "input.txt";

    let contents = fs::read_to_string(file_path)
        .expect("Should have been able to read the file");

    let result_1 = part1(contents.clone());
    print!("The answer to Part 1 is {}\n",result_1);

    let result_2 = part2(contents);
    print!("The answer to Part 2 is {}\n",result_2);
}

#[derive(Debug, PartialEq, PartialOrd, Copy, Clone)]
enum HANDTYPE {

    FiveOfAKind,
    FourOfAKind,
    FullHouse,
    ThreeOfAKind,
    TwoPair,
    OnePair,
    HighCard

}

// For each Race, 0 is the race time while 1 is the race record
struct Hand(String, u32, u32, HANDTYPE, String);

fn part1(input:String) -> u32 {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    let mut hands =  vec![];
    // Read hands
    for line in split_input {
        let split:Vec<&str> = line.split(" ").collect();

        // Also sort hand
        let hand_string = split[0].to_owned();
        let mut chars: Vec<char> = hand_string.chars().collect();
        chars.sort_by(|a, b| cmp_card(a,b));
        let hand_string_sorted = String::from_iter(chars);

        let hand_bid:u32 = split[1].parse().unwrap();
        let hand_rank = 0;

        let hand_type = hand_type(&hand_string_sorted);

        hands.push(Hand(hand_string_sorted,hand_bid,hand_rank,hand_type,hand_string));
    }

    // Work out each hand
    let mut current_rank = 0;
    let mut new_hands: Vec<Hand> =  vec![];
    while hands.len() != 0  {
        // Figure out the lowest each time and assign a rank
        current_rank += 1;

        let mut lowest_hand: &Hand = &hands[0];
        for hand in &hands {
            if is_lower_hand(hand, lowest_hand) {
                lowest_hand = hand;
            }
        }

        new_hands.push(Hand(lowest_hand.0.to_owned(), lowest_hand.1, current_rank, HANDTYPE::HighCard,lowest_hand.4.to_owned()));

        for i in 0..hands.len() {
            if hands[i].0 == lowest_hand.0 {
                hands.remove(i);
                break;
            }
        }
    }

    for ranked_hand in new_hands {
        result += ranked_hand.1 * ranked_hand.2;
    }

    return result;
}

fn is_lower_hand(first:&Hand,second:&Hand) -> bool {

    if first.3 < second.3 {
        return false;
    }
    else if first.3 > second.3 {
        return  true;
    }
    else {
        // These hands are equal, need to check which is bigger
        let mut chars_first: Vec<char> = first.4.chars().collect();
        let mut chars_second: Vec<char> = second.4.chars().collect();

        for c in 0..5 {
            if card_value(chars_first[c]) < card_value(chars_second[c]) {
                return false
            }
            else if card_value(chars_first[c]) > card_value(chars_second[c]) {
                return true
            }
        }
        return false;
    }
}

fn card_value(card:char) -> usize {
    let cards = vec!['A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'];
    let val = cards.into_iter().position(|c| c == card).unwrap();

    return val;
}

fn cmp_card(first:&char,second:&char) -> Ordering {
    if first == second {
        return Ordering::Equal;
    }

    let cards = vec!['A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'];
    let val_first = cards.into_iter().position(|c| c == *first).unwrap();
    let cards2 = vec!['A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'];
    let val_second = cards2.into_iter().position(|c| c == *second).unwrap();

    if val_first < val_second {
        return Ordering::Less;
    }
    else {
        return Ordering::Greater;
    }
}

// Return the Hand Type
fn hand_type(hand:&String) -> HANDTYPE {

    // Check hands
    // Praise God for Regex!
    let re_five_of_a_kind: Regex = Regex::new(r"([AKQJT98765432])\1{4}").unwrap();
    if re_five_of_a_kind.is_match(&hand).unwrap() {
        return HANDTYPE::FiveOfAKind;
    }

    let re_four_of_a_kind: Regex = Regex::new(r"([AKQJT98765432])\1{3}").unwrap();
    if re_four_of_a_kind.is_match(&hand).unwrap() {
        return HANDTYPE::FourOfAKind;
    }

    // Note, this regex also works for five of a kind
    // But that shoudn't matter as we should already have checked that by now
    let re_full_house: Regex = Regex::new(r"([AKQJT98765432])\1{2}([AKQJT98765432])\2{1}").unwrap();
    let re_full_house_alt: Regex = Regex::new(r"([AKQJT98765432])\1{1}([AKQJT98765432])\2{2}").unwrap();
    if re_full_house.is_match(&hand).unwrap() || re_full_house_alt.is_match(&hand).unwrap() {
        return HANDTYPE::FullHouse;
    }

    // Note, this regex also works for a few others above, again shoudn't matter
    let re_three_of_a_kind: Regex = Regex::new(r"([AKQJT98765432])\1{2}").unwrap();
    if re_three_of_a_kind.is_match(&hand).unwrap() {
        return HANDTYPE::ThreeOfAKind;
    }

    let re_two_pair: Regex = Regex::new(r"([AKQJT98765432])\1{1}.*([AKQJT98765432])\2{1}").unwrap();
    if re_two_pair.is_match(&hand).unwrap() {
        return HANDTYPE::TwoPair;
    }

    let re_one_pair: Regex = Regex::new(r"([AKQJT98765432])\1{1}").unwrap();
    if re_one_pair.is_match(&hand).unwrap() {
        return HANDTYPE::OnePair;
    }

    return HANDTYPE::HighCard
}

// Part two makes Jokers into Wild cards
fn part2(input:String) -> u32 {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    let mut hands =  vec![];
    // Read hands
    for line in split_input {
        let split:Vec<&str> = line.split(" ").collect();

        // Also sort hand
        let hand_string = split[0].to_owned();
        let mut chars: Vec<char> = hand_string.chars().collect();
        chars.sort_by(|a, b| cmp_card_part_2(a,b));
        let hand_string_sorted = String::from_iter(chars);

        let hand_bid:u32 = split[1].parse().unwrap();
        let hand_rank = 0;

        let hand_type = hand_type_part_2(&hand_string_sorted);

        println!("Processed Hand {}",hand_string);

        hands.push(Hand(hand_string,hand_bid,hand_rank,hand_type,hand_string_sorted));
    }

    // Work out each hand
    let mut current_rank = 0;
    let mut new_hands: Vec<Hand> =  vec![];
    while hands.len() != 0  {
        // Figure out the lowest each time and assign a rank
        current_rank += 1;

        let mut lowest_hand: &Hand = &hands[0];
        for hand in &hands {
            if is_lower_hand_part_2(hand, lowest_hand) {
                lowest_hand = hand;
            }
        }

        new_hands.push(Hand(lowest_hand.0.to_owned(), lowest_hand.1, current_rank, lowest_hand.3,lowest_hand.4.to_owned()));

        for i in 0..hands.len() {
            if hands[i].0 == lowest_hand.0 {
                hands.remove(i);
                break;
            }
        }
    }

    for ranked_hand in new_hands {
        result += ranked_hand.1 * ranked_hand.2;
    }

    return result;
}

fn card_value_part_2(card:char) -> usize {
    let cards = vec!['A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2','J'];
    let val = cards.into_iter().position(|c| c == card).unwrap();

    return val;
}

fn cmp_card_part_2(first:&char,second:&char) -> Ordering {
    if first == second {
        return Ordering::Equal;
    }

    let cards = vec!['A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2','J'];
    let val_first = cards.into_iter().position(|c| c == *first).unwrap();
    let cards2 = vec!['A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2','J'];
    let val_second = cards2.into_iter().position(|c| c == *second).unwrap();

    if val_first < val_second {
        return Ordering::Less;
    }
    else {
        return Ordering::Greater;
    }
}

fn is_lower_hand_part_2(first:&Hand,second:&Hand) -> bool {

    if first.3 < second.3 {
        return false;
    }
    else if first.3 > second.3 {
        return  true;
    }
    else {
        // These hands are equal, need to check which is bigger
        let mut chars_first: Vec<char> = first.0.chars().collect();
        let mut chars_second: Vec<char> = second.0.chars().collect();

        for c in 0..5 {
            if card_value_part_2(chars_first[c]) < card_value_part_2(chars_second[c]) {
                return false
            }
            else if card_value_part_2(chars_first[c]) > card_value_part_2(chars_second[c]) {
                return true
            }
        }
        return false;
    }
}

fn hand_type_part_2(hand:&String) -> HANDTYPE {

    // We need to check if we have a Joker and then check each possible hand type we could have
    // Return the best one
    let do_we_have_joker: Regex = Regex::new(r"J").unwrap();
    if do_we_have_joker.is_match(&hand).unwrap() {
        // For each Joker we need to work out all possible hands
        // Then find the best possible hand type
        return hand_type_joker(hand);
    }
    else {
        return hand_type(hand);
    }
}

// We have a joker in this hand
// Work out every possible combination
fn hand_type_joker(hand:&String) -> HANDTYPE {

    if hand == "JJJJJ" || hand == "JJ8JJ" {
        return HANDTYPE::FiveOfAKind
    }

    let current_hand = hand_type(hand);

    // Get number of jokers
    let mut num_jokers = 0;
    for c in hand.chars() {
        if c =='J' {
            num_jokers += 1;
        }
    }

    if num_jokers == 0 {
        return current_hand;
    }

    // For each joker, we want to create a new Hand, replacing a joker with a value from
    let cards = vec!["A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2"];
    let mut best_hand = current_hand;
    for i in 0..cards.len() {
        let new_hand = hand.replacen("J", cards[i], 1);

        // Also need to resort string for hand_type
        let mut chars: Vec<char> = new_hand.chars().collect();
        chars.sort_by(|a, b| cmp_card(a,b));
        let hand_string_sorted = String::from_iter(chars);


        if hand_type_joker(&hand_string_sorted) < best_hand {
            best_hand = hand_type_joker(&hand_string_sorted);
        }
    }

    return best_hand;
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_p1_basic() {
        let test = String::from("32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483");
        let answer = part1(test);
        assert_eq!(answer, 6440);
    }

    #[test]
    fn test_p2_basic() {
        let test = String::from("32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483");
        let answer = part2(test);
        assert_eq!(answer, 5905);
    }

    #[test]
    fn test_p2_test_edge() {
        let test = String::from("KJQKK");
        let answer = hand_type_part_2(&test);
        assert_eq!(answer, HANDTYPE::FourOfAKind);
    }

}