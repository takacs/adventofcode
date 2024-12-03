use std::collections::HashMap;

fn main() {
    let input = include_str!("./day01.in");
    let output1 = part1(input);
    let output2 = part2(input);
    dbg!(output1);
    dbg!(output2);
}

fn part1(input: &str) -> String {
    let lines: Vec<&str> = input.split("\n").collect();
    let mut l: Vec<i32> = Vec::new();
    let mut r: Vec<i32> = Vec::new();

    for line in lines.iter() {
        let nums: Vec<&str> = line.split_whitespace().collect();
        if nums.len() < 2 {
            break;
        }
        l.push(nums[0].parse().unwrap());
        r.push(nums[1].parse().unwrap());
    }

    l.sort();
    r.sort();
    let mut total = 0;
    for i in 0..l.len() {
        total += (l[i] - r[i]).abs()
    }
    total.to_string()
}

fn part2(input: &str) -> String {
    let lines: Vec<&str> = input.split("\n").collect();
    let mut l: Vec<i32> = Vec::new();
    let mut r: Vec<i32> = Vec::new();

    for line in lines.iter() {
        let nums: Vec<&str> = line.split_whitespace().collect();
        if nums.len() < 2 {
            break;
        }
        l.push(nums[0].parse().unwrap());
        r.push(nums[1].parse().unwrap());
    }

    let mut counter: HashMap<&i32, i32> = HashMap::new();
    for num in r.iter() {
        if !counter.contains_key(num) {
            counter.insert(num, 0);
        }
        *counter.get_mut(num).unwrap() += 1;
    }

    let mut total = 0;
    for num in l.iter_mut() {
        let mut multiplier: i32 = 0;
        if counter.contains_key(num) {
            multiplier = counter[num];
        }
        total += *num * multiplier;
    }
    total.to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let result = part1(
            "3   4
4   3
2   5
1   3
3   9
3   3",
        );
        assert_eq!(result, "11".to_string());
    }

    #[test]
    fn test_part2() {
        let result = part2(
            "3   4
4   3
2   5
1   3
3   9
3   3",
        );
        assert_eq!(result, "31".to_string());
    }
}
