use std::fs;
use array2d::Array2D;
use std::collections::VecDeque;

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
// 0 is char
// 1 is x position
// 2 is y position
#[derive(Copy, Clone)]
struct Node(char,usize,usize);

fn part1(input:String) -> usize {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    // First expand galaxy
    let galaxy: Vec<String> = expand_galaxy(split_input);

    // Now find shortest path between each pair of galaxys.
    let mut rows = vec![];

    for n in 0..galaxy.len() {
        let input_chars: Vec<char> = galaxy[n].chars().collect();

        let mut add_line = vec![];
        for char in 0..input_chars.len()  {
            add_line.push(input_chars[char]);
        }
        rows.push(add_line)
    }

    let galaxy_grid = Array2D::from_rows(&rows).unwrap();

    // Find each galaxy
    let mut galaxy_spaces: Vec<Node> = vec![];

    for x in 0..galaxy_grid.column_len() {
        for y in 0..galaxy_grid.row_len() {
            if galaxy_grid.get(x, y).unwrap() == &'#' {
                galaxy_spaces.push(Node('#', x, y))
            }
        }
    }

    while galaxy_spaces.len() != 0 {
        let check_galaxy = galaxy_spaces.pop().unwrap();

        for galaxy in &galaxy_spaces {
            // Find shortest route
            let mut visited = vec![];
            visited.push(check_galaxy);
            let queue: VecDeque<Node> = VecDeque::from([check_galaxy]);
            let path = bfs(&galaxy_grid,check_galaxy,galaxy,visited,queue);

            result += path.len();
        }
    }

    return result.try_into().unwrap();
}

// BFS to find shortest route from A to B
fn bfs(grid:&Array2D<char>,current_node:Node,target_node:&Node,mut visited:Vec<Node>,mut queue:VecDeque<Node>) -> Vec<Node> {

    while queue.len() != 0 {
        let next_node = queue.pop_front().unwrap();

        if next_node.1 == target_node.1 && next_node.2 == target_node.2 {
            return visited;
        }

        // Neighbours
        let mut neighbours = get_neighbours(grid,&visited, next_node);

        for n in neighbours {
            queue.push_front(n);
            visited.push(n);
        }
    }

    return visited;
}

fn get_neighbours(grid: &Array2D<char>, visited: &[Node], current_node: Node) -> Vec<Node> {

    let mut neighbours = vec![];

    // Four possible directions
    // Up
    if current_node.1 > 0 {
        neighbours.push(Node(*grid.get(current_node.1-1, current_node.2).unwrap(), current_node.1-1, current_node.2,))
    }

    // Right
    if current_node.2 < grid.row_len()-1 {
        neighbours.push(Node(*grid.get(current_node.1, current_node.2+1).unwrap(), current_node.1, current_node.2+1))
    }
    // Down
    if current_node.1 < grid.column_len()-1 {
        neighbours.push(Node(*grid.get(current_node.1+1, current_node.2).unwrap(), current_node.1+1, current_node.2))
    }
    // Left
    if current_node.2 > 0 {
        neighbours.push(Node(*grid.get(current_node.1, current_node.2-1).unwrap(), current_node.1, current_node.2-1))
    }

    // Remove all that are already in visited
    let mut neighbours_filter = vec![];
    for n in neighbours {
        let mut is_visited = false;
        for v in visited {
            if n.0 == v.0 && n.1 == v.1 && n.2 == v.2 {
                is_visited = true;
            }
        }
        if !is_visited {
            neighbours_filter.push(n)
        }
    }

    return neighbours_filter;
}

fn expand_galaxy(split_input: Vec<&str>) -> Vec<String> {
    // First expand galaxy
    let mut expanded_galaxy: Vec<String> = vec![];

    // Expand rows
    for line in split_input {
        if line.find("#") == None {
            // Expand
            expanded_galaxy.push(line.to_string());
            let mut expand = String::new();
            for i in 0..line.len() {
                expand.push_str(".");
            }
            expanded_galaxy.push(expand); 
        } else {
            expanded_galaxy.push(line.to_string());
        }
    }

    // Transpose
    let mut transposed_galaxy: Vec<String> = vec![];
    for column in 0..expanded_galaxy[0].len() {
        let mut transpose_line = String::new();
        for line_num in 0..expanded_galaxy.len() {
            let chars: Vec<char> = expanded_galaxy[line_num].chars().collect();
            transpose_line.push_str(&chars[column].to_string());
        }
        transposed_galaxy.push(transpose_line);
    }

    expanded_galaxy.clear();

    // Then expand columns
    for line in transposed_galaxy {
        if line.find("#") == None {
            // Expand
            expanded_galaxy.push(line.to_string());
            let mut expand = String::new();
            for i in 0..line.len() {
                expand.push_str(".");
            }
            expanded_galaxy.push(expand); 
        } else {
            expanded_galaxy.push(line.to_string());
        }
    }

    // Transpose again
    let mut transposed_galaxy_2: Vec<String> = vec![];
    for column in 0..expanded_galaxy[0].len() {
        let mut transpose_line = String::new();
        for line_num in 0..expanded_galaxy.len() {
            let chars: Vec<char> = expanded_galaxy[line_num].chars().collect();
            transpose_line.push_str(&chars[column].to_string());
        }
        transposed_galaxy_2.push(transpose_line);
    }

    return transposed_galaxy_2;
}

fn part2(input:String) -> usize {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    return result.try_into().unwrap();
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_p1_example() {
        let test = String::from("...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....");
        let answer = part1(test);
        assert_eq!(answer, 374);
    }

    #[test]
    fn test_p2_example() {
        let test = String::from("");
        let answer = part2(test);
        assert_eq!(answer, 4);
    }

}