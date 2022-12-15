#![allow(dead_code)]

mod util {
  use std::{path::Path, fs::{File}, io::{self, BufRead}};
  pub fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>> where P: AsRef<Path>, {
      let file = File::open(filename)?;
      Ok(io::BufReader::new(file).lines())
  }
}

mod calorie_counting {
  use crate::util;
  pub fn count_calories(file_path: &str) -> i32 {
    let mut most_calories = [0,0,0];
    let mut cur_calories: i32 = 0;
    if let Ok(lines) = util::read_lines(file_path) {
      for line in lines {
        if let Ok(calorie) = line {
          if let Ok(int_calorie) = calorie.parse::<i32>(){
            cur_calories += int_calorie;
          } else {
            for check_most in most_calories.iter_mut() {
              if cur_calories > *check_most {
                *check_most = cur_calories;
              }
              break;
            }
            most_calories.sort();
            cur_calories = 0;
          }
        }
      }
    }
    most_calories.iter().sum()
  }
}

mod rps {
  use crate::util;
  fn play(lhs: char, rhs: char) -> i32 {
    let mut score = 0;
    let norm_rhs = match rhs {
      'X' => 'A',
      'Y' => 'B',
      'Z' => 'C',
      _ => '\0'
    };
    score += match rhs {
      'X' => 1,
      'Y' => 2,
      'Z' => 3,
      _ => 0
    };
    if lhs == norm_rhs {
      score += 3;
    } else {
      if lhs == 'A' {
        score += match norm_rhs {
                  'B' => 6,
                  'C' => 0,
                  _ => 0
                };
      }
      if lhs == 'B' {
        score += match norm_rhs {
                  'A' => 0,
                  'C' => 6,
                  _ => 0
                };
      }
      if lhs == 'C' {
        score += match norm_rhs {
                  'A' => 6,
                  'B' => 0,
                  _ => 0
                };
      }
    }
    score
  }
  fn play_reverse(lhs: char, rhs: char) -> i32 {
    let mut score = match rhs {
      'X' => 0,
      'Y' => 3,
      'Z' => 6,
      _ => 0
    };
    if lhs == 'A' {
      score += match rhs {
                'X' => 3,
                'Y' => 1,
                'Z' => 2,
                _ => 0
              };
    }
    if lhs == 'B' {
      score += match rhs {
                'X' => 1,
                'Y' => 2,
                'Z' => 3,
                _ => 0
              };
    }
    if lhs == 'C' {
      score += match rhs {
                'X' => 2,
                'Y' => 3,
                'Z' => 1,
                _ => 0
              };
    }
    score
    }
  

  pub fn rock_paper_sissors(file_path: &str) -> i32 {
    let mut total_score = 0;
    if let Ok(lines) = util::read_lines(file_path) {
      for line in lines {
        if let Ok(game) = line {
          let hands: Vec<char> = game.chars().collect();
          total_score += play_reverse(hands[0], hands[2]);
        }
      }
    }
    total_score
  }
}

mod rucksack {
use crate::util;

  fn compute_sack(sack: &String) -> u32 {
    let halves = sack.split_at(sack.len() / 2);
    let mut shared = '\0';
    for c in halves.0.chars() {
      if halves.1.contains(c) {
        shared = c;
      }
    }
    let mut to_dec = shared as u32;
    if to_dec >= 97 {
      to_dec -= 96;
    }
    if (65..91).contains(&to_dec) {
      to_dec -= 38;
    }
    to_dec
  }

  pub fn parse_rucksacks(file_path: &str) -> u32 {
    let mut total = 0;
    if let Ok(lines) = util::read_lines(file_path) {
      for line in lines {
        if let Ok(raw_sack) = line {
          total += compute_sack(&raw_sack);
        }
      }
    }
    total
  }

  fn compute_group(group: &[String; 3]) -> u32 {
    let mut val = '\0';
    for c in group[0].chars() {
      if group[1].contains(c) && group[2].contains(c) {
        val = c;
      }
    }
    let mut to_dec = val as u32;
    if to_dec >= 97 {
      to_dec -= 96;
    }
    if (65..91).contains(&to_dec) {
      to_dec -= 38;
    }
    to_dec
  }

