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
    if to_dec >= 65 && to_dec < 91 {
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
    if to_dec >= 65 && to_dec < 91 {
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

fn main() {
    //println!("{}", calorie_counting::count_calories("data/calorie_counting.txt"));
    //println!("{}", rps::rock_paper_sissors("data/rock_paper_sissors.txt"))
    //println!("{}", rucksack::parse_rucksacks("data/rucksack.txt"))
    println!("{}", rucksack::parse_rucksack_groups("data/rucksack.txt"))
}
