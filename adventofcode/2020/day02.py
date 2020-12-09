import sys

def countRange(inp):
    total = 0
    for line in inp:
        l = line.split()
        low = int(l[0])
        hi = int(l[1])
        word = l[-1]
        letter = l[2]
        count = word.count(letter) 
        if low <= count <= hi:
            total += 1
    return total

def countMatch(inp):
    total = 0
    for line in inp:
        l = line.split()
        low = int(l[0])-1
        hi = int(l[1])-1
        word = l[-1]
        letter = l[2]
        if word[low] == letter and word[hi] == letter:
            continue
        if word[hi] == letter:
            total += 1
        if word[low] == letter:
            total += 1
    return total

inp = open(sys.argv[1]).read().strip().split('\n')
print(countRange(inp))
print(countMatch(inp))
