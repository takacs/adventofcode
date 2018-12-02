from collections import Counter

# part1
def checksum(fn):
    two, three =  0,0
    with open(fn, 'r') as f:
        for line in f:
            linevals = set(Counter(line).values())
            if 2 in linevals:
                two += 1
            if 3 in linevals:
                three += 1
    return two*three

# part2
def stringdiff(s1, s2):
    return sum([1 for s1,s2 in zip(s1,s2) if s1 != s2])

def readfile(fn):
    with open(fn, 'r')  as f:
        ls = [line[:-1] for line in f]
    return ls

def matchingchars(s1, s2):
    match = [s1 for s1, s2 in zip(list(s1), list(s2)) if s1 == s2]
    return ''.join(match)

def findclose(fn):
    ids =  readfile(fn)
    diffs = []
    for i in ids:
        for j in ids:
            if stringdiff(i, j) == 1:
                return matchingchars(i,j)

print(checksum('day2in.txt'))
print(findclose('day2in.txt'))
