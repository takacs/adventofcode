fn main() {
    let input = include_str!("./day02.in");
    let output1 = part1(input);
    dbg!(output1);
    let output2 = part2(input);
    dbg!(output2);
}

fn part1(input: &str) -> String {
    let lines: Vec<&str> = input.split("\n").collect();

    let mut safecount = 0;
    for line in lines.iter() {
        let nums: Vec<i32> = line
            .split_whitespace()
            .filter_map(|s| s.parse::<i32>().ok())
            .collect();
        if nums.len() == 0 {
            continue;
        }
        if is_safe(nums.clone()) {
            safecount += 1
        }
    }

    safecount.to_string()
}

fn is_safe(nums: Vec<i32>) -> bool {
    let increasing: bool = nums[0] < nums[1];
    for i in 1..nums.len() {
        let diff: i32 = nums[i] - nums[i - 1];
        if diff.abs() > 3 || diff.abs() == 0 {
            return false;
        }

        if increasing && diff < 0 {
            return false;
        }

        if !increasing && diff > 0 {
            return false;
        }
    }
    return true;
}

fn part2(input: &str) -> String {
    let lines: Vec<&str> = input.split("\n").collect();

    let mut safecount = 0;
    for line in lines.iter() {
        let nums: Vec<i32> = line
            .split_whitespace()
            .filter_map(|s| s.parse::<i32>().ok())
            .collect();
        if nums.len() == 0 {
            continue;
        }
        if is_safe2(nums.clone()) {
            safecount += 1
        } else {
        }
    }

    safecount.to_string()
}

fn is_safe2(nums: Vec<i32>) -> bool {
    if is_safe(nums.clone()) {
        return true;
    }

    for i in 0..nums.len() {
        let mut nums_rm = nums.clone();
        nums_rm.remove(i);

        if is_safe(nums_rm) {
            return true;
        }
    }

    false
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let result = part1(
            "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9",
        );
        assert_eq!(result, "2".to_string());
    }

    #[test]
    fn test_part2() {
        let result = part2(
            "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9",
        );
        assert_eq!(result, "4".to_string());
    }

    #[test]
    fn test_part2_first_remove() {
        let result = part2("10 1 2 3 4");
        assert_eq!(result, "1".to_string());
    }

    #[test]
    fn test_part2_last_remove() {
        let result = part2("1 2 3 4 0");
        assert_eq!(result, "1".to_string());
    }

    #[test]
    fn test_part2_unsafe_found() {
        let result = part2("76 73 71 68 66 65 59");
        assert_eq!(result, "1".to_string());
    }

    #[test]
    fn test_part2_middle_remove() {
        let result = part2("1 5 2 3 4");
        assert_eq!(result, "1".to_string());
    }

    #[test]
    fn test_part2_complex_case() {
        let result = part2("1 2 5 4 6 7 3");
        assert_eq!(result, "0".to_string());
    }

    #[test]
    fn test_part2_first_remove_2() {
        let result = part2("27 24 25 26 28 31 34");
        assert_eq!(result, "1".to_string());
    }
}
