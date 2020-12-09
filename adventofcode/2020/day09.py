import fileinput
import sys

def twoSum(vals, target):
    d = {}
    for val in vals:
        if val in d.keys():
            return True
        else:
             d[target-val] = None
    return False

def findInv(inp, presize=26):
    i = presize
    while True:
        vals = inp[i-presize:i]
        target = inp[i]
        if not twoSum(vals, target):
            return inp[i]
        else:
            i += 1
    return None

def contMinMax(inp, target):
    for i in range(len(inp)):
        cont = []
        if inp[i] == target:
           continue
        else:
            while sum(cont) < target:
                cont.append(inp[i])
                i += 1
        if sum(cont) == target:
            return min(cont) + max(cont)
    return None

inp = list(map(int, open(sys.argv[1]).readlines()))
invalid = findInv(inp, 26)
print(invalid)
print(contMinMax(inp, invalid))
