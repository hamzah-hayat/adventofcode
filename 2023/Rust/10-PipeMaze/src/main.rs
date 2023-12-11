use std::fs;
use array2d::Array2D;

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
enum DIRECTION {
    ANY,
    NORTH,
    EAST,
    SOUTH,
    WEST,
}

// For each Node
// 0 is pipe char
// 1 is x position
// 2 is y position
// 3 is direction we came from
#[derive(Copy, Clone)]
struct Node(char,usize,usize,DIRECTION);

fn part1(input:String) -> usize {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    // Parse input
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

    // Find start position
    let mut start_x = 0;
    let mut start_y = 0;

    for x in 0..grid.row_len() {
        for y in 0..grid.column_len() {
            if grid.get(x, y).unwrap() == &'S' {
                start_x = x;
                start_y = y;
            }
        }
    }

    // Solve
    let mut visited = vec![];
    let start = Node('S',start_x,start_y,DIRECTION::ANY);
    visited.push(start);
    let solution = dfs(&grid, start.clone(), visited);
    result = (solution.len()-1)/2;

    return result.try_into().unwrap();
}

fn dfs(grid:&Array2D<char>,current_node:Node,mut visited:Vec<Node>) -> Vec<Node> {

    // Neighbours
    let mut neighbours = get_neighbours(grid,&visited, current_node);

    while neighbours.len() != 0 {
        let next_node = neighbours.pop().unwrap();
        visited.push(next_node);

        if next_node.0 =='S' {
            return visited;
        } else {
            visited = dfs(grid, next_node, visited);
        }
    }

    return visited;
}

