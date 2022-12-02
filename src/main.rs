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

fn main() {
    //println!("{}", calorie_counting::count_calories("data/calorie_counting.txt"));
    println!("{}", rps::rock_paper_sissors("data/rock_paper_sissors.txt"))
}
