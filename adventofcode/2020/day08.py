import fileinput
import sys

def run(ops):
    acc = 0
    i = 0
    iss = []
    while i < len(ops):
        iss.append(i)
        op, val = ops[i].split()
        val = int(val)
        if op == 'acc':
            acc += val
            i += 1
        elif op == 'jmp':
            i += val
        else:
            i += 1
    return acc

def findInfinite(ops):
    acc = 0
    i = 0
    iss = []
    while i < len(ops):
        if i in iss:
            return acc 
        iss.append(i)
        op, val = ops[i].split()
        val = int(val)
        if op == 'acc':
            acc += val
            i += 1
        elif op == 'jmp':
            i += val
        else:
            i += 1
    return -1
    
def findCorrupted(ops):
    for i, op in enumerate(ops):
        op, val = ops[i].split()
        if op == 'jmp':
            opsCop = ops.copy()
            opsCop[i] = 'nop ' + val
            if findInfinite(opsCop) < 0:
                return run(opsCop)
        if op == 'nop':
            opsCop = ops.copy()
            opsCop[i] = 'jmp ' + val
            if findInfinite(opsCop) < 0:
                return run(opsCop)
    return None

inp = open(sys.argv[1]).read().strip().split('\n')
print(findInfinite(inp))
print(findCorrupted(inp))
