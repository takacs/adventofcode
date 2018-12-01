# part1
def solve(fn):
    count = 0

    with open(fn, 'r') as f:
        for line in f:
            v, val = line[0], int(line[1:])
            if v == '+':
                count += val
            if v == '-':
                count -= val
    return count

# part2
def solve2(fn, countls, count):

    with open(fn, 'r') as f:
        for line in f:
            print(count)
            v, val = line[0], int(line[1:])
            if v == '+':
                count += val
            if v == '-':
                count -= val
            if count in countls:
                return count
            countls.append(count)

    return solve2(fn, countls, count)

print(solve('day1in1.txt'))
print(solve2('day1in2.txt', [0], 0))