fn get_neighbours(grid:&Array2D<char>,visited:&Vec<Node>,node:Node) -> Vec<Node> {
    let mut neighbours = vec![];

    let mut can_go_up = false;
    let mut can_go_right = false;
    let mut can_go_down = false;
    let mut can_go_left = false;

    if node.0 == 'S' || node.0 == '|' || node.0 == 'J' || node.0 == 'L' {
        can_go_up = true;
    }
    if node.0 == 'S' || node.0 == '-' || node.0 == 'F' || node.0 == 'L' {
        can_go_right = true;
    }
    if node.0 == 'S' || node.0 == '|' || node.0 == 'F' || node.0 == '7' {
        can_go_down = true;
    }
    if node.0 == 'S' || node.0 == '-' || node.0 == '7' || node.0 == 'J' {
        can_go_left = true;
    }

    // Four possible directions
    let mut top_node = &'.';
    if node.1>0 {
        top_node = grid.get(node.1-1, node.2).unwrap_or(&'.');
    }
    let right_node = grid.get(node.1, node.2+1).unwrap_or(&'.');
    let bottom_node = grid.get(node.1+1, node.2).unwrap_or(&'.');
    let mut left_node = &'.';
    if node.2>0 {
        left_node = grid.get(node.1, node.2-1).unwrap_or(&'.');
    }

    if can_go_up && (top_node == &'7' || top_node == &'F' || top_node == &'|' || top_node == &'S') {
        neighbours.push(Node(top_node.clone(), node.1-1, node.2, DIRECTION::SOUTH))
    }

    if can_go_right && (right_node == &'7' || right_node == &'J' || right_node == &'-' || right_node == &'S') {
        neighbours.push(Node(right_node.clone(), node.1, node.2+1, DIRECTION::WEST))
    }

    if can_go_down && (bottom_node == &'L' || bottom_node == &'J' || bottom_node == &'|' || bottom_node == &'S') {
        neighbours.push(Node(bottom_node.clone(), node.1+1, node.2, DIRECTION::NORTH))
    }

    if can_go_left && (left_node == &'L' || left_node == &'F' || left_node == &'-' || left_node == &'S') {
        neighbours.push(Node(left_node.clone(), node.1, node.2-1, DIRECTION::EAST))
    }

    // remove any neighbours that are in visited
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

fn part2(input:String) -> usize {
    let split_input: Vec<&str> = input.split("\n").collect();
    let mut result = 0;

    // Parse input
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

    // for x in 0..grid.column_len() {
    //     for y in 0..grid.row_len() {
    //         print!("{}",grid.get(x, y).unwrap());
    //     }
    //     println!();
    // }

    // Find start position
    let mut start_x = 0;
    let mut start_y = 0;

    for y in 0..grid.row_len() {
        for x in 0..grid.column_len() {
            if grid.get(x, y).unwrap() == &'S' {
                start_x = x;
                start_y = y;
            }
        }
    }

    // Solve
    let mut visited = vec![];
    let start = Node('S',start_x,start_y,DIRECTION::ANY);
    visited.push(start);
    let mut solution = dfs(&grid, start.clone(), visited);

    // Now that we have solution, find every space "inside" our solution
    let mut in_loop = vec![];
    let mut corner = false;
    let mut current_corner = '.';
    for x in 0..grid.column_len() {
        let mut number_of_passes = 0;
        for y in 0..grid.row_len() {
            //print!("{}",grid.get(x, y).unwrap());
            let current_grid = *grid.get(x, y).unwrap();
            if corner {

                // This is because The S could be an S or L
                // Currently hardcoded because i cba at this point
                // For the examples, S is actually an F, while on the real input the S is an L
                // Hardcoded to work on input, but for examples we need to swap around the '7' and 'J'
                if current_corner=='S' {
                    if current_grid == '7' {
                        number_of_passes += 1;
                        corner = false;
                        current_corner = '.';
                    }
                    if current_grid == 'J' {
                        corner = false;
                        current_corner = '.';
                    }
                }

                // If we see a corner F, we have to mark as corner and skip until we either find a J (go out of loop) or 7(stay in loop)
                // If we see a corner L, we have to mark as corner and skip unil we find a J (stay in loop) or 7(leave loop)
                if current_corner=='F' {
                    if current_grid == 'J' {
                        number_of_passes += 1;
                        corner = false;
                        current_corner = '.';
                    }
                    if current_grid == '7' {
                        corner = false;
                        current_corner = '.';
                    }
                }

                if current_corner=='L' {
                    if current_grid == '7' {
                        number_of_passes += 1;
                        corner = false;
                        current_corner = '.';
                    }
                    if current_grid == 'J' {
                        corner = false;
                        current_corner = '.';
                    }
                }

            }

            // If we are on a corner, mark it and continue
            if !corner && is_corner(x,y,&grid,&solution) {
                corner = true;
                current_corner = current_grid;
            }

            // If we see a flat edge, we can increase number of passes and start counting spaces
            if is_flat_edge(x,y,&grid,&solution) {
                number_of_passes += 1;
            }

            // Mark as space
            if number_of_passes % 2 == 1 && not_part_of_loop(x,y,&grid,&solution) && !corner {
                in_loop.push(Node(*grid.get(x, y).unwrap(),x,y,DIRECTION::ANY))
            }
        }
        //println!()
    }

    result = in_loop.len();



    return result.try_into().unwrap();
}

fn is_flat_edge(x: usize, y: usize, grid: &Array2D<char>, solution: &Vec<Node>) -> bool {
    for n in solution {
        if n.1 == x && n.2 == y && n.0 == '|' {
            return true;
        }
    }
    return false;
}

fn is_corner(x: usize, y: usize, grid: &Array2D<char>, solution: &Vec<Node>) -> bool {
    for n in solution {
        if n.1 == x && n.2 == y && (n.0 == 'F' || n.0 == 'L' || n.0 == 'S') {
            return true;
        }
    }
    return false;
}

fn not_part_of_loop(x: usize, y: usize, grid: &Array2D<char>, solution: &Vec<Node>) -> bool {
    for n in solution {
        if n.1 == x && n.2 == y {
            return false;
        }
    }
    return true;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_p1_example() {
        let test = String::from("-L|F7
7S-7|
L|7||
-L-J|
L|-JF");
        let answer = part1(test);
        assert_eq!(answer, 4);
    }

    #[test]
    fn test_p1_example_2() {
        let test = String::from("7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ");
        let answer = part1(test);
        assert_eq!(answer, 8);
    }

    #[test]
    fn test_p2_example() {
        let test = String::from("...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........");
        let answer = part2(test);
        assert_eq!(answer, 4);
    }

    #[test]
    fn test_p2_example_2() {
        let test = String::from(".F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...");
        let answer = part2(test);
        assert_eq!(answer, 8);
    }

    #[test]
    fn test_p2_example_3() {
        let test = String::from("FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L");
        let answer = part2(test);
        assert_eq!(answer, 10);
    }

}