  pub fn parse_rucksack_groups(file_path: &str) -> u32 {
    let mut total = 0;
    if let Ok(lines) = util::read_lines(file_path) {
      let mut group = [String::new(), String::new(), String::new()];
      let mut index = 0;
      for line in lines {
        if let Ok(raw_sack) = line {
          group[index % 3] = raw_sack;
        }
        if (index + 1)% 3 == 0 {
          total += compute_group(&group)
        }
        index += 1;
      }
    }
    total
  }

}

mod camp_cleanup {
    use crate::util;

  pub fn check_pairs(file_path: &str) -> u32 {
    let mut total = 0;
    if let Ok(lines) = util::read_lines(file_path) {
      for line in lines {
        if let Ok(pair_str) = line {
          let pairs: Vec<&str> = pair_str.split(',').collect();
          let pair1: Vec<&str> = pairs[0].split('-').collect();
          let pair2: Vec<&str> = pairs[1].split('-').collect();

          let lhs_start = pair1[0].parse::<u32>().unwrap();
          let lhs_end = pair1[1].parse::<u32>().unwrap();

          let rhs_start = pair2[0].parse::<u32>().unwrap();
          let rhs_end = pair2[1].parse::<u32>().unwrap();

          if pair_contains_at_all([lhs_start, lhs_end], [rhs_start, rhs_end]) {
            total += 1;
          }
        }
      }
    }
    total
  }

  fn pair_contains(lhs: [u32; 2], rhs: [u32; 2]) -> bool {
    if (lhs[0] <= rhs[0] && lhs[1] >= rhs[1]) ||  (rhs[0] <= lhs[0] && rhs[1] >= lhs[1]) {
      return true;
    }
    false
  }
  
  fn pair_contains_at_all(lhs: [u32; 2], rhs: [u32; 2]) -> bool {
    if (lhs[0] <= rhs[0] && lhs[1] >= rhs[0]) || (lhs[0] <= rhs[1] && lhs[1] >= rhs[1]) || 
      (rhs[0] <= lhs[0] && rhs[1] >= lhs[0])  || (rhs[0] <= lhs[1] && rhs[1] >= lhs[1]) {
      return true;
    }
    false
  }
}

mod supply_stacks {
    use crate::util;

  pub fn move_many_crates(file_path: &str) -> Vec<char>{
    let mut stacks = [vec!['J', 'H', 'G', 'M', 'Z', 'N', 'T', 'F'],
                                      vec!['V', 'W' , 'J'],
                                      vec!['G', 'V', 'L', 'J', 'B' , 'T', 'H'],
                                      vec!['B', 'P', 'J', 'N', 'C', 'D', 'V', 'L'],
                                      vec!['F', 'W', 'S', 'M', 'P', 'R', 'G'],
                                      vec!['G', 'H', 'C', 'F', 'B', 'N', 'V', 'M'],
                                      vec!['D', 'H', 'G', 'M', 'R'],
                                      vec!['H', 'N', 'M', 'V', 'Z', 'D'],
                                      vec!['G', 'N', 'F', 'H']
                    ];

    // Skip first 10 lines
    let mut line_index = 0;
    if let Ok(lines) = util::read_lines(file_path) {
      for line in lines {
        if line_index < 10 {
          line_index += 1;
          continue;
        }
        if let Ok(move_str) = line {
          let tokens: Vec<&str> = move_str.split(' ').collect();
          let to_move = tokens[1].parse::<u32>().unwrap();
          let orig_size = stacks[tokens[3].parse::<usize>().unwrap() - 1].len();
          let mut temp:Vec<_> = stacks[tokens[3].parse::<usize>().unwrap() - 1].drain(((orig_size-to_move as usize))..).collect();
          stacks[tokens[5].parse::<usize>().unwrap() - 1].append(&mut temp);
        }
      }
    }
    let mut last_elements:Vec<char> = vec![];
    for x in 0..stacks.len() {
      last_elements.push(*stacks[x].last().unwrap());
    }
  last_elements
  }
  pub fn move_crates(file_path: &str) -> Vec<char>{
    let mut stacks = [vec!['J', 'H', 'G', 'M', 'Z', 'N', 'T', 'F'],
                                      vec!['V', 'W' , 'J'],
                                      vec!['G', 'V', 'L', 'J', 'B' , 'T', 'H'],
                                      vec!['B', 'P', 'J', 'N', 'C', 'D', 'V', 'L'],
                                      vec!['F', 'W', 'S', 'M', 'P', 'R', 'G'],
                                      vec!['G', 'H', 'C', 'F', 'B', 'N', 'V', 'M'],
                                      vec!['D', 'H', 'G', 'M', 'R'],
                                      vec!['H', 'N', 'M', 'V', 'Z', 'D'],
                                      vec!['G', 'N', 'F', 'H']
                    ];

    // Skip first 10 lines
    let mut line_index = 0;
    if let Ok(lines) = util::read_lines(file_path) {
      for line in lines {
        if line_index < 10 {
          line_index += 1;
          continue;
        }
        if let Ok(move_str) = line {
          let tokens: Vec<&str> = move_str.split(' ').collect();
          for _x in 0..tokens[1].parse::<u32>().unwrap() {
            let temp = stacks[tokens[3].parse::<usize>().unwrap() - 1].pop().unwrap();
            stacks[tokens[5].parse::<usize>().unwrap() - 1].push(temp);
          }
        }
      }
    }
    let mut last_elements:Vec<char> = vec![];
    for x in 0..stacks.len() {
      last_elements.push(*stacks[x].last().unwrap());
    }
  last_elements
  }
}

