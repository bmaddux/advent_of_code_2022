import math

with open("puzzle_input.txt") as infile:
  sum = 0
  for line in infile:
    header = line.split(":")[0]
    numbers = line.split(":")[1]
    lucky_nums = [x.strip() for x in numbers.split("|")[0].strip().split(" ") if len(x) > 0]
    our_nums = [x.strip() for x in numbers.split("|")[1].strip().split(" ") if len(x) > 0]

    pow = 0
    for num in lucky_nums:
      if num in our_nums:
        pow += 1
    if pow > 0:
      value = int(math.pow(2, pow - 1))
      sum += int(math.pow(2, pow - 1))
  
  print(sum)
