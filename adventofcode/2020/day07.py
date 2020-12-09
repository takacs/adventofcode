from collections import defaultdict, deque
import sys

def parse(inp):
    parents = defaultdict(list)
    contains = defaultdict(list)
    for line in inp:
        info = line.strip().split()
        parent = info[0] + info[1]
        i = 2
        while i < len(info):
            val = info[i]
            child = info[i+1] + info[i+2]
            i += 3
            parents[child].append(parent)
            contains[parent].append((child, int(val)))
    return parents, contains

def validBFS(parents, color):
    valpar = []
    fridge = deque([color])
    while len(fridge) != 0:
        curr = fridge.pop()
        if curr not in valpar:
            valpar.append(curr)
        for p in parents[curr]:
            fridge.append(p)
    return len(valpar) - 1

def totalBags(contains, color):
    total = 0
    for (bag, val) in contains[color]:
        total += val*totalBags(contains, bag)
    return total

inp = open(sys.argv[1]).read().replace('contains','').replace('bags', '').replace('bag', '').replace('.', '').replace('contain','').replace(',', '').replace(' no other', '').strip().split('\n')

parents, contains = parse(inp)
print(validBFS(parents, 'shinygold'))
print(totalBags(contains, 'shinygold'))