mod tuning_trouble{
  use crate::util;
  use std::collections::HashSet;

  pub fn lock_on(file_path: &str) -> usize {
    if let Ok(lines) = util::read_lines(file_path) {
      for line in lines {
        if let Ok(buffer) = line {
          let mut start_index = 0;
          while start_index + 13 <= buffer.len() {
            let slice = &buffer[start_index..start_index+14];
            let mut hset = HashSet::new();
            for c in slice.chars() {
              hset.insert(c);
            }
            if hset.len() == 14 {
              return start_index + 14;
            } else {
              start_index += 1;
            }
          }
        }
      }
    }
  0
  }
}

mod no_space_left {
  use std::collections::HashMap;

use crate::util;
  #[derive(Debug)]
  struct Dir {
    name: String,
    files: Vec<u32>,
    children: HashMap<String, usize>,
    parent: usize
  }

  fn get_dir_size(dir_index: usize, filesystem: &Vec<Dir>) -> u32 {
    let dir = &filesystem[dir_index];
    if dir.children.is_empty() {
      return dir.files.iter().sum();
    } else {
      let mut size = dir.files.iter().sum();
      for (_name, index) in &dir.children {
        size += get_dir_size(*index, filesystem);
      }
      size 
    }
  }

  pub fn create_file_system(file_path: &str) -> u32 {
    let mut file_index: usize = 1;
    let mut filesystem: Vec<Dir> = vec![Dir {
      name: String::from("/"),
      files: vec![],
      children: HashMap::new(),
      parent: 0
    }];
    let mut cwd: usize = 0;
    if let Ok(lines) = util::read_lines(file_path) {
      for line in lines {
        let line_str = line.unwrap();
        let tokens: Vec<&str> = line_str.split(' ').collect();
        if line_str == "$ cd /" {
          cwd = 0;
          continue;
        }
        // Commands [ls, cd]
        if line_str.starts_with('$') {
          if tokens[1] == "cd" {
            if tokens[2] == ".." {
              cwd = filesystem[cwd].parent;
            } else {
              cwd = filesystem[cwd].children[tokens[2]];
            }
          } else if tokens[1] == "ls" {
            continue;
          }
        }
        // result of an ls [files, dirs] 
        else {
          let wd = &mut filesystem[cwd];
          if line_str.starts_with("dir") {
            wd.children.insert(String::from(tokens[1]), file_index);
            filesystem.push(Dir {
              name: String::from(tokens[1]),
              files: vec![],
              children: HashMap::new(),
              parent: cwd
            });
            file_index += 1;
          } else {
            wd.files.push(tokens[0].parse::<u32>().unwrap());
          }
        }
      }
    }
    let mut file_sizes: Vec<u32> = vec![];
    let mut smallest_to_del = std::u32::MAX;
    let free_size =70000000 - get_dir_size(0, &filesystem);
    let addiontal_size_to_free = 30000000 - free_size;
    println!("additional size to free {}", addiontal_size_to_free);
    // Traverse filesystem 
    for i in 0..filesystem.len() {
      let s = get_dir_size(i, &filesystem);
      if s <= 100000 {
        file_sizes.push(s);
      }
      if s >= addiontal_size_to_free && s < smallest_to_del {
        smallest_to_del = s;
      }
    }
    println!("smallest to del {}", smallest_to_del);
    println!("{}", get_dir_size(0, &filesystem));
    file_sizes.iter().sum()
  }

}
  mod treetop_tree_house {
    use crate::util;
    
    fn check_visibility(forest: &Vec<Vec<u32>>, x: usize, y: usize) -> bool {
      if x == 0 || y == 0 {
        return true;
      }
      let cur_height = forest[y][x];
      
      let mut taller_than_right = true;
      for right in x+1..forest[y].len()  {
        taller_than_right &= cur_height > forest[y][right];
      }
      let mut taller_than_left = true;
      for left in (0..x).rev() {
        taller_than_left &= cur_height > forest[y][left];
      }
      let mut taller_than_above = true;
      for above in (0..y).rev() {
        taller_than_above &= cur_height > forest[above][x];
      }
      let mut taller_than_below = true;
      for below in y+1..forest.len() {
        taller_than_below &= cur_height > forest[below][x];
      }
      taller_than_above || taller_than_below || taller_than_left || taller_than_right
    }

    fn count_visible_trees(forest: &Vec<Vec<u32>>) -> u32 {
      // now check each tree for visibility
      let num_rows = forest.len();
      let num_cols = forest[0].len();
      let mut visible_trees = 0;
      for y in 0..num_rows {
        for x in 0..num_cols{
          if check_visibility(forest, x, y) {
            visible_trees += 1;
          }
        }
      }
      visible_trees
    }

    fn compute_scenic_score(forest: &Vec<Vec<u32>>, x: usize, y: usize) -> u32 {
      if x == 0 || y == 0 || x == forest[0].len()-1 || y == forest.len()-1{
        return 0;
      }
      let (mut max_up, mut max_down, mut max_left, mut max_right) = (1, 1, 1, 1);
      let cur_height = forest[y][x];
      for right in x+1..forest[y].len()-1 {
        match cur_height > forest[y][right] {
          true => max_right += 1,
          false => break
        }
      }
      for left in (1..x).rev() {
        match cur_height > forest[y][left] {
          true => max_left += 1,
          false => break
        }
      }
      for down in y+1..forest.len()-1 {
        match cur_height > forest[down][x] {
          true => max_down += 1,
          false => break
        }
      }
      for up in (1..y).rev() {
        match cur_height > forest[up][x] {
          true => max_up += 1,
          false => break
        }
      }
      
      max_up * max_down * max_left * max_right
    }

    fn find_max_scenic_score(forest: &Vec<Vec<u32>>) -> u32 {
      let num_rows = forest.len();
      let num_cols = forest[0].len();
      let mut max_scenic_score: u32 = 0;
      for y in 0..num_rows {
        for x in 0..num_cols{
          let scenic_score = compute_scenic_score(forest, x, y);
          if scenic_score > max_scenic_score {
            max_scenic_score = scenic_score;
          }
        }
      }
      max_scenic_score
    }
    pub fn run_trees(file_path: &str) -> u32 {
      let mut forest:Vec<Vec<u32>> = vec![];
      if let Ok(lines) = util::read_lines(file_path) {
        for line in lines {
          let mut row: Vec<u32> = vec![];
          let line_str = line.unwrap();
          for c in line_str.chars() {
            row.push(c.to_digit(10).unwrap());
          }
          forest.push(row);
        }
        find_max_scenic_score(&forest)
      } else {
        panic!("Error reading {}", file_path);
      }


    }
  }

  mod rope_bridge {
    use crate::util;
    use std::collections::HashSet;

    #[derive(Clone, Copy, Debug)]
    struct End {
      x: i32,
      y: i32
    }

    impl End {
      fn new() -> Self {
        End {
          x: 0,
          y: 0
        }
      }
    }

    fn move_rope(rope: &mut [End; 10], direction: &str, magnitude: i32, visited_locations: &mut HashSet<(i32, i32)>) {
      for _ in 0..magnitude {
        match direction {
          "R" => rope[0].x += 1,
          "D" => rope[0].y += 1,
          "L" => rope[0].x -= 1,
          "U" => rope[0].y -= 1,
          _ => {},
        }
        for rope_index in 1..rope.len() {
          let (head_list, tail_list) = rope.split_at_mut(rope_index);
          let head = head_list.last_mut().unwrap();
          let tail = tail_list.first_mut().unwrap();
          // Same row, different column
          if (head.x - tail.x).abs() > 1 && head.y == tail.y {
            match head.x > tail.x {
              true => tail.x += 1,
              false => tail.x -= 1,
            }
          }
          // Same column, differnt row
          else if (head.y - tail.y).abs() > 1 && head.x == tail.x {
            match head.y > tail.y {
              true => tail.y += 1,
              false => tail.y -= 1,
            }
          } 
          // If either is further away in row or column by 2, then the other dimension must
          // also be adjustable
          else if (head.x - tail.x).abs() > 1 || (head.y - tail.y).abs() > 1 {
            match head.y > tail.y {
              true => tail.y += 1,
              false => tail.y -= 1,
            }
            match head.x > tail.x {
              true => tail.x += 1,
              false => tail.x -= 1,
            }
          } else if (head.y - tail.y).abs() <= 1 && (head.x - tail.x).abs() <= 1 {
            // noop
          } else {
            panic!("rope index {}, node before {:?}, node at index {:?}", rope_index, head, tail);
          }
        }
        visited_locations.insert((rope.last().unwrap().x, rope.last().unwrap().y));
      }
    }

    pub fn rope_path(file_path: &str) -> usize {
      let mut rope: [End; 10] = [End::new(); 10];
      let mut visted_locations:HashSet<(i32, i32)> = HashSet::new();
      if let Ok(lines) = util::read_lines(file_path) {
        for line in lines {
          let line_str = line.unwrap();
          let tokens: Vec<&str> = line_str.split_whitespace().collect();
          let direction = tokens[0];
          let magnitude = tokens[1].parse::<i32>().unwrap();
          move_rope(&mut rope, direction, magnitude, &mut visted_locations);
        }
      }
      visted_locations.len()
    }
  }

  mod crt {
    use crate::util;

    fn run_cycle(cycle: &mut i32, x: i32, values: &mut Vec<i32> , crt: &mut [[char; 40]; 6]) {
      const CHECK_CYCLES: [i32; 6] = [20, 60, 100, 140, 180, 220];
      if CHECK_CYCLES.contains(cycle) {
        values.push(*cycle * x);
      }
      let (horizontal, vertical) = get_coordinates_from_cycle(*cycle);
      if x.abs_diff(horizontal) <= 1 {
        crt[vertical as usize][horizontal as usize] = '#';
      }
      *cycle += 1;
    }

    fn get_coordinates_from_cycle(cycle: i32) -> (i32, i32) {
      let y = (cycle - 1) / 40;
      let x = (cycle - 1) % 40;
      (x, y)
    }

    pub fn signal_strength(file_path: &str) -> i32 {
      let mut cycle = 1;
      let mut X = 1;
      let mut values: Vec<i32> = vec![];
      let mut crt: [[char; 40]; 6] = [['.'; 40]; 6];
      if let Ok(lines) = util::read_lines(file_path) {
        for line in lines.flatten() {
          let tokens: Vec<&str> = line.split_whitespace().collect();
          if tokens[0] == "noop" {
            run_cycle(&mut cycle, X, &mut values, &mut crt);
          } else if tokens[0] == "addx" {
            run_cycle(&mut cycle, X, &mut values, &mut crt);
            run_cycle(&mut cycle, X, &mut values, &mut crt);
            X += tokens[1].parse::<i32>().unwrap();
          } else {
            panic!("Unrecognized instruction {}", tokens[0]);
          }
        }
      }
    for y in 0..6 {
      for x in 0..40 {
        print!("{}" , crt[y][x]);
      }
      println!("");
    }
    values.iter().sum()
    }
  }

  mod monkey_middle {
    use std::collections::{VecDeque, HashMap};

    use crate::util;

    struct Monkey {
      operation: Box<dyn Fn(&i64) -> i64>,
      test: Box<dyn Fn(&i64) -> usize>,
      test_val: i64,
      items: VecDeque<i64>,
      inspections: i32,
    }

    impl Monkey {
      fn new() -> Self{
        Monkey{
          operation: Box::new(|_| 0 ),
          test: Box::new(|_| 0),
          test_val: 0,
          items: VecDeque::new(),
          inspections: 0,
        }
      }
    }

    pub fn run_monkeys(file_path: &str) -> u128 {
      let mut monkeys: Vec<Monkey> = vec![];
      let (mut test, mut val_true, mut val_false) = (0 as i64, 0 as usize, 0 as usize);
      if let Ok(lines) = util::read_lines(file_path) {
        for line in lines.flatten() {
          if line.starts_with("Monkey") {
            monkeys.push(Monkey::new());
          }
          else if line.trim().starts_with("Starting") {
            monkeys.last_mut().unwrap().items = line.split_whitespace().skip(2).map(|x| x.replace(",", "").parse::<i64>().unwrap()).collect();
          }
          else if line.trim().starts_with("Operation") {
            if line.contains('*') {
              if line.ends_with("old") {
                monkeys.last_mut().unwrap().operation = Box::new(move |x| x * x);
              } else {
                monkeys.last_mut().unwrap().operation = Box::new(move |x| x * line.split_whitespace().last().unwrap().parse::<i64>().unwrap())
              }
            } else if line.contains('+') {
              monkeys.last_mut().unwrap().operation = Box::new(move |x| x + line.split_whitespace().last().unwrap().parse::<i64>().unwrap())
            }
          } else if line.trim().starts_with("Test") {
            test = line.split_whitespace().last().unwrap().parse::<i64>().unwrap();
          } else if line.trim().starts_with("If true") {
            val_true = line.split_whitespace().last().unwrap().parse::<usize>().unwrap();
          } else if line.trim().starts_with("If false") {
            val_false = line.split_whitespace().last().unwrap().parse::<usize>().unwrap();
          } else {
            monkeys.last_mut().unwrap().test = Box::new(move |x| match x % test == 0 {
              true => { val_true},
              false => {val_false}
            } );
            monkeys.last_mut().unwrap().test_val = test;
          }
        }
        monkeys.last_mut().unwrap().test = Box::new(move |x| match x % test == 0 {
          true => {val_true},
          false => {val_false}
        } );
        monkeys.last_mut().unwrap().test_val = test;
      }
      // Begin processing rounds
      let mut divisors: Vec<i64> = vec![];
      for m in &monkeys {
        divisors.push(m.test_val);
      }
      let mut cache: HashMap<i64, i64> = HashMap::new();
      for l in 0..10000 {
        println!("{}", l);
        for monkey_index in 0..monkeys.len() {
          let mut toss: Vec<(usize, i64)> = vec![];
          let monkey = &mut monkeys[monkey_index];
          if monkey.items.len() == 0 {
            continue;
          } else {
            while !monkey.items.is_empty() {
              let item = monkey.items.front_mut().unwrap();
              *item = (monkey.operation)(item);

              let index = (monkey.test)(item);
              if cache.contains_key(item) {
                *item = cache[item];
              } else {
                let mut remainders: Vec<i64> = vec![];

                for d in &divisors {
                  remainders.push(*item % d);
                }
              
                let max_divisor = *divisors.iter().max().unwrap();
                let max_index = divisors.iter().position(|&r| r == max_divisor).unwrap();
                let mut checknum: i64 = max_divisor + remainders[max_index];
                while checknum < *item {
                  let mut substitute = true;
                  for i in 0..divisors.len() {
                    substitute &= checknum % divisors[i] == remainders[i];
                    if !substitute {
                      break;
                    }
                  }
                  if substitute {
                    cache.insert(*item, checknum);
                    *item = checknum;
                    break;
                  }
                  checknum += max_divisor;
                }
              }

              toss.push((index, monkey.items.pop_front().unwrap()));
              monkey.inspections += 1;
            }
          }
          for (index, value) in toss {
            let dest_monkey = &mut monkeys[index];
            dest_monkey.items.push_back(value);
          }
        }
      }
      let mut inspections: Vec<i32> = vec![];
      for monkey in monkeys {
        inspections.push(monkey.inspections);
      }
      inspections.sort();
      inspections.reverse();
      println!("{:?}", inspections);
      inspections[0] as u128 * inspections[1] as u128
      // answer is 29703395016
      // [172863, 171832, 167043, 167042, 160272, 12961, 12942, 12599]
    }
  }

  mod distress_signal {
    use crate::util;
    use serde_json::Value;

    enum State {
      True,
      False,
      Continue
    }

    fn check_pairs(packet1: &Value, packet2: &Value) -> State {
      println!("{} {}", packet1, packet2);
      if packet1.is_array() && packet2.is_array() {
        let array1 = packet1.as_array().unwrap();
        let array2 = packet2.as_array().unwrap();
        for (index, item) in array1.into_iter().enumerate() {
          if index >= array2.len() {
            return State::False;
          }
          match check_pairs(item, &array2[index]) {
            State::True => return State::True,
            State::False => return State::False,
            State::Continue => {},
          }
        }
        if array1.len() < array2.len() {
          return State::True;
        }
      }
      if packet1.is_i64() && packet2.is_i64() {
        match packet1.as_i64().unwrap().cmp(&packet2.as_i64().unwrap()) {
          std::cmp::Ordering::Less =>return State::True,
          std::cmp::Ordering::Greater =>return State::False,
          std::cmp::Ordering::Equal => return State::Continue,
        }
      }
      if packet1.is_array() && packet2.is_i64() {
        let temp =Value::Array(Vec::from([packet2.clone()]));
        return check_pairs(packet1, &temp);
      }
      if packet1.is_i64() && packet2.is_array() {
        let temp =Value::Array(Vec::from([packet1.clone()]));
        return check_pairs(&temp, packet2);
      }
      panic!()
    }

    pub fn run_check_pairs(file_path: &str) -> i32 {
      let mut pair_counter = 0;
      let mut line_counter = 0;
      let mut packet1: Value = Value::Null;
      let mut packet2: Value = Value::Null;
      if let Ok(lines) = util::read_lines(file_path) {
        for line in lines.flatten() {
          if line_counter % 3 == 0 {
            packet1 = serde_json::from_str(&line).unwrap();
          }
          if line_counter % 3 == 1 {
            packet2 = serde_json::from_str(&line).unwrap();
          }
          if line_counter % 3 == 2 {
            match check_pairs(&packet1, &packet2)  {
              State::True => {
                println!("True");
                pair_counter += (line_counter + 1) / 3;
              },
              State::False => {
                println!("False");
              },
              _ => panic!()
            }
          }
          line_counter += 1;
        }
      }
      pair_counter
    }

  }

fn main() {
    //println!("{}", calorie_counting::count_calories("data/calorie_counting.txt"));
    //println!("{}", rps::rock_paper_sissors("data/rock_paper_sissors.txt"))
    //println!("{}", rucksack::parse_rucksacks("data/rucksack.txt"))
    //println!("{}", rucksack::parse_rucksack_groups("data/rucksack.txt"))
    //println!("{}", camp_cleanup::check_pairs("data/camp_cleanup.txt"))
    //println!("{:?}", supply_stacks::move_many_crates("data/supply_stacks.txt"))
    //println!("{}", tuning_trouble::lock_on("data/tuning_trouble.txt"));
    //println!("{}", no_space_left::create_file_system("data/no_space_left.txt"));
    //println!("{}", treetop_tree_house::run_trees("data/treetop_tree_house.txt"));
    //println!("{}", crt::signal_strength("data/signal_strength.txt"));
    //println!("{}", rope_bridge::rope_path("data/rope_bridge.txt"));
    //println!("{}", monkey_middle::run_monkeys("data/monkey_in_the_middle.txt"))
    println!("{}", distress_signal::run_check_pairs("data/distress_signal.txt"))
}
