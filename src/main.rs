mod calorie_counting {
  use std::{path::Path, fs::{File}, io::{self, BufRead}};

  fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>> where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
  }

  pub fn count_calories(file_path: &str) -> i32 {
    let mut most_calories = [0,0,0];
    let mut cur_calories: i32 = 0;
    if let Ok(lines) = read_lines(file_path) {
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

fn main() {
    println!("{}", calorie_counting::count_calories("data/calorie_counting.txt"));
}
