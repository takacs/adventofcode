use regex::Regex;

fn main() {
    let input = include_str!("./day03.in");
    let output1 = part1(input);
    dbg!(output1);
    let output2 = part2(input);
    dbg!(output2);
}

fn part1(input: &str) -> String {
    let re = Regex::new(r"mul\(\d+,\d+\)").unwrap();
    let matches: Vec<&str> = re.find_iter(input).map(|mat| mat.as_str()).collect();

    let mut tot: i32 = 0;
    for o in matches {
        tot += extract_nums(o)
    }

    tot.to_string()
}

fn extract_nums(s: &str) -> i32 {
    let s = s.replace("mul(", "");
    let s = s.replace(")", "");
    let ops: Vec<&str> = s.split(',').filter(|s| !s.is_empty()).collect();
    let o1: i32 = ops[0].parse().unwrap();
    let o2: i32 = ops[1].parse().unwrap();

    o1 * o2
}

fn part2(input: &str) -> String {
    let re = Regex::new(r"mul\(\d+,\d+\)").unwrap();
    // creates (start, multiply)
    let muls: Vec<(usize, &str)> = re
        .find_iter(input)
        .map(|mat| (mat.start(), mat.as_str()))
        .collect();
    let mut mix = 0;

    let re = Regex::new(r"do\(\)").unwrap();
    let dos: Vec<usize> = re.find_iter(input).map(|mat| mat.start()).collect();
    let mut doix = 0;

    let re = Regex::new(r"don't\(\)").unwrap();
    let nos: Vec<usize> = re.find_iter(input).map(|mat| mat.start()).collect();
    let mut noix = 0;

    let mut tot: i32 = 0;
    let mut calc = true;
    for i in 0..input.len() {
        if mix < muls.len() && i == muls[mix].0 {
            if calc {
                tot += extract_nums(muls[mix].1);
            }
            mix += 1;
        } else if doix < dos.len() && i == dos[doix] {
            calc = true;
            doix += 1;
        } else if noix < nos.len() && i == nos[noix] {
            calc = false;
            noix += 1;
        }
    }

    tot.to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let result =
            part1("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))");
        assert_eq!(result, "161".to_string());
    }

    #[test]
    fn test_part2() {
        let result =
            part2("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))");
        assert_eq!(result, "48".to_string());
    }
}
