import sys
from functools import reduce
from operator import mul

def treeTrav(inp, rv, cv):
    tree = [list(row) for row in inp]
    count = 0
    row, col = 0, 0
    while row < len(tree):
        if tree[row][col%len(tree[0])] == '#':
            count += 1
        row += rv
        col += cv
    return count

inp = open(sys.argv[1]).read().strip().split('\n')
print(treeTrav(inp, 1, 3))
slopes = [(1,1), (3,1), (5,1), (7,1), (1,2)]
print(reduce(mul, [treeTrav(inp, r, c) for c, r in slopes]))